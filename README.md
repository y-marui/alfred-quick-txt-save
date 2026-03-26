# Alfred Quick Text Save

> **This is the English (reference) version.**
> For the Japanese canonical version, see [README-jp.md](README-jp.md).

> Save clipboard or selected text to a .txt file with one keystroke.

[![CI](https://github.com/y-marui/alfred-quick-txt-save/actions/workflows/ci.yml/badge.svg)](https://github.com/y-marui/alfred-quick-txt-save/actions/workflows/ci.yml)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)

## Features

- Save clipboard text to a file with a single Alfred command
- Configurable save directory (default: `~/Downloads`)
- Auto-generated filename with today's date (`quick_save_YYYYMMDD.txt`)
- Custom filename support — with or without extension
- macOS notification on save

## Requirements

- Alfred 5 (Powerpack required for Script Filter)
- Python 3.9+

## Installation

1. Download the latest `.alfredworkflow` from [Releases](https://github.com/y-marui/alfred-quick-txt-save/releases)
2. Double-click to install in Alfred

## Usage

### Save clipboard text

```
save              → ~/Downloads/quick_save_20260326.txt
save notes        → ~/Downloads/notes.txt
save notes.md     → ~/Downloads/notes.md
```

Copy text to your clipboard, then type `save` in Alfred and press Enter.

| Key | Action |
|---|---|
| ↩ Enter | Save clipboard to the shown path |

### Change save directory

```
save dir ~/Desktop
save dir /Users/you/Documents/notes
```

The setting is persisted and used for all future saves.

### View current settings

```
wf config
```

## Quick Start (developers)

```bash
git clone https://github.com/y-marui/alfred-quick-txt-save
cd alfred-quick-txt-save

# Install dev dependencies
make install

# Simulate Alfred locally
make run Q="save"
make run Q="save mynotes"
make run Q="save dir ~/Desktop"

# Run tests
make test

# Build workflow package
make build
# → dist/alfred-quick-txt-save-0.1.0.alfredworkflow
```

## Project Structure

```
alfred-quick-txt-save/
├── src/
│   ├── alfred/         # Alfred SDK (response, router, cache, config, logger, safe_run)
│   └── app/            # Application layer (commands, services)
├── workflow/           # Alfred package (info.plist, scripts/)
│   └── scripts/
│       ├── entry.py    # Script Filter entrypoint
│       └── save_text.py  # Run Script: reads clipboard and writes file
├── tests/              # pytest test suite
├── scripts/            # build.sh, dev.sh, release.sh, vendor.sh
└── docs/               # Architecture, development, and usage documentation
```

## Documentation

| Document | Description |
|---|---|
| [docs/architecture.md](docs/architecture.md) | Full architecture and layer design |
| [docs/development.md](docs/development.md) | Adding commands, managing dependencies, release |
| [docs/usage.md](docs/usage.md) | End-user usage guide |

## Release

```bash
# 1. Bump version in pyproject.toml
# 2. Tag and push
git tag v1.2.3
git push --tags
# GitHub Actions builds .alfredworkflow and creates a GitHub Release
```

## License

MIT — see [LICENSE](LICENSE)
