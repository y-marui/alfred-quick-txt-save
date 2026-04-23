# File Map

_最終更新: 2026-04-23_

## Entry Point

| ファイル | 役割 | 主な依存先 |
|---|---|---|
| `workflow/scripts/entry.py` | Alfred が実行する唯一のエントリーポイント | `src/app/core.py` |

## Alfred SDK (`src/alfred/`)

| ファイル | 役割 | 主な依存先 |
|---|---|---|
| `src/alfred/response.py` | Script Filter JSON レスポンス生成 | — |
| `src/alfred/router.py` | クエリキーワードをコマンドにディスパッチ | `src/alfred/response.py` |
| `src/alfred/cache.py` | ファイルベースキャッシュ | — |
| `src/alfred/config.py` | Alfred 環境変数の読み込み | — |
| `src/alfred/logger.py` | ロギングユーティリティ | — |
| `src/alfred/safe_run.py` | 未捕捉例外をキャッチしてエラーアイテムを返す | `src/alfred/response.py` |

## Application Layer (`src/app/`)

| ファイル | 役割 | 主な依存先 |
|---|---|---|
| `src/app/core.py` | ルーターへのコマンド登録 | `src/alfred/router.py`、各コマンドモジュール |
| `src/app/commands/` | 各コマンドのハンドラ（`handle(args: str) -> None`） | `src/alfred/response.py`、各サービス |
