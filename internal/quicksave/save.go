package quicksave

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

// SaveText writes text to path, creating parent directories as needed, and
// posts a notification confirming the save. If text is empty, it posts a
// "Clipboard is empty" notification and writes nothing. Returns whether a
// file was written.
func SaveText(path, text string) (bool, error) {
	if text == "" {
		notify("Quick Save", "Clipboard is empty.")
		return false, nil
	}

	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return false, fmt.Errorf("quicksave: creating save directory: %w", err)
	}
	if err := os.WriteFile(path, []byte(text), 0o644); err != nil {
		return false, fmt.Errorf("quicksave: writing file: %w", err)
	}

	notify("Quick Save", fmt.Sprintf("Saved to %s", filepath.Base(path)))
	return true, nil
}

// Save reads the clipboard's current plain text and saves it to path via
// SaveText.
func Save(path string) error {
	text, err := readClipboard()
	if err != nil {
		return fmt.Errorf("quicksave: reading clipboard: %w", err)
	}
	_, err = SaveText(path, text)
	return err
}

func readClipboard() (string, error) {
	out, err := exec.Command("pbpaste").Output()
	if err != nil {
		return "", err
	}
	return string(out), nil
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
