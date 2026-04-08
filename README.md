# Dev Charter

> **This is the reference (English) version.**
> For the canonical (Japanese) version, see [README-jp.md](README-jp.md).

[![License: CC BY 4.0](https://img.shields.io/badge/License-CC%20BY%204.0-lightgrey.svg)](LICENSE)
[![check-charter CI](https://github.com/y-marui/dev-charter/actions/workflows/check-charter.yml/badge.svg)](https://github.com/y-marui/dev-charter/actions/workflows/check-charter.yml)

Shared development charter for AI-assisted software projects.

This repository defines common philosophy, architecture principles,
and development rules used across projects.

## Documents

| File | Description |
|---|---|
| [PRINCIPLES.md](PRINCIPLES.md) | Development philosophy, design and architecture principles |
| [CODE_STYLE.md](CODE_STYLE.md) | Code style guide |
| [AI_COLLABORATION_RULES.md](AI_COLLABORATION_RULES.md) | AI collaboration rules and role assignments |
| [AI_CONTEXT_HIERARCHY.md](AI_CONTEXT_HIERARCHY.md) | AI context priority hierarchy |
| [AI_TOOL_SETUP.md](AI_TOOL_SETUP.md) | AI context file structure spec (AI_CONTEXT.md and agent config files) |
| [LANGUAGE_POLICY.md](LANGUAGE_POLICY.md) | Language policy (canonical = Japanese) |
| [LOCALIZATION_POLICY.md](LOCALIZATION_POLICY.md) | Localization and supported languages |
| [PROJECT_LIFECYCLE.md](PROJECT_LIFECYCLE.md) | Project lifecycle and team structure |
| [UI_GUIDELINES.md](UI_GUIDELINES.md) | UI guidelines, color palette, iconography |
| [MONETIZATION_POLICY.md](MONETIZATION_POLICY.md) | Monetization policy and platform-specific guidelines |
| [SECURITY_POLICY.md](SECURITY_POLICY.md) | Security policy and hook configuration reference |
| [LEGAL_POLICY.md](LEGAL_POLICY.md) | License selection criteria and templates (Closed / MIT / AGPL·GPL·LGPL) |
| [topics/CI_POLICY.md](topics/CI_POLICY.md) | CI job design and Branch Protection Ruleset |
| [topics/GITHUB_SETTINGS.md](topics/GITHUB_SETTINGS.md) | GitHub repository settings review (Ruleset, Sponsors) |
| [topics/GITHUB_CONTRIBUTING.md](topics/GITHUB_CONTRIBUTING.md) | Issue, PR, CONTRIBUTING.md, PR template, and Quasi-CLA (for OSS) |
| [topics/TEMPLATE_README_GUIDELINES.md](topics/TEMPLATE_README_GUIDELINES.md) | GitHub template repository README guidelines (environment, language, LICENSE, required sections) |
| [topics/PROJECT_README_GUIDELINES.md](topics/PROJECT_README_GUIDELINES.md) | README setup guide for projects created from a template |

## How to Use

1. Pull dev-charter into `docs/dev-charter/` via `git subtree`
2. Have the AI read the charter and generate `AI_CONTEXT.md` and agent config files at the project root
3. After charter updates, run `git subtree pull` and have the AI sync the context files

See [AI_TOOL_SETUP.md](AI_TOOL_SETUP.md) for the structure spec.

## Install (git subtree)

```
git remote add dev-charter https://github.com/y-marui/dev-charter
git fetch dev-charter
git subtree add --prefix=docs/dev-charter dev-charter main --squash
```

After installing, follow the [Install Checklist](INSTALL_CHECKLIST.md).

## Update

If the `dev-charter` remote is not set up (e.g., after cloning the project), add it first:

```
git remote add dev-charter https://github.com/y-marui/dev-charter
git subtree pull --prefix=docs/dev-charter dev-charter main --squash
```

After updating, follow the [Update Checklist](UPDATE_CHECKLIST.md).

## Makefile helper

```
update-charter:
	git remote | grep -q '^dev-charter$$' || \
	  git remote add dev-charter https://github.com/y-marui/dev-charter
	git fetch dev-charter
	git subtree pull --prefix=docs/dev-charter dev-charter main --squash
```

## Version Check (CI)

Add `.github/workflows/dev-charter-check.yml` to your project to automatically
check for updates weekly and open a PR when a new version is available.

```yaml
name: Dev Charter
on:
  schedule:
    - cron: "23 3 * * 1"  # Every Monday at 03:23 UTC — change to your own random minute/hour/day-of-week
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

> **Note:** If your repository has Branch Protection rules that prevent direct pushes,
> add a bypass rule for the GitHub Actions bot
> (Settings > Rules > Rulesets > Bypass list > GitHub Actions).

## Badge for Adopting Projects

Place this badge in your project README to show dev-charter update health.

### Workflow Status Badge

Shows whether dev-charter is up to date. Requires `fail_if_outdated: true` in the workflow (see above).

```markdown
[![Charter Check](https://github.com/{owner}/{repo}/actions/workflows/dev-charter-check.yml/badge.svg)](https://github.com/{owner}/{repo}/actions/workflows/dev-charter-check.yml)
```

Replace `{owner}` and `{repo}` with your GitHub organization and repository name.

| State | Status Badge |
|---|---|
| Not installed / CI not set up | red (VERSION not found) |
| Installed, up to date | green |
| Installed, outdated | red |

---

*This document has a Japanese canonical version [README-jp.md](README-jp.md). Update both in the same commit when editing.*
