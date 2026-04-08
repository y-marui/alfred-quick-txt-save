# Dev Charter (開発憲章)

> **このファイルは正本（日本語版）です。**
> 英語版（参照）は [README.md](README.md) を参照してください。

[![License: CC BY 4.0](https://img.shields.io/badge/License-CC%20BY%204.0-lightgrey.svg)](LICENSE)
[![check-charter CI](https://github.com/y-marui/dev-charter/actions/workflows/check-charter.yml/badge.svg)](https://github.com/y-marui/dev-charter/actions/workflows/check-charter.yml)

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
| [AI_TOOL_SETUP.md](AI_TOOL_SETUP.md) | AI コンテキストファイルの構成仕様（AI_CONTEXT.md・各ツール設定ファイル） |
| [LANGUAGE_POLICY.md](LANGUAGE_POLICY.md) | 言語ポリシー（正本＝日本語） |
| [LOCALIZATION_POLICY.md](LOCALIZATION_POLICY.md) | ローカライゼーションポリシー |
| [PROJECT_LIFECYCLE.md](PROJECT_LIFECYCLE.md) | プロジェクトライフサイクルと体制 |
| [UI_GUIDELINES.md](UI_GUIDELINES.md) | UI ガイドライン・カラー・アイコン |
| [MONETIZATION_POLICY.md](MONETIZATION_POLICY.md) | マネタイズポリシーとプラットフォーム別方針 |
| [SECURITY_POLICY.md](SECURITY_POLICY.md) | セキュリティポリシーとフック設定リファレンス |
| [LEGAL_POLICY.md](LEGAL_POLICY.md) | ライセンス選択基準・テンプレート（Closed / MIT / AGPL・GPL・LGPL） |
| [topics/CI_POLICY.md](topics/CI_POLICY.md) | CI job設計方針・Branch Protection Ruleset |
| [topics/GITHUB_SETTINGS.md](topics/GITHUB_SETTINGS.md) | GitHub リポジトリ設定確認（Ruleset・Sponsors） |
| [topics/GITHUB_CONTRIBUTING.md](topics/GITHUB_CONTRIBUTING.md) | Issue・PR・CONTRIBUTING.md・PRテンプレート・準CLA（OSS向け） |
| [topics/TEMPLATE_README_GUIDELINES.md](topics/TEMPLATE_README_GUIDELINES.md) | GitHub テンプレートリポジトリの README 設計規約（開発環境・言語・LICENSE・必須セクション） |
| [topics/PROJECT_README_GUIDELINES.md](topics/PROJECT_README_GUIDELINES.md) | テンプレートから作成したプロジェクトの README 整備手順 |

## How to Use

1. `git subtree` で `docs/dev-charter/` に取り込む
2. AI に dev-charter を読ませ、プロジェクトルートに `AI_CONTEXT.md` と AI ツール設定ファイルを生成させる
3. 憲章が更新されたら `git subtree pull` 後、AI にコンテキストファイルを追従させる

構成仕様は [AI_TOOL_SETUP.md](AI_TOOL_SETUP.md) を参照。

## Install (git subtree)

```
git remote add dev-charter https://github.com/y-marui/dev-charter
git fetch dev-charter
git subtree add --prefix=docs/dev-charter dev-charter main --squash
```

インストール後、[インストールチェックリスト](INSTALL_CHECKLIST.md) に従って対応してください。

## Update

`dev-charter` リモートが未設定の場合（プロジェクトを clone した直後など）は先に追加する：

```
git remote add dev-charter https://github.com/y-marui/dev-charter
git subtree pull --prefix=docs/dev-charter dev-charter main --squash
```

更新後、[更新チェックリスト](UPDATE_CHECKLIST.md) に従って対応してください。

## Makefile Helper

```
update-charter:
	git remote | grep -q '^dev-charter$$' || \
	  git remote add dev-charter https://github.com/y-marui/dev-charter
	git fetch dev-charter
	git subtree pull --prefix=docs/dev-charter dev-charter main --squash
```

## Version Check (CI)

`.github/workflows/dev-charter-check.yml` をプロジェクトに追加すると、
毎週自動で最新バージョンを確認し、古い場合は update PR を作成します。

```yaml
name: Dev Charter
on:
  schedule:
    - cron: "23 3 * * 1"  # 毎週月曜 3:23 UTC — minute/hour/day-of-week はランダムな値に変更してください
  workflow_dispatch:

jobs:
  check:
    name: Check
    uses: y-marui/dev-charter/.github/workflows/check-charter.yml@main
    with:
      fail_if_outdated: true
    permissions:
      contents: write
      pull-requests: write
```

> **Note:** Branch Protection で direct push が禁止されている場合は、
> GitHub Actions bot の bypass rule を追加してください
> （Settings > Rules > Rulesets > Bypass list > GitHub Actions）。

## Badge for Adopting Projects

プロジェクトの README にこのバッジを追加すると、dev-charter の更新状態を可視化できます。

### Workflow Status Badge

dev-charter が最新かどうかを表示します。バッジが機能するには上記ワークフローに `fail_if_outdated: true` が必要です。

```markdown
[![Charter Check](https://github.com/{owner}/{repo}/actions/workflows/dev-charter-check.yml/badge.svg)](https://github.com/{owner}/{repo}/actions/workflows/dev-charter-check.yml)
```

`{owner}` と `{repo}` を自分のリポジトリのオーナー名・リポジトリ名に置き換えてください。

| 状態 | Status Badge |
|---|---|
| 未導入 / CI 未設定 | 赤（VERSION not found） |
| 導入済み・最新 | 緑 |
| 導入済み・更新必要 | 赤 |

---

*この文書には英語版 [README.md](README.md) があります。編集時は同一コミットで更新してください。*
