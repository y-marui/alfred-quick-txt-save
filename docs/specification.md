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

- If clipboard is empty, nothing is written and a "Clipboard is empty" notification is posted.
- If the resolved path already exists, appends `(1)`, `(2)`, … before the extension until a free name is found.
- A macOS notification confirms the save, or reports the failure if the write itself fails
  (e.g. an unwritable save directory).
- The Script Filter (`list` subcommand) previews the resolved path. The actual write happens
  in the `write` subcommand after Alfred passes the path as `arg` and the clipboard's text as
  the `text` variable (via an Arguments and Variables node using Alfred's `{clipboard}`
  placeholder — the binary itself never reads the pasteboard).

---

## Configuration

Configuration is managed through Alfred Preferences → Workflows → Quick Text Save →
**Configure Workflow**, and read by the binary as environment variables.

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
Alfred (Script Filter node, keyword "save")
  │  query string (filename, or empty)
  ▼
cmd/quick-txt-save-alfred list [filename]
  │
  ▼
internal/quicksavecmd.List()      ← resolves path, returns Script Filter items
  │
  ▼
Alfred (Arguments and Variables node)
  │  argument = {query} (the resolved path, passed through unchanged)
  │  sets variable text = {clipboard}
  ▼
Alfred (Run Script node)
  │  arg = resolved file path; env var text = clipboard's plain text
  ▼
cmd/quick-txt-save-alfred write <path>
  │
  ▼
internal/quicksave.SaveText()     ← writes file, sends macOS notification (or reports failure)
```

## Error Handling

- Any panic in the `list` subcommand's dispatch is recovered in
  `cmd/quick-txt-save-alfred/main.go`, which emits a single `Workflow Error` result item —
  a Script Filter step must always print valid JSON.
- The `write` subcommand has no JSON contract (it's a Run Script action); an internal error is
  printed to stderr and the process exits non-zero.

## Constraints

- Script Filter response time target: **< 100 ms** (compiled Go binary, no I/O beyond a
  single `os.Stat`/env var read per invocation).
- All Script Filter output must go through `scriptfilter.Response.Write()` — never
  `fmt.Print*` directly.
- `cmd/quick-txt-save-alfred` contains no business logic; it only parses `os.Args`, recovers
  panics (for `list`), and writes what `internal/quicksavecmd`/`internal/quicksave` return.
