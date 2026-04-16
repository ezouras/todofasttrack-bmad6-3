---
stepsCompleted:
  - step-01-init
  - step-02-context
  - step-03-starter
  - step-04-decisions
  - step-05-patterns
  - step-06-structure
  - step-07-validation
  - step-08-complete
inputDocuments:
  - _bmad-output/planning-artifacts/prd.md
workflowType: 'architecture'
lastStep: 8
status: 'complete'
completedAt: '2026-04-15'
project_name: 'toDoFastTrack'
user_name: 'Evelynzouras'
date: '2026-04-15'
---

# Architecture Decision Document

_This document builds collaboratively through step-by-step discovery. Sections are appended as we work through each architectural decision together._

## Project Context Analysis

### Requirements Overview

**Functional Requirements:**
49 FRs across 9 capability areas: User Account Management, Onboarding & Goal Setup, Daily Planning, Capacity Management, Feedback & Reinforcement, Notifications, Cross-Platform & Sync, Subscription & Billing, and Landing Page & Discoverability. The capacity management and cross-platform sync areas carry the highest architectural weight — they touch every other capability area and have strict performance requirements.

**Non-Functional Requirements:**
- *Performance:* Web load <3s, mobile launch <2s, UI actions <300ms, cross-platform sync <5s, capacity calculation <1s
- *Security:* AES-256 at rest, TLS 1.2+ in transit, bcrypt passwords, JWT + refresh rotation, per-user data isolation, Stripe delegation for payment data, 30-day deletion SLA
- *Scalability:* 10,000 users without architectural changes, per-user access patterns, stateless API
- *Accessibility:* WCAG 2.1 AA (web), iOS HIG + Android Material (mobile)
- *Integration:* Stripe (idempotent webhooks), Google/Apple OAuth (server-side validation), APNs/FCM (non-blocking delivery), offline sync (conflict resolution protocol)

**Scale & Complexity:**
- Primary domain: Full-stack (Web + Mobile + API + Database)
- Complexity level: Medium — real-time sync, offline-first mobile, multi-provider auth, and adaptive capacity model add meaningful scope beyond a standard CRUD app
- Estimated architectural components: ~8 discrete service boundaries

### Technical Constraints & Dependencies

- **Capacity model is server-side only** — must be consistent across devices and survive app reinstalls; client devices cannot own this calculation
- **Offline-first on mobile** — today's todos and goals must be locally accessible without network; sync on reconnect using last-write-wins for todo status, server-authoritative for capacity data
- **Web session vs mobile token** — landing page uses cookie-based sessions with auto-login detection; mobile app uses token-based auth stored securely on device
- **Monorepo** — shared types, API client, and utilities between React (web) and React Native (mobile) to reduce duplication
- **Web-first launch** — mobile public release follows 4-6 weeks after web; architecture must support this phased rollout without requiring parallel infrastructure
- **Subscription gating** — all user-facing features require active trial or paid subscription; subscription state must be reliably propagated and checked

### Cross-Cutting Concerns Identified

- **Authentication & Authorization:** Spans all three layers (web, mobile, API); multi-provider (email, Google, Apple); different session strategies per platform
- **Subscription state:** Must be checked consistently across web and mobile; Stripe webhook handling must be idempotent to prevent state drift
- **Per-user data isolation:** Enforced at the database query level — no cross-user data access permitted at any layer
- **Offline sync conflict resolution:** Last-write-wins for mutable user data (todo status, edits); server-authoritative for computed data (capacity model, streaks)
- **Push notification delivery:** Non-blocking — APNs/FCM failures must not affect core app function; delivery status logged for monitoring
- **Wellness language constraints:** Application layer must enforce observational (not prescriptive) copy in all capacity and goal-alignment messaging

## Starter Template Evaluation

### Primary Technology Domain

Full-stack monorepo: React SPA (web) + Expo (mobile) + Go REST API + PostgreSQL, with Turborepo as the monorepo orchestrator.

### Selected Stack

| Layer | Technology | Rationale |
|---|---|---|
| Monorepo | Turborepo + pnpm | Official Expo recommendation, fast caching, supports Go + TS |
| Web | Vite + React + TypeScript | SPA requirement, Vercel-native, fastest dev experience |
| Mobile | Expo SDK 52 + React Native + TypeScript | Simplifies push notifications, EAS Build, store submission |
| API | Go + Gin | Most popular Go framework, best learning resources for frontend-focused dev |
| ORM | GORM | Most approachable for solo dev, handles migrations, fast prototyping |
| Database | PostgreSQL | Per PRD requirement |
| Deployment (web) | Vercel | Per user preference |
| Deployment (API) | Railway | Native Go support, managed PostgreSQL add-on, simplest solo-dev deployment |
| Deployment (mobile) | Expo EAS Build | App Store + Play Store submission |

**Type sharing across Go ↔ TypeScript:** Go API generates an OpenAPI spec (via `swaggo/swag`). TypeScript types auto-generated from spec using `openapi-typescript` into `packages/types`. This keeps types in sync without manual duplication.

### Initialization Commands

```bash
# 1. Bootstrap monorepo
npx create-turbo@latest todoFastTrack --package-manager pnpm

# 2. Web app
npm create vite@latest apps/web -- --template react-ts

# 3. Mobile app
npx create-expo-app apps/mobile

# 4. Go API
mkdir apps/api && cd apps/api
go mod init github.com/<username>/todofasttrack-api
go get github.com/gin-gonic/gin
go get gorm.io/gorm
go get gorm.io/driver/postgres

# 5. Shared types package
mkdir packages/types
```

**Note:** Epic 1, Story 1 must be monorepo initialization using the above commands.

### Architectural Decisions Provided by Starter

**Language & Runtime:** TypeScript (strict mode) for web + mobile; Go 1.22+ for API

**Build Tooling:** Vite (web), Expo Metro (mobile), Go build toolchain (API); Turborepo orchestrates all three

**Code Organization:** Monorepo with `apps/` (web, mobile, api) and `packages/` (types, potentially shared UI components)

**Development Experience:** Hot reload on all three layers simultaneously via Turborepo; pnpm workspace for JS dependency management; `go.work` for Go module management within monorepo

## Core Architectural Decisions

### Decision Priority Analysis

**Critical Decisions (Block Implementation):**
- Authentication provider: Clerk
- State management: TanStack Query + Zustand
- Styling system: Tailwind CSS + NativeWind v4
- Real-time sync mechanism: Server-Sent Events (SSE)
- Routing: React Router v7 (web) + Expo Router (mobile)

**Important Decisions (Shape Architecture):**
- API versioning: `/api/v1/` prefix
- Error handling: standardized Go error responses + TypeScript error types
- Database migrations: GORM AutoMigrate (dev) + GORM migrate (prod)
- Logging: Go `slog` (structured logging)
- CI/CD: GitHub Actions

**Deferred Decisions (Post-MVP):**
- Rate limiting (add post-launch once traffic patterns are known)
- CDN/edge caching (Vercel handles web; API caching deferred)
- Monitoring/APM (add after first paying users)

### Data Architecture

**Database:** PostgreSQL (managed via Railway add-on)

**ORM & Migrations:** GORM v2
- `AutoMigrate` in development for schema iteration
- GORM `Migrate` with versioned migration files in staging/production
- Schema-first: Go structs define the data model; GORM generates SQL

**Caching:** None at MVP — PostgreSQL with proper indexes is sufficient for 10,000 users. Per-user query patterns mean cache hit rates would be low anyway. Revisit with Redis post-MVP if capacity model queries become a bottleneck.

**Data Validation:** Go struct tags for GORM constraints + custom validation middleware in Gin for request body validation

### Authentication & Security

**Provider:** Clerk (`github.com/clerk/clerk-sdk-go/v2` — last updated Jan 2026)

**Rationale:** Clerk handles email/password, Google OAuth, and Apple Sign-In out of the box. React and Expo SDKs available for frontend. Go SDK v2 provides HTTP middleware for Gin that validates Clerk session JWTs automatically. Eliminates ~2-3 weeks of custom auth implementation for a solo developer.

**Session Strategy:**
- Web: Clerk session cookie (managed by Clerk's React SDK + Clerk middleware on API)
- Mobile: Clerk Expo SDK stores session token securely; passed as `Authorization: Bearer <token>` header
- Landing page auto-login: Clerk's `<SignedIn>` / `<SignedOut>` components handle redirect logic

**API Security:**
- All API routes except `/health` protected by Clerk Gin middleware
- Subscription status checked as secondary middleware after auth (validates trial/paid state from database)
- Per-user data isolation enforced at query level: all DB queries scoped to `user_id` from Clerk session claims
- HTTPS enforced (Railway + Vercel both terminate TLS)

### API & Communication Patterns

**Design:** REST API with consistent patterns across all endpoints

**Versioning:** `/api/v1/` prefix on all routes. Breaking changes increment to `/api/v2/` (not anticipated for MVP)

**Error Response Standard:**
```json
{
  "error": {
    "code": "CAPACITY_MODEL_UNAVAILABLE",
    "message": "Capacity estimate not available — still learning your pace",
    "details": {}
  }
}
```

**Real-time Sync:** Server-Sent Events (SSE)
- Clients connect to `GET /api/v1/stream` after login
- API pushes events when todo/goal state changes affecting the user
- SSE is unidirectional (server → client); mutations use standard REST endpoints
- Reconnection handled automatically by `EventSource` API (web) and equivalent on Expo
- Graceful degradation: if SSE connection drops, client polls on next user action

**OpenAPI / Type Generation:**
- `swaggo/swag` generates OpenAPI spec from Go annotations
- `openapi-typescript` generates TypeScript types in `packages/types` from spec
- Run as part of CI pipeline to keep types in sync

**Push Notifications:** Triggered by Go API via APNs (Apple) and FCM (Google) after relevant events (capacity exceeded, goal nudge threshold met, daily reminder time)

### Frontend Architecture

**State Management:**
- **TanStack Query v5** (v5.99.0) — server state management: todos, goals, capacity data, subscription status. Handles caching, background refetching, and optimistic updates.
- **Zustand** — local UI state: modal open/close, form state, notification preferences UI, offline queue indicator

**Styling:**
- **Web:** Tailwind CSS v4 + shadcn/ui component primitives
- **Mobile:** NativeWind v4 (Tailwind utility classes for React Native) — shared class names where possible for design consistency
- Design tokens defined once in `packages/types` and referenced in both platforms

**Routing:**
- **Web:** React Router v7 (Vite-native, no SSR needed)
- **Mobile:** Expo Router v3 (file-based routing built on React Navigation)

**API Client:**
- Auto-generated TypeScript client from OpenAPI spec (via `openapi-fetch` or similar)
- Wraps TanStack Query — queries and mutations typed end-to-end

**Offline Strategy (Mobile):**
- TanStack Query's `persistQueryClient` plugin + AsyncStorage for caching today's todos/goals locally
- Mutations queued when offline using TanStack Query's `onMutate` optimistic update pattern
- Sync on reconnect: flush queued mutations, then invalidate queries to pull fresh server state

### Infrastructure & Deployment

**Environments:** Development → Staging → Production (three Railway environments for API + DB; Vercel preview deployments for web)

**CI/CD:** GitHub Actions
- On PR: lint, typecheck, Go tests, generate OpenAPI spec, check types in sync
- On merge to `main`: deploy to staging automatically
- On release tag: deploy to production (manual trigger)

**Deployment Targets:**
- Web: Vercel (automatic from GitHub, preview URLs on PRs)
- API: Railway (Docker-based Go deployment, auto-deploys from `main`)
- Database: Railway PostgreSQL add-on (same project as API)
- Mobile: Expo EAS Build → App Store / Play Store (manual submission)

**Logging:** Go `slog` (structured JSON logging) → Railway log aggregation. Web errors: Vercel function logs. Mobile: Expo crash reporting.

**Environment Configuration:**
- Secrets managed in Railway (API) and Vercel (web) dashboards — never committed
- `packages/config` holds shared environment variable schemas with TypeScript validation (using `zod`)

### Decision Impact Analysis

**Implementation Sequence:**
1. Monorepo scaffold (Turborepo + pnpm + go.work)
2. PostgreSQL schema + GORM models
3. Clerk integration (Go middleware + React/Expo SDKs)
4. Core REST API endpoints (Gin routes)
5. Web SPA (Vite + React Router + TanStack Query + Tailwind)
6. Mobile app (Expo + Expo Router + NativeWind)
7. SSE real-time sync
8. Push notifications (APNs + FCM via Go)
9. Stripe subscription flow
10. CI/CD pipeline (GitHub Actions)

**Cross-Component Dependencies:**
- Clerk auth must be working before any protected API or UI can be built
- PostgreSQL schema + GORM models must exist before API endpoints
- OpenAPI spec generation must run before TypeScript client is generated
- SSE requires authenticated API connection — depends on Clerk middleware

## Implementation Patterns & Consistency Rules

### Naming Patterns

**Database Naming Conventions (GORM + PostgreSQL):**
- Tables: plural `snake_case` — `users`, `todos`, `goals`, `daily_summaries`, `notification_preferences`
- Columns: `snake_case` — `user_id`, `effort_points`, `created_at`, `updated_at`
- Foreign keys: `{table_singular}_id` — `user_id`, `goal_id`
- Indexes: `idx_{table}_{column}` — `idx_todos_user_id`, `idx_todos_created_at`
- GORM struct tags: `gorm:"column:user_id"` (explicit, never rely on convention inference)

**API Naming Conventions (Gin REST):**
- Routes: plural `kebab-case` resources — `/api/v1/todos`, `/api/v1/goals`, `/api/v1/daily-summaries`
- URL parameters: `:id` format — `/api/v1/todos/:id`
- Query parameters: `snake_case` — `?user_id=`, `?created_after=`
- HTTP verbs: POST=create, GET=read, PUT=replace, PATCH=update, DELETE=delete

**Go Code Naming:**
- Structs: `PascalCase` — `Todo`, `Goal`, `DailySummary`
- JSON struct tags: `snake_case` — `json:"effort_points"`, `json:"created_at"`
- Handler functions: `{Verb}{Resource}` — `CreateTodo`, `ListTodos`, `UpdateGoal`
- Package names: lowercase singular — `handler`, `service`, `repository`, `model`

**TypeScript / React Code Naming:**
- Components: `PascalCase` — `TodoCard`, `GoalBadge`, `CapacityMeter`
- Files: `PascalCase.tsx` for components — `TodoCard.tsx`, `CapacityMeter.tsx`
- Hooks: `use` prefix — `useTodos`, `useCapacity`, `useGoals`
- Utility files: `camelCase.ts` — `formatPoints.ts`, `dateHelpers.ts`
- TanStack Query keys: arrays with resource + scope — `['todos', userId, date]`, `['goals', userId]`
- Zustand stores: `use{Name}Store` — `useUIStore`, `useOfflineQueueStore`

### Structure Patterns

**Project Organization:**
```
apps/
  api/
    cmd/server/        # entry point
    internal/
      handler/         # Gin route handlers (one file per resource)
      service/         # business logic
      repository/      # GORM database queries
      model/           # GORM structs
      middleware/      # Clerk auth, subscription check
    docs/              # swaggo-generated OpenAPI spec
  web/
    src/
      features/        # feature-based: features/todos/, features/goals/
        {feature}/
          components/  # feature-specific components
          hooks/       # feature-specific TanStack Query hooks
          {Feature}Page.tsx
      components/      # shared UI components
      lib/             # shared utilities
      routes/          # React Router v7 route definitions
  mobile/
    app/               # Expo Router file-based routes
    features/          # mirrors web: features/todos/, features/goals/
      {feature}/
        components/
        hooks/
    components/        # shared UI components
packages/
  types/               # auto-generated from OpenAPI spec (do not hand-edit)
```

**Test File Location:**
- Go: `{file}_test.go` co-located — `handler/todo_handler_test.go`
- TypeScript: `{File}.test.tsx` co-located — `TodoCard.test.tsx`
- No separate `__tests__` directories

### Format Patterns

**API Success Response:**
```json
// Single resource
{ "data": { "id": "...", "effort_points": 3 } }

// Collection
{ "data": [...], "meta": { "total": 42 } }
```

**API Error Response:**
```json
{
  "error": {
    "code": "TODO_NOT_FOUND",
    "message": "Todo not found",
    "details": {}
  }
}
```

**Error Code Format:** `SCREAMING_SNAKE_CASE` noun phrases — `TODO_NOT_FOUND`, `CAPACITY_MODEL_UNAVAILABLE`, `SUBSCRIPTION_REQUIRED`

**Date/Time Format:** ISO 8601 strings in all API payloads — `"2026-04-15T09:00:00Z"`. Never Unix timestamps in JSON.

**JSON Field Naming:** `snake_case` in all API responses (Go struct tags enforce this). TypeScript types generated from OpenAPI spec will use `snake_case` and match exactly — no conversion layer.

### Communication Patterns

**SSE Event Structure:**
```json
{
  "event": "todo.updated",
  "data": { "id": "...", "status": "complete" },
  "timestamp": "2026-04-15T09:00:00Z"
}
```
- Event names: `{resource}.{verb}` dot notation — `todo.created`, `goal.updated`, `capacity.recalculated`
- Always include `timestamp` on SSE events

**TanStack Query Invalidation Pattern:**
- After mutations, invalidate by resource key — `queryClient.invalidateQueries({ queryKey: ['todos', userId] })`
- Never invalidate all queries (`queryClient.invalidateQueries()`) — too broad
- Optimistic updates required for: todo complete/incomplete toggle, todo reorder

**Zustand Store Pattern:**
- One store for all transient UI state (`useUIStore`) — modal state, notification prefs draft
- One store for offline queue (`useOfflineQueueStore`) — pending mutations when offline
- No store for server data — that's TanStack Query's domain

### Process Patterns

**Error Handling:**
- Go handlers: return early on error, log with `slog.Error`, return standardized JSON error
- React: TanStack Query `onError` callbacks for user-visible errors; React Error Boundary for render failures
- Clerk auth failures: redirect to `/login` (web) or auth screen (mobile) — never show raw 401 to user
- SSE disconnection: silent reconnect via `EventSource` retry; no user-visible error unless 3+ consecutive failures

**Loading States:**
- Use TanStack Query's `isPending` / `isFetching` — never hand-roll loading booleans for server data
- Skeleton UI (not spinners) for initial page load — `TodoCard` skeleton, `GoalBadge` skeleton
- Optimistic updates for toggles — no loading state shown for instant-feel interactions

**Capacity Model Messaging (Wellness Language Rule):**
- All copy observational, never prescriptive: "Your recent average is 18 points" ✅ vs "You should do 18 points" ❌
- "Learning your pace" messaging for days 1-10, never "not enough data" or error language
- Encouragement triggers: completion ≥ 90% of capacity estimate

### Enforcement Guidelines

**All AI Agents MUST:**
- Use `snake_case` JSON tags on all Go structs — never rely on GORM/JSON default inference
- Scope every database query to `user_id` from Clerk session claims — never query without user isolation
- Import TypeScript types exclusively from `packages/types` — never hand-write types that duplicate the OpenAPI schema
- Use TanStack Query for all server state — never `useState` + `useEffect` + `fetch` pattern
- Check subscription middleware response before any data endpoint returns — subscription gate is server-enforced

**Anti-Patterns to Avoid:**
```go
// ❌ Never query without user_id scope
db.Find(&todos)

// ✅ Always scope to authenticated user
db.Where("user_id = ?", clerkUserID).Find(&todos)
```

```typescript
// ❌ Never hand-write types that mirror API shapes
interface Todo { effort_points: number }

// ✅ Import from generated types
import type { Todo } from '@repo/types'

// ❌ Never manage server state in useState
const [todos, setTodos] = useState([])
useEffect(() => { fetch('/api/v1/todos').then(...) }, [])

// ✅ Use TanStack Query
const { data: todos } = useQuery({ queryKey: ['todos', userId], queryFn: fetchTodos })
```

## Project Structure & Boundaries

### Complete Project Directory Structure

```
todoFastTrack/                          # Turborepo monorepo root
├── .github/
│   └── workflows/
│       ├── ci.yml                      # PR: lint, typecheck, Go tests, OpenAPI sync check
│       └── deploy.yml                  # main→staging, tag→production
├── .gitignore
├── turbo.json                          # Turborepo pipeline config
├── pnpm-workspace.yaml
├── package.json                        # root workspace package.json
├── go.work                             # Go workspace for monorepo
├── go.work.sum
│
├── apps/
│   ├── api/                            # Go REST API (Railway)
│   │   ├── go.mod
│   │   ├── go.sum
│   │   ├── Dockerfile
│   │   ├── .env.example
│   │   ├── cmd/
│   │   │   └── server/
│   │   │       └── main.go             # entry point, Gin setup, route registration
│   │   ├── docs/                       # swaggo-generated OpenAPI spec (auto-generated)
│   │   │   ├── swagger.json
│   │   │   └── swagger.yaml
│   │   └── internal/
│   │       ├── middleware/
│   │       │   ├── auth.go             # Clerk JWT validation middleware
│   │       │   └── subscription.go     # subscription gating middleware
│   │       ├── model/                  # GORM structs (source of truth for DB schema)
│   │       │   ├── user.go             # User (Clerk user_id, subscription state)
│   │       │   ├── goal.go             # Goal (name, user_id, created_at)
│   │       │   ├── todo.go             # Todo (title, effort_points, goal_id, wellness_category, date, status)
│   │       │   ├── daily_summary.go    # DailySummary (date, points_planned, points_completed)
│   │       │   └── notification_pref.go # NotificationPreference (type, enabled, time)
│   │       ├── repository/             # GORM database queries (all scoped to user_id)
│   │       │   ├── user_repo.go
│   │       │   ├── goal_repo.go
│   │       │   ├── todo_repo.go
│   │       │   ├── daily_summary_repo.go
│   │       │   └── notification_pref_repo.go
│   │       ├── service/                # business logic
│   │       │   ├── capacity_service.go # capacity model calculation (server-authoritative)
│   │       │   ├── goal_service.go     # goal nudge logic, streak calculation
│   │       │   ├── todo_service.go     # todo CRUD, carry-forward logic
│   │       │   ├── notification_service.go # APNs + FCM dispatch (non-blocking)
│   │       │   └── stripe_service.go   # Stripe subscription lifecycle
│   │       └── handler/                # Gin route handlers
│   │           ├── auth_handler.go     # POST /api/v1/auth/webhook (Clerk user sync)
│   │           ├── goal_handler.go     # CRUD /api/v1/goals
│   │           ├── todo_handler.go     # CRUD /api/v1/todos
│   │           ├── capacity_handler.go # GET /api/v1/capacity
│   │           ├── summary_handler.go  # GET /api/v1/daily-summaries
│   │           ├── notification_handler.go # GET/PATCH /api/v1/notification-preferences
│   │           ├── subscription_handler.go # GET /api/v1/subscription, POST /api/v1/stripe/webhook
│   │           ├── stream_handler.go   # GET /api/v1/stream (SSE)
│   │           └── health_handler.go   # GET /health (unauthenticated)
│   │
│   ├── web/                            # Vite + React SPA (Vercel)
│   │   ├── package.json
│   │   ├── vite.config.ts
│   │   ├── tsconfig.json
│   │   ├── tailwind.config.ts
│   │   ├── index.html
│   │   ├── .env.example
│   │   └── src/
│   │       ├── main.tsx                # React entry point, Clerk provider, Router, QueryClient
│   │       ├── App.tsx                 # route definitions
│   │       ├── routes/                 # React Router v7 route tree
│   │       │   ├── index.tsx           # landing page route (FR46-FR49)
│   │       │   ├── auth.tsx            # auth callback route
│   │       │   └── app.tsx             # authenticated app shell (subscription guard)
│   │       ├── features/
│   │       │   ├── landing/            # FR46-FR49: landing page, SEO
│   │       │   │   ├── LandingPage.tsx
│   │       │   │   └── components/
│   │       │   │       ├── HeroSection.tsx
│   │       │   │       └── PricingSection.tsx
│   │       │   ├── onboarding/         # FR9-FR12: goal setup, discovery flow
│   │       │   │   ├── OnboardingPage.tsx
│   │       │   │   ├── hooks/
│   │       │   │   │   └── useOnboarding.ts
│   │       │   │   └── components/
│   │       │   │       ├── GoalDiscoveryFlow.tsx
│   │       │   │       └── GoalSetupForm.tsx
│   │       │   ├── goals/              # FR11: goal management
│   │       │   │   ├── GoalsPage.tsx
│   │       │   │   ├── hooks/
│   │       │   │   │   └── useGoals.ts
│   │       │   │   └── components/
│   │       │   │       ├── GoalCard.tsx
│   │       │   │       └── GoalBadge.tsx
│   │       │   ├── todos/              # FR13-FR21: daily planning
│   │       │   │   ├── TodayPage.tsx
│   │       │   │   ├── hooks/
│   │       │   │   │   └── useTodos.ts
│   │       │   │   └── components/
│   │       │   │       ├── TodoCard.tsx
│   │       │   │       ├── TodoForm.tsx
│   │       │   │       └── WellnessCategoryPicker.tsx
│   │       │   ├── capacity/           # FR22-FR27: capacity model, history
│   │       │   │   ├── hooks/
│   │       │   │   │   └── useCapacity.ts
│   │       │   │   └── components/
│   │       │   │       ├── CapacityMeter.tsx
│   │       │   │       ├── CapacityLearningState.tsx
│   │       │   │       └── HistoryChart.tsx
│   │       │   ├── reinforcement/      # FR28-FR31: feedback, streaks
│   │       │   │   └── components/
│   │       │   │       ├── CompletionCelebration.tsx
│   │       │   │       └── GoalStreakBadge.tsx
│   │       │   ├── notifications/      # FR32-FR36: notification preferences
│   │       │   │   ├── NotificationSettingsPage.tsx
│   │       │   │   └── hooks/
│   │       │   │       └── useNotificationPrefs.ts
│   │       │   ├── subscription/       # FR41-FR45, FR8: billing management
│   │       │   │   ├── SubscriptionPage.tsx
│   │       │   │   └── hooks/
│   │       │   │       └── useSubscription.ts
│   │       │   └── account/            # FR1-FR8: auth, account settings, deletion
│   │       │       ├── AccountPage.tsx
│   │       │       └── hooks/
│   │       │           └── useAccount.ts
│   │       ├── components/             # shared UI components
│   │       │   ├── ui/                 # shadcn/ui primitives
│   │       │   ├── SkeletonCard.tsx
│   │       │   └── ErrorBoundary.tsx
│   │       ├── lib/
│   │       │   ├── api-client.ts       # openapi-fetch client (typed from packages/types)
│   │       │   ├── query-client.ts     # TanStack Query client config
│   │       │   ├── sse-client.ts       # SSE connection manager (FR37-FR40)
│   │       │   └── date-helpers.ts
│   │       └── store/
│   │           ├── ui-store.ts         # Zustand: modal state, notification prefs draft
│   │           └── offline-queue-store.ts # Zustand: pending mutations (web fallback)
│   │
│   └── mobile/                         # Expo SDK 52 + React Native (EAS Build)
│       ├── package.json
│       ├── app.json                    # Expo config
│       ├── babel.config.js
│       ├── tsconfig.json
│       ├── .env.example
│       ├── app/                        # Expo Router file-based routes
│       │   ├── _layout.tsx             # root layout, Clerk provider, QueryClient
│       │   ├── index.tsx               # redirect: onboarding or today
│       │   ├── (auth)/
│       │   │   ├── sign-in.tsx         # FR1-FR4
│       │   │   └── sign-up.tsx
│       │   ├── (onboarding)/
│       │   │   └── index.tsx           # FR9-FR12
│       │   └── (app)/
│       │       ├── _layout.tsx         # authenticated tab navigator (subscription guard)
│       │       ├── today.tsx           # FR13-FR21: daily planning (default tab)
│       │       ├── goals.tsx           # FR11: goal management
│       │       ├── history.tsx         # FR27: historical data
│       │       └── settings.tsx        # FR6, FR8, FR36: account + notification prefs
│       ├── features/                   # mirrors web feature structure
│       │   ├── todos/
│       │   │   ├── hooks/
│       │   │   │   └── useTodos.ts
│       │   │   └── components/
│       │   │       ├── TodoCard.tsx
│       │   │       └── TodoForm.tsx
│       │   ├── goals/
│       │   │   ├── hooks/
│       │   │   │   └── useGoals.ts
│       │   │   └── components/
│       │   │       └── GoalBadge.tsx
│       │   ├── capacity/
│       │   │   ├── hooks/
│       │   │   │   └── useCapacity.ts
│       │   │   └── components/
│       │   │       ├── CapacityMeter.tsx
│       │   │       └── CapacityLearningState.tsx
│       │   ├── onboarding/
│       │   │   └── components/
│       │   │       └── GoalDiscoveryFlow.tsx
│       │   └── notifications/
│       │       └── hooks/
│       │           └── useNotificationPrefs.ts
│       ├── components/
│       │   ├── SkeletonCard.tsx
│       │   └── ErrorBoundary.tsx
│       ├── lib/
│       │   ├── api-client.ts
│       │   ├── query-client.ts         # TanStack Query + AsyncStorage persistence
│       │   ├── sse-client.ts
│       │   └── push-notifications.ts  # Expo notification registration (APNs/FCM)
│       └── store/
│           ├── ui-store.ts
│           └── offline-queue-store.ts
│
└── packages/
    └── types/                          # auto-generated — DO NOT HAND-EDIT
        ├── package.json
        ├── index.ts
        └── generated/
            └── api.ts                  # openapi-typescript output from docs/swagger.json
```

### Architectural Boundaries

**API Boundaries:**
- All API routes under `/api/v1/` prefix, protected by Clerk middleware (except `/health`)
- Subscription middleware runs after auth on all data routes — returns `403 SUBSCRIPTION_REQUIRED` if inactive
- SSE stream (`GET /api/v1/stream`) maintains persistent connection per authenticated user
- Stripe webhooks (`POST /api/v1/stripe/webhook`) validated by Stripe signature header before processing
- Clerk webhooks (`POST /api/v1/auth/webhook`) create/sync user records in local DB on account creation

**Component Boundaries:**
- `packages/types` is the only place TypeScript API types live — web and mobile both import from here
- Feature-specific hooks own their TanStack Query subscriptions — no cross-feature query sharing
- `lib/sse-client.ts` owns the SSE connection lifecycle; features subscribe to events via a listener pattern
- Zustand stores hold only UI state — never replicate server data that TanStack Query owns

**Data Boundaries:**
- Capacity model lives entirely in `service/capacity_service.go` — no client-side calculation ever
- `repository/` layer is the only code that touches the database — `service/` calls repos, never `db` directly
- All repository methods require a `userID string` parameter — enforced by function signatures
- `packages/types/generated/` is owned by the CI pipeline — regenerated on every PR from the Go OpenAPI spec

### Requirements to Structure Mapping

| FR Category | API Handler | Web Feature | Mobile Feature |
|---|---|---|---|
| User Account Management (FR1-8) | `auth_handler.go`, `subscription_handler.go` | `features/account/` | `app/(auth)/`, settings tab |
| Onboarding & Goal Setup (FR9-12) | `goal_handler.go` | `features/onboarding/` | `app/(onboarding)/` |
| Daily Planning (FR13-21) | `todo_handler.go` | `features/todos/` | `features/todos/` |
| Capacity Management (FR22-27) | `capacity_handler.go`, `summary_handler.go` | `features/capacity/` | `features/capacity/` |
| Feedback & Reinforcement (FR28-31) | `goal_handler.go` (streak calc) | `features/reinforcement/` | `features/capacity/components/` |
| Notifications (FR32-36) | `notification_handler.go`, `notification_service.go` | `features/notifications/` | `features/notifications/` |
| Cross-Platform & Sync (FR37-40) | `stream_handler.go` | `lib/sse-client.ts` | `lib/sse-client.ts`, `store/offline-queue-store.ts` |
| Subscription & Billing (FR41-45) | `subscription_handler.go`, `stripe_service.go` | `features/subscription/` | settings tab |
| Landing Page (FR46-49) | n/a (static SPA) | `features/landing/`, `routes/index.tsx` | n/a |

### Integration Points

**Internal Communication:**
- Web/Mobile → API: typed `openapi-fetch` client (`lib/api-client.ts`), all calls authenticated via Clerk token
- API → Database: GORM repository pattern — handlers call services, services call repositories
- API → Clients (real-time): SSE push on mutation events (`todo.updated`, `capacity.recalculated`)
- API → Push services: `notification_service.go` calls APNs/FCM in a goroutine — failures logged, never block response

**External Integrations:**
- **Clerk**: JWT validation in `middleware/auth.go`; webhook syncs user creation to local `users` table
- **Stripe**: `stripe_service.go` creates subscriptions; webhook at `/api/v1/stripe/webhook` updates state idempotently
- **APNs/FCM**: called from `notification_service.go` in a goroutine — failures logged, never block API response
- **OpenAPI → TypeScript**: CI generates `packages/types/generated/api.ts` from `apps/api/docs/swagger.json`

**Data Flow (daily planning):**
1. User adds todo → `POST /api/v1/todos` → `todo_handler` → `todo_service` → `todo_repo` → PostgreSQL
2. API returns saved todo → TanStack Query cache updated (optimistic update confirmed)
3. SSE event `todo.created` pushed to all connected clients for same user
4. Other device receives SSE → TanStack Query invalidates `['todos', userId, date]` → refetch

### Development Workflow Integration

**Local Development:**
```bash
pnpm dev  # Turborepo runs all three in parallel:
          # apps/web:    vite dev server → localhost:5173
          # apps/mobile: expo start      → localhost:8081
          # apps/api:    go run ./cmd/server → localhost:8080
```

**Type Sync (run after Go model changes):**
```bash
pnpm run generate:types
# 1. swag init (apps/api) → docs/swagger.json
# 2. openapi-typescript docs/swagger.json → packages/types/generated/api.ts
```

**CI Pipeline (GitHub Actions on PR):**
1. Go: `go vet`, `go test ./...`
2. TypeScript: `tsc --noEmit` (web + mobile)
3. Generate OpenAPI spec, check `packages/types` is in sync (fail if diff)
4. Lint: `golangci-lint` (Go), `eslint` (TS)

## Architecture Validation Results

### Coherence Validation ✅

**Decision Compatibility:**
All technology choices verified compatible:
- Clerk Go SDK v2 ↔ Gin middleware: direct integration, no conflicts
- TanStack Query v5 ↔ React 18 + Expo SDK 52: both support concurrent mode
- NativeWind v4 ↔ Expo SDK 52: officially supported pairing
- Expo Router v3 ↔ Expo SDK 52: bundled together, no version mismatch risk
- GORM v2 ↔ PostgreSQL: native driver, no adapter layer needed
- swaggo/swag ↔ openapi-typescript: standard OpenAPI 3.0 spec output, compatible with any consumer
- Turborepo ↔ go.work: orthogonal tools — Turborepo orchestrates JS pipeline, go.work manages Go modules; no conflict

**Pattern Consistency:**
- Repository pattern (Go) aligns with GORM's intended usage
- Feature-based folder structure (web/mobile) consistent with TanStack Query's per-resource query key design
- SSE unidirectional model consistent with REST mutation pattern — no bidirectional confusion
- Zustand (UI state) + TanStack Query (server state) separation is clear and non-overlapping

**Structure Alignment:**
- `packages/types` as single source of TypeScript types enforces the Go-first type ownership model
- Feature directories mirror between web and mobile — consistent mental model for agents working across platforms
- `repository/` signatures requiring `userID string` structurally enforce the per-user isolation NFR

### Requirements Coverage Validation ✅

**All 49 FRs covered:**

| FR Category | Coverage | Notes |
|---|---|---|
| User Account Management (FR1-8) | ✅ | Clerk handles FR1-4 natively; FR7 deletion — immediate hard delete approach |
| Onboarding & Goal Setup (FR9-12) | ✅ | goal_handler + onboarding features |
| Daily Planning (FR13-21) | ✅ | todo_handler + carry-forward in todo_service |
| Capacity Management (FR22-27) | ✅ | capacity_service (server-authoritative) |
| Feedback & Reinforcement (FR28-31) | ✅ | goal_service for streaks; reinforcement components |
| Notifications (FR32-36) | ✅ | notification_service goroutine, APNs/FCM |
| Cross-Platform & Sync (FR37-40) | ✅ | SSE + offline queue + AsyncStorage |
| Subscription & Billing (FR41-45) | ✅ | stripe_service + idempotent webhook handler |
| Landing Page (FR46-49) | ✅ | features/landing, Clerk SignedIn/SignedOut for FR49 |

**All 21 NFRs covered:**

| NFR Area | Coverage | Notes |
|---|---|---|
| Performance (NFR1-5) | ✅ | Vite/Vercel (NFR1), optimistic updates (NFR3), SSE (NFR4), server-only capacity (NFR5) |
| Security (NFR6-11) | ✅ | Railway TLS + at-rest encryption (NFR6), Clerk bcrypt (NFR7-8), repo signatures (NFR9), Stripe delegation (NFR10) |
| Scalability (NFR12-14) | ✅ | Stateless API + per-user indexes |
| Accessibility (NFR15-17) | ✅ | shadcn/ui + NativeWind; implementation concern |
| Integration (NFR18-21) | ✅ | Stripe idempotent webhook, Clerk server-side OAuth, goroutine notifications, offline sync |

### Gap Analysis Results

**Account Deletion Implementation (FR7, NFR11) — implementation detail, not an architectural gap:**
The 30-day data purge SLA (GDPR/CCPA) means deletion must *complete* within 30 days — not that it is delayed. MVP approach: immediate hard delete in a single database transaction in `auth_handler.go`:
1. Delete all user data across all tables in one transaction
2. Revoke Clerk session via Clerk Go SDK
3. Return 204

This is simpler than soft-delete + scheduled job and fully satisfies the compliance requirement at MVP scale.

**No critical architectural gaps identified.**

### Architecture Completeness Checklist

**✅ Requirements Analysis**
- [x] 49 FRs across 9 capability areas analyzed
- [x] 21 NFRs with specific targets documented
- [x] Technical constraints identified (server-side capacity, offline-first, monorepo, web-first launch)
- [x] Cross-cutting concerns mapped (auth, subscription state, per-user isolation, offline sync, push delivery, wellness language)

**✅ Architectural Decisions**
- [x] Critical decisions documented with verified versions
- [x] Technology stack fully specified (Go 1.22+, Expo SDK 52, TanStack Query v5.99.0, etc.)
- [x] Integration patterns defined (Clerk, Stripe, APNs/FCM, SSE)
- [x] Performance considerations addressed (optimistic updates, no MVP caching, per-user indexes)

**✅ Implementation Patterns**
- [x] Naming conventions for DB, API, Go, and TypeScript established
- [x] Structure patterns defined with test co-location rules
- [x] Communication patterns specified (SSE events, TanStack Query keys, Zustand scope)
- [x] Process patterns documented (error handling, loading states, wellness language rule)

**✅ Project Structure**
- [x] Complete directory structure defined to file level
- [x] Component boundaries established (types ownership, repo signatures, SSE client)
- [x] 9 FR categories mapped to specific files and directories
- [x] Integration points fully documented

### Architecture Readiness Assessment

**Overall Status: READY FOR IMPLEMENTATION**

**Confidence Level: High** — all critical decisions are documented with verified versions, all 49 FRs have explicit architectural homes, and consistency rules are structurally enforced (compiler-checked type signatures, generated types pipeline).

**Key Strengths:**
- Type safety enforced end-to-end via Go → OpenAPI → TypeScript pipeline — entire class of runtime type errors eliminated
- Per-user isolation enforced by repository function signatures — compiler-checked, not a convention to remember
- Capacity model fully server-side from day one — no refactoring risk as user base grows
- Offline-first mobile architecture handles the primary trust-eroding scenario before it becomes a user problem
- Clerk eliminates ~2-3 weeks of auth implementation — well-suited for solo developer

**Areas for Future Enhancement (post-MVP):**
- Redis caching for capacity model queries if PostgreSQL becomes a bottleneck at scale
- Rate limiting on API endpoints (deferred until traffic patterns are known)
- APM/monitoring tooling (deferred until first paying users)
- Shared `packages/ui` for cross-platform component primitives (if design divergence becomes a maintenance burden)

### Implementation Handoff

**First implementation step:** Monorepo scaffold per the initialization commands in the Starter Template section. Epic 1, Story 1 must be this scaffold.

**AI Agent Guidelines:**
- All architectural questions answered by this document — do not invent decisions not documented here
- `packages/types` is read-only at the agent level; regenerate via `pnpm run generate:types` after Go model changes
- Every database query must pass `userID` — if a query doesn't require a `userID`, question whether it belongs in the hot path
- Wellness language rule applies to all user-facing copy, not just capacity messaging
