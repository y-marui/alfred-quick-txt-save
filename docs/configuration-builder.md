# Configuration Builder

Alfred's **Configure Workflow** panel provides a user-facing UI to set workflow variables.
These map directly to environment variables read by `src/app/services/save_service.py`.

## Variables

| Variable | Label (in UI) | Type | Default | Description |
|---|---|---|---|---|
| `save_dir` | Save Directory | Text | _(empty → `~/Downloads`)_ | Directory where saved files are written. Supports `~` expansion. |
| `file_prefix` | Filename Prefix | Text | `quick_save` | Prefix for auto-generated filenames (when no filename is given). |
| `file_ext` | Default Extension | Text | `.txt` | Extension appended when the user does not provide one. Leading `.` is added automatically if omitted. |

## Variable Priority

All variables are read from `os.environ` at call time:

1. Alfred workflow variable (set via Configure Workflow) — highest priority
2. Hardcoded default in `save_service.py` — fallback

## Accessing the Panel

Alfred Preferences → Workflows → Quick Text Save → **Configure Workflow** (the `[x]` button).

## info.plist

Workflow variables are declared in `workflow/info.plist` under the `variables` key.
When adding a new variable:
1. Add it to `workflow/info.plist` → `variables`
2. Read it in the appropriate service with `os.environ.get("var_name", "default")`
3. Document it in this file and in `docs/specification.md`
