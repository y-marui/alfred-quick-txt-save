"""save command - save clipboard or selected text to a .txt file.

Usage in Alfred:  save [filename]
                  save dir <path>

Examples:
  save              → saves to ~/Downloads/quick_save_20260326.txt
  save notes        → saves to ~/Downloads/notes.txt
  save notes.md     → saves to ~/Downloads/notes.md
  save dir ~/Desktop  → sets save directory to ~/Desktop
"""

from __future__ import annotations

from alfred.logger import get_logger
from alfred.response import item, output
from app.services.save_service import get_save_dir, resolve_save_path, set_save_dir

log = get_logger(__name__)


def handle(args: str) -> None:
    """Show save destination preview or update the save directory."""
    log.debug("save command: args=%r", args)

    parts = args.strip().split(None, 1)
    subcommand = parts[0].lower() if parts else ""

    if subcommand == "dir":
        _handle_set_dir(parts[1] if len(parts) > 1 else "")
        return

    filename = args.strip() or None
    path = resolve_save_path(filename)
    save_dir = get_save_dir()

    output(
        [
            item(
                title=f"Save to {path.name}",
                subtitle=str(path),
                arg=str(path),
                uid="save-clipboard",
                variables={"save_filepath": str(path)},
            ),
            item(
                title=f"Save directory: {save_dir}",
                subtitle='Type "save dir <path>" to change',
                valid=False,
                uid="save-dir-info",
            ),
        ]
    )


def _handle_set_dir(path_arg: str) -> None:
    """Set the save directory."""
    if not path_arg.strip():
        save_dir = get_save_dir()
        output(
            [
                item(
                    title=f"Current save directory: {save_dir}",
                    subtitle='Type "save dir <path>" to change',
                    valid=False,
                )
            ]
        )
        return

    resolved = set_save_dir(path_arg.strip())
    output(
        [
            item(
                title=f"Save directory set to {resolved}",
                subtitle=str(resolved),
                valid=False,
            )
        ]
    )
