package quicksave

import (
	"os"
	"path/filepath"
	"testing"
)

func TestSaveTextEmptyWritesNothing(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "out.txt")

	wrote, message, err := SaveText(path, "")
	if err != nil {
		t.Fatalf("SaveText: %v", err)
	}
	if wrote {
		t.Error("wrote = true, want false for empty text")
	}
	if want := "Clipboard is empty."; message != want {
		t.Errorf("message = %q, want %q", message, want)
	}
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		t.Errorf("expected %s not to exist, stat err = %v", path, err)
	}
}

func TestSaveTextWritesContent(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "out.txt")

	wrote, message, err := SaveText(path, "hello world")
	if err != nil {
		t.Fatalf("SaveText: %v", err)
	}
	if !wrote {
		t.Error("wrote = false, want true for non-empty text")
	}
	if want := "Saved to out.txt"; message != want {
		t.Errorf("message = %q, want %q", message, want)
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

	if _, _, err := SaveText(path, "content"); err != nil {
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

func TestSaveTextWriteFailureReturnsError(t *testing.T) {
	dir := t.TempDir()
	// blocker is a file, not a directory, so MkdirAll for a path underneath
	// it must fail.
	blocker := filepath.Join(dir, "blocker")
	if err := os.WriteFile(blocker, []byte("x"), 0o644); err != nil {
		t.Fatalf("setup: %v", err)
	}
	path := filepath.Join(blocker, "out.txt")

	wrote, message, err := SaveText(path, "content")
	if err == nil {
		t.Fatal("SaveText: expected error when parent path is not a directory")
	}
	if wrote {
		t.Error("wrote = true, want false on failure")
	}
	if want := "Failed to save out.txt"; message != want {
		t.Errorf("message = %q, want %q", message, want)
	}
}
