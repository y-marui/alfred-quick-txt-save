# Architecture

## Overview

An Alfred Workflow (Go): `cmd/quick-txt-save-alfred` is the single universal (amd64+arm64)
binary `workflow/info.plist` invokes. The "save" Script Filter node runs it as
`list [filename]` to preview the resolved destination path. Pressing Enter passes the item's
`arg` through an Arguments and Variables node, which also sets a `text` variable to Alfred's
`{clipboard}` placeholder — the binary itself never reads the pasteboard. The binary then
runs again as `write <path>` (a Run Script action), reading `text` from the environment,
writing it to `path`, and posting a macOS notification.

## Entry Points

- `cmd/quick-txt-save-alfred/main.go` — subcommands `list` and `write`, one per
  `workflow/info.plist` node's script (see the package doc comment in `main.go`).
- An `alfred.workflow.utility.argument` node sits between the two Script Filter/Run Script
  steps in `workflow/info.plist`, carrying `{clipboard}` into the `write` step's `text`
  variable while passing the resolved path through unchanged as `{query}`/`$1`.

## Directory Structure

| Directory | Role |
|---|---|
| `internal/scriptfilter/` | Alfred Script Filter JSON response types |
| `internal/quicksave/` | Path resolution, collision avoidance, and text-to-file write logic — Alfred-independent, stdlib only |
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
  └─ write <path> ─────┐   (text = $text env var, set upstream from {clipboard})
                        ▼
                  internal/quicksave.SaveText()  ← writes file, notifies (success/empty/failure)
```

Dependency direction: `cmd → quicksavecmd → quicksave`, and `quicksavecmd → scriptfilter`.
`internal/quicksave` never imports `internal/scriptfilter` — it has no Alfred JSON concerns.
It also never reads the clipboard itself — that's `workflow/info.plist`'s job via `{clipboard}`.

## Why the file write and notification stay in Go

Two other pieces of `internal/quicksave` were considered for a move to native Alfred
objects (alongside the clipboard read, which did move — see above) and deliberately kept in
Go:

- **Writing the file** — Alfred's native "Write File" output object writes relative to one
  of three workflow-scoped folders (data/cache/workflow folder); its documentation does not
  confirm it can target an arbitrary absolute path outside that sandbox. This workflow must
  write to whatever directory the user configures via `save_dir` (any folder, anywhere), so
  the write stays in `SaveText`. If this is ever re-verified in Alfred directly (build a
  throwaway Write File node, point it at e.g. `~/Desktop/test.txt`, confirm it actually lands
  there), this constraint may no longer apply.
- **Posting the notification** — Alfred's `alfred.workflow.action.script` (Run Script) node
  has no way to branch a downstream connection on the script's exit code (only on modifier
  keys held). A native Post Notification node reached unconditionally after the Run Script
  step would show "Saved" even when the write failed. Keeping the notification in
  `SaveText`, decided in the same process that performs the write, is what lets it correctly
  distinguish success, empty clipboard, and failure.

## Key Dependencies

| Library / Module | Purpose |
|---|---|
| `internal/scriptfilter` | Script Filter JSON output |
| `osascript` (macOS system binary) | Save/empty-clipboard/failure notification (`internal/quicksave`) |

No third-party Go dependencies (`go.mod` has no `require` block).
