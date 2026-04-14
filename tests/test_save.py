"""Tests for save command and save service."""

from __future__ import annotations

import json
from pathlib import Path

import pytest

from app.commands import save_cmd
from app.services import save_service


class TestSaveService:
    def test_resolve_save_path_default(self) -> None:
        path = save_service.resolve_save_path()
        assert path.parent == Path.home() / "Downloads"
        assert path.name.startswith("quick_save_")
        assert path.suffix == ".txt"

    def test_default_filename_includes_time(self) -> None:
        path = save_service.resolve_save_path()
        # quick_save_YYYYMMDD_HHMMSS  →  quick(1) + save(1) + date(1) + time = 3 underscores
        assert path.stem.count("_") == 3

    def test_resolve_save_path_with_filename(self) -> None:
        path = save_service.resolve_save_path("notes")
        assert path.name == "notes.txt"

    def test_resolve_save_path_with_extension(self) -> None:
        path = save_service.resolve_save_path("notes.md")
        assert path.name == "notes.md"

    def test_get_save_dir_uses_env_var(
        self, tmp_path: Path, monkeypatch: pytest.MonkeyPatch
    ) -> None:
        monkeypatch.setenv("save_dir", str(tmp_path))
        assert save_service.get_save_dir() == tmp_path

    def test_get_save_dir_env_var_expanduser(self, monkeypatch: pytest.MonkeyPatch) -> None:
        monkeypatch.setenv("save_dir", "~/Documents")
        assert save_service.get_save_dir() == Path.home() / "Documents"

    def test_get_save_dir_default_when_env_empty(self, monkeypatch: pytest.MonkeyPatch) -> None:
        monkeypatch.setenv("save_dir", "")
        assert save_service.get_save_dir() == Path.home() / "Downloads"

    def test_file_prefix_env_var(self, monkeypatch: pytest.MonkeyPatch) -> None:
        monkeypatch.setenv("file_prefix", "note")
        path = save_service.resolve_save_path()
        assert path.name.startswith("note_")

    def test_file_prefix_empty_uses_default(self, monkeypatch: pytest.MonkeyPatch) -> None:
        monkeypatch.setenv("file_prefix", "")
        path = save_service.resolve_save_path()
        assert path.name.startswith("quick_save_")

    def test_file_ext_env_var(self, monkeypatch: pytest.MonkeyPatch) -> None:
        monkeypatch.setenv("file_ext", ".md")
        path = save_service.resolve_save_path()
        assert path.suffix == ".md"

    def test_file_ext_env_var_without_dot(self, monkeypatch: pytest.MonkeyPatch) -> None:
        monkeypatch.setenv("file_ext", "md")
        path = save_service.resolve_save_path()
        assert path.suffix == ".md"

    def test_file_ext_empty_uses_default(self, monkeypatch: pytest.MonkeyPatch) -> None:
        monkeypatch.setenv("file_ext", "")
        path = save_service.resolve_save_path()
        assert path.suffix == ".txt"

    def test_file_ext_applied_to_custom_name(self, monkeypatch: pytest.MonkeyPatch) -> None:
        monkeypatch.setenv("file_ext", ".md")
        path = save_service.resolve_save_path("notes")
        assert path.name == "notes.md"

    def test_resolve_save_path_uses_env_dir(
        self, tmp_path: Path, monkeypatch: pytest.MonkeyPatch
    ) -> None:
        monkeypatch.setenv("save_dir", str(tmp_path))
        path = save_service.resolve_save_path("test")
        assert path.parent == tmp_path
        assert path.name == "test.txt"

    def test_unique_path_no_collision(
        self, tmp_path: Path, monkeypatch: pytest.MonkeyPatch
    ) -> None:
        monkeypatch.setenv("save_dir", str(tmp_path))
        path = save_service.resolve_save_path("report")
        assert path.name == "report.txt"

    def test_unique_path_one_collision(
        self, tmp_path: Path, monkeypatch: pytest.MonkeyPatch
    ) -> None:
        (tmp_path / "report.txt").touch()
        monkeypatch.setenv("save_dir", str(tmp_path))
        path = save_service.resolve_save_path("report")
        assert path.name == "report (1).txt"

    def test_unique_path_multiple_collisions(
        self, tmp_path: Path, monkeypatch: pytest.MonkeyPatch
    ) -> None:
        (tmp_path / "report.txt").touch()
        (tmp_path / "report (1).txt").touch()
        (tmp_path / "report (2).txt").touch()
        monkeypatch.setenv("save_dir", str(tmp_path))
        path = save_service.resolve_save_path("report")
        assert path.name == "report (3).txt"


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

    def test_arg_is_full_path(
        self, tmp_path: Path, monkeypatch: pytest.MonkeyPatch, capsys: pytest.CaptureFixture[str]
    ) -> None:
        monkeypatch.setenv("save_dir", str(tmp_path))
        save_cmd.handle("out")
        data = json.loads(capsys.readouterr().out)
        assert data["items"][0]["arg"] == str(tmp_path / "out.txt")

    def test_second_item_shows_save_dir(
        self, tmp_path: Path, monkeypatch: pytest.MonkeyPatch, capsys: pytest.CaptureFixture[str]
    ) -> None:
        monkeypatch.setenv("save_dir", str(tmp_path))
        save_cmd.handle("")
        data = json.loads(capsys.readouterr().out)
        assert str(tmp_path) in data["items"][1]["title"]
        assert data["items"][1]["valid"] is False
