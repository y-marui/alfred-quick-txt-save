# Alfred Quick Text Save

> **This is the English (reference) version.**
> For the Japanese canonical version, see [README-jp.md](README-jp.md).

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)
[![CI](https://github.com/y-marui/alfred-quick-txt-save/actions/workflows/ci.yml/badge.svg)](https://github.com/y-marui/alfred-quick-txt-save/actions/workflows/ci.yml)
[![Charter Check](https://github.com/y-marui/alfred-quick-txt-save/actions/workflows/dev-charter-check.yml/badge.svg)](https://github.com/y-marui/alfred-quick-txt-save/actions/workflows/dev-charter-check.yml)

Save clipboard or selected text to a .txt file with one keystroke.

## Requirements

- Alfred 5 (Powerpack required for Script Filter)
- Python 3.11+

## Setup

1. Download the latest `.alfredworkflow` from [Releases](https://github.com/y-marui/alfred-quick-txt-save/releases)
2. Double-click to install in Alfred

## Usage

### Save clipboard text

Copy text to your clipboard, then type `save` in Alfred and press Enter.

```
save              -> ~/Downloads/quick_save_20260326_143012.txt
save notes        -> ~/Downloads/notes.txt
save notes.md     -> ~/Downloads/notes.md
```

| Key | Action |
|---|---|
| Enter | Save clipboard to the shown path |

The second line under each result shows the full destination path.
The save directory is created automatically if it does not exist.

### Configuration

Open Alfred Preferences → Workflows → Quick Text Save → **Configure Workflow**.

| Setting | Description | Default |
|---|---|---|
| Save Directory | Directory where files are saved | `~/Downloads` |
| Filename Prefix | Prefix for auto-generated filenames | `quick_save` |
| Default Extension | Extension appended when none is specified | `.txt` |

## Notes

**Nothing was saved** — Check that the clipboard is not empty before running `save`.

**File saved to wrong directory** — Run `save dir <path>` to update the save directory,
or check `wf config` for the current `save_dir` setting.

**Workflow not responding** — Open Alfred's debugger (⌘D) or check
`~/Library/Logs/Alfred/Workflow/<bundle-id>.log`.

## Documentation

| Document | Description |
|---|---|
| [DEVELOPING.md](DEVELOPING.md) | Build, test, lint, release, and AI workflow |
| [docs/architecture.md](docs/architecture.md) | Layer design and data flow |
| [docs/specification.md](docs/specification.md) | Commands, config, and behavior |

## License

MIT — see [LICENSE](LICENSE)
