# AGENTS

## Project
- Name: `neocex`
- Language: Go (`go1.24`)
- Entrypoint: `serve.go`
- Runtime port: `3000`

## Goals
- Keep changes small, testable, and scoped.
- Prefer improving reliability and observability over large refactors.
- Preserve existing package layout unless a task requires structural changes.

## Workflow
- Read context first with `rg --files` and targeted `sed -n`.
- Use `make build` for fast compile checks of the main server.
- Use `go test ./...` when touching shared packages.
- Do not remove or revert unrelated uncommitted changes.

## Coding conventions
- Follow Go idioms and keep functions focused.
- Keep exported APIs stable unless explicitly requested.
- Add comments only when behavior is non-obvious.
- Return concrete errors with context where practical.

## Suggested commands
- `make build`
- `make run`
- `make fmt`
- `make vet`
- `make test`

## Notes
- Current workspace may include in-progress compile issues in some packages.
- When that blocks a task, validate impacted paths first and report the exact failing files.
