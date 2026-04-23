# Developing

Developer guide for building, testing, and extending alfred-quick-txt-save.

For contribution guidelines (branching, PR process, commit format), see [CONTRIBUTING.md](CONTRIBUTING.md).

## Prerequisites

- macOS (required for Alfred)
- Python 3.11+
- Alfred 5 with Powerpack
- `jq` (optional, for pretty-printed dev output): `brew install jq`
- `gh` CLI (required for releases): `brew install gh`

## Setup

```bash
git clone https://github.com/y-marui/alfred-quick-txt-save
cd alfred-quick-txt-save
make install
```

## Daily Workflow

### Simulate Alfred locally

```bash
make run Q="save"
make run Q="save mynotes"
make run Q="save dir ~/Desktop"
make run Q="search foo"
make run Q="open repo"
make run Q="config"
make run Q=""
```

`scripts/dev.sh` sets all `alfred_workflow_*` env vars to temp directories, calls
`workflow/scripts/entry.py` with your query, and pretty-prints the JSON output.

### Run tests

```bash
make test          # fast
make test-cov      # with coverage
```

### Lint and format

```bash
make lint          # check
make format        # auto-fix
make typecheck     # mypy
```

## Adding a New Command

1. Create `src/app/commands/my_cmd.py` with `handle(args: str) -> None`
2. Register in `src/app/core.py`: `router.register("my")(my_cmd.handle)`
3. Add tests in `tests/test_commands.py`
4. Update `docs/specification.md` and `workflow/info.plist` keyword help

## Adding a Third-Party Dependency

1. Add to `requirements.txt`
2. Run `make vendor`
3. Import in your code — the vendor path is added by `entry.py`

## Building the Package

```bash
make build
```

Output: `dist/<name>-<version>.alfredworkflow`

Install during development: double-click the `.alfredworkflow` file,
or drag it into Alfred Preferences → Workflows.

## Testing in Alfred

1. Build: `make build`
2. Install: open `dist/*.alfredworkflow`
3. Open Alfred, type your keyword

To iterate quickly, you can also point Alfred's workflow directory directly
at the `workflow/` folder during development (see Alfred docs on workflow
symlinks), but the `make run` simulator is usually faster.

## Releasing

```bash
# 1. Update version in pyproject.toml
# 2. Update CHANGELOG.md
# 3. Commit
git add pyproject.toml CHANGELOG.md
git commit -m "chore: release v1.2.3"

# 4. Tag
git tag v1.2.3
git push origin main --tags

# GitHub Actions automatically builds and releases.
# To release manually:
make build
make release
```

## AI Development Workflow

### Claude Code (major features, refactoring, tests)

Claude Code reads `CLAUDE.md` at the project root for context. Use it for:
- Implementing new commands and services
- Refactoring existing code
- Writing test suites
- Reviewing architecture decisions

### GitHub Copilot (bug fixes, inline completions)

Copilot works best for:
- Fixing small bugs inline
- Completing repetitive boilerplate
- Suggesting type annotations

### Gemini CLI (documentation)

Use Gemini CLI for:
- Generating/updating `README.md`
- Writing `CHANGELOG.md` entries from git log

Example:
```bash
gemini "Update README.md based on the current source code in src/"
gemini "Generate CHANGELOG entry for commits since v1.2.3"
```

## Security

### Automated Checks

The following hooks run on every commit (pre-commit) and in CI (`security` job):

| Hook | What it detects |
|---|---|
| `gitleaks` (`.gitleaks.toml`) | Hardcoded secrets, API keys, local absolute paths |
| `detect-private-key` | SSH/TLS private key headers |
| `no-commit-dotenv` | `.env` files accidentally staged |
| `check-added-large-files` | Files over 500 KB |

Do **not** skip hooks with `--no-verify`.

### Development Security Rules

- Never store secrets in `workflow/info.plist` or committed files; use Alfred's built-in encrypted keychain instead.
- Alfred query strings are passed to `entry.py` — do not interpolate them into shell commands or SQL without sanitization.
- Keep vendored packages in `workflow/vendor/` up-to-date; Dependabot monitors `.github/workflows/` automatically.

For vulnerability reporting, see [SECURITY.md](SECURITY.md).
