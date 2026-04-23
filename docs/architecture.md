# Architecture

## Overview

Alfred Quick Text Save is a layered Alfred 5 workflow. Alfred's Script Filter passes a query
to `entry.py`, which dispatches it through an SDK router to pure-Python command handlers.
Handlers resolve paths via services and return Script Filter JSON; a Run Script node writes the file.

## Entry Points

- `workflow/scripts/entry.py` — Alfred executes this file directly (Script Filter node). Sets `sys.path`, calls `safe_run(main)`. No business logic.
- `workflow/scripts/save_text.py` — Alfred executes this file (Run Script node) after the user confirms. Reads clipboard (`pbpaste`), writes to the resolved path, sends macOS notification.

## Directory Structure

| Directory | Role |
|---|---|
| `src/alfred/` | Alfred SDK: response builder, router, cache, config, logger, safe_run |
| `src/app/commands/` | One module per Alfred command. Each exports `handle(args: str) -> None` |
| `src/app/services/` | Business logic: path resolution, config read/write |
| `src/app/clients/` | External IO wrappers (HTTP, system calls) |
| `workflow/` | Alfred package: `info.plist`, scripts, icon |
| `tests/` | pytest test suite (pure Python, no Alfred dependency) |
| `scripts/` | Shell scripts: `build.sh`, `dev.sh`, `release.sh`, `vendor.sh` |

## Layers

```
Alfred
  │  query string
  ▼
workflow/scripts/entry.py          ← Alfred boundary (UI layer)
  │
  ▼
src/alfred/safe_run.py             ← exception safety
  │
  ▼
src/app/core.py                    ← wires router to commands
  │
  ▼
src/alfred/router.py               ← dispatches to command handlers
  │
  ├─ save   → src/app/commands/save_cmd.py
  ├─ open   → src/app/commands/open_cmd.py
  ├─ config → src/app/commands/config_cmd.py
  ├─ help   → src/app/commands/help_cmd.py
  └─ search → src/app/commands/search.py (default fallback)
              │
              ▼
          src/app/services/        ← business logic (path, config)
              │
              ▼
          src/app/clients/         ← IO / external APIs
```

Dependency direction: `commands → services → clients`. Layers must not be skipped.

## Key Dependencies

| Library / Module | Purpose |
|---|---|
| `src/alfred/response.py` | Script Filter JSON output |
| `src/alfred/cache.py` | TTL disk cache via `alfred_workflow_cache` |
| `src/alfred/config.py` | Persistent config via `alfred_workflow_data` |
| `pytest` | Test suite |
| `ruff` | Linter + formatter |
| `mypy` | Static type checking |
