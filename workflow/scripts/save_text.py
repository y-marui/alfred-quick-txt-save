#!/usr/bin/env python3
"""Save clipboard text to a file.

Called by Alfred's Run Script node after the 'save' Script Filter.
Receives the destination file path as $1 (Alfred passes the selected item's arg).

Alfred run script command:
    python3 scripts/save_text.py "$1"
"""

from __future__ import annotations

import subprocess
import sys
from pathlib import Path


def main() -> None:
    if len(sys.argv) < 2 or not sys.argv[1].strip():
        sys.exit("save_text.py: file path argument is required")

    filepath = Path(sys.argv[1]).expanduser()

    # Read from clipboard
    result = subprocess.run(["pbpaste"], capture_output=True, check=True)
    text = result.stdout.decode("utf-8")

    if not text:
        subprocess.run(
            [
                "osascript",
                "-e",
                'display notification "Clipboard is empty." with title "Quick Save"',
            ],
            check=False,
        )
        return

    filepath.parent.mkdir(parents=True, exist_ok=True)
    filepath.write_text(text, encoding="utf-8")

    subprocess.run(
        [
            "osascript",
            "-e",
            f'display notification "Saved to {filepath.name}" with title "Quick Save"',
        ],
        check=False,
    )


if __name__ == "__main__":
    main()
