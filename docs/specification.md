# Specification

Functional specification and behavior definition for alfred-quick-txt-save.

## Commands

### `save` — Save Clipboard to File

**Trigger:** Type `save` in Alfred (optionally followed by a filename).

**Behavior:**

| Input | Result |
|---|---|
| `save` | Generates `{prefix}_{YYYYMMDD_HHMMSS}{ext}` in save directory |
| `save <name>` (no extension) | Saves as `<name>{ext}` in save directory |
| `save <name>.<ext>` | Saves as `<name>.<ext>` in save directory |

- If clipboard is empty, nothing is written and no error is raised.
- If the resolved path already exists, appends `(1)`, `(2)`, … before the extension until a free name is found.
- A macOS notification confirms the save.
- The Script Filter (entry.py) previews the resolved path. The actual write happens in `save_text.py` after Alfred passes the path as arg.

### `wf config` — View / Reset Configuration

**Trigger:** `wf config` or `wf config reset`.

| Input | Result |
|---|---|
| `wf config` | Shows current configuration values |
| `wf config reset` | Clears all stored configuration |

### `open` — Open Resources

**Trigger:** `wf open <target>`

| Target | Action |
|---|---|
| `repo` | Opens the GitHub repository in the default browser |

### `search` — Default Fallback

Any unrecognized keyword is routed to `search`. Used as the default dispatch fallback.

### `help` — Show Help

**Trigger:** `wf help`

Lists available commands with brief descriptions.

---

## Configuration

Configuration is stored via Alfred's `alfred_workflow_data` directory and managed through
Alfred Preferences → Workflows → Quick Text Save → **Configure Workflow**.

| Variable | Description | Default |
|---|---|---|
| `save_dir` | Directory where files are saved | `~/Downloads` |
| `file_prefix` | Prefix for auto-generated filenames | `quick_save` |
| `file_ext` | Extension appended when none is specified | `.txt` |

Priority: Alfred workflow variable (Configure Workflow) → default value.

---

## File Naming Rules

Auto-generated filenames (when no filename is given) follow:

```
{file_prefix}_{YYYYMMDD_HHMMSS}{file_ext}
```

Example: `quick_save_20260326_143012.txt`

Collision avoidance: if the target path exists, the suffix `(N)` is inserted before the extension:

```
quick_save_20260326_143012.txt
quick_save_20260326_143012 (1).txt
quick_save_20260326_143012 (2).txt
```

---

## Data Flow

```
Alfred (Script Filter node)
  │  keyword + query string
  ▼
workflow/scripts/entry.py     ← sets sys.path, calls safe_run(main)
  │
  ▼
src/app/core.py               ← router.dispatch(query)
  │
  ▼
src/alfred/router.py          ← splits "save notes" → command="save", args="notes"
  │
  ▼
src/app/commands/save_cmd.py  ← resolves path, returns Script Filter items
  │
  ▼
Alfred (Run Script node)
  │  arg = resolved file path
  ▼
workflow/scripts/save_text.py ← reads pbpaste, writes file, sends macOS notification
```
