---
stepsCompleted:
  - step-01-validate-prerequisites
inputDocuments:
  - _bmad-output/planning-artifacts/prd.md
  - _bmad-output/planning-artifacts/architecture.md
  - _bmad-output/planning-artifacts/ux-design-specification.md
---

# Tend - Epic Breakdown

## Overview

This document provides the complete epic and story breakdown for Tend, decomposing the requirements from the PRD, UX Design Specification, and Architecture into implementable stories.

## Requirements Inventory

### Functional Requirements

FR1: Users can create an account with email and password
FR2: Users can create an account using Google Sign-In
FR3: Users can create an account using Apple Sign-In
FR4: Users can log in to an existing account
FR5: Returning authenticated users are automatically redirected to the app from the landing page
FR6: Users can log out of their account
FR7: Users can delete their account and all associated data
FR8: Users can manage their subscription (view status, cancel, reactivate)
FR9: New users are guided through a goal-setting flow to establish up to 3 long-term goals
FR10: Users who don't know their goals can access a guided discovery flow that prompts reflection questions to help identify them
FR11: Users can name, edit, and delete their long-term goals at any time
FR12: Users can skip the guided discovery flow and set goals manually
FR13: Users can create todos for the current day
FR14: Users can assign an effort point value to each todo
FR15: Users can tag each todo with one of their long-term goals (or mark it as untagged)
FR16: Users can assign each todo to a wellness category (exercise, fun/hobby, rest)
FR17: Users can mark todos as complete
FR18: Users can edit or delete todos
FR19: Users can reorder todos within their daily list
FR20: Users must include at least one wellness-category todo before their day plan is considered complete
FR21: Users can carry incomplete todos forward to the next day
FR22: The system tracks effort points completed by a user each day over time
FR23: The system displays a "learning" state for the first 7-10 days with messaging explaining it is building a capacity baseline
FR24: The system generates a daily capacity estimate once sufficient history exists (minimum 5 days)
FR25: The system alerts users when planned daily points exceed their capacity estimate
FR26: The system encourages users to add tasks when planned total is significantly below capacity estimate
FR27: Users can view their historical daily point completion data
FR28: The system delivers positive reinforcement messages when a user completes a todo
FR29: The system delivers positive reinforcement when a user meets or exceeds their daily capacity target
FR30: The system notifies users when a long-term goal has not been touched in 3 or more days
FR31: The system tracks and displays goal-activity streaks (consecutive days with at least one goal-tagged todo completed)
FR32: Users receive a configurable morning push notification to plan their day
FR33: Users receive a push notification when their planned list exceeds their capacity estimate
FR34: Users receive a push notification when a long-term goal has not been worked on for 3+ days
FR35: Users receive a push notification with positive reinforcement when they complete a strong day
FR36: Users can configure notification preferences (enable/disable each type, set notification time)
FR37: Users can access their account, goals, and todos from both the web app and the mobile app
FR38: Changes made on one platform appear on other connected platforms within 5 seconds
FR39: Users can view and interact with their todo list while offline on mobile
FR40: Changes made offline sync automatically to the server when connectivity is restored
FR41: New users receive a 30-day free trial with full feature access
FR42: Users are prompted to enter payment details at trial end
FR43: The system charges users $5/month via Stripe upon trial completion
FR44: Users receive notification of upcoming billing and failed payment attempts
FR45: Users can cancel their subscription and retain access until the end of the billing period
FR46: The public landing page is optimized for search engine indexing
FR47: The landing page allows new visitors to sign up
FR48: The landing page allows existing users to log in
FR49: The landing page detects authenticated returning users and redirects them to the app

### NonFunctional Requirements

NFR1 (Performance): Initial web app load under 3 seconds on standard broadband
NFR2 (Performance): React Native app launch under 2 seconds on mid-range devices
NFR3 (Performance): UI response to user actions under 300ms
NFR4 (Performance): Cross-platform sync latency under 5 seconds
NFR5 (Performance): Capacity model server response under 1 second after day plan submission
NFR6 (Security): All user data encrypted at rest (AES-256) and in transit (TLS 1.2+)
NFR7 (Security): Passwords stored using bcrypt — plaintext never stored or logged
NFR8 (Security): JWT tokens expire after 24 hours; refresh tokens rotate on use
NFR9 (Security): User data isolated — no cross-user data access possible
NFR10 (Security): Payment data fully delegated to Stripe — never touches application servers
NFR11 (Security): Account deletion purges all user data within 30 days
NFR12 (Scalability): System supports up to 10,000 users without architectural changes
NFR13 (Scalability): Database queries optimized for per-user access — no global table scans in hot paths
NFR14 (Scalability): Stateless API — horizontally scalable post-MVP
NFR15 (Accessibility): Web app meets WCAG 2.1 AA — keyboard navigable, screen reader compatible, sufficient color contrast
NFR16 (Accessibility): Mobile app follows iOS HIG and Android Material accessibility guidelines
NFR17 (Accessibility): Push notification content meaningful without visual context
NFR18 (Integration): Stripe subscription lifecycle webhooks processed idempotently
NFR19 (Integration): Google/Apple Sign-In OAuth 2.0 tokens validated server-side
NFR20 (Integration): APNs/FCM delivery failures logged; non-blocking to core app function
NFR21 (Integration): Offline sync uses last-write-wins for todo status; capacity model server-authoritative on reconnect

### Additional Requirements

- **Monorepo scaffold is Epic 1 Story 1** — Turborepo + pnpm workspace + go.work initialization using: `npx create-turbo@latest tend --package-manager pnpm`, Vite React TS web app, Expo mobile app, Go API module, `packages/types` package
- Three deployment environments required before launch: development, staging, production (Railway for API/DB, Vercel for web)
- GitHub Actions CI/CD pipeline: PR checks (go vet, go test, tsc --noEmit, eslint, OpenAPI sync check); auto-deploy to staging on merge to main; manual tag deploy to production
- OpenAPI type generation pipeline: `swaggo/swag` annotations → `docs/swagger.json` → `openapi-typescript` → `packages/types/generated/api.ts`; run after every Go model change; CI fails if types out of sync
- All database repository methods require `userID string` parameter — enforced at function signature level, never query without user scope
- SSE endpoint `GET /api/v1/stream` for real-time cross-platform sync — unidirectional server→client push on todo/goal/capacity mutations
- Clerk webhook `POST /api/v1/auth/webhook` syncs user creation to local `users` table
- Stripe webhook `POST /api/v1/stripe/webhook` processes subscription lifecycle events idempotently; validated by Stripe signature header
- APNs/FCM notification dispatch in goroutines — failures logged via `slog.Error`, never block API response
- GORM AutoMigrate for development; versioned GORM migration files for staging/production
- Offline-first mobile: TanStack Query `persistQueryClient` + AsyncStorage; mutations queued offline; sync flush on reconnect; last-write-wins for todo status, server-authoritative for capacity data
- Capacity model calculated server-side only in `service/capacity_service.go` — no client-side calculation ever
- All API routes under `/api/v1/` prefix; subscription middleware runs after Clerk auth middleware on all data routes

### UX Design Requirements

UX-DR1: Implement Tend's design token system — 14 semantic color tokens (primary #E8A0A0, accent #F0B090, background #FDFAF7, surface #FFF8F5, warning #D4A843, success #7BAF8A, goal slot colors lavender/sage/peach, etc.), 6-role typography scale (Plus Jakarta Sans, display through caption), 4px-base spacing scale — defined once in `packages/config`, consumed by Tailwind config (web) and NativeWind (mobile)
UX-DR2: Implement `CapacityBar` custom component — 4 states: `learning` (days 1-10, "Learning your pace" message), `on-track` (blush fill), `over-capacity` (amber fill + permissive inline note), `first-reveal` (animated one-time entrance at day 5-7 with specific message). `role="progressbar"`, `aria-valuenow`, `aria-valuemax`
UX-DR3: Implement `SizeChip` component — inline XS/S/M/L/XL effort size selector; opacity scale (30%→100%) on primary color; M pre-selected; `role="radiogroup"` + `role="radio"` per chip; arrow key navigation; `aria-label="Effort: [size]"`
UX-DR4: Implement `GoalChip` component — 3 variants: `selector` (creation row, tap-to-select one of 3 goals + None), `pill` (day header with colored dot + name), `badge` (todo row). Colored dot always paired with goal name text. `aria-label="[Goal name] tag"`
UX-DR5: Implement `WellnessIcon` component — 3 icon buttons (exercise/fun-hobby/rest); single-select; tapping again deselects; always optional; `aria-label` on each; text label always visible alongside icon
UX-DR6: Implement `TodoRow` component — states: `active`, `complete` (55% opacity + strikethrough via 300ms animation), `dragging`; anatomy: checkbox + title + chip row (SizeChip + GoalChip badge + WellnessIcon badge); full row tappable (not just checkbox); long-press reorder on mobile, drag handle on hover desktop
UX-DR7: Implement `TodoCreationRow` — inline progressive disclosure: collapsed dashed "Add a todo…" row → expands (200ms slide-down) → title input → SizeChip → GoalChip → WellnessIcon → Add button; capacity bar updates live on size selection; ESC collapses without saving; Return submits; no modal, no new screen
UX-DR8: Implement `GoalNudgeMessage` component — inline reinforcement beneath relevant todo row after completing a goal-tagged todo; variants: `goal-completion`, `day-complete`; fades in 150ms, fades out 300ms after 4s or next user action; `role="status"`, `aria-live="polite"`; copy always names specific goal/task, never generic
UX-DR9: Implement `CarryForwardPrompt` — morning prompt showing yesterday's incomplete todos one at a time; Keep / Skip per todo; focus trapped until dismissed; keyboard accessible; appears on first login of day when incomplete todos exist
UX-DR10: Implement `DaySummaryScreen` — optional end-of-day wrap-up triggered by "Wrap up today"; shows points completed + goals touched + warm specific sign-off; copy specific to day's data, never generic; day also closes automatically at midnight if not explicitly wrapped up
UX-DR11: Implement `CapacityRevealCard` — one-time first capacity estimate reveal at day 5-7 on app open; animated capacity bar entrance + specific message ("You tend to accomplish around X points on a good day") + dismiss; shown exactly once, permanently dismissed after; must feel like a reveal moment, not a data update
UX-DR12: Implement Direction A Airy Column layout — single centered column, max-width 600px, `mx-auto`, warm cream background (#FDFAF7); todo rows as soft surface cards (#FFF8F5) with `rounded-2xl`, 14px padding, 8px gap between rows; top navigation bar on desktop (logo left, avatar right); no sidebar at MVP; dashed inline "Add a todo…" row at bottom of list (no floating action button on web)
UX-DR13: Implement responsive layout — desktop-first; 5 Tailwind breakpoints (sm 320-599px full-width, md 600px+ centered column, lg/xl/2xl progressively wider whitespace); column fills 92% on mobile web; `px-4 md:px-0` padding pattern; minimum supported width 320px
UX-DR14: Implement WCAG 2.1 AA accessibility — semantic HTML (`<main>`, `<nav>`, `<section>`, `<button>`, `<ul>`/`<li>` for todo list); visible 2px focus rings (blush rose tone, `offset`); skip-to-content link (visually hidden until focused); ARIA live regions for capacity/reinforcement messages (`aria-live="polite"`); `aria-describedby` on todo items associating size/goal chips; native `<input type="checkbox">` with custom CSS; all interactive elements `<button>` never `<div onClick>`; minimum 44×44px touch targets
UX-DR15: Implement `prefers-reduced-motion` support — all entry/exit animations (todo completion, creation row expand, GoalNudgeMessage, CapacityBar) fallback to opacity-only transitions when OS preference set; implemented as a shared animation utility
UX-DR16: Implement micro-interaction animations — todo completion: checkmark fills `color-success` over 200ms, title fade+strikethrough 300ms; capacity bar width animates 400ms ease-in-out on todo add/remove/resize; creation row slide-down 200ms eased; GoalNudgeMessage fade-in 150ms / fade-out 300ms; no confetti, sparkles, or sound
UX-DR17: Enforce wellness language rules in all user-facing copy — capacity messages observational only ("You tend to accomplish around X"); warnings always permissive ("that's okay" / "feel free to"); reinforcement always specific (names goal or task); empty states warm not instructional; no exclamation points in system messages; no red anywhere (amber #D4A843 for warnings only); these rules apply to push notification copy as well
UX-DR18: Implement React Native accessibility — `accessibilityLabel` on all interactive elements; `accessibilityRole` declared (button/checkbox/text); `accessibilityLiveRegion` for reinforcement toasts; minimum 44pt touch targets enforced via `minHeight`/`minWidth` style rules; swipe-to-carry-forward has long-press menu tap-based alternative

### FR Coverage Map

{{requirements_coverage_map}}

## Epic List

{{epics_list}}
