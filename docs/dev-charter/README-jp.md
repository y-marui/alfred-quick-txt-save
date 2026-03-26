# Dev Charter (開発憲章)

> **このファイルは正本（日本語版）です。**
> 英語版（参照）は [README.md](README.md) を参照してください。

AI支援ソフトウェアプロジェクトのための共有開発憲章。

このリポジトリは、プロジェクト横断的に使用される共通の哲学、アーキテクチャ原則、
および開発ルールを定義します。

## Documents

| ファイル | 内容 |
|---|---|
| [PRINCIPLES.md](PRINCIPLES.md) | 開発哲学・デザイン・アーキテクチャ原則 |
| [CODE_STYLE.md](CODE_STYLE.md) | コードスタイル |
| [AI_COLLABORATION_RULES.md](AI_COLLABORATION_RULES.md) | AI 協働ルールと役割分担 |
| [AI_CONTEXT_HIERARCHY.md](AI_CONTEXT_HIERARCHY.md) | AI コンテキスト優先階層 |
| [LANGUAGE_POLICY.md](LANGUAGE_POLICY.md) | 言語ポリシー（正本＝日本語） |
| [LOCALIZATION_POLICY.md](LOCALIZATION_POLICY.md) | ローカライゼーションポリシー |
| [PROJECT_LIFECYCLE.md](PROJECT_LIFECYCLE.md) | プロジェクトライフサイクルと体制 |
| [UI_GUIDELINES.md](UI_GUIDELINES.md) | UI ガイドライン・カラー・アイコン |
| [MONETIZATION_POLICY.md](MONETIZATION_POLICY.md) | マネタイズポリシーとプラットフォーム別方針 |
| [SECURITY_POLICY.md](SECURITY_POLICY.md) | セキュリティポリシーとフック設定リファレンス |
| [LICENSE_POLICY.md](LICENSE_POLICY.md) | ライセンス選択基準・テンプレート（Closed / MIT / AGPL・GPL・LGPL） |
| [topics/CI_POLICY.md](topics/CI_POLICY.md) | CI job設計方針・Branch Protection Ruleset |
| [topics/GITHUB_CONTRIBUTING.md](topics/GITHUB_CONTRIBUTING.md) | Issue・PR・CONTRIBUTING.md・PRテンプレート・準CLA（OSS向け） |
| [topics/TEMPLATE_README_GUIDELINES.md](topics/TEMPLATE_README_GUIDELINES.md) | GitHub テンプレートリポジトリの README 設計規約（開発環境・言語・LICENSE・必須セクション） |
| [topics/PROJECT_README_GUIDELINES.md](topics/PROJECT_README_GUIDELINES.md) | テンプレートから作成したプロジェクトの README 整備手順 |

## How to Use

### 1. Add to Your Project
`git subtree` を使い、`docs/dev-charter/` に取り込む。ファイルはコミットし、`.gitignore` しない。これにより、オフラインでも参照可能でバージョン管理された状態を保てる。

### 2. Generate AI_CONTEXT.md at Project Setup
新規プロジェクトの開始時に、AI に dev-charter を読ませ、プロジェクトルートの `AI_CONTEXT.md` を生成・編集させる。これにより、憲章をプロジェクト固有のコンテキストファイルにコンパイルする。

### 3. Configure AI Tools to Auto-Load AI_CONTEXT.md

AI ツールごとに設定ファイルを作成し、セッション開始時に自動で読み込まれるようにする。

**Claude Code** — プロジェクトルートに `CLAUDE.md` を作成：

```
@AI_CONTEXT.md
```

**Gemini CLI** — プロジェクトルートに `GEMINI.md` を作成：

```
@AI_CONTEXT.md
```

自動読み込みが未サポートの場合は使用時に手動で渡すこと。自動読み込み機能が追加された際はこの手順を更新すること。

**GitHub Copilot** — `.github/copilot-instructions.md` に `AI_CONTEXT.md` を参照する旨を記載し、個別設定のみ追記する。

### 4. When the Charter Is Updated
`git subtree pull` 後、AI に差分を確認させ、必要に応じて `AI_CONTEXT.md` を更新する。

## Install (git subtree)

```
git remote add dev-charter https://github.com/y-marui/dev-charter
git fetch dev-charter
git subtree add --prefix=docs/dev-charter dev-charter main --squash
```

インストール後、以下のプロンプトを Claude Code に貼り付けて実行すると、
プロジェクト固有の `AI_CONTEXT.md` が生成されます。

```
docs/dev-charter/ 内の全ファイルを読み、このプロジェクトを調査した上で、
プロジェクトルートに AI_CONTEXT.md を生成または更新してください。
AI_CONTEXT.md は、以後のセッションで AI が唯一参照するコンテキストファイルです。

## Step 1: Explore

以下を読んで把握してください：

- 憲章：docs/dev-charter/*.md（全ファイル）
- プロジェクト概要：README、既存ドキュメント
- 技術スタック：使用言語・FW・バージョン（package.json、go.mod、Podfile 等）
- 既存規約：.editorconfig、lint 設定、pre-commit 設定、CI 設定
- 既存 AI 設定：CLAUDE.md、AI_CONTEXT.md（存在する場合は内容を確認し、差分を反映して更新）

## Step 2: Generate AI_CONTEXT.md

以下のセクション構成で作成してください。
憲章全文は転記せず、このプロジェクトの技術スタック・ワークフローに
直接関係する内容のみを抽出・要約すること。
関係しない憲章ドキュメントはスキップしてよい。

### Project Overview
目的・技術スタック（言語・FW・バージョン）・主要ディレクトリ一覧

### Applied Charter Principles
このプロジェクトの開発・運用フローに直接影響する原則とルール

### Project-Specific Rules
憲章に含まれない既存規約、または憲章を上書き・補足する事項

### AI Tool Assignments
Claude Code / Copilot / Gemini CLI の担当範囲
（未使用ツールは省略し、実際の体制に合わせて記載）

### Prohibited Actions
セキュリティ制約・スコープ外変更の禁止事項

## Notes

- 不明点・確認事項は作業前に 1 回まとめて質問する
- 憲章と既存規約が矛盾する場合は矛盾点を列挙し、優先順位をユーザーに確認してから進める
- 生成後にコミットしない（ユーザーが確認してから行う）
```

## Update

```
git subtree pull --prefix=docs/dev-charter dev-charter main --squash
```

更新後、以下のプロンプトを Claude Code に貼り付けて実行すると、
AI_CONTEXT.md が最新の憲章に追従します。

```
docs/dev-charter/ 内の全ファイルを読み、現在の AI_CONTEXT.md と比較して、
憲章の変更が影響するセクションのみ更新してください。

## Steps

1. docs/dev-charter/*.md を読む
2. AI_CONTEXT.md を読む
3. 差分を特定し、影響するセクションのみ書き換える

## Notes

- プロジェクト全体の再調査は不要
- AI_CONTEXT.md が存在しない場合はインストール時のプロンプトを使うこと
- 憲章の変更がプロジェクト固有ルールと矛盾する場合は、矛盾点を列挙してユーザーに確認する
- 更新後にコミットしない（ユーザーが確認してから行う）
```

## Makefile Helper

```
update-charter:
	git subtree pull --prefix=docs/dev-charter dev-charter main --squash
```

---

*この文書には英語版 [README.md](README.md) があります。編集時は同一コミットで更新してください。*
