"""Tests for save command and save service."""

from __future__ import annotations

import json
from pathlib import Path

import pytest

from app.commands import save_cmd
from app.services import save_service


class TestSaveService:
    def test_resolve_save_path_default(
        self, tmp_path: Path, monkeypatch: pytest.MonkeyPatch
    ) -> None:
        monkeypatch.setattr(save_service, "_config", save_service.Config())
        path = save_service.resolve_save_path()
        assert path.parent == Path.home() / "Downloads"
        assert path.name.startswith("quick_save_")
        assert path.suffix == ".txt"

    def test_resolve_save_path_with_filename(self) -> None:
        path = save_service.resolve_save_path("notes")
        assert path.name == "notes.txt"

    def test_resolve_save_path_with_extension(self) -> None:
        path = save_service.resolve_save_path("notes.md")
        assert path.name == "notes.md"

    def test_resolve_save_path_uses_configured_dir(self, tmp_path: Path) -> None:
        save_service.set_save_dir(str(tmp_path))
        path = save_service.resolve_save_path("test")
        assert path.parent == tmp_path
        assert path.name == "test.txt"

    def test_set_save_dir_persists(self, tmp_path: Path) -> None:
        resolved = save_service.set_save_dir(str(tmp_path))
        assert resolved == tmp_path
        assert save_service.get_save_dir() == tmp_path


class TestSaveCommand:
    def test_no_args_shows_default_filename(self, capsys: pytest.CaptureFixture[str]) -> None:
        save_cmd.handle("")
        data = json.loads(capsys.readouterr().out)
        items = data["items"]
        assert len(items) == 2
        assert items[0]["title"].startswith("Save to quick_save_")
        assert items[0]["arg"].endswith(".txt")

    def test_custom_filename_no_ext(self, capsys: pytest.CaptureFixture[str]) -> None:
        save_cmd.handle("myfile")
        data = json.loads(capsys.readouterr().out)
        assert data["items"][0]["title"] == "Save to myfile.txt"

    def test_custom_filename_with_ext(self, capsys: pytest.CaptureFixture[str]) -> None:
        save_cmd.handle("notes.md")
        data = json.loads(capsys.readouterr().out)
        assert data["items"][0]["title"] == "Save to notes.md"

    def test_arg_is_full_path(self, tmp_path: Path, capsys: pytest.CaptureFixture[str]) -> None:
        save_service.set_save_dir(str(tmp_path))
        save_cmd.handle("out")
        data = json.loads(capsys.readouterr().out)
        assert data["items"][0]["arg"] == str(tmp_path / "out.txt")

    def test_dir_subcommand_no_path_shows_current(
        self, tmp_path: Path, capsys: pytest.CaptureFixture[str]
    ) -> None:
        save_service.set_save_dir(str(tmp_path))
        save_cmd.handle("dir")
        data = json.loads(capsys.readouterr().out)
        assert str(tmp_path) in data["items"][0]["title"]
        assert data["items"][0]["valid"] is False

    def test_dir_subcommand_sets_directory(
        self, tmp_path: Path, capsys: pytest.CaptureFixture[str]
    ) -> None:
        new_dir = tmp_path / "saves"
        new_dir.mkdir()
        save_cmd.handle(f"dir {new_dir}")
        data = json.loads(capsys.readouterr().out)
        assert str(new_dir) in data["items"][0]["title"]
        assert save_service.get_save_dir() == new_dir
