package main

import (
	"bytes"
	"encoding/json"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"testing"
)

// TestMain saves the real clipboard's plain-text content once before any
// test in this package runs, and restores it once after they all finish,
// so a local run doesn't clobber the developer's clipboard.
func TestMain(m *testing.M) {
	original, _ := exec.Command("pbpaste").Output()
	code := m.Run()
	restoreClipboard(original)
	os.Exit(code)
}

func restoreClipboard(text []byte) {
	cmd := exec.Command("pbcopy")
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return
	}
	if err := cmd.Start(); err != nil {
		return
	}
	_, _ = stdin.Write(text)
	_ = stdin.Close()
	_ = cmd.Wait()
}

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

func TestWriteSavesRealClipboard(t *testing.T) {
	requireMacOS(t)
	bin := buildBinary(t)

	setClipboard(t, "quick-txt-save-alfred integration test content")

	tmp := t.TempDir()
	path := filepath.Join(tmp, "out.txt")

	_, stderr, code := runBinary(t, bin, os.Environ(), "write", path)
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
}

func requireMacOS(t *testing.T) {
	t.Helper()
	if runtime.GOOS != "darwin" {
		t.Skip("clipboard integration is macOS-only")
	}
	for _, bin := range []string{"pbcopy", "pbpaste", "osascript"} {
		if _, err := exec.LookPath(bin); err != nil {
			t.Skipf("%s not found", bin)
		}
	}
}

func setClipboard(t *testing.T, text string) {
	t.Helper()
	cmd := exec.Command("pbcopy")
	stdin, err := cmd.StdinPipe()
	if err != nil {
		t.Fatalf("pbcopy stdin: %v", err)
	}
	if err := cmd.Start(); err != nil {
		t.Fatalf("pbcopy start: %v", err)
	}
	if _, err := stdin.Write([]byte(text)); err != nil {
		t.Fatalf("pbcopy write: %v", err)
	}
	if err := stdin.Close(); err != nil {
		t.Fatalf("pbcopy close: %v", err)
	}
	if err := cmd.Wait(); err != nil {
		t.Fatalf("pbcopy wait: %v", err)
	}
}
