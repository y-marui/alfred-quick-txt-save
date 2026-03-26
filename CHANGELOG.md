# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added

- `save [filename]` command: saves clipboard text to a .txt file via Alfred Script Filter
- `save dir <path>` subcommand: sets and persists the save directory
- `save_service`: path resolution with configurable save directory (default `~/Downloads`)
- `save_text.py`: standalone Run Script that reads `pbpaste` and writes the file
- macOS notification on successful save (via `osascript`)
- Auto-generated filename `quick_save_YYYYMMDD.txt` when no filename is given
- Auto-appended `.txt` extension when filename has no extension

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

[Unreleased]: https://github.com/y-marui/alfred-quick-txt-save/compare/v0.1.0...HEAD
[0.1.0]: https://github.com/y-marui/alfred-quick-txt-save/releases/tag/v0.1.0
