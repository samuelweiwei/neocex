# neocex

A Go backend service scaffold for the `neocex` exchange project.

## Tech stack
- Go `1.24`
- Fiber v2
- GORM + PostgreSQL

## Quick start

```bash
make setup
make build
make run
```

Server default URL: `http://localhost:3000`

## Project commands

```bash
make help   # list all targets
make build  # build binary from serve.go
make run    # run locally
make fmt    # format Go files
make vet    # static checks
make test   # run tests
make tidy   # tidy go.mod/go.sum
make clean  # remove build artifacts
```

## Codex project notes
- Agent workflow and contribution rules are documented in `AGENTS.md`.
- Existing uncommitted/local changes are preserved.
- If `go test ./...` fails, fix package-level compile errors first, then rerun checks.

## Structure
- `serve.go`: local startup entrypoint
- `server/`, `router/`, `internal/`: application layers
- `config/`, `global/`, `utils/`: shared configuration and helpers
