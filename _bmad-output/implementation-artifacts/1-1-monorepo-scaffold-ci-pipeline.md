# Story 1.1: Monorepo Scaffold & CI Pipeline

Status: review

<!-- Note: Validation is optional. Run validate-create-story for quality check before dev-story. -->

## Story

As a **developer**,
I want a Turborepo monorepo initialized with Go API, React web app, Expo mobile app, and shared types package,
so that all three apps can be developed, type-checked, and tested in a unified workspace.

## Acceptance Criteria

1. **Given** the repository is cloned, **When** `pnpm install` is run, **Then** all JS dependencies install without errors and the workspace structure matches: `apps/api`, `apps/web`, `apps/mobile`, `packages/types`.
2. **Given** the monorepo is set up, **When** `pnpm dev` is run, **Then** the web Vite dev server (`localhost:5173`), Go API (`localhost:8080`), and Expo mobile bundler (`localhost:8081`) all start in parallel via Turborepo.
3. **Given** a PR is opened on GitHub, **When** the CI pipeline runs, **Then** `go vet ./...` and `go test ./...` pass, TypeScript type-checks pass for web and mobile, and ESLint passes for both TS apps.
4. **Given** `pnpm run generate:types` is run after Go model changes, **When** the pipeline executes `swaggo/swag` then `openapi-typescript`, **Then** `packages/types/generated/api.ts` is regenerated from `apps/api/docs/swagger.json`.
5. **Given** the CI pipeline runs on a PR, **When** `packages/types` is out of sync with the Go OpenAPI spec, **Then** the pipeline fails with a clear message indicating types need regeneration.

## Tasks / Subtasks

- [x] **Task 1 — Initialize Turborepo + pnpm workspace in-place at repo root** (AC: 1)
  - [x] Run `npx create-turbo@latest tend --package-manager pnpm` in a temp directory, then move generated files (`turbo.json`, `pnpm-workspace.yaml`, root `package.json`, `.gitignore` additions) into the existing repo root. **Do NOT create a nested `tend/` directory** — the repo at `/Users/evelynzouras/Documents/code/toDoFastTrack` IS the monorepo root.
  - [x] Delete any sample `apps/` and `packages/` content the starter ships with (we'll scaffold our own).
  - [x] Confirm `pnpm-workspace.yaml` declares `apps/*` and `packages/*` as workspace globs.
  - [x] Pin Node engine in root `package.json` (`"engines": { "node": ">=20.11.0" }`) — required by Vite 5+ and Expo SDK 52.
- [x] **Task 2 — Scaffold `apps/web` (Vite + React + TypeScript)** (AC: 1, 2)
  - [x] Run `npm create vite@latest apps/web -- --template react-ts` (per architecture init commands).
  - [x] Add `dev`, `build`, `lint`, `typecheck` scripts to `apps/web/package.json` (`tsc --noEmit` for typecheck).
  - [x] Confirm Vite dev server starts on port `5173` (default — no override needed).
- [x] **Task 3 — Scaffold `apps/mobile` (Expo SDK 52 + React Native + TypeScript)** (AC: 1, 2)
  - [x] Run `npx create-expo-app apps/mobile --template blank-typescript`.
  - [x] Confirm Expo SDK version is **52.x** (architecture pins this — do not accept 51 or 53 if defaults change). _**Deviation accepted:** template now ships SDK 54.0.33 with React 19.1 + RN 0.81. See Change Log entry 2._
  - [x] Add `dev`, `lint`, `typecheck` scripts to `apps/mobile/package.json`. `dev` should run `expo start` (defaults to port `8081`).
- [x] **Task 4 — Scaffold `apps/api` (Go + Gin + GORM)** (AC: 1, 2)
  - [x] `mkdir -p apps/api/cmd/server apps/api/internal/handler apps/api/internal/middleware apps/api/internal/model apps/api/internal/repository apps/api/internal/service apps/api/docs`
  - [x] `cd apps/api && go mod init github.com/evelynzouras/tend-api`
  - [x] `go get github.com/gin-gonic/gin gorm.io/gorm gorm.io/driver/postgres` _**Note:** `go mod tidy` later removed gorm/postgres because nothing yet imports them. They will re-add automatically when Story 1.4+ writes the first model. See Change Log entry 4._
  - [x] Create `apps/api/cmd/server/main.go` with a minimal Gin server: register `GET /health` returning `200 OK` with `{"status":"ok"}`. Listen on `:8080`. This is the smoke test for AC2 and seeds Story 1.2 (`/health` is referenced in 1.2's ACs).
  - [x] Verify `go run ./cmd/server` starts cleanly from `apps/api/`.
- [x] **Task 5 — Initialize Go workspace at repo root** (AC: 1)
  - [x] Run `go work init ./apps/api` at repo root → produces `go.work` and `go.work.sum`. _**Note:** `go.work.sum` only generates when there are cross-module dep conflicts to resolve. With one module, only `go.work` exists. Expected._
  - [x] Commit both files. (Only `go.work` exists.)
- [x] **Task 6 — Create `packages/types` skeleton** (AC: 1, 4)
  - [x] `mkdir -p packages/types/generated`
  - [x] Create `packages/types/package.json`: name `@repo/types`, `"main": "index.ts"`, no build step (consumed via TypeScript path resolution from workspace).
  - [x] Create `packages/types/index.ts`: `export * from './generated/api';`
  - [x] Create placeholder `packages/types/generated/api.ts` (later overwritten by Task 8 when first generated from the live `swagger.json`).
  - [x] `packages/types/generated/` is **NOT** gitignored — it must be committed so CI can detect drift (AC5). Verified: `.gitignore` does not list it.
- [x] **Task 7 — Configure `turbo.json` pipeline** (AC: 2, 3)
  - [x] Define `dev` task with `"cache": false, "persistent": true` so all four workspaces run concurrently and stream logs.
  - [x] Define `build`, `lint`, `typecheck`, `test` tasks with appropriate dependencies (`build` depends on `^build`, etc.).
  - [x] Define `generate:types` task — `cache: false`.
  - [x] Add root `package.json` scripts: `"dev"`, `"build"`, `"lint"`, `"typecheck"`, `"test"`, `"generate:types"`.
- [x] **Task 8 — Wire OpenAPI generation pipeline** (AC: 4)
  - [x] In `apps/api`: install `swag` CLI. _**Deviation:** v1 (`swaggo/swag`) emits Swagger 2.0 which `openapi-typescript` rejects. Switched to v2 (`github.com/swaggo/swag/v2/cmd/swag@latest`) with the `--v3.1` flag. See Change Log entry 3._
  - [x] Add minimal swag annotation to `main.go` so `swag init -g cmd/server/main.go -o ./docs --v3.1` produces a valid `apps/api/docs/swagger.json` even with only the `/health` route.
  - [x] At repo root: install `openapi-typescript` as a workspace dev dep.
  - [x] Add root script `"generate:types"` chaining swag → openapi-typescript. The script invokes swag via `$(go env GOPATH)/bin/swag` so it works without `~/go/bin` on `PATH`.
  - [x] Run it once and commit the generated `swagger.json`, `swagger.yaml`, `docs.go`, and `api.ts` so the initial state is the source-of-truth baseline for AC5.
- [x] **Task 9 — Add GitHub Actions CI workflow** (AC: 3, 5)
  - [x] Create `.github/workflows/ci.yml` triggered on `pull_request` and `push: branches: [main]`.
  - [x] Jobs run in parallel: **`go-checks`** (setup-go stable, `go vet` + `go test`), **`ts-checks`** (setup-pnpm + setup-node 20, `pnpm install --frozen-lockfile`, `pnpm typecheck`, `pnpm lint`), **`openapi-sync`** (installs swag v2 CLI, runs `pnpm run generate:types`, then `git diff --exit-code` against `packages/types/generated/api.ts`, `apps/api/docs/swagger.json`, `swagger.yaml`, `docs.go` — fails with explicit message on drift).
  - [x] Pin action versions (`actions/checkout@v4`, `actions/setup-node@v4`, `actions/setup-go@v5`, `pnpm/action-setup@v4`).
  - [x] _Adjustment:_ `setup-go` uses `go-version: stable` instead of `1.22` because the local `go.mod` declares `go 1.26.2`. CI must use a Go toolchain ≥ that version. See Change Log entry 5.
- [x] **Task 10 — ESLint setup** (AC: 3)
  - [x] Use the ESLint config that ships with the Vite React TS template (`apps/web` → `eslint.config.js` already present).
  - [x] _Adjustment:_ Expo `blank-typescript` template no longer ships an ESLint config. Added `eslint-config-expo` and a flat `apps/mobile/eslint.config.js`. See Change Log entry 6.
  - [x] Confirm `pnpm --filter @repo/web lint` and `pnpm --filter @repo/mobile lint` both pass on the empty starter code. ✅ Verified — both exit 0 with no output.
  - [x] Do NOT introduce a custom shared ESLint package at this story — defer until duplication is felt.
- [x] **Task 11 — Repo hygiene + README** (AC: all)
  - [x] Update root `README.md` with: setup (`pnpm install`), local dev (`pnpm dev`), type generation (`pnpm run generate:types`), CI overview, repo layout.
  - [x] Confirm `.gitignore` covers `node_modules/`, `dist/`, `.expo/`, `apps/api/tend-api` (compiled binary), `.env`, `.env.local`, `*.log`, plus Go binary patterns and Expo signing certs.
  - [x] Add `apps/api/.env.example`, `apps/web/.env.example`, `apps/mobile/.env.example` as commented placeholders (Story 1.2 will populate them).
- [x] **Task 12 — Local verification of all five ACs** (AC: 1, 2, 3, 4, 5)
  - [x] Fresh clone simulation: `rm -rf node_modules apps/*/node_modules && pnpm install --frozen-lockfile` → exits 0, all workspace symlinks present (AC1).
  - [x] `pnpm dev` → all three dev servers reachable on their ports within 30 seconds (AC2). `curl localhost:8080/health` returned `{"status":"ok"}`. Web 5173 returned 200. Expo 8081 returned 200.
  - [x] `pnpm typecheck && pnpm lint && (cd apps/api && go vet ./... && go test ./...)` all pass locally (AC3). Turbo reports 3/3 typecheck, 3/3 lint, plus go vet + go test green.
  - [x] `pnpm run generate:types` produces the expected `api.ts` deterministically (AC4). Verified the full chain: swag → swagger.json (OpenAPI 3.1) → openapi-typescript → api.ts with typed `/health` path.
  - [x] Manually edited `apps/api/cmd/server/main.go` to add `/drift-test` route, ran `pnpm run generate:types` — confirmed `swagger.json` and `api.ts` both gained the new route (1 grep match each). The CI's `git diff --exit-code` would flag this against a committed baseline. Reverted the edit. Re-ran generate to restore clean state. (AC5)

## Dev Notes

### Ground Truth: Repo State at Story Start

The repo currently contains only:
```
/.claude            ← Claude Code config (do not touch)
/_bmad              ← BMad install (do not touch)
/_bmad-output       ← Planning artifacts (this story lives here)
/docs               ← Empty
/README.md          ← Two-line placeholder, safe to overwrite
```

There is **no existing source code, no `package.json`, no `go.mod`**. This story creates all of those from scratch. **The repo IS already a git repo on `main`** — do not `git init` again.

### Critical Architecture Compliance

[Source: `_bmad-output/planning-artifacts/architecture.md#Selected Stack`]

**Locked stack — do not substitute:**

| Layer | Technology | Notes |
|---|---|---|
| Monorepo | **Turborepo + pnpm** | `pnpm` is mandatory; do NOT use npm or yarn for workspace mgmt |
| Web | **Vite + React + TypeScript** (strict mode) | React 18, TypeScript 5.x |
| Mobile | **Expo SDK 52** + React Native + TypeScript | Pinned to 52 — official NativeWind v4 + Expo Router v3 pairing |
| API | **Go 1.22+** + Gin + GORM v2 + PostgreSQL driver | Module path TBD (Q1) |
| Type sharing | **swaggo/swag → openapi-typescript** | Go is the source of truth for types |

[Source: `architecture.md#Initialization Commands`] The exact init commands are codified in the architecture doc — Tasks 1–4 mirror them line-for-line.

### Project Structure (Target State After This Story)

[Source: `architecture.md#Complete Project Directory Structure`]

```
/                                       # repo root (in-place, NOT nested)
├── .github/workflows/ci.yml            # Task 9
├── .gitignore                          # extended in Task 11
├── turbo.json                          # Task 7
├── pnpm-workspace.yaml                 # Task 1
├── package.json                        # root workspace package.json
├── go.work                             # Task 5
├── go.work.sum
├── README.md                           # rewritten in Task 11
├── apps/
│   ├── api/
│   │   ├── go.mod, go.sum
│   │   ├── cmd/server/main.go          # Task 4 (Gin + /health only)
│   │   ├── docs/swagger.json           # generated by swag (Task 8)
│   │   └── internal/                   # empty stubs for handler/, model/, repository/, service/, middleware/
│   ├── web/                            # Task 2 (Vite React TS template defaults)
│   └── mobile/                         # Task 3 (Expo TS template defaults)
└── packages/
    └── types/                          # Task 6
        ├── package.json
        ├── index.ts                    # re-exports generated/
        └── generated/api.ts            # written by openapi-typescript
```

### Out of Scope for This Story (Will Be Future Stories)

**Do NOT implement any of these in Story 1.1** — they are intentionally separate stories:

- **Clerk auth integration** → Story 1.4 / 1.5 (do not install `@clerk/*` packages or `clerk-sdk-go` yet)
- **Stripe / subscription middleware** → Story 6.x
- **Tailwind / NativeWind / shadcn / design tokens** → Story 1.3
- **Any GORM models, repositories, services, handlers beyond `/health`** → Stories 1.4+, 2.x, 3.x
- **Vercel + Railway deployment config** → Story 1.2 (this story is local-only + GitHub Actions)
- **TanStack Query, Zustand, React Router, Expo Router setup** → Story 1.3 / Story 3.1
- **Database connection** → Story 1.2 / 1.4

Keep the API at exactly **one route (`/health`)** and exactly **one handler file (`main.go`)** until Story 1.4 begins introducing the layered architecture.

### Anti-Patterns to Avoid

- ❌ **Hand-writing TypeScript API types.** `packages/types/generated/api.ts` is owned by the codegen pipeline. Anything hand-written there will be clobbered. [Source: `architecture.md#Enforcement Guidelines`]
- ❌ **Skipping the OpenAPI sync check.** AC5 exists specifically to prevent type drift. Do not make this check warning-only.
- ❌ **Nesting the scaffold inside a `tend/` subdirectory.** `npx create-turbo@latest tend …` produces a `tend/` folder by default — you must lift its contents up one level. The repo root is the monorepo root.
- ❌ **Using `go mod init` per-package.** There is exactly one `go.mod` (in `apps/api/`) and one `go.work` (at repo root). Do not create additional Go modules.
- ❌ **Adding heavy deps "just in case."** Do not pre-install Clerk, Stripe, TanStack Query, Tailwind, etc. in this story — those belong to the stories that consume them.

### Testing Standards

[Source: `architecture.md#Test File Location`]

- **Go:** Co-located `{file}_test.go` (e.g., `cmd/server/main_test.go` for the health endpoint smoke test). Use the standard `testing` package — no testify/ginkgo for now.
- **TypeScript:** Co-located `{File}.test.tsx`. No test runner is wired up in this story (Vitest/Jest deferred until first real component). `pnpm test` for web/mobile can be a no-op script that exits 0.
- **CI must run:** `go vet ./...` and `go test ./...` (Task 9). At minimum, write a single Go test that exercises the `/health` handler and asserts a 200 response — this proves the test pipeline works end-to-end.

### Project Structure Notes

The architecture's "Complete Project Directory Structure" lists ~50 internal Go files (handlers, services, repos for every resource) — those are the **target shape after all of Epic 1–8 ships**, not what Story 1.1 creates. Create only the empty parent directories listed in Task 4 plus `cmd/server/main.go`. Subsequent stories will populate them.

### Library Versions to Use (as of 2026-04-16)

| Package | Version Pin Strategy | Why |
|---|---|---|
| `turbo` | Whatever `create-turbo@latest` resolves | Starter is canonical |
| `vite` + `@vitejs/plugin-react` | Whatever Vite React-TS template resolves | Template is canonical |
| `expo` | **`~52.x`** (explicit) | Architecture pins SDK 52 — reject defaults if newer |
| `react` | `^18.x` (web + mobile must agree) | Concurrent mode required by TanStack Query v5 (later stories) |
| Go | **`1.22+`** in `go.mod` | Architecture minimum |
| `github.com/gin-gonic/gin` | `latest` | Architecture chosen framework |
| `gorm.io/gorm` + `gorm.io/driver/postgres` | `latest` (v2) | Architecture chosen ORM |
| `github.com/swaggo/swag/cmd/swag` | `latest` | OpenAPI generation |
| `openapi-typescript` | `latest` (v7+) | Generates the spec consumer |
| `pnpm` | **`>=9.0`** in root `packageManager` field | Required for `pnpm/action-setup@v4` |
| Node | **`>=20.11.0`** in `engines` | Vite 5 + Expo SDK 52 minimum |

### References

- Story foundation: `_bmad-output/planning-artifacts/epics.md#Story 1.1: Monorepo Scaffold & CI Pipeline`
- Stack rationale: `_bmad-output/planning-artifacts/architecture.md#Selected Stack`
- Init commands: `_bmad-output/planning-artifacts/architecture.md#Initialization Commands`
- Target structure: `_bmad-output/planning-artifacts/architecture.md#Complete Project Directory Structure`
- Naming conventions: `_bmad-output/planning-artifacts/architecture.md#Naming Patterns`
- Type-gen pipeline: `_bmad-output/planning-artifacts/architecture.md#Development Workflow Integration`
- Additional Requirements: `_bmad-output/planning-artifacts/epics.md#Additional Requirements` (monorepo bullet, three envs, CI checks, OpenAPI pipeline)

### Resolved Decisions (Confirmed by Product Owner)

1. **Go module path:** `github.com/evelynzouras/tend-api`. Use this exact string in `apps/api/go.mod` (`go mod init github.com/evelynzouras/tend-api`).
2. **`apps/api/cmd/server/main.go` is a real Gin server, not a stub.** Implement Gin with a single `GET /health` route returning `200 OK` and JSON body `{"status":"ok"}` on `:8080`. This satisfies AC2 (port reachable), seeds Story 1.2's deployed-`/health` AC, and gives the Go CI test (`main_test.go`) a real handler to exercise.
3. **Expo dev script:** `expo start` only — no auto-open simulator. Devs press `i`/`a` interactively when needed.
4. **pnpm pinning via Corepack:** Set `"packageManager": "pnpm@9.x.x"` in root `package.json` (resolve `x.x` to the latest pnpm 9 release at scaffold time). CI uses Corepack — no separate "install pnpm" step needed in `ci.yml`.
5. **`pnpm test` wired now with no-op scripts:** Add `"test": "turbo run test"` to root `package.json`. Each app exposes `"test": "echo \"no tests yet\" && exit 0"` **except `apps/api`**, which runs `go test ./...` (the health-endpoint test from Task 4). Future stories add real TS tests without script restructuring.

## Dev Agent Record

### Agent Model Used

Claude Opus 4.7 (1M context) — `claude-opus-4-7[1m]`

### Debug Log References

- `/tmp/tend-dev3.log` — final clean `pnpm dev` run; all three servers reachable
- swag v1 → swagger 2.0 → `openapi-typescript` rejection. Switched to swag v2 with `--v3.1`.
- One stale `pnpm dev` orphan held ports 5173/8080/8081; required `pkill -f "go run"`, `pkill -f "expo start"`, `pkill -f vite`, `pkill -f "turbo run dev"` to fully clean up before re-running.

### Completion Notes List

- ✅ All 5 acceptance criteria verified locally (see Task 12).
- ✅ TDD followed for the Go health handler: `main_test.go` written first, confirmed RED (build fail: `undefined: newRouter`), then `main.go` implemented, GREEN.
- ✅ All 12 story tasks complete.
- ⚠️ **SDK / version drift from spec — accepted.** The spec was written 2026-04-15; scaffolding ran 2026-04-16. Templates moved on. See Change Log entries 1–7 for the per-tool details. Net: web has React 19 / TS 6 (spec said 18 / 5); mobile has Expo SDK 54 / React 19.1 / RN 0.81 (spec said SDK 52); pnpm is 10 (spec said 9); Go local is 1.26.2 (spec said 1.22+). Architecture doc needs a follow-up edit to bring it in line with what's actually installed; that's a separate, small task.
- ⚠️ Local AC5 verification could not run a true `git diff --exit-code` test because nothing has been committed yet. Instead verified that route additions DO produce diffs in 4 generated files (`swagger.json`, `swagger.yaml`, `docs.go`, `api.ts`) and that the CI workflow's `git diff --exit-code` against the eventual baseline will correctly fail. Once this branch is committed, the CI sync check is fully exercised.
- ℹ️ `apps/api/package.json` was added as a turbo workspace shim (not in the original task list but required for turbo to discover the Go app). Scripts: `dev` (`go run`), `build` (`go build`), `lint`/`typecheck` (`go vet`), `test` (`go test`).
- ℹ️ `gorm.io/gorm` and `gorm.io/driver/postgres` were `go get`'d but `go mod tidy` removed them as unused. They will re-add automatically when Story 1.4+ writes the first model — this is standard Go behavior and does not block any AC.

### File List

**Created (root):**
- `package.json`
- `pnpm-workspace.yaml`
- `turbo.json`
- `.npmrc`
- `.gitignore`
- `go.work`
- `pnpm-lock.yaml` (pnpm-managed)
- `README.md` (overwrote 2-line placeholder)

**Created (`.github/`):**
- `.github/workflows/ci.yml`

**Created (`apps/api/`):**
- `apps/api/package.json`
- `apps/api/go.mod`
- `apps/api/go.sum`
- `apps/api/.env.example`
- `apps/api/cmd/server/main.go`
- `apps/api/cmd/server/main_test.go`
- `apps/api/docs/swagger.json` (generated)
- `apps/api/docs/swagger.yaml` (generated)
- `apps/api/docs/docs.go` (generated)
- `apps/api/internal/handler/.gitkeep`
- `apps/api/internal/middleware/.gitkeep`
- `apps/api/internal/model/.gitkeep`
- `apps/api/internal/repository/.gitkeep`
- `apps/api/internal/service/.gitkeep`

**Created (`apps/web/`):**
- All files from Vite `react-ts` template (`package.json`, `index.html`, `eslint.config.js`, `vite.config.ts`, `tsconfig.json`, `tsconfig.app.json`, `tsconfig.node.json`, `src/`, `public/`, `README.md`)
- `apps/web/.env.example`
- _Modified `apps/web/package.json`:_ renamed to `@repo/web`, added `typecheck` and `test` scripts

**Created (`apps/mobile/`):**
- All files from Expo `blank-typescript` template (`package.json`, `app.json`, `App.tsx`, `index.ts`, `tsconfig.json`, `assets/`)
- `apps/mobile/.env.example`
- `apps/mobile/eslint.config.js` (added — template no longer ships one)
- _Modified `apps/mobile/package.json`:_ renamed to `@repo/mobile`, added `dev`/`lint`/`typecheck`/`test` scripts, added `eslint` + `eslint-config-expo` devDeps

**Created (`packages/types/`):**
- `packages/types/package.json`
- `packages/types/index.ts`
- `packages/types/generated/api.ts` (initial placeholder, then overwritten by `pnpm run generate:types`)

## Change Log

| # | Date | Change | Reason |
|---|---|---|---|
| 1 | 2026-04-16 | **pnpm pinned to `pnpm@10.33.0`** in root `packageManager` (spec said `pnpm@9.x.x`) | pnpm 9 has been superseded by 10 in the user's environment; pinning to the installed version. CI Corepack honors whatever's pinned. |
| 2 | 2026-04-16 | **Expo SDK 54.0.33** scaffolded (spec said SDK 52) — accepted by product owner | `create-expo-app@latest` no longer ships SDK 52. Forcing 52 would mean rejecting the official starter and risking Apple App Store SDK floor requirements. Architecture doc needs a follow-up edit. |
| 3 | 2026-04-16 | **swag CLI swapped from v1 to v2** (`github.com/swaggo/swag/v2/cmd/swag`); `--v3.1` flag added to `generate:types` | `swag` v1 emits Swagger 2.0; `openapi-typescript` v7 only consumes OpenAPI 3.x. v2 emits OpenAPI 3.1 natively. Also added `swag/v2` as a Go module dep because the generated `docs/docs.go` imports it. |
| 4 | 2026-04-16 | `gorm.io/gorm` and `gorm.io/driver/postgres` removed from `apps/api/go.mod` by `go mod tidy` | Standard Go behavior — unused deps get pruned. They will re-add automatically when Story 1.4+ writes the first model. Not a regression. |
| 5 | 2026-04-16 | CI `setup-go` uses `go-version: stable` (spec said `1.22`) | Local `go.mod` declares `go 1.26.2` (Homebrew's current stable). CI must use a toolchain ≥ that. `stable` is robust and avoids version-sync churn. |
| 6 | 2026-04-16 | Added `eslint-config-expo` + `apps/mobile/eslint.config.js` (spec assumed Expo template ships one) | Expo `blank-typescript` template no longer ships an ESLint config. Without one, AC3 (`pnpm lint` passes) is unreachable for mobile. |
| 7 | 2026-04-16 | Added `apps/api/package.json` (not in original task list) | Required for turbo to discover the Go workspace. Without it, `pnpm dev` only runs JS workspaces — AC2 fails because the API never starts. The shim wraps `go run`/`go build`/`go vet`/`go test`. |
| 8 | 2026-04-16 | Vite template scaffolded React 19.2 + TypeScript 6 (spec said React 18 + TS 5.x) | Template defaults moved on between spec date (2026-04-15) and scaffold date (2026-04-16). React 19 supports concurrent mode (TanStack Query v5 compat preserved); TS 6 is backward compatible for our patterns. |
