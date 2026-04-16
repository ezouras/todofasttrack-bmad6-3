---
stepsCompleted:
  - step-01-init
  - step-02-context
  - step-03-starter
  - step-04-decisions
inputDocuments:
  - _bmad-output/planning-artifacts/prd.md
workflowType: 'architecture'
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
