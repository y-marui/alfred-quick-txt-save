// Package quicksave resolves the destination path for the "save" command
// and writes the clipboard's text to it.
package quicksave

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const (
	defaultDirName = "Downloads"
	defaultPrefix  = "quick_save"
	defaultExt     = ".txt"
)

// SaveDir returns the save directory.
//
// Priority: Alfred workflow variable save_dir (set via Configure Workflow)
// → ~/Downloads.
func SaveDir() string {
	if raw := strings.TrimSpace(os.Getenv("save_dir")); raw != "" {
		return expandHome(raw)
	}
	home, _ := os.UserHomeDir()
	return filepath.Join(home, defaultDirName)
}

// filePrefix returns the filename prefix from env var file_prefix or the
// default.
func filePrefix() string {
	if raw := strings.TrimSpace(os.Getenv("file_prefix")); raw != "" {
		return raw
	}
	return defaultPrefix
}

// fileExt returns the default file extension from env var file_ext or the
// default, ensuring a leading dot.
func fileExt() string {
	raw := strings.TrimSpace(os.Getenv("file_ext"))
	if raw == "" {
		return defaultExt
	}
	if !strings.HasPrefix(raw, ".") {
		raw = "." + raw
	}
	return raw
}

// ResolvePath returns a non-colliding save path for filename.
//
//   - No filename: generates "{prefix}_YYYYMMDD_HHMMSS{ext}".
//   - Filename without an extension: appends the configured default
//     extension.
//   - If the resolved path already exists, appends " (1)", " (2)", …
//     before the extension until a free name is found.
func ResolvePath(filename string) string {
	dir := SaveDir()
	filename = strings.TrimSpace(filename)
	if filename == "" {
		ts := time.Now().Format("20060102_150405")
		filename = fmt.Sprintf("%s_%s%s", filePrefix(), ts, fileExt())
	} else if !strings.Contains(filepath.Base(filename), ".") {
		filename += fileExt()
	}
	return uniquePath(filepath.Join(dir, filename))
}

// uniquePath returns path if it does not exist, otherwise "stem (N).ext".
func uniquePath(path string) string {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return path
	}
	dir, base := filepath.Split(path)
	ext := filepath.Ext(base)
	stem := strings.TrimSuffix(base, ext)
	for i := 1; ; i++ {
		candidate := filepath.Join(dir, fmt.Sprintf("%s (%d)%s", stem, i, ext))
		if _, err := os.Stat(candidate); os.IsNotExist(err) {
			return candidate
		}
	}
}

// expandHome expands a leading "~" or "~/..." in path using the current
// user's home directory. Paths not starting with "~" are returned as-is.
func expandHome(path string) string {
	if path == "~" {
		home, _ := os.UserHomeDir()
		return home
	}
	if strings.HasPrefix(path, "~/") {
		home, _ := os.UserHomeDir()
		return filepath.Join(home, path[2:])
	}
	return path
}
