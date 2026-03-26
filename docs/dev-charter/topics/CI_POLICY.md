# CI Policy

## Job Design

**CIのjob構成とRuleset設定を分離し、Ruleset管理を最小化する。**

- 複数jobの場合：最後に `build` jobを配置し、`needs` で全依存を定義する
- 単一jobの場合：そのjobを `build` と命名する
- Ruleset設定：`build` のみ指定（全リポジトリ共通）

この方針により、jobを追加してもRulesetの変更が不要になる。

### Job Pattern

**複数job（推奨）：**

```yaml
name: CI
on:
  push:
    branches: [main]
  pull_request:

jobs:
  security:
    name: Security scan (pre-commit)
    # pre-commit run --all-files を実行する
    # gitleaks 等のシークレット検知フックを含める
    # gitleaks はフル履歴が必要なため fetch-depth: 0 を指定する
    # ...

  lint:
    name: Lint
    # ...

  test:
    name: Test
    # ...

  build:
    name: Build
    needs: [security, lint, test]
    # ...
```

`build` が全jobの集約点となる。いずれかが失敗すると `build` がskipされ、マージ不可になる。

**単一job：** そのjobを `build` と命名する。

**ビルド成果物が存在しない場合（純粋なライブラリ等）：** `build` jobでインストール可能性を検証する。

```yaml
build:
  name: Installability check
  needs: [security, lint, test]
  steps:
    - uses: actions/checkout@v4
    - run: pip install -e .
    - run: python -c "import mypackage"
```

### Artifact Retention

| 対象 | 保持期間（目安） |
|---|---|
| PR | 短期（例：7日） |
| main | 長期（例：30日） |

## Dependabot

`.github/dependabot.yml` の導入を検討する。依存パッケージがあるプロジェクトでは自動でアップデートPRを作成し、脆弱性対応を省力化できる。ドキュメントのみのリポジトリや依存パッケージが存在しないプロジェクトでは不要。

## Branch Protection (Ruleset)

`main` ブランチに以下のRulesetを適用する（全リポジトリ共通）：

```
Name: main-protection
Target: main
Enforcement: Active

Rules:
☑ Require a pull request before merging
  └ Required approvals: 0（個人開発）/ 1以上（複数人）
☑ Require status checks to pass before merging
  └ Status checks: build
☑ Require conversation resolution before merging
☑ Block force pushes
☑ Restrict deletions
```
