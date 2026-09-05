package quicksave

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestResolvePathDefault(t *testing.T) {
	home, _ := os.UserHomeDir()
	path := ResolvePath("")
	if got := filepath.Dir(path); got != filepath.Join(home, "Downloads") {
		t.Errorf("dir = %q, want %q", got, filepath.Join(home, "Downloads"))
	}
	if !strings.HasPrefix(filepath.Base(path), "quick_save_") {
		t.Errorf("base = %q, want quick_save_ prefix", filepath.Base(path))
	}
	if filepath.Ext(path) != ".txt" {
		t.Errorf("ext = %q, want .txt", filepath.Ext(path))
	}
}

func TestDefaultFilenameIncludesTime(t *testing.T) {
	path := ResolvePath("")
	stem := strings.TrimSuffix(filepath.Base(path), filepath.Ext(path))
	// quick_save_YYYYMMDD_HHMMSS -> quick(1) + save(1) + date(1) + time = 3 underscores
	if got := strings.Count(stem, "_"); got != 3 {
		t.Errorf("underscore count = %d, want 3 (stem=%q)", got, stem)
	}
}

func TestResolvePathWithFilename(t *testing.T) {
	path := ResolvePath("notes")
	if got := filepath.Base(path); got != "notes.txt" {
		t.Errorf("base = %q, want notes.txt", got)
	}
}

func TestResolvePathWithExtension(t *testing.T) {
	path := ResolvePath("notes.md")
	if got := filepath.Base(path); got != "notes.md" {
		t.Errorf("base = %q, want notes.md", got)
	}
}

func TestGetSaveDirUsesEnvVar(t *testing.T) {
	tmp := t.TempDir()
	t.Setenv("save_dir", tmp)
	if got := SaveDir(); got != tmp {
		t.Errorf("SaveDir() = %q, want %q", got, tmp)
	}
}

func TestGetSaveDirEnvVarExpanduser(t *testing.T) {
	home, _ := os.UserHomeDir()
	t.Setenv("save_dir", "~/Documents")
	want := filepath.Join(home, "Documents")
	if got := SaveDir(); got != want {
		t.Errorf("SaveDir() = %q, want %q", got, want)
	}
}

func TestGetSaveDirDefaultWhenEnvEmpty(t *testing.T) {
	home, _ := os.UserHomeDir()
	t.Setenv("save_dir", "")
	want := filepath.Join(home, "Downloads")
	if got := SaveDir(); got != want {
		t.Errorf("SaveDir() = %q, want %q", got, want)
	}
}

func TestFilePrefixEnvVar(t *testing.T) {
	t.Setenv("file_prefix", "note")
	path := ResolvePath("")
	if !strings.HasPrefix(filepath.Base(path), "note_") {
		t.Errorf("base = %q, want note_ prefix", filepath.Base(path))
	}
}

func TestFilePrefixEmptyUsesDefault(t *testing.T) {
	t.Setenv("file_prefix", "")
	path := ResolvePath("")
	if !strings.HasPrefix(filepath.Base(path), "quick_save_") {
		t.Errorf("base = %q, want quick_save_ prefix", filepath.Base(path))
	}
}

func TestFileExtEnvVar(t *testing.T) {
	t.Setenv("file_ext", ".md")
	path := ResolvePath("")
	if filepath.Ext(path) != ".md" {
		t.Errorf("ext = %q, want .md", filepath.Ext(path))
	}
}

func TestFileExtEnvVarWithoutDot(t *testing.T) {
	t.Setenv("file_ext", "md")
	path := ResolvePath("")
	if filepath.Ext(path) != ".md" {
		t.Errorf("ext = %q, want .md", filepath.Ext(path))
	}
}

func TestFileExtEmptyUsesDefault(t *testing.T) {
	t.Setenv("file_ext", "")
	path := ResolvePath("")
	if filepath.Ext(path) != ".txt" {
		t.Errorf("ext = %q, want .txt", filepath.Ext(path))
	}
}

func TestFileExtAppliedToCustomName(t *testing.T) {
	t.Setenv("file_ext", ".md")
	path := ResolvePath("notes")
	if got := filepath.Base(path); got != "notes.md" {
		t.Errorf("base = %q, want notes.md", got)
	}
}

func TestResolvePathUsesEnvDir(t *testing.T) {
	tmp := t.TempDir()
	t.Setenv("save_dir", tmp)
	path := ResolvePath("test")
	if got := filepath.Dir(path); got != tmp {
		t.Errorf("dir = %q, want %q", got, tmp)
	}
	if got := filepath.Base(path); got != "test.txt" {
		t.Errorf("base = %q, want test.txt", got)
	}
}

func TestUniquePathNoCollision(t *testing.T) {
	tmp := t.TempDir()
	t.Setenv("save_dir", tmp)
	path := ResolvePath("report")
	if got := filepath.Base(path); got != "report.txt" {
		t.Errorf("base = %q, want report.txt", got)
	}
}

func TestUniquePathOneCollision(t *testing.T) {
	tmp := t.TempDir()
	touch(t, filepath.Join(tmp, "report.txt"))
	t.Setenv("save_dir", tmp)
	path := ResolvePath("report")
	if got := filepath.Base(path); got != "report (1).txt" {
		t.Errorf("base = %q, want report (1).txt", got)
	}
}

func TestUniquePathMultipleCollisions(t *testing.T) {
	tmp := t.TempDir()
	touch(t, filepath.Join(tmp, "report.txt"))
	touch(t, filepath.Join(tmp, "report (1).txt"))
	touch(t, filepath.Join(tmp, "report (2).txt"))
	t.Setenv("save_dir", tmp)
	path := ResolvePath("report")
	if got := filepath.Base(path); got != "report (3).txt" {
		t.Errorf("base = %q, want report (3).txt", got)
	}
}

func touch(t *testing.T, path string) {
	t.Helper()
	if err := os.WriteFile(path, nil, 0o644); err != nil {
		t.Fatalf("touch %s: %v", path, err)
	}
}
