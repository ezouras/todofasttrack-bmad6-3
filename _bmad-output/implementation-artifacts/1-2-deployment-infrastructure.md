# Story 1.2: Deployment Infrastructure

Status: in-progress

<!-- Note: Validation is optional. Run validate-create-story for quality check before dev-story. -->

## Story

As a **developer**,
I want the web app deployed to Vercel and the API + PostgreSQL deployed to Railway with development, staging, and production environments,
so that changes can be tested in deployed environments before reaching production.

## Acceptance Criteria

1. **Given** code is merged to `main`, **When** the CI deploy job runs, **Then** the web app auto-deploys to the Vercel staging URL and the API auto-deploys to Railway staging.
2. **Given** a PR is opened, **When** Vercel processes it, **Then** a unique preview URL is generated for the web app.
3. **Given** the Railway project is configured with staging and production environments, **When** the Go API starts in either environment, **Then** `GET /health` returns `200 OK` at the deployed URL.
4. **Given** Railway staging and production are provisioned, **When** the API connects, **Then** it can read from and write to the managed PostgreSQL add-on in both environments.
5. **Given** secrets (Clerk keys, Stripe keys, DB URL), **When** configured in Railway and Vercel environment dashboards, **Then** they are available to the apps at runtime and are not present in the git repository.

## Tasks / Subtasks

- [ ] **Task 1 — Containerize the Go API** (AC: 1, 3)
  - [ ] Create `apps/api/Dockerfile` using a multi-stage build:
    - Stage 1 (`builder`): `golang:1.26-alpine`, `WORKDIR /src`, copy `go.mod`/`go.sum`, `go mod download`, copy source, `CGO_ENABLED=0 go build -ldflags="-s -w" -o /out/tend-api ./cmd/server`.
    - Stage 2 (runtime): `gcr.io/distroless/static-debian12` (or `alpine:3.20` if you need a shell — distroless is preferred for size + attack surface), `COPY --from=builder /out/tend-api /tend-api`, `EXPOSE 8080`, `ENTRYPOINT ["/tend-api"]`.
  - [ ] Add `apps/api/.dockerignore` to keep the build context minimal: `docs/` (regenerated, not needed in image), `*_test.go`, `.env*`, `tend-api` (local binary), `internal/*/.gitkeep`, `.git/`.
  - [ ] Verify locally: `docker build -t tend-api:local apps/api && docker run --rm -p 8080:8080 tend-api:local` → `curl localhost:8080/health` returns `{"status":"ok"}`. Image size goal: ≤25 MB (distroless static).
- [ ] **Task 2 — Provision Railway project + environments** (AC: 1, 3, 4)
  - [ ] **Manual step (you, in Railway dashboard):** Create a new Railway project named `tend`. Add the GitHub repo. Create three environments: `development`, `staging`, `production`. See **Open Question Q1** if a project already exists.
  - [ ] Add a **PostgreSQL add-on** to each of the staging and production environments (development can either use a separate add-on or developers can run postgres locally — see Q4).
  - [ ] In the Railway dashboard, configure the Go service:
    - **Source:** GitHub repo, root directory `apps/api`, Dockerfile-based build (Railway auto-detects `apps/api/Dockerfile`).
    - **Branch trigger:** `staging` env deploys on push to `main`; `production` env deploys on tag matching `v*` (or via Railway's "deploy from commit" UI when you cut a release — see Q3).
  - [ ] Configure environment variables in **each** Railway environment:
    - `DATABASE_URL` — Railway auto-provides this when a Postgres add-on is attached (variable reference `${{Postgres.DATABASE_URL}}`).
    - `PORT` — Railway sets this automatically; the existing `apps/api/cmd/server/main.go` already reads `os.Getenv("PORT")` and falls back to `:8080`. No code change needed.
    - `GIN_MODE=release` — silences Gin's debug logging in deployed envs (per the deferred review finding from Story 1.1).
- [ ] **Task 3 — Provision Vercel project** (AC: 1, 2)
  - [ ] **Manual step (you, in Vercel dashboard):** Create a new Vercel project from the GitHub repo. Set:
    - **Root directory:** `apps/web`
    - **Framework preset:** Vite
    - **Build command:** (Vercel auto-detects from preset — `pnpm build`; verify it picks up workspace context correctly)
    - **Install command:** `pnpm install --frozen-lockfile` (run from repo root via Vercel's "include files outside root" setting — see Q5 for monorepo-install nuance)
    - **Output directory:** `dist` (Vite default)
  - [ ] Confirm Vercel's default behavior gives you preview URLs on PRs and a stable production URL on `main` merges (this is Vercel's default — no extra config required for AC2).
  - [ ] Add `apps/web/vercel.json` only if necessary to override defaults (e.g., SPA history-mode rewrites: `{ "rewrites": [{ "source": "/(.*)", "destination": "/index.html" }] }` so React Router (later stories) doesn't 404 on deep links). **Recommended to add now** — it costs nothing today and prevents a confusing bug when routes land in Story 1.3 / 3.1.
- [ ] **Task 4 — Wire env-var contracts at the application layer** (AC: 5)
  - [ ] In `apps/web/src/`, add a tiny `lib/env.ts` that reads `import.meta.env.VITE_API_BASE_URL` (with a clear runtime error if missing in production builds). Vite exposes only `VITE_*`-prefixed envs to the bundle. **Do not introduce zod yet** — the architecture's "`packages/config` with zod schemas" is an Epic-level pattern; for one variable, a 5-line type-safe getter is enough. Defer the zod package to when the env surface grows past ~3 vars (see deferred-work.md for tracking).
  - [ ] Update `apps/web/.env.example` to add `VITE_API_BASE_URL=http://localhost:8080` as the local-dev default (commented out — devs uncomment when needed).
  - [ ] Update `apps/api/.env.example` to add `# DATABASE_URL=postgres://user:password@localhost:5432/tend?sslmode=disable` (commented placeholder — Story 1.4 wires the actual GORM connection).
  - [ ] Update `apps/mobile/.env.example` with `EXPO_PUBLIC_API_BASE_URL=http://localhost:8080` (commented).
  - [ ] **Do NOT add Clerk or Stripe env vars yet** — same reasoning as Story 1.1's review finding (P6); those land with the stories that consume them.
- [ ] **Task 5 — Deploy workflow (GitHub Actions for staging auto-deploy on `main`)** (AC: 1)
  - [ ] **Decision per architecture:** Vercel and Railway both auto-deploy from GitHub natively when their projects are wired up correctly (Tasks 2 and 3). **No GitHub Actions workflow is strictly required** for staging deploys — the cloud providers handle it.
  - [ ] **Therefore:** instead of a `deploy.yml`, add a small **`.github/workflows/deploy-status.yml`** that posts a comment on each PR with the Vercel preview URL (Vercel does this natively too via their bot, so this is optional) — **defer** this nice-to-have unless you want to override Vercel's default bot.
  - [ ] **Required:** add a `deployments` permissions check job in CI to confirm both providers are wired (described in Task 7).
- [ ] **Task 6 — Production deploy trigger** (AC: 1)
  - [ ] **Per architecture spec:** "On release tag: deploy to production (manual trigger)". The recommended pattern:
    - **Vercel:** in dashboard, set the "Production Branch" to a tag pattern like `v*` OR keep `main` as production and treat Vercel preview deployments as "staging" (simpler — see Q3).
    - **Railway:** in dashboard, set the production environment to deploy from a tag pattern OR trigger manually via Railway UI / `railway up` CLI.
  - [ ] **Most pragmatic for this story:** treat Vercel's "production" as Tend's staging (auto from `main`); use Railway's UI "Promote to production" button when you cut a release. Document the manual promotion step in `README.md`. Don't build a custom GitHub Actions release workflow until the team grows past one developer.
- [ ] **Task 7 — Health check verification + smoke test** (AC: 3, 4)
  - [ ] After deploying staging via Tasks 2–6, manually verify:
    - `curl https://<railway-staging-url>/health` returns `{"status":"ok"}` (AC3).
    - Railway's logs show the Go server's `slog` JSON output on startup (`{"time":"...","level":"INFO","msg":"starting tend-api","addr":":8080"}`).
    - Vercel staging URL loads the Vite SPA without errors (AC1, partial AC2).
    - Open a draft PR with a trivial change to `apps/web/`; confirm Vercel posts a preview URL (full AC2).
  - [ ] AC4 verification (DB read/write) is partially deferred to Story 1.4 (when GORM models exist). For this story, prove the DB connection is **reachable** by adding a minimal `/health` enhancement OR a one-shot test:
    - **Option A (recommended):** add a sibling `/health/db` endpoint that opens a DB connection from `DATABASE_URL`, runs `SELECT 1`, returns `{"status":"ok","db":"reachable"}` — small Gin handler + minimal database/sql import. Strips later when GORM lands.
    - **Option B:** SSH into the Railway service and run `psql $DATABASE_URL -c 'SELECT 1'` manually. Cheaper but doesn't survive as a regression check.
    - See **Open Question Q6** for which to pick.
  - [ ] Document URLs in `README.md` (replace `localhost` with the staging URLs in the "deployed environments" subsection).
- [ ] **Task 8 — README + deferred-work updates** (AC: all)
  - [ ] Update root `README.md`:
    - New "Deployments" section: list all 3 environments with URLs, who triggers each, and where to watch logs.
    - Add a "Production release" subsection with the manual Railway promotion steps + `git tag v0.x.x && git push --tags` recipe.
    - Note that secrets live in Railway/Vercel dashboards — never `.env` in the repo.
  - [ ] Add the API base URL note to the local-dev section: `pnpm dev` runs on `localhost:8080`; staging swaps to `https://<railway-staging>.up.railway.app`.
  - [ ] Move any newly identified deferred items to `_bmad-output/implementation-artifacts/deferred-work.md` (e.g., custom domains, the Vercel preview-URL bot, zod env schema, automatic production-deploy-on-tag).
- [ ] **Task 9 — Local verification of all five ACs** (AC: 1, 2, 3, 4, 5)
  - [ ] **AC1 (auto-deploy on `main`):** push a trivial commit to `main`; observe both Vercel and Railway deployments triggered automatically. Both succeed.
  - [ ] **AC2 (Vercel preview URL on PR):** open a draft PR; observe Vercel bot comment with preview URL within ~2 minutes. URL loads the SPA.
  - [ ] **AC3 (`/health` 200 in deployed envs):** `curl` against both staging and production URLs (after promoting); both return `{"status":"ok"}`.
  - [ ] **AC4 (DB read/write):** per Q6 resolution, either `/health/db` returns `db:"reachable"` from both staging and production, or a manual `psql SELECT 1` succeeds.
  - [ ] **AC5 (secrets isolated):** `git grep -i clerk\|stripe\|database_url`-style audit returns only `.env.example` template files and docs — no real secret values committed. Railway/Vercel dashboards show all required secrets configured.

## Dev Notes

### Ground Truth: Repo State at Story Start (Post-1.1)

[Source: Story 1.1 File List + Change Log entries]

After Story 1.1 ships:
- **Monorepo scaffold complete:** Turborepo + pnpm workspaces, with `apps/api` (Go + Gin), `apps/web` (Vite + React 19.2 + TS 6), `apps/mobile` (Expo SDK 54 + RN 0.81), `packages/types` (auto-generated from OpenAPI).
- **API has exactly one route: `GET /health` → `200 {"status":"ok"}`** in `apps/api/cmd/server/main.go`. Listens on `:8080`, honors `PORT` env override. Uses `slog` JSON logging — Railway will index this natively.
- **CI workflow exists** at `.github/workflows/ci.yml` (3 parallel jobs: go-checks, ts-checks, openapi-sync). Hardened with concurrency group, permissions block, pinned action versions.
- **No Dockerfile yet.** No `vercel.json` yet. No `deploy.yml` yet. No DB connection code anywhere. No Clerk, no Stripe, no auth middleware.
- **Three `.env.example` files exist** (api/web/mobile) — all currently bare placeholders post-Story 1.1 review (Clerk/Stripe keys were stripped).
- **`packages/config` with zod schemas** (architecture-mentioned) does **NOT exist yet** — Task 4 explicitly defers it. See Open Question Q5.

### Critical Architecture Compliance

[Source: `_bmad-output/planning-artifacts/architecture.md#Infrastructure & Deployment`]

**Locked deployment topology — do not substitute:**

| Layer | Provider | Notes |
|---|---|---|
| Web | **Vercel** | Auto-detects Vite, ships preview URLs on PRs natively |
| API | **Railway** | Docker-based Go deployment; auto-deploys from `main` |
| Database | **Railway PostgreSQL add-on** | Same project as API; `DATABASE_URL` injected via `${{Postgres.DATABASE_URL}}` reference |
| Mobile (deferred) | Expo EAS Build | NOT this story — Story 8.5 |

**Three environments required:** development → staging → production. Architecture says "auto-deploy to staging on merge to main; manual tag deploy to production" (line 99 of `epics.md` Additional Requirements).

**Secrets contract:** "Secrets managed in Railway (API) and Vercel (web) dashboards — never committed" (architecture line 246).

### Out of Scope for This Story (Will Be Future Stories)

**Do NOT implement any of these in Story 1.2** — they belong elsewhere:

- **Database schema, GORM models, migrations** → Story 1.4+ when models land. Task 7's optional `/health/db` only proves connectivity, not schema.
- **Clerk, Stripe, any other 3rd-party API integration** → respective stories (1.4, 6.x).
- **`packages/config` with zod env schema** → defer until env surface is bigger than ~3 vars; see deferred-work.md.
- **Custom domain (e.g., `api.tend.app`)** → defer; default Railway/Vercel subdomains are fine for staging+production until launch.
- **Mobile app distribution (EAS Build, App Store)** → Story 8.5.
- **Production deploy automation via GitHub Actions** → defer; manual Railway "Promote" UI is fine for solo dev (see Resolved Decision #3).
- **Monitoring / APM / error tracking (Sentry, Datadog)** → architecture explicitly defers ("add after first paying users", line 142).

### Anti-Patterns to Avoid

- ❌ **Committing real secret values.** Even in `.env.example`, never paste real Clerk/Stripe/DB credentials. The audit step in Task 9 catches this.
- ❌ **Hardcoding Railway/Vercel URLs in code.** Use env vars (`VITE_API_BASE_URL` on web, `EXPO_PUBLIC_API_BASE_URL` on mobile). The deployed values live only in Vercel/Expo dashboards, never in source.
- ❌ **Skipping the `apps/api/.dockerignore`.** Without it, your `node_modules`, `.git`, and any local binaries get copied into the build context — slows builds 10x and can leak secrets.
- ❌ **Adding a `deploy.yml` GitHub Actions workflow that re-implements what Vercel/Railway already do.** Both auto-deploy from GitHub; a custom workflow adds maintenance burden and another secret-leak surface.
- ❌ **Pinning the Docker base image to `:latest`.** Pin major versions (`golang:1.26-alpine`, `gcr.io/distroless/static-debian12`) so future Go releases don't silently break builds.
- ❌ **Building the Docker image without `CGO_ENABLED=0`.** Distroless static images don't have glibc; CGO-built binaries will segfault on startup.
- ❌ **Adding zod / a `packages/config` package early.** The architecture mentions it as an Epic-level pattern; for Story 1.2's 1–2 env vars, a 5-line plain-TS getter beats pulling in a dep + new workspace.

### Testing Standards

[Source: `architecture.md#Test File Location`; Story 1.1 Testing Standards]

- **Go:** Co-located `{file}_test.go`. If you add `/health/db` (Task 7 Option A), write `health_db_test.go` that uses `httptest` + an in-memory or short-circuit DB. Don't require a live Postgres for unit tests — that's an integration concern.
- **TypeScript:** No test runner wired yet (still deferred per Story 1.1). The new `apps/web/src/lib/env.ts` is small enough to skip a test for; if you write one, use `import.meta.env` mocking via Vite's `vi.stubEnv`.
- **CI:** existing `.github/workflows/ci.yml` already runs `go vet` + `go test` + typecheck + lint + OpenAPI sync. Story 1.2 should not add new jobs — let Vercel and Railway own deploy verification.

### Project Structure Notes

After this story, the repo gains:
```
apps/api/
├── Dockerfile                    # Task 1 (multi-stage, distroless)
├── .dockerignore                 # Task 1
└── cmd/server/main.go            # (optional) +/health/db handler — Task 7 Option A
apps/web/
├── src/lib/env.ts                # Task 4 (typed env getter)
└── vercel.json                   # Task 3 (SPA rewrites)
README.md                         # Task 8 (Deployments section)
```

No new packages. No `packages/config`. No new GitHub Actions workflows. Architecture compliance: this matches the Complete Project Directory Structure (architecture.md line ~444) — `apps/api/Dockerfile` and `apps/api/.env.example` are explicitly listed there.

### Library Versions to Use (as of 2026-04-16)

| Package / Image | Version | Why |
|---|---|---|
| `golang:1.26-alpine` | exact | Matches local Go install (Story 1.1 Change Log entry 5) |
| `gcr.io/distroless/static-debian12` | exact | Stable runtime base; ~2 MB; no shell = smaller attack surface |
| `database/sql` (stdlib) | stdlib | For Task 7 Option A `/health/db`; do NOT pull GORM yet |
| `github.com/lib/pq` | latest | Plain Postgres driver for `database/sql` (Task 7 Option A only). Or `github.com/jackc/pgx/v5/stdlib` if you prefer pgx. **Defer GORM** (Story 1.4 owns it). |

### References

- Story foundation: `_bmad-output/planning-artifacts/epics.md#Story 1.2: Deployment Infrastructure`
- Architecture deployment topology: `_bmad-output/planning-artifacts/architecture.md#Infrastructure & Deployment`
- Architecture target structure: `_bmad-output/planning-artifacts/architecture.md#Complete Project Directory Structure` (lines around 460 — apps/api/Dockerfile listed there)
- Additional Requirements: `_bmad-output/planning-artifacts/epics.md#Additional Requirements` (3 envs required, GitHub Actions CI/CD, secret management)
- Previous story: `_bmad-output/implementation-artifacts/1-1-monorepo-scaffold-ci-pipeline.md` — read the **Change Log** + **Review Findings** sections for inherited deviations (SDK 54, Go 1.26, pnpm 10) and deferred work that may resurface here (e.g., the Gin `ReleaseMode` deferred item is now actionable via `GIN_MODE=release` env in Task 2).
- Deferred work pulled forward into 1.2: `_bmad-output/implementation-artifacts/deferred-work.md` — specifically the Gin `ReleaseMode` item.

### Resolved Decisions (Confirmed by Product Owner, 2026-04-16)

1. **Q1+Q2 — Neither Railway nor Vercel project exists yet.** Code work proceeds first (Dockerfile, vercel.json, env getter, `/health/db`). After the code lands, dev HALTS and provides the user with a dashboard checklist for both providers. User pastes URLs back; dev wires the README and runs AC verification.
2. **Q3 — Production deploy trigger: `main → prod` with manual Railway promote.** Vercel auto-deploys main as production. PR previews serve as de-facto staging URLs for web. Railway has its own "staging" environment that auto-deploys from main; production is promoted manually via Railway UI when ready to release. **No GitHub Actions release workflow** is built — that complexity is reserved for when team grows beyond solo.
3. **Q4 — Local Postgres via Docker for dev.** Devs run `docker run -d --name tend-pg -p 5432:5432 -e POSTGRES_PASSWORD=tend -e POSTGRES_USER=tend -e POSTGRES_DB=tend postgres:16`. Railway only provisions Postgres add-ons for **staging + production** (NOT a third dev env). Architecture's "three environments" requirement is satisfied as: local dev → Railway staging → Railway production.
4. **Q5 — Vercel uses dashboard "Root Directory: apps/web" + "Include files outside of the Root Directory" toggle.** No root-level `vercel.json` orchestration. Only `apps/web/vercel.json` is added, and it contains exactly one thing: SPA history-mode rewrites (`{ "rewrites": [{ "source": "/(.*)", "destination": "/index.html" }] }`) to keep React Router (Story 1.3+) from 404'ing on deep links.
5. **Q6 — Add `/health/db` endpoint** using stdlib `database/sql` + `github.com/lib/pq`. ~30 lines in `apps/api/cmd/server/main.go` (or a new sibling file). Reads `DATABASE_URL` env, opens a connection, runs `SELECT 1`, returns `{"status":"ok","db":"reachable"}` on success or `503` on DB error. Stays useful after GORM lands (it's a connectivity probe, separate from ORM concerns) — Story 1.4 doesn't need to remove it.

## Dev Agent Record

### Agent Model Used

_To be filled by dev-story agent (model + version)._

### Debug Log References

### Completion Notes List

### File List

## Change Log

| # | Date | Change | Reason |
|---|---|---|---|
| _to be populated by dev-story_ | | | |
