# Tend

A warm, calm daily planner that adapts to your real capacity.

This repo is a Turborepo monorepo containing the web app, mobile app, Go API, and shared types package.

## Prerequisites

- Node.js >= 20.11.0
- pnpm >= 9 (this repo pins pnpm via Corepack — `corepack enable` once)
- Go >= 1.22 (locally installed; `apps/api/go.mod` floor is currently `1.26.2`)
- Docker (for local Postgres + building the API image)
- The `swag` CLI for OpenAPI generation: `go install github.com/swaggo/swag/v2/cmd/swag@v2.0.0-rc5`
  - The `pnpm run generate:types` script invokes it by absolute path, so you don't need `$(go env GOPATH)/bin` on `PATH`.

## Setup

```bash
pnpm install
```

## Local development

### Start Postgres locally (one-time)

```bash
docker run -d --name tend-pg \
  -p 5432:5432 \
  -e POSTGRES_USER=tend -e POSTGRES_PASSWORD=tend -e POSTGRES_DB=tend \
  postgres:16
```

Add to `apps/api/.env` (copy from `.env.example`):
```
DATABASE_URL=postgres://tend:tend@localhost:5432/tend?sslmode=disable
```

`docker stop tend-pg` to pause; `docker start tend-pg` to resume. Data persists across restarts unless you `docker rm tend-pg`.

### Run all three apps in parallel

```bash
pnpm dev
```

| App | URL |
|---|---|
| Web (Vite) | http://localhost:5173 |
| API (Gin) | http://localhost:8080 |
| Mobile (Expo) | http://localhost:8081 (Expo dev tools — press `i` / `a` to open simulator) |

API liveness checks:
- `curl http://localhost:8080/health` → `{"status":"ok"}`
- `curl http://localhost:8080/health/db` → `{"status":"ok","db":"reachable"}` (returns 503 if Postgres isn't running)

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
  api/                  Go REST API (Gin + GORM later)
    Dockerfile          Multi-stage distroless build (~24 MB image)
  web/                  Vite + React + TypeScript SPA
    vercel.json         SPA history-mode rewrites for React Router
  mobile/               Expo SDK 54 + React Native
packages/
  types/                Auto-generated TypeScript API types — DO NOT EDIT BY HAND
scripts/                Repo-level shell scripts (e.g., generate-types.sh)
.github/workflows/      CI pipelines
```

## Deployments

Three environments. Local Postgres for dev (above). Railway for staging + production API/DB. Vercel for staging + production web.

| Environment | Web | API | Trigger |
|---|---|---|---|
| Local | `pnpm dev` → `localhost:5173` | `pnpm dev` → `localhost:8080` | Manual |
| Staging | _Vercel staging URL — paste after first deploy_ | _Railway staging URL — paste after first deploy_ | Auto on push to `main` |
| Production | _Vercel production URL — paste after promotion_ | _Railway production URL — paste after promotion_ | **Manual promote** in Railway UI |

Where to look when something breaks:
- **Web build / runtime errors:** Vercel project → Deployments → select build → Logs
- **API runtime errors:** Railway project → API service → Deployments → Logs (Go `slog` JSON output)
- **DB connection issues:** Railway project → Postgres add-on → Connect tab; or `curl https://<api-url>/health/db`

### Production release

Solo-dev process — no GitHub Actions automation:

1. Confirm staging is healthy: `curl https://<api-staging>/health` and `/health/db` both 200.
2. Open Railway dashboard → API service → **Promote staging deployment to production**.
3. Vercel auto-deploys `main` to production already; nothing needed for web.
4. Tag the commit for the audit trail: `git tag v0.x.x && git push --tags`.
5. Verify production: `curl https://<api-production>/health` and `/health/db` both 200; load Vercel production URL.

If production breaks: Railway has a one-click **"Roll back to previous deployment"**. Vercel same.

### Secrets

All secrets live in Railway (API) and Vercel (web) project dashboards. Never commit `.env` files. Each `.env.example` lists the *names* of the env vars an app needs (with safe local defaults where applicable) but never real values.
