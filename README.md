# Alfred Quick Text Save

> **This is the English (reference) version.**
> For the Japanese canonical version, see [README-jp.md](README-jp.md).

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)
[![CI](https://github.com/y-marui/alfred-quick-txt-save/actions/workflows/ci.yml/badge.svg)](https://github.com/y-marui/alfred-quick-txt-save/actions/workflows/ci.yml)
[![Charter Check](https://github.com/y-marui/alfred-quick-txt-save/actions/workflows/dev-charter-check.yml/badge.svg)](https://github.com/y-marui/alfred-quick-txt-save/actions/workflows/dev-charter-check.yml)

Save clipboard or selected text to a .txt file with one keystroke.

## Requirements

- Alfred 5 (Powerpack required for Script Filter)
- Python 3.9+

## Setup

### End User

1. Download the latest `.alfredworkflow` from [Releases](https://github.com/y-marui/alfred-quick-txt-save/releases)
2. Double-click to install in Alfred

### Developer

```bash
git clone https://github.com/y-marui/alfred-quick-txt-save
cd alfred-quick-txt-save

make install           # Install dev dependencies
make run Q="save"      # Simulate Alfred locally
make test              # Run tests
make build             # Build .alfredworkflow -> dist/
```

| Command | Description |
|---|---|
| `make install` | Install dev dependencies |
| `make lint` | Run ruff linter |
| `make format` | Auto-format with ruff |
| `make typecheck` | Run mypy type checker |
| `make test` | Run tests |
| `make build` | Build .alfredworkflow package |
| `make run Q="..."` | Simulate Alfred with query |

## Project Structure

```
alfred-quick-txt-save/
├── src/
│   ├── alfred/         # Alfred SDK (response, router, cache, config, logger, safe_run)
│   └── app/            # Application layer (commands, services)
├── workflow/           # Alfred package (info.plist, scripts/)
│   └── scripts/
│       ├── entry.py      # Script Filter entrypoint
│       └── save_text.py  # Run Script: reads clipboard and writes file
├── tests/              # pytest test suite
├── scripts/            # build.sh, dev.sh, release.sh, vendor.sh
└── docs/               # Architecture, development, and usage documentation
```

## Usage

### Save clipboard text

```
save              -> ~/Downloads/quick_save_20260326_143012.txt
save notes        -> ~/Downloads/notes.txt
save notes.md     -> ~/Downloads/notes.md
```

Copy text to your clipboard, then type `save` in Alfred and press Enter.

| Key | Action |
|---|---|
| Enter | Save clipboard to the shown path |

### Configuration

Open Alfred Preferences -> Workflows -> Quick Text Save -> **Configure Workflow**.

| Setting | Description | Default |
|---|---|---|
| Save Directory | Directory where files are saved | `~/Downloads` |
| Filename Prefix | Prefix for auto-generated filenames | `quick_save` |
| Default Extension | Extension appended when none is specified | `.txt` |
| Use uv (if available) | Use `uv run` for Python scripts with non-standard libraries | Enabled |

## Documentation

| Document | Description |
|---|---|
| [docs/architecture.md](docs/architecture.md) | Full architecture and layer design |
| [docs/development.md](docs/development.md) | Adding commands, managing dependencies, release |
| [docs/usage.md](docs/usage.md) | End-user usage guide |

## License

MIT — see [LICENSE](LICENSE)
