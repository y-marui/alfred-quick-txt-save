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


def _get_file_prefix() -> str:
    """Return the filename prefix from env var ``file_prefix`` or the default."""
    return os.environ.get("file_prefix", "").strip() or _DEFAULT_PREFIX


def _get_file_ext() -> str:
    """Return the default file extension from env var ``file_ext`` or the default.

    Ensures the value starts with a leading dot.
    """
    raw = os.environ.get("file_ext", "").strip()
    if raw and not raw.startswith("."):
        raw = "." + raw
    return raw or _DEFAULT_EXT


def resolve_save_path(filename: str | None = None) -> Path:
    """Return a non-colliding save path for *filename*.

    - No filename: generates ``{prefix}_YYYYMMDD_HHMMSS{ext}``.
    - Filename without extension: appends the configured default extension.
    - If the resolved path already exists, appends ``(1)``, ``(2)``, …
      before the extension until a free name is found.
    """
    save_dir = get_save_dir()
    if not filename:
        ts = datetime.now().strftime("%Y%m%d_%H%M%S")
        filename = f"{_get_file_prefix()}_{ts}{_get_file_ext()}"
    elif "." not in Path(filename).name:
        filename = filename + _get_file_ext()
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
