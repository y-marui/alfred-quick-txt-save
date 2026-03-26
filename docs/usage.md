# Usage

## Quick Start

Copy text to your clipboard, then open Alfred and type `save`.

## Commands

### Save clipboard text

```
save
save <filename>
```

Saves the current clipboard contents to a file.

- No argument: generates `quick_save_YYYYMMDD.txt` in the configured save directory
- With filename (no extension): appends `.txt` automatically
- With filename and extension: uses as-is

**Examples:**

```
save                → ~/Downloads/quick_save_20260326.txt
save meeting        → ~/Downloads/meeting.txt
save meeting.md     → ~/Downloads/meeting.md
```

| Key | Action |
|---|---|
| ↩ Enter | Save clipboard to the shown path |

A macOS notification confirms the save. If the clipboard is empty, nothing is written.

### Change save directory

Open Alfred Preferences → Workflows → Quick Text Save → **Configure Workflow**.

Set the **Save Directory** field to any path (e.g. `~/Desktop`). Leave it blank to use `~/Downloads`.

The current save directory is always shown as the second line when `save` is invoked.

## Tips

- The second line shown under each `save` result is the full destination path.
- The save directory is created automatically if it does not exist.
- To quickly save selected text from any app: copy it first (⌘C), then run `save`.

## Troubleshooting

**Nothing was saved**
- Check that the clipboard is not empty before running `save`.

**File saved to wrong directory**
- Run `save dir <path>` to update the save directory.
- Run `wf config` to verify the current `save_dir` setting.

**Workflow not responding**
- Check Alfred's debugger: open Alfred → ⌘D
- Check logs: `~/Library/Logs/Alfred/Workflow/<bundle-id>.log`
