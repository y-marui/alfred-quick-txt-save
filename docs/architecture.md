# Architecture

## Overview

An Alfred Workflow (Go): `cmd/quick-txt-save-alfred` is the single universal (amd64+arm64)
binary `workflow/info.plist` invokes. The "save" Script Filter node runs it as
`list [filename]` to preview the resolved destination path; pressing Enter runs it again as
`write <path>` (a Run Script action), which reads the clipboard's plain text, writes it to
`path`, and posts a macOS notification.

## Entry Points

- `cmd/quick-txt-save-alfred/main.go` — subcommands `list` and `write`, one per
  `workflow/info.plist` node's script (see the package doc comment in `main.go`).

## Directory Structure

| Directory | Role |
|---|---|
| `internal/scriptfilter/` | Alfred Script Filter JSON response types |
| `internal/quicksave/` | Path resolution, collision avoidance, and clipboard-to-file write logic — Alfred-independent, stdlib only |
| `internal/quicksavecmd/` | Builds the Script Filter response for the `list` step from `internal/quicksave` |
| `cmd/quick-txt-save-alfred/` | The binary Alfred invokes; dispatches to `internal/quicksavecmd`/`internal/quicksave` and prints Script Filter JSON |
| `workflow/` | Alfred package: `info.plist`, icon |
| `scripts/` | `build-workflow.sh`, `extract-changelog.sh` |

## Layers

```
Alfred
  │  "save" keyword + query string, or Enter on a result
  ▼
cmd/quick-txt-save-alfred/main.go   ← Alfred boundary; argv dispatch only
  │
  ├─ list [filename] ─┐
  │                    ▼
  │              internal/quicksavecmd.List()   ← builds Script Filter response
  │                    │
  │                    ▼
  │              internal/quicksave.ResolvePath()/.SaveDir()
  │
  └─ write <path> ─────┐
                        ▼
                  internal/quicksave.Save()      ← reads clipboard, writes file, notifies
```

Dependency direction: `cmd → quicksavecmd → quicksave`, and `quicksavecmd → scriptfilter`.
`internal/quicksave` never imports `internal/scriptfilter` — it has no Alfred JSON concerns.

## Key Dependencies

| Library / Module | Purpose |
|---|---|
| `internal/scriptfilter` | Script Filter JSON output |
| `pbpaste`/`pbcopy` (macOS system binaries) | Clipboard read (`internal/quicksave`) |
| `osascript` (macOS system binary) | Save/empty-clipboard notification (`internal/quicksave`) |

No third-party Go dependencies (`go.mod` has no `require` block).
