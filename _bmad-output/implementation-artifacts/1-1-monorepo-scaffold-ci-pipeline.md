# Story 1.1: Monorepo Scaffold & CI Pipeline

Status: ready-for-dev

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

- [ ] **Task 1 — Initialize Turborepo + pnpm workspace in-place at repo root** (AC: 1)
  - [ ] Run `npx create-turbo@latest tend --package-manager pnpm` in a temp directory, then move generated files (`turbo.json`, `pnpm-workspace.yaml`, root `package.json`, `.gitignore` additions) into the existing repo root. **Do NOT create a nested `tend/` directory** — the repo at `/Users/evelynzouras/Documents/code/toDoFastTrack` IS the monorepo root.
  - [ ] Delete any sample `apps/` and `packages/` content the starter ships with (we'll scaffold our own).
  - [ ] Confirm `pnpm-workspace.yaml` declares `apps/*` and `packages/*` as workspace globs.
  - [ ] Pin Node engine in root `package.json` (`"engines": { "node": ">=20.11.0" }`) — required by Vite 5+ and Expo SDK 52.
- [ ] **Task 2 — Scaffold `apps/web` (Vite + React + TypeScript)** (AC: 1, 2)
  - [ ] Run `npm create vite@latest apps/web -- --template react-ts` (per architecture init commands).
  - [ ] Add `dev`, `build`, `lint`, `typecheck` scripts to `apps/web/package.json` (`tsc --noEmit` for typecheck).
  - [ ] Confirm Vite dev server starts on port `5173` (default — no override needed).
- [ ] **Task 3 — Scaffold `apps/mobile` (Expo SDK 52 + React Native + TypeScript)** (AC: 1, 2)
  - [ ] Run `npx create-expo-app apps/mobile --template blank-typescript`.
  - [ ] Confirm Expo SDK version is **52.x** (architecture pins this — do not accept 51 or 53 if defaults change).
  - [ ] Add `dev`, `lint`, `typecheck` scripts to `apps/mobile/package.json`. `dev` should run `expo start` (defaults to port `8081`).
- [ ] **Task 4 — Scaffold `apps/api` (Go + Gin + GORM)** (AC: 1, 2)
  - [ ] `mkdir -p apps/api/cmd/server apps/api/internal/handler apps/api/internal/middleware apps/api/internal/model apps/api/internal/repository apps/api/internal/service apps/api/docs`
  - [ ] `cd apps/api && go mod init <module-path>` — see **Open Question Q1** for module path resolution.
  - [ ] `go get github.com/gin-gonic/gin gorm.io/gorm gorm.io/driver/postgres`
  - [ ] Create `apps/api/cmd/server/main.go` with a minimal Gin server: register `GET /health` returning `200 OK` with `{"status":"ok"}`. Listen on `:8080`. This is the smoke test for AC2 and seeds Story 1.2 (`/health` is referenced in 1.2's ACs).
  - [ ] Verify `go run ./cmd/server` starts cleanly from `apps/api/`.
- [ ] **Task 5 — Initialize Go workspace at repo root** (AC: 1)
  - [ ] Run `go work init ./apps/api` at repo root → produces `go.work` and `go.work.sum`.
  - [ ] Commit both files.
- [ ] **Task 6 — Create `packages/types` skeleton** (AC: 1, 4)
  - [ ] `mkdir -p packages/types/generated`
  - [ ] Create `packages/types/package.json`: name `@repo/types`, `"main": "index.ts"`, no build step (consumed via TypeScript path resolution from workspace).
  - [ ] Create `packages/types/index.ts`: `export * from './generated/api';`
  - [ ] Create placeholder `packages/types/generated/api.ts`: `// AUTO-GENERATED — DO NOT EDIT BY HAND. Run `pnpm run generate:types`.` followed by `export {};`
  - [ ] Add `packages/types/generated/` is **NOT** gitignored — it must be committed so CI can detect drift (AC5).
- [ ] **Task 7 — Configure `turbo.json` pipeline** (AC: 2, 3)
  - [ ] Define `dev` task with `"cache": false, "persistent": true` so all three apps run concurrently and stream logs.
  - [ ] Define `build`, `lint`, `typecheck`, `test` tasks with appropriate dependencies (`build` depends on `^build`, etc.).
  - [ ] Define `generate:types` task — `cache: false`, runs the swag → openapi-typescript chain (see Task 8).
  - [ ] Add root `package.json` scripts: `"dev": "turbo run dev"`, `"build": "turbo run build"`, `"lint": "turbo run lint"`, `"typecheck": "turbo run typecheck"`, `"test": "turbo run test"`.
- [ ] **Task 8 — Wire OpenAPI generation pipeline** (AC: 4)
  - [ ] In `apps/api`: `go install github.com/swaggo/swag/cmd/swag@latest`
  - [ ] Add minimal swag annotation to `main.go` so `swag init -g cmd/server/main.go -o ./docs` produces a valid `apps/api/docs/swagger.json` even with only the `/health` route.
  - [ ] At repo root: `pnpm add -D -w openapi-typescript` (workspace-level dev dep).
  - [ ] Add root script `"generate:types": "cd apps/api && swag init -g cmd/server/main.go -o ./docs && cd ../.. && openapi-typescript apps/api/docs/swagger.json -o packages/types/generated/api.ts"`.
  - [ ] Run it once and commit the generated `swagger.json` and `api.ts` so the initial state is the source-of-truth baseline for AC5.
- [ ] **Task 9 — Add GitHub Actions CI workflow** (AC: 3, 5)
  - [ ] Create `.github/workflows/ci.yml` triggered on `pull_request` and `push: branches: [main]`.
  - [ ] Jobs (run in parallel where independent):
    - **`go-checks`**: setup-go (1.22+), `cd apps/api && go vet ./... && go test ./...`
    - **`ts-checks`**: setup-node (20.x), setup-pnpm, `pnpm install --frozen-lockfile`, `pnpm typecheck`, `pnpm lint`
    - **`openapi-sync`**: setup-go + setup-node + setup-pnpm, install swag CLI, `pnpm install --frozen-lockfile`, run `pnpm run generate:types`, then `git diff --exit-code packages/types/generated/api.ts apps/api/docs/swagger.json` — if diff is non-empty, fail with message: `"❌ packages/types is out of sync with Go OpenAPI spec. Run 'pnpm run generate:types' locally and commit the changes."`
  - [ ] Pin action versions (`actions/checkout@v4`, `actions/setup-node@v4`, `actions/setup-go@v5`, `pnpm/action-setup@v4`).
- [ ] **Task 10 — ESLint setup** (AC: 3)
  - [ ] Use the ESLint configs that ship with the Vite React TS template (`apps/web`) and Expo template (`apps/mobile`) — both produce a working `eslint.config.js` (flat config) out of the box.
  - [ ] Confirm `pnpm --filter @repo/web lint` and `pnpm --filter @repo/mobile lint` both pass on the empty starter code.
  - [ ] Do NOT introduce a custom shared ESLint package at this story — defer until duplication is felt.
- [ ] **Task 11 — Repo hygiene + README** (AC: all)
  - [ ] Update root `README.md` with: setup (`pnpm install`), local dev (`pnpm dev`), type generation (`pnpm run generate:types`), CI overview.
  - [ ] Confirm `.gitignore` covers `node_modules/`, `dist/`, `.expo/`, `apps/api/tend-api` (compiled binary), `.env`, `.env.local`, `*.log`.
  - [ ] Add `apps/api/.env.example`, `apps/web/.env.example`, `apps/mobile/.env.example` as empty placeholders (Story 1.2 will populate them).
- [ ] **Task 12 — Local verification of all five ACs** (AC: 1, 2, 3, 4, 5)
  - [ ] Fresh clone simulation: `rm -rf node_modules apps/*/node_modules && pnpm install` → exits 0 (AC1).
  - [ ] `pnpm dev` → all three dev servers reachable on their ports within 30 seconds (AC2). Curl `localhost:8080/health` returns 200.
  - [ ] `pnpm typecheck && pnpm lint && (cd apps/api && go vet ./... && go test ./...)` all pass locally (AC3).
  - [ ] `pnpm run generate:types` produces the expected `api.ts` deterministically (AC4).
  - [ ] Manually edit `apps/api/cmd/server/main.go` to add a no-op route, run the CI sync check command from Task 9 — verify it fails. Revert the edit. (AC5)

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

_To be filled by dev-story agent (model + version)._

### Debug Log References

### Completion Notes List

### File List
