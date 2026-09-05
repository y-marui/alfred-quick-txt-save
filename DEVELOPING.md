# Developing

Developer guide for building, testing, and extending alfred-quick-txt-save.

For contribution guidelines (branching, PR process, commit format), see [CONTRIBUTING.md](CONTRIBUTING.md).

## Prerequisites

- macOS (required for Alfred)
- Go (see `go.mod` for the toolchain version)
- Alfred 5 with Powerpack
- `jq` (optional, for pretty-printed dev output): `brew install jq`
- `gh` CLI (required for releases): `brew install gh`

## Development Setup

```bash
git clone https://github.com/y-marui/alfred-quick-txt-save
cd alfred-quick-txt-save
go build ./...
```

## Daily Workflow

### Simulate Alfred locally

```bash
go run ./cmd/quick-txt-save-alfred list ""            # default filename preview
go run ./cmd/quick-txt-save-alfred list "mynotes"
go run ./cmd/quick-txt-save-alfred list "notes.md"
text="hello" go run ./cmd/quick-txt-save-alfred write /tmp/out.txt  # writes "hello" to /tmp/out.txt
```

Pipe through `jq` for pretty-printed JSON:

```bash
go run ./cmd/quick-txt-save-alfred list "" | jq .
```

### Run tests

```bash
make test          # go test ./...
```

### Lint and format

```bash
make lint          # gofmt -l + go vet
make fmt           # gofmt -w (auto-fix)
```

## Adding a New Command

This workflow currently ships a single command (`save`). To add another Script Filter
command:

1. Add the domain logic to a new `internal/<domain>/` package (stdlib only, Alfred-independent,
   unit-testable).
2. Add an `internal/<domain>cmd/` package with a function returning `scriptfilter.Response`,
   following `internal/quicksavecmd/quicksavecmd.go`'s shape.
3. Register a new subcommand in `cmd/quick-txt-save-alfred/main.go`'s `switch` statement.
4. Add tests for both new packages.
5. Add a Script Filter (and, if it has a side effect, a Run Script) node to
   `workflow/info.plist`, wired to the new subcommand.
6. Update `docs/specification.md`, `README.md`/`README-jp.md`, and `CHANGELOG.md`.

## Building the Package

```bash
make build-workflow
```

Output: `dist/<name>-<version>.alfredworkflow`

Install during development: double-click the `.alfredworkflow` file,
or drag it into Alfred Preferences → Workflows.

## Testing in Alfred

1. Build: `make build-workflow`
2. Install: open `dist/*.alfredworkflow`
3. Open Alfred, type `save`

During rapid iteration you can symlink `workflow/` to Alfred's workflow directory,
but `go run ./cmd/quick-txt-save-alfred list "query"` is usually faster for logic changes.

## Naming Conventions

| Scope | Convention | Example |
|---|---|---|
| Go packages | short, lowercase, no underscores | `quicksave`, `quicksavecmd`, `scriptfilter` |
| Exported functions / types | `PascalCase` | `Save`, `Response`, `Item` |
| Unexported functions / variables | `camelCase` | `readClipboard`, `defaultPrefix` |
| Alfred command names | lowercase | `"save"` |
| Alfred variable names | `lowercase_with_underscores` | `save_dir`, `file_prefix`, `file_ext` |
| Commit messages | Conventional Commits | `feat:`, `fix:`, `docs:`, `chore:` |
| Branch names | `feat/`, `fix/`, `docs/`, `chore/`, `work/` | `feat/add-open-browser` |

## Code Style

- **Formatter:** `gofmt`. CI enforces this (`make lint`).
- **Linter:** `go vet`.
- **Comments:** Write *why*, not *what*. Do not comment self-evident code.
- **No debug prints:** Remove all stray `fmt.Print*` statements before committing;
  the only writer to stdout is `scriptfilter.Response.Write`.
- **No third-party dependencies** unless clearly justified — keep `go.mod` dependency-free.

## Releasing

```bash
# 1. Update version in workflow/info.plist
# 2. Update CHANGELOG.md
git add workflow/info.plist CHANGELOG.md
git commit -m "chore: release v1.2.3"

# 3. Tag and push
git tag v1.2.3
git push origin main --tags
# GitHub Actions builds .alfredworkflow and creates a GitHub Release
```

## Commit Guidelines

- Commit per **feature unit**, after confirming it works.
- **No WIP commits** — do not commit code that does not run.
- **No `--no-verify`** — never skip pre-commit hooks.

### Commit Message Format

Follow [Conventional Commits](https://www.conventionalcommits.org/):

```
feat: add clipboard copy action
fix: collision suffix off-by-one
chore: update Go toolchain to 1.28
docs: update README usage section
refactor: simplify quicksavecmd dispatch logic
```

## Pull Request Checklist

- [ ] `make lint` passes
- [ ] `make test` passes
- [ ] `make build-workflow` succeeds
- [ ] New commands have tests
- [ ] `README.md`/`README-jp.md` updated if user-facing changes
- [ ] `CHANGELOG.md` entry added under `[Unreleased]`

## Code Review Guidelines

**Reviewers check for:**
- Architectural constraints respected (no business logic in `cmd/quick-txt-save-alfred`)
- No hardcoded absolute paths (use `$HOME` / env vars)
- No debug prints in production code
- No Unicode emoji in Alfred result item `title` / `subtitle`
- Tests cover the new or changed behavior
- Alfred env variables managed via Config Builder, not `variables` key

**Security-sensitive changes** (auth, encryption, data access) require explicit
security review before merge.

**Self-review:** Individual contributors open a PR and self-review before merging
to `main`.

## AI Development Workflow

### Claude Code (major features, refactoring, tests)

Claude Code reads `CLAUDE.md` at the project root for context. Use it for:
- Implementing new commands and internal packages
- Refactoring existing code
- Writing test suites
- Reviewing architecture decisions

### GitHub Copilot (bug fixes, inline completions)

Copilot works best for:
- Fixing small bugs inline
- Completing repetitive boilerplate
- Suggesting Go idioms

### Gemini CLI (documentation)

Use Gemini CLI for:
- Generating/updating `README.md`
- Writing `CHANGELOG.md` entries from git log

Example:
```bash
gemini "Update README.md based on the current source code in cmd/ and internal/"
gemini "Generate CHANGELOG entry for commits since v1.0.0"
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
| `gofmt` / `go-vet` | Formatting and suspicious Go constructs |

Do **not** skip hooks with `--no-verify`.

### Development Security Rules

- Never store secrets in `workflow/info.plist` or committed files; use Alfred's built-in encrypted keychain instead.
- Alfred query strings are passed to `cmd/quick-txt-save-alfred`; do not interpolate them into shell commands without sanitization.
- This project has no third-party Go dependencies by design; Dependabot monitors `.github/workflows/` automatically.

For vulnerability reporting, see [SECURITY.md](SECURITY.md).
