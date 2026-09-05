# Contributing

Thank you for contributing!

## Before you start

- Check existing issues and PRs to avoid duplicate work.
- For large changes, open an issue first to discuss the approach.

## Development setup

See [DEVELOPING.md](DEVELOPING.md) for full setup instructions.

```bash
git clone https://github.com/y-marui/alfred-quick-txt-save
cd alfred-quick-txt-save
go build ./...
```

## Making changes

1. Create a branch: `git checkout -b feat/my-feature`
2. Make your changes
3. Run checks:

```bash
make lint
make test
make build-workflow
```

4. Test in Alfred: `make build-workflow` → double-click the `.alfredworkflow`
5. Open a PR using the template

## Code style

- `gofmt` + `go vet` enforced by CI
- Keep the dependency-free `go.mod` — avoid adding third-party packages

## Commit guidelines

- Commit per **feature unit**, after confirming it works.
- **No WIP commits** — do not commit code that does not run.

### Commit message format

Follow [Conventional Commits](https://www.conventionalcommits.org/):

```
feat: add clipboard copy action
fix: collision suffix off-by-one
chore: update Go toolchain to 1.28
docs: add examples to usage.md
refactor: simplify quicksavecmd dispatch logic
```

## Pull Request checklist

- [ ] `make lint` passes
- [ ] `make test` passes
- [ ] `make build-workflow` succeeds
- [ ] New commands have tests
- [ ] `docs/specification.md` updated if user-facing changes
- [ ] `CHANGELOG.md` entry added under `[Unreleased]`
