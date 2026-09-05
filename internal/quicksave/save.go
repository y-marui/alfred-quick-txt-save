package quicksave

import (
	"fmt"
	"os"
	"path/filepath"
)

// SaveText writes text to path, creating parent directories as needed, and
// returns a message describing the outcome for the caller to surface (e.g.
// via a native Alfred Post Notification node reading a workflow variable):
//   - text empty: "Clipboard is empty." and writes nothing.
//   - write succeeds: "Saved to <basename>".
//   - write fails: "Failed to save <basename>" — a Run Script action's
//     stderr isn't shown to the user, so this is the only way a write
//     failure (e.g. an unwritable save directory) becomes visible.
//
// Returns whether a file was written.
func SaveText(path, text string) (bool, string, error) {
	if text == "" {
		return false, "Clipboard is empty.", nil
	}

	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return false, fmt.Sprintf("Failed to save %s", filepath.Base(path)), fmt.Errorf("quicksave: creating save directory: %w", err)
	}
	if err := os.WriteFile(path, []byte(text), 0o644); err != nil {
		return false, fmt.Sprintf("Failed to save %s", filepath.Base(path)), fmt.Errorf("quicksave: writing file: %w", err)
	}

	return true, fmt.Sprintf("Saved to %s", filepath.Base(path)), nil
}
