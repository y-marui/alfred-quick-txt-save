// Package quicksavecmd builds the Alfred Script Filter response for the
// "save" keyword.
package quicksavecmd

import (
	"fmt"
	"path/filepath"

	"github.com/y-marui/alfred-quick-txt-save/internal/quicksave"
	"github.com/y-marui/alfred-quick-txt-save/internal/scriptfilter"
)

// List returns the Script Filter preview for the "save" keyword: the
// resolved destination path (Enter to save) and a second, non-actionable
// row showing the current save directory.
func List(filename string) scriptfilter.Response {
	path := quicksave.ResolvePath(filename)
	dir := quicksave.SaveDir()

	return scriptfilter.Response{Items: []scriptfilter.Item{
		{
			UID:       "save-clipboard",
			Title:     fmt.Sprintf("Save to %s", filepath.Base(path)),
			Subtitle:  path,
			Arg:       path,
			Valid:     scriptfilter.BoolPtr(true),
			Variables: map[string]string{"save_filepath": path},
		},
		{
			UID:      "save-dir-info",
			Title:    fmt.Sprintf("Save directory: %s", dir),
			Subtitle: "Change via Alfred Preferences → Workflows → Configure Workflow",
			Valid:    scriptfilter.BoolPtr(false),
		},
	}}
}
