# Alfred Quick Text Save

> **これは日本語版（正本）です。**
> 英語版（参照）は [README.md](README.md) を参照してください。

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)
[![CI](https://github.com/y-marui/alfred-quick-txt-save/actions/workflows/ci.yml/badge.svg)](https://github.com/y-marui/alfred-quick-txt-save/actions/workflows/ci.yml)
[![Charter Check](https://github.com/y-marui/alfred-quick-txt-save/actions/workflows/dev-charter-check.yml/badge.svg)](https://github.com/y-marui/alfred-quick-txt-save/actions/workflows/dev-charter-check.yml)

クリップボードまたは選択テキストをワンキーで .txt ファイルに保存する Alfred ワークフロー。

## 動作要件

- Alfred 5（Script Filter には Powerpack が必要）
- Python 3.11+

## セットアップ

1. [Releases](https://github.com/y-marui/alfred-quick-txt-save/releases) から最新の `.alfredworkflow` をダウンロード
2. ダブルクリックして Alfred にインストール

## 使い方

### クリップボードのテキストを保存

テキストをクリップボードにコピーし、Alfred で `save` と入力して Enter を押します。

```
save              -> ~/Downloads/quick_save_20260326_143012.txt
save notes        -> ~/Downloads/notes.txt
save notes.md     -> ~/Downloads/notes.md
```

| キー | 操作 |
|---|---|
| Enter | 表示されたパスにクリップボードを保存 |

各結果の2行目に保存先フルパスが表示されます。
保存ディレクトリが存在しない場合は自動的に作成されます。

### 設定

Alfred Preferences → Workflows → Quick Text Save → **Configure Workflow** を開きます。

| 設定項目 | 説明 | デフォルト値 |
|---|---|---|
| Save Directory | ファイルの保存先ディレクトリ | `~/Downloads` |
| Filename Prefix | 自動生成ファイル名のプレフィックス | `quick_save` |
| Default Extension | 拡張子未指定時に付加される拡張子 | `.txt` |

## 注意事項

**保存されない** — `save` を実行する前にクリップボードが空でないことを確認してください。

**別のディレクトリに保存された** — `save dir <path>` で保存先を更新するか、
`wf config` で現在の `save_dir` 設定を確認してください。

**ワークフローが反応しない** — Alfred のデバッガー（⌘D）を開くか、
`~/Library/Logs/Alfred/Workflow/<bundle-id>.log` を確認してください。

## ドキュメント

| ドキュメント | 内容 |
|---|---|
| [DEVELOPING.md](DEVELOPING.md) | ビルド・テスト・リント・リリース・AI ワークフロー |
| [docs/architecture.md](docs/architecture.md) | レイヤー設計とデータフロー |
| [docs/specification.md](docs/specification.md) | コマンド・設定・動作仕様 |

## ライセンス

MIT — [LICENSE](LICENSE) を参照
