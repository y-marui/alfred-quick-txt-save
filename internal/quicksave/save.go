package quicksave

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

// SaveText writes text to path, creating parent directories as needed, and
// posts a notification describing the outcome:
//   - text empty: posts "Clipboard is empty" and writes nothing.
//   - write succeeds: posts "Saved to <basename>".
//   - write fails: posts "Failed to save <basename>" — a Run Script
//     action's stderr isn't shown to the user, so this is the only way a
//     write failure (e.g. an unwritable save directory) becomes visible.
//
// Returns whether a file was written.
func SaveText(path, text string) (bool, error) {
	if text == "" {
		notify("Quick Save", "Clipboard is empty.")
		return false, nil
	}

	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		notify("Quick Save", fmt.Sprintf("Failed to save %s", filepath.Base(path)))
		return false, fmt.Errorf("quicksave: creating save directory: %w", err)
	}
	if err := os.WriteFile(path, []byte(text), 0o644); err != nil {
		notify("Quick Save", fmt.Sprintf("Failed to save %s", filepath.Base(path)))
		return false, fmt.Errorf("quicksave: writing file: %w", err)
	}

	notify("Quick Save", fmt.Sprintf("Saved to %s", filepath.Base(path)))
	return true, nil
}

// notify posts a macOS notification. Errors are ignored — a failed
// notification must never stop or fail the save itself (mirrors the
// Python original's check=False).
func notify(title, message string) {
	script := fmt.Sprintf("display notification %s with title %s", quoteAppleScript(message), quoteAppleScript(title))
	_ = exec.Command("osascript", "-e", script).Run()
}

// quoteAppleScript renders s as an AppleScript string literal. Go's %q
// escaping (backslash and double-quote) happens to produce valid
// AppleScript string syntax for the plain titles/messages this package
// generates.
func quoteAppleScript(s string) string {
	return fmt.Sprintf("%q", s)
}
