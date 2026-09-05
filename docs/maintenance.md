# Maintenance Prompts

Prompts for a local LLM (read-only) to draft updates to `docs/` files. The
primary AI reviews and verifies against real files before saving anything.

## Update docs/architecture.md

```
Review this project's directory structure and entry points.

Steps:
1. List top-level directories and their purpose
2. Identify entry points (main packages, executables)
3. List key external dependencies (from go.mod)
4. Propose an update to docs/architecture.md following this format:

# Architecture

## Overview
<!-- 3 lines max -->

## Entry Points
- `path/file` — description

## Directory Structure
| Directory | Role |
|---|---|

## Key Dependencies
| Library / Module | Purpose |
|---|---|

Note: keep Overview to 3 lines max. Don't list file-level detail (that's
file-map.md's job). List only key dependencies, not every one.
```

## Update docs/file-map.md

```
For the files touched in this session, propose additions to docs/file-map.md
following this format:

# File Map

_Last updated: YYYY-MM-DD_

## [Module / Feature Name]
| File | Role | Key Dependencies |
|---|---|---|

Note: don't try to cover every file — only ones actually read or edited this
session. Don't list unexplored files. Update the "Last updated" date.
```
