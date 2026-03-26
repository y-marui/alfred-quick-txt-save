"""save command - save clipboard or selected text to a .txt file.

Usage in Alfred:  save [filename]

Examples:
  save              → saves to ~/Downloads/quick_save_20260326_143012.txt
  save notes        → saves to ~/Downloads/notes.txt
  save notes.md     → saves to ~/Downloads/notes.md

Save directory is configured via Alfred Preferences →
  Workflows → [Quick Text Save] → Configure Workflow.
"""

from __future__ import annotations

from alfred.logger import get_logger
from alfred.response import item, output
from app.services.save_service import get_save_dir, resolve_save_path

log = get_logger(__name__)


def handle(args: str) -> None:
    """Show save destination preview."""
    log.debug("save command: args=%r", args)

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
                subtitle="Change via Alfred Preferences → Workflows → Configure Workflow",
                valid=False,
                uid="save-dir-info",
            ),
        ]
    )
