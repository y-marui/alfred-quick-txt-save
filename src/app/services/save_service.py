"""Save service - resolves file paths for the save command."""

from __future__ import annotations

import os
from datetime import datetime
from pathlib import Path

_DEFAULT_DIR = Path.home() / "Downloads"
_DEFAULT_PREFIX = "quick_save"
_DEFAULT_EXT = ".txt"


def get_save_dir() -> Path:
    """Return the save directory.

    Priority: Alfred workflow variable ``save_dir`` (set via Configuration
    Builder or Environment Variables tab) → ``~/Downloads``.
    """
    raw = os.environ.get("save_dir", "").strip()
    if raw:
        return Path(raw).expanduser()
    return _DEFAULT_DIR


def resolve_save_path(filename: str | None = None) -> Path:
    """Return a non-colliding save path for *filename*.

    - No filename: generates ``quick_save_YYYYMMDD_HHMMSS.txt``.
    - Filename without extension: appends ``.txt``.
    - If the resolved path already exists, appends ``(1)``, ``(2)``, …
      before the extension until a free name is found.
    """
    save_dir = get_save_dir()
    if not filename:
        ts = datetime.now().strftime("%Y%m%d_%H%M%S")
        filename = f"{_DEFAULT_PREFIX}_{ts}{_DEFAULT_EXT}"
    elif "." not in Path(filename).name:
        filename = filename + _DEFAULT_EXT
    return _unique_path(save_dir / filename)


def _unique_path(path: Path) -> Path:
    """Return *path* if it does not exist, otherwise ``stem (N).suffix``."""
    if not path.exists():
        return path
    stem = path.stem
    suffix = path.suffix
    counter = 1
    while True:
        candidate = path.with_name(f"{stem} ({counter}){suffix}")
        if not candidate.exists():
            return candidate
        counter += 1
