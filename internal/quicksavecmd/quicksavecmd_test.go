package quicksavecmd

import (
	"path/filepath"
	"strings"
	"testing"
)

func TestNoArgsShowsDefaultFilename(t *testing.T) {
	resp := List("")
	if len(resp.Items) != 2 {
		t.Fatalf("got %d items, want 2", len(resp.Items))
	}
	if !strings.HasPrefix(resp.Items[0].Title, "Save to quick_save_") {
		t.Errorf("title = %q, want quick_save_ prefix", resp.Items[0].Title)
	}
	if !strings.HasSuffix(resp.Items[0].Arg, ".txt") {
		t.Errorf("arg = %q, want .txt suffix", resp.Items[0].Arg)
	}
}

func TestCustomFilenameNoExt(t *testing.T) {
	resp := List("myfile")
	if got := resp.Items[0].Title; got != "Save to myfile.txt" {
		t.Errorf("title = %q, want %q", got, "Save to myfile.txt")
	}
}

func TestCustomFilenameWithExt(t *testing.T) {
	resp := List("notes.md")
	if got := resp.Items[0].Title; got != "Save to notes.md" {
		t.Errorf("title = %q, want %q", got, "Save to notes.md")
	}
}

func TestArgIsFullPath(t *testing.T) {
	tmp := t.TempDir()
	t.Setenv("save_dir", tmp)

	resp := List("out")
	want := filepath.Join(tmp, "out.txt")
	if got := resp.Items[0].Arg; got != want {
		t.Errorf("arg = %q, want %q", got, want)
	}
}

func TestSecondItemShowsSaveDir(t *testing.T) {
	tmp := t.TempDir()
	t.Setenv("save_dir", tmp)

	resp := List("")
	if !strings.Contains(resp.Items[1].Title, tmp) {
		t.Errorf("title = %q, want it to contain %q", resp.Items[1].Title, tmp)
	}
	if resp.Items[1].Valid == nil || *resp.Items[1].Valid {
		t.Errorf("second item Valid = %v, want explicit false", resp.Items[1].Valid)
	}
}
