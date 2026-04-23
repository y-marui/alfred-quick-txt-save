# UI Design

Alfred Quick Text Save presents results through Alfred's Script Filter JSON format.

## Script Filter Items

All items are returned via `alfred.response.output()` as a list of Alfred Script Filter items.
Unicode emoji are not used in `title` or `subtitle` per UI guidelines (use ASCII symbols or no decoration).

### `save` command

| Field | Value |
|---|---|
| `title` | `Save to <filename>` |
| `subtitle` | Full resolved path (e.g. `~/Downloads/quick_save_20260326_143012.txt`) |
| `arg` | Full resolved path — passed to the Run Script node |
| `uid` | `save-clipboard` |

A second informational item shows the current save directory:

| Field | Value |
|---|---|
| `title` | `Save directory: <path>` |
| `subtitle` | `Change via Alfred Preferences → Workflows → Configure Workflow` |
| `valid` | `false` (not selectable) |
| `uid` | `save-dir-info` |

### `wf config` command

Shows current configuration key-value pairs as individual items, each with:
- `title`: setting name and value
- `valid`: `false` (display only)

A "Reset" action item with `valid: true` triggers `wf config reset`.

### Error items

When an unhandled exception occurs, `safe_run()` catches it and returns a single error item:
- `title`: `Error: <exception type>`
- `subtitle`: exception message
- `valid`: `false`

## Icon

`workflow/icon.png` — single PNG icon used for all items. Light/dark mode is handled by Alfred.

## Keyboard Actions

| Key | Action |
|---|---|
| ↩ Enter | Confirm save / execute action |

No modifier key actions are currently defined.
