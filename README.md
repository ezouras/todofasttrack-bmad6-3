# Tend

A warm, calm daily planner that adapts to your real capacity.

This repo is a Turborepo monorepo containing the web app, mobile app, Go API, and shared types package.

## Prerequisites

- Node.js >= 20.11.0
- pnpm >= 9 (this repo pins pnpm via Corepack — `corepack enable` once)
- Go >= 1.22
- The `swag` CLI for OpenAPI generation: `go install github.com/swaggo/swag/v2/cmd/swag@latest`
  - Make sure `$(go env GOPATH)/bin` is on your `PATH`, or rely on `pnpm run generate:types` which invokes it by absolute path.

## Setup

```bash
pnpm install
```

## Local development

Run all three apps in parallel:

```bash
pnpm dev
```

| App | URL |
|---|---|
| Web (Vite) | http://localhost:5173 |
| API (Gin) | http://localhost:8080 |
| Mobile (Expo) | http://localhost:8081 (Expo dev tools — press `i` / `a` to open simulator) |

API liveness check: `curl http://localhost:8080/health` → `{"status":"ok"}`.

## Type generation

The Go API is the source of truth for API types. After changing any Go model or handler, regenerate the TypeScript types:

```bash
pnpm run generate:types
```

This runs `swag init` to produce `apps/api/docs/swagger.json`, then `openapi-typescript` to produce `packages/types/generated/api.ts`. Commit both. CI will fail if these are out of sync.

## Other commands

| Command | What it does |
|---|---|
| `pnpm build` | Build all apps |
| `pnpm typecheck` | `tsc --noEmit` across web + mobile |
| `pnpm lint` | ESLint across web + mobile |
| `pnpm test` | Test runners across all workspaces (Go test runs in `apps/api`; web + mobile are no-ops until tests land) |
| `cd apps/api && go test ./...` | Go tests directly |
| `cd apps/api && go vet ./...` | Go static checks |

## CI

GitHub Actions runs three jobs in parallel on every PR and every push to `main`:

1. **`go-checks`** — `go vet` + `go test` against `apps/api`
2. **`ts-checks`** — `pnpm typecheck` + `pnpm lint` for web + mobile
3. **`openapi-sync`** — regenerates types from the Go OpenAPI spec; fails if `packages/types/generated/api.ts`, `swagger.json`, `swagger.yaml`, or `docs.go` would change

Workflow lives at `.github/workflows/ci.yml`.

## Repo layout

```
apps/
  api/                  Go REST API (Gin + GORM)
  web/                  Vite + React + TypeScript SPA
  mobile/               Expo SDK 54 + React Native
packages/
  types/                Auto-generated TypeScript API types — DO NOT EDIT BY HAND
.github/workflows/      CI pipelines
```
