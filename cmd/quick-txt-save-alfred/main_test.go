package main

import (
	"bytes"
	"encoding/json"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

func buildBinary(t *testing.T) string {
	t.Helper()
	bin := filepath.Join(t.TempDir(), "quick-txt-save-alfred")
	cmd := exec.Command("go", "build", "-o", bin, ".")
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("go build: %v\n%s", err, out)
	}
	return bin
}

func runBinary(t *testing.T, bin string, env []string, args ...string) (stdout, stderr string, exitCode int) {
	t.Helper()
	cmd := exec.Command(bin, args...)
	if env != nil {
		cmd.Env = env
	}
	var outBuf, errBuf bytes.Buffer
	cmd.Stdout = &outBuf
	cmd.Stderr = &errBuf
	err := cmd.Run()
	code := 0
	if exitErr, ok := err.(*exec.ExitError); ok {
		code = exitErr.ExitCode()
	} else if err != nil {
		t.Fatalf("running binary: %v", err)
	}
	return outBuf.String(), errBuf.String(), code
}

type scriptFilterItem struct {
	UID       string            `json:"uid"`
	Title     string            `json:"title"`
	Subtitle  string            `json:"subtitle"`
	Arg       string            `json:"arg"`
	Valid     *bool             `json:"valid"`
	Variables map[string]string `json:"variables"`
}

type scriptFilterResponse struct {
	Items []scriptFilterItem `json:"items"`
}

type workflowVariablesEnvelope struct {
	Alfredworkflow struct {
		Variables map[string]string `json:"variables"`
	} `json:"alfredworkflow"`
}

func decodeMessage(t *testing.T, stdout string) string {
	t.Helper()
	var env workflowVariablesEnvelope
	if err := json.Unmarshal([]byte(stdout), &env); err != nil {
		t.Fatalf("Unmarshal(%q): %v", stdout, err)
	}
	return env.Alfredworkflow.Variables["message"]
}

func TestListPrintsValidJSON(t *testing.T) {
	bin := buildBinary(t)
	tmp := t.TempDir()
	env := append(os.Environ(), "save_dir="+tmp)

	stdout, stderr, code := runBinary(t, bin, env, "list", "notes.md")
	if code != 0 {
		t.Fatalf("exit code = %d, stderr = %s", code, stderr)
	}

	var resp scriptFilterResponse
	if err := json.Unmarshal([]byte(stdout), &resp); err != nil {
		t.Fatalf("Unmarshal(%q): %v", stdout, err)
	}
	if len(resp.Items) != 2 {
		t.Fatalf("got %d items, want 2", len(resp.Items))
	}
	want := filepath.Join(tmp, "notes.md")
	if resp.Items[0].Arg != want {
		t.Errorf("arg = %q, want %q", resp.Items[0].Arg, want)
	}
}

func TestMissingSubcommandPrintsErrorJSON(t *testing.T) {
	bin := buildBinary(t)

	stdout, _, code := runBinary(t, bin, nil)
	if code == 0 {
		t.Error("exit code = 0, want non-zero")
	}
	var resp scriptFilterResponse
	if err := json.Unmarshal([]byte(stdout), &resp); err != nil {
		t.Fatalf("Unmarshal(%q): %v", stdout, err)
	}
	if len(resp.Items) != 1 || resp.Items[0].Title != "Workflow Error" {
		t.Errorf("items = %+v, want a single Workflow Error item", resp.Items)
	}
}

func TestWriteRequiresPathArgument(t *testing.T) {
	bin := buildBinary(t)

	_, stderr, code := runBinary(t, bin, nil, "write")
	if code == 0 {
		t.Error("exit code = 0, want non-zero")
	}
	if stderr == "" {
		t.Error("expected a stderr message when path argument is missing")
	}
}

func TestWriteUsesTextEnvVar(t *testing.T) {
	bin := buildBinary(t)
	tmp := t.TempDir()
	path := filepath.Join(tmp, "out.txt")
	env := append(os.Environ(), "text=quick-txt-save-alfred integration test content")

	stdout, stderr, code := runBinary(t, bin, env, "write", path)
	if code != 0 {
		t.Fatalf("exit code = %d, stderr = %s", code, stderr)
	}
	got, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("ReadFile: %v", err)
	}
	if string(got) != "quick-txt-save-alfred integration test content" {
		t.Errorf("content = %q, want %q", got, "quick-txt-save-alfred integration test content")
	}
	if want := "Saved to out.txt"; decodeMessage(t, stdout) != want {
		t.Errorf("message = %q, want %q", decodeMessage(t, stdout), want)
	}
}

func TestWriteEmptyTextEnvVarWritesNothing(t *testing.T) {
	bin := buildBinary(t)
	tmp := t.TempDir()
	path := filepath.Join(tmp, "out.txt")

	stdout, stderr, code := runBinary(t, bin, os.Environ(), "write", path)
	if code != 0 {
		t.Fatalf("exit code = %d, stderr = %s", code, stderr)
	}
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		t.Errorf("expected %s not to exist, stat err = %v", path, err)
	}
	if want := "Clipboard is empty."; decodeMessage(t, stdout) != want {
		t.Errorf("message = %q, want %q", decodeMessage(t, stdout), want)
	}
}

func TestWriteFailurePrintsFailureMessage(t *testing.T) {
	bin := buildBinary(t)
	tmp := t.TempDir()
	// blocker is a file, not a directory, so the write's parent-dir creation
	// fails.
	blocker := filepath.Join(tmp, "blocker")
	if err := os.WriteFile(blocker, []byte("x"), 0o644); err != nil {
		t.Fatalf("setup: %v", err)
	}
	path := filepath.Join(blocker, "out.txt")
	env := append(os.Environ(), "text=content")

	stdout, _, _ := runBinary(t, bin, env, "write", path)
	if want := "Failed to save out.txt"; decodeMessage(t, stdout) != want {
		t.Errorf("message = %q, want %q", decodeMessage(t, stdout), want)
	}
}
