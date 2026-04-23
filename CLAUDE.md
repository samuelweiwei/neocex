# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Commands

```bash
make build        # compile binary to bin/neocex
make run          # go run ./serve.go (port 3000)
make fmt          # gofmt -w all .go files
make vet          # go vet ./...
make test         # go test ./...
make tidy         # go mod tidy
```

Always run `make build` after any change to catch compile errors quickly. Run `go test ./...` when touching shared packages (`utils/`, `global/`, `config/`).

## Architecture

The project is a Go exchange backend using **Fiber v2** (not Gin — do not introduce `gin.Context` in new code), **GORM + PostgreSQL**, and **zap** for structured logging.

### Layer flow

```
serve.go  →  server/initialize/  →  router/  →  internal/api/v1/  →  internal/service/  →  internal/models/
```

- **`serve.go`** — local dev entrypoint (`main`). Starts a Fiber app directly on `:3000`. The production path goes through `server/core/server.go → server/initialize/`.
- **`server/initialize/`** — `Routers()` builds the Fiber app and registers route groups; `initializeBizRouter` wires domain routers from `router/`.
- **`router/<domain>/`** — each domain has a `RouteGroup` and router structs that mount paths and point to API handlers.
- **`internal/api/v1/<domain>/`** — Fiber handler functions; call service layer, read request body, call `utils.GetUserID(f)` for auth.
- **`internal/service/<domain>/`** — business logic structs. Aggregated via `ServiceGroup` in `internal/service/enter.go`.
- **`internal/models/`** — GORM model structs (`contract/`, `user/`) and request DTOs (`contract/request/`, `global/req/`).

### Registration pattern (all three layers use the same singleton pattern)

Every layer exposes a single `var XxxGroupApp = new(XxxGroup)` aggregate and a `XxxGroup` struct composed of sub-group structs. Adding a new domain means:
1. Create `router/<domain>/enter.go` with a `RouteGroup` struct.
2. Add it to `router/RouterGroup` in `router/enter.go`.
3. Create `internal/api/v1/<domain>/enter.go` with an `ApiGroup` struct.
4. Add it to `internal/api/v1/ApiGroup`.
5. Create `internal/service/<domain>/enter.go` with a `ServiceGroup` struct.
6. Add it to `internal/service/ServiceGroup`.
7. Wire in `server/initialize/biz_router_init.go`.

### JWT / Auth (`utils/`)

- `utils/jwt.go` — `JWT` struct; construct with `NewJWT(&global.GVA_CONF.JWT)`. Methods: `GenerateToken(BaseClaims)`, `ParseToken(string)`, `RefreshToken(string)`. Sentinel errors: `TokenExpired`, `TokenNotValidYet`, `TokenMalformed`, `TokenInvalid`.
- `utils/auth.go` — Fiber helpers: `GetClaims(f)`, `GetToken(f)`, `SetToken(f, token, maxAge)`, `GetUserID(f)`. Claims type is `req.CustomerClaims` from `internal/models/global/req/jwt.go`.
- JWT config (`config/jwt.go`) stores `SigningKey`, `BufferTime` (duration string e.g. `"1h"`), `ExpiredTime` (e.g. `"24h"`), `Issuer`. Loaded into `global.GVA_CONF.JWT`.

### Globals (`global/global.go`)

`GVA_CONF` (config), `GVA_DB` (gorm), `GVA_REDIS`, `GVA_LOG`, `GVA_JWT` — all package-level vars. Initialize before starting handlers.

### Logging & i18n

- `logging.Logger` — package-level `*zap.Logger`, initialized via `init()`.
- `i18n.T("error.key", data)` — returns translated string; falls back to the key if not found. Translation files live in `i18n/*.json`.

## Conventions

- Use `jwt/v5` (`github.com/golang-jwt/jwt/v5`) throughout; `jwt/v4` is a transitive dependency only.
- Handler signatures are always `func(f *fiber.Ctx) error`.
- Service methods take typed arguments (e.g. `userID uint`), not framework context objects.
- `utils.GetUserID(f)` extracts user ID from a Fiber context via JWT claims.
