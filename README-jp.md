# Alfred Quick Text Save

> **これは日本語版（正本）です。**
> 英語版（参照）は [README.md](README.md) を参照してください。

> クリップボードまたは選択テキストをワンキーで .txt ファイルに保存する Alfred ワークフロー。

[![CI](https://github.com/y-marui/alfred-quick-txt-save/actions/workflows/ci.yml/badge.svg)](https://github.com/y-marui/alfred-quick-txt-save/actions/workflows/ci.yml)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)

## 特徴

- クリップボードのテキストを Alfred コマンド一つでファイルに保存
- 保存先ディレクトリを設定可能（デフォルト: `~/Downloads`）
- 日付付きファイル名を自動生成（`quick_save_YYYYMMDD.txt`）
- カスタムファイル名対応（拡張子あり・なし両方）
- 保存完了を macOS 通知で通知

## 動作要件

- Alfred 5（Script Filter には Powerpack が必要）
- Python 3.9+

## インストール

1. [Releases](https://github.com/y-marui/alfred-quick-txt-save/releases) から最新の `.alfredworkflow` をダウンロード
2. ダブルクリックして Alfred にインストール

## 使い方

### クリップボードのテキストを保存

```
save              → ~/Downloads/quick_save_20260326.txt
save notes        → ~/Downloads/notes.txt
save notes.md     → ~/Downloads/notes.md
```

テキストをクリップボードにコピーし、Alfred で `save` と入力して Enter を押します。

| キー | 操作 |
|---|---|
| ↩ Enter | 表示されたパスにクリップボードを保存 |

### 保存先ディレクトリを変更

Alfred Preferences → Workflows → Quick Text Save → **Configure Workflow** を開き、**Save Directory** フィールドに保存先を入力します。空白のままにすると `~/Downloads` が使用されます。

## クイックスタート（開発者）

```bash
git clone https://github.com/y-marui/alfred-quick-txt-save
cd alfred-quick-txt-save

# 開発用依存関係をインストール
make install

# Alfred をローカルでシミュレート
make run Q="save"
make run Q="save mynotes"
make run Q="save dir ~/Desktop"

# テストを実行
make test

# ワークフローパッケージをビルド
make build
# → dist/alfred-quick-txt-save-0.1.0.alfredworkflow
```

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

## ドキュメント

| ドキュメント | 内容 |
|---|---|
| [docs/architecture.md](docs/architecture.md) | アーキテクチャ全体設計 |
| [docs/development.md](docs/development.md) | コマンド追加・依存関係管理・リリース手順 |
| [docs/usage.md](docs/usage.md) | エンドユーザー向け利用ガイド |

## リリース手順

```bash
# 1. pyproject.toml のバージョンを更新
# 2. タグを付けてプッシュ
git tag v1.2.3
git push --tags
# GitHub Actions が .alfredworkflow をビルドして GitHub Release を作成
```

## ライセンス

MIT — [LICENSE](LICENSE) を参照
