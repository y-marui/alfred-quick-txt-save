# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Changed

- The save-outcome notification (saved / clipboard empty / write failed) is now posted by a
  native Post Notification node reading a `{message}` workflow variable, instead of the Go
  binary shelling out to `osascript`. The binary is now stdlib-only — no external process is
  invoked at all.

## [1.0.0] - 2026-09-05

### Added

- Rewrite in Go (`cmd/`+`internal/` layout, matching sibling `alfred-*` workflows)
- `save [filename]` command ported 1:1 from the Python prototype: path resolution
  (`{prefix}_{YYYYMMDD_HHMMSS}{ext}` default filename, extension auto-append, collision
  avoidance via `(1)`, `(2)`, …), clipboard-to-file write, macOS notification on save, on an
  empty clipboard, or on a write failure
- Single universal (amd64+arm64) static binary — no Python/uv runtime or vendored
  dependencies required
- Clipboard text is read via Alfred's own `{clipboard}` placeholder (an Arguments and
  Variables node), not by shelling out to `pbpaste` — the binary never touches the
  pasteboard itself

### Removed

- The unused `alfred-workflow-template` scaffold (`search`/`open`/`config`/`help` commands,
  the `wf` keyword Script Filter node, `ExampleService`/`ApiClient`) — never customized past
  the template defaults and not part of this workflow's actual product surface

## [0.1.0] - 2024-01-01

### Added

- Initial release based on Alfred Workflow Template
- Alfred SDK: `response`, `cache`, `config`, `logger`, `router`, `safe_run`
- Command-based UX: `search`, `open`, `config`, `help`
- Vendor packaging via `scripts/vendor.sh`
- Build pipeline via `scripts/build.sh`
- GitHub Actions CI (lint, test, build)
- GitHub Actions Release (tag → `.alfredworkflow` → GitHub Release)
- Full pytest test suite

[Unreleased]: https://github.com/y-marui/alfred-quick-txt-save/compare/v1.0.0...HEAD
[1.0.0]: https://github.com/y-marui/alfred-quick-txt-save/releases/tag/v1.0.0
[0.1.0]: https://github.com/y-marui/alfred-quick-txt-save/releases/tag/v0.1.0
