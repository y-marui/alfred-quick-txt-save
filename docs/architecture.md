# Architecture

## Overview

An Alfred Workflow (Go): `cmd/quick-txt-save-alfred` is the single universal (amd64+arm64)
binary `workflow/info.plist` invokes. The "save" Script Filter node runs it as
`list [filename]` to preview the resolved destination path. Pressing Enter passes the item's
`arg` through an Arguments and Variables node, which also sets a `text` variable to Alfred's
`{clipboard}` placeholder — the binary itself never reads the pasteboard. The binary then
runs again as `write <path>` (a Run Script action), reading `text` from the environment,
writing it to `path`, and printing Alfred's workflow-variables JSON envelope with a `message`
describing the outcome. A native Post Notification node downstream, reached unconditionally,
displays `{message}` — the binary itself never calls `osascript`.

## Entry Points

- `cmd/quick-txt-save-alfred/main.go` — subcommands `list` and `write`, one per
  `workflow/info.plist` node's script (see the package doc comment in `main.go`).
- An `alfred.workflow.utility.argument` node sits between the two Script Filter/Run Script
  steps in `workflow/info.plist`, carrying `{clipboard}` into the `write` step's `text`
  variable while passing the resolved path through unchanged as `{query}`/`$1`.
- An `alfred.workflow.output.notification` node sits after the `write` step, reached
  unconditionally, displaying the `{message}` variable that step's stdout sets.

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
                  internal/quicksave.SaveText()  ← writes file, returns outcome message
                        │
                        ▼
                  main.writeVariables()  ← prints {"alfredworkflow":{"variables":{"message":…}}}
                        │
                        ▼ (native, unconditional)
                  Post Notification node  ← displays {message}
```

Dependency direction: `cmd → quicksavecmd → quicksave`, and `quicksavecmd → scriptfilter`.
`internal/quicksave` never imports `internal/scriptfilter` — it has no Alfred JSON concerns.
It also never reads the clipboard itself — that's `workflow/info.plist`'s job via `{clipboard}`.

## Why the file write stays in Go

Alfred's native "Write File" output object writes relative to one of three workflow-scoped
folders (data/cache/workflow folder); its documentation does not confirm it can target an
arbitrary absolute path outside that sandbox. This workflow must write to whatever directory
the user configures via `save_dir` (any folder, anywhere), so the write stays in `SaveText`.
If this is ever re-verified in Alfred directly (build a throwaway Write File node, point it
at e.g. `~/Desktop/test.txt`, confirm it actually lands there), this constraint may no longer
apply — see [`docs/alfred-workflow-notes/workflow-object-schema.md`](../../docs/alfred-workflow-notes/workflow-object-schema.md)'s
Write File entry, which carries the same open question for every project in this ecosystem.

## Why the notification moved to native Alfred

The notification used to be posted from Go via `exec.Command("osascript", "-e", "display
notification …")`, on the reasoning that Alfred's `alfred.workflow.action.script` (Run
Script) node has no way to branch a downstream connection on the script's exit code — a
native Post Notification node reached unconditionally after `write` would always fire, so it
looked like it couldn't distinguish success from failure.

That reasoning conflated two different things: *branching to a different node* (which really
does require exit-code-based routing Alfred doesn't have) and *changing what one always-run
node displays* (which doesn't). `write` no longer decides whether to notify — it always
prints Alfred's workflow-variables JSON envelope (`{"alfredworkflow":{"variables":{"message":
"…"}}}`) with the right text for whichever outcome occurred, and a single Post Notification
node, wired unconditionally after `write`, displays `{message}`. `SaveText` now returns that
message string instead of calling `notify()`; `main.go`'s `writeVariables` is the only place
that talks to Alfred's JSON contract for this step, mirroring how `internal/quicksave` never
touches `internal/scriptfilter` for the `list` step.

## Key Dependencies

| Library / Module | Purpose |
|---|---|
| `internal/scriptfilter` | Script Filter JSON output |

No third-party Go dependencies (`go.mod` has no `require` block), and no external process is
shelled out to — the binary is stdlib-only.
