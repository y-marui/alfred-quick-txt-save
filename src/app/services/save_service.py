"""Save service - resolves file paths for the save command."""

from __future__ import annotations

from datetime import datetime
from pathlib import Path

from alfred.config import Config

_config = Config()
_DEFAULT_DIR = Path.home() / "Downloads"
_DEFAULT_PREFIX = "quick_save"
_DEFAULT_EXT = ".txt"


def get_save_dir() -> Path:
    """Return the configured save directory, falling back to ~/Downloads."""
    raw = _config.get("save_dir")
    if raw:
        return Path(raw).expanduser()
    return _DEFAULT_DIR


def set_save_dir(path: str) -> Path:
    """Persist *path* as the save directory and return the resolved Path."""
    resolved = Path(path).expanduser()
    _config.set("save_dir", str(resolved))
    return resolved


def resolve_save_path(filename: str | None = None) -> Path:
    """Return the full save path for *filename*.

    If *filename* is omitted or empty, generates ``quick_save_YYYYMMDD.txt``.
    If *filename* has no extension, ``.txt`` is appended.
    """
    save_dir = get_save_dir()
    if not filename:
        date_str = datetime.now().strftime("%Y%m%d")
        filename = f"{_DEFAULT_PREFIX}_{date_str}{_DEFAULT_EXT}"
    elif "." not in Path(filename).name:
        filename = filename + _DEFAULT_EXT
    return save_dir / filename
