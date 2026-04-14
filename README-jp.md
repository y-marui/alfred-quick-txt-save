# Alfred Quick Text Save

> **これは日本語版（正本）です。**
> 英語版（参照）は [README.md](README.md) を参照してください。

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)
[![CI](https://github.com/y-marui/alfred-quick-txt-save/actions/workflows/ci.yml/badge.svg)](https://github.com/y-marui/alfred-quick-txt-save/actions/workflows/ci.yml)
[![Charter Check](https://github.com/y-marui/alfred-quick-txt-save/actions/workflows/dev-charter-check.yml/badge.svg)](https://github.com/y-marui/alfred-quick-txt-save/actions/workflows/dev-charter-check.yml)

クリップボードまたは選択テキストをワンキーで .txt ファイルに保存する Alfred ワークフロー。

## 動作要件

- Alfred 5（Script Filter には Powerpack が必要）
- Python 3.9+

## セットアップ

### エンドユーザー

1. [Releases](https://github.com/y-marui/alfred-quick-txt-save/releases) から最新の `.alfredworkflow` をダウンロード
2. ダブルクリックして Alfred にインストール

### 開発者

```bash
git clone https://github.com/y-marui/alfred-quick-txt-save
cd alfred-quick-txt-save

make install           # 開発用依存関係をインストール
make run Q="save"      # Alfred をローカルでシミュレート
make test              # テストを実行
make build             # .alfredworkflow をビルド -> dist/
```

| コマンド | 説明 |
|---|---|
| `make install` | 開発用依存関係をインストール |
| `make lint` | ruff でリント |
| `make format` | ruff でフォーマット |
| `make typecheck` | mypy で型チェック |
| `make test` | テスト実行 |
| `make build` | .alfredworkflow パッケージをビルド |
| `make run Q="..."` | クエリを指定して Alfred をシミュレート |

## プロジェクト構造

```
alfred-quick-txt-save/
├── src/
│   ├── alfred/         # Alfred SDK (response, router, cache, config, logger, safe_run)
│   └── app/            # アプリケーション層 (commands, services)
├── workflow/           # Alfred パッケージ (info.plist, scripts/)
│   └── scripts/
│       ├── entry.py      # Script Filter エントリーポイント
│       └── save_text.py  # Run Script: クリップボードを読んでファイルに書き込む
├── tests/              # pytest テストスイート
├── scripts/            # build.sh, dev.sh, release.sh, vendor.sh
└── docs/               # アーキテクチャ・開発・利用ドキュメント
```

## 使い方

### クリップボードのテキストを保存

```
save              -> ~/Downloads/quick_save_20260326_143012.txt
save notes        -> ~/Downloads/notes.txt
save notes.md     -> ~/Downloads/notes.md
```

テキストをクリップボードにコピーし、Alfred で `save` と入力して Enter を押します。

| キー | 操作 |
|---|---|
| Enter | 表示されたパスにクリップボードを保存 |

### 設定

Alfred Preferences -> Workflows -> Quick Text Save -> **Configure Workflow** を開きます。

| 設定項目 | 説明 | デフォルト値 |
|---|---|---|
| Save Directory | ファイルの保存先ディレクトリ | `~/Downloads` |
| Filename Prefix | 自動生成ファイル名のプレフィックス | `quick_save` |
| Default Extension | 拡張子未指定時に付加される拡張子 | `.txt` |
| Use uv (if available) | 非標準ライブラリ使用時に `uv run` でスクリプトを実行する | 有効 |

## ドキュメント

| ドキュメント | 内容 |
|---|---|
| [docs/architecture.md](docs/architecture.md) | アーキテクチャ全体設計 |
| [docs/development.md](docs/development.md) | コマンド追加・依存関係管理・リリース手順 |
| [docs/usage.md](docs/usage.md) | エンドユーザー向け利用ガイド |

## ライセンス

MIT — [LICENSE](LICENSE) を参照
