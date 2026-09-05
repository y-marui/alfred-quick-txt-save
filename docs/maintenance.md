# Docs Maintenance

ローカル LLM で `docs/` ファイルをメンテナンスするためのプロンプト集。

---

## docs/architecture.md を更新する

```
このプロジェクトの docs/architecture.md を更新してください。

手順：
1. ディレクトリ構造を確認する: ls -R cmd/ internal/ workflow/
2. 主要ファイルを読む（cmd/quick-txt-save-alfred/main.go、internal/quicksavecmd/quicksavecmd.go 等）
3. 既存の docs/architecture.md を読む（存在する場合）
4. 以下のフォーマットで docs/architecture.md を上書き保存する:

# Architecture

## Overview
<!-- プロジェクト全体像を3行以内で記述 -->

## Entry Points
- `パス/ファイル` — 説明

## Directory Structure
| ディレクトリ | 役割 |
|---|---|
| `internal/quicksave/` | ... |

## Key Dependencies
| ライブラリ / モジュール | 用途 |
|---|---|
| `go.mod` 参照 | ... |

注意：ファイルレベルの詳細は記載しない（file-map.md に委譲）。主要な依存のみ列挙する。
```

---

## docs/file-map.md を更新する

```
このプロジェクトの docs/file-map.md を更新してください。

手順：
1. 最近変更されたファイルを確認する: git diff --name-only HEAD~5 HEAD
2. 変更があったファイルと関連ファイルを読む
3. 既存の docs/file-map.md を読む
4. 以下のフォーマットで docs/file-map.md を上書き保存する（未探索ファイルは記載しない）:

# File Map

_最終更新: YYYY-MM-DD_

## [モジュール / 機能名]
| ファイル | 役割 | 主な依存先 |
|---|---|---|
| `internal/foo/foo.go` | 説明 | `internal/bar` |

注意：全ファイルを網羅しなくてよい。AI が参照・編集したファイルを順次追記する。
更新のたびに「最終更新」日付を更新すること。
```
