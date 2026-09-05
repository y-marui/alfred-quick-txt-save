package quicksave

import (
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

func requireMacOS(t *testing.T) {
	t.Helper()
	if runtime.GOOS != "darwin" {
		t.Skip("quicksave clipboard integration is macOS-only")
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

func TestSaveTextEmptyWritesNothing(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "out.txt")

	wrote, err := SaveText(path, "")
	if err != nil {
		t.Fatalf("SaveText: %v", err)
	}
	if wrote {
		t.Error("wrote = true, want false for empty text")
	}
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		t.Errorf("expected %s not to exist, stat err = %v", path, err)
	}
}

func TestSaveTextWritesContent(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "out.txt")

	wrote, err := SaveText(path, "hello world")
	if err != nil {
		t.Fatalf("SaveText: %v", err)
	}
	if !wrote {
		t.Error("wrote = false, want true for non-empty text")
	}
	got, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("ReadFile: %v", err)
	}
	if string(got) != "hello world" {
		t.Errorf("content = %q, want %q", got, "hello world")
	}
}

func TestSaveTextCreatesParentDirectories(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "nested", "sub", "out.txt")

	if _, err := SaveText(path, "content"); err != nil {
		t.Fatalf("SaveText: %v", err)
	}
	got, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("ReadFile: %v", err)
	}
	if string(got) != "content" {
		t.Errorf("content = %q, want %q", got, "content")
	}
}

func TestSaveReadsRealClipboard(t *testing.T) {
	requireMacOS(t)
	setClipboard(t, "clipboard integration test content")

	dir := t.TempDir()
	path := filepath.Join(dir, "out.txt")

	if err := Save(path); err != nil {
		t.Fatalf("Save: %v", err)
	}
	got, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("ReadFile: %v", err)
	}
	if string(got) != "clipboard integration test content" {
		t.Errorf("content = %q, want %q", got, "clipboard integration test content")
	}
}

func TestSaveEmptyClipboardWritesNothing(t *testing.T) {
	requireMacOS(t)
	setClipboard(t, "")

	dir := t.TempDir()
	path := filepath.Join(dir, "out.txt")

	if err := Save(path); err != nil {
		t.Fatalf("Save: %v", err)
	}
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		t.Errorf("expected %s not to exist, stat err = %v", path, err)
	}
}
