# File Map

_最終更新: 2026-09-05_（クリップボード読み取りをAlfredネイティブの `{clipboard}` プレースホルダーへ移管）

## Entry Point

| ファイル | 役割 | 主な依存先 |
|---|---|---|
| `cmd/quick-txt-save-alfred/main.go` | Alfred が実行する唯一のバイナリ。サブコマンド (`list`/`write`) ディスパッチのみ | `internal/quicksavecmd`、`internal/quicksave`、`internal/scriptfilter` |

## Alfred JSON (`internal/scriptfilter/`)

| ファイル | 役割 | 主な依存先 |
|---|---|---|
| `internal/scriptfilter/scriptfilter.go` | Script Filter JSON レスポンス型・エンコード | — |

## Application Layer (`internal/`)

| ファイル | 役割 | 主な依存先 |
|---|---|---|
| `internal/quicksave/quicksave.go` | 保存先パス解決（`save_dir`/`file_prefix`/`file_ext` 環境変数、衝突回避） | — |
| `internal/quicksave/save.go` | ファイル書き込み・macOS 通知（保存成功/クリップボード空/書き込み失敗）。クリップボードの読み取りは `workflow/info.plist` の Arguments and Variables ノード（`{clipboard}` → `text` 変数）が担い、このパッケージは `text` を受け取るのみ | — |
| `internal/quicksavecmd/quicksavecmd.go` | `list` サブコマンド用の Script Filter レスポンス生成 | `internal/quicksave`、`internal/scriptfilter` |
