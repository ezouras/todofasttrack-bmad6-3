---
stepsCompleted:
  - step-01-validate-prerequisites
  - step-02-design-epics
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

FR1: Epic 1 — Email/password account creation
FR2: Epic 1 — Google Sign-In
FR3: Epic 1 — Apple Sign-In
FR4: Epic 1 — Log in
FR5: Epic 1 — Auto-redirect authenticated returning users
FR6: Epic 1 — Log out
FR7: Epic 1 — Account deletion
FR8: Epic 6 — Subscription management
FR9: Epic 2 — Guided goal-setting flow
FR10: Epic 2 — Guided discovery flow
FR11: Epic 2 — Edit/delete goals
FR12: Epic 2 — Skip discovery, set goals manually
FR13: Epic 3 — Create todos
FR14: Epic 3 — Assign effort point value (XS–XL)
FR15: Epic 3 — Tag todo to goal
FR16: Epic 3 — Assign wellness category
FR17: Epic 3 — Mark todo complete
FR18: Epic 3 — Edit/delete todos
FR19: Epic 3 — Reorder todos
FR20: Epic 3 — Wellness todo required for day plan completion
FR21: Epic 3 — Carry forward incomplete todos
FR22: Epic 4 — Track daily effort points over time
FR23: Epic 4 — Learning state (days 1-10)
FR24: Epic 4 — Generate capacity estimate (5+ days)
FR25: Epic 4 — Alert when plan exceeds capacity
FR26: Epic 4 — Encourage adding tasks when significantly under capacity
FR27: Epic 4 — View historical point completion data
FR28: Epic 5 — Positive reinforcement on todo completion
FR29: Epic 5 — Positive reinforcement on capacity target met
FR30: Epic 5 — Goal inactivity nudge (3+ days)
FR31: Epic 5 — Goal-activity streaks
FR32: Epic 7 — Morning planning push notification
FR33: Epic 7 — Push notification when plan exceeds capacity
FR34: Epic 7 — Push notification when goal untouched 3+ days
FR35: Epic 7 — Push notification for strong day reinforcement
FR36: Epic 7 — Configure notification preferences
FR37: Epic 8 — Access from web and mobile
FR38: Epic 8 — Cross-platform sync within 5 seconds
FR39: Epic 8 — Offline mobile todo access
FR40: Epic 8 — Offline changes sync on reconnect
FR41: Epic 6 — 30-day free trial
FR42: Epic 6 — Payment details at trial end
FR43: Epic 6 — $5/month Stripe charge
FR44: Epic 6 — Billing notifications (upcoming, failed)
FR45: Epic 6 — Cancel and retain access through billing period
FR46: Epic 1 — SEO-optimized landing page
FR47: Epic 1 — Landing page sign-up
FR48: Epic 1 — Landing page log-in
FR49: Epic 1 — Landing page authenticated user redirect

## Epic List

### Epic 1: Foundation, Authentication & Landing Page
Users can discover Tend on the web, create an account, log in, and access the app as a returning user. The monorepo is initialized as the first story.
**FRs covered:** FR1, FR2, FR3, FR4, FR5, FR6, FR7, FR46, FR47, FR48, FR49
**UX-DRs covered:** UX-DR1, UX-DR12, UX-DR13, UX-DR14

### Epic 2: Goal Setup & Onboarding
New users establish up to 3 long-term goals through guided discovery or direct entry. Returning users can edit or delete goals at any time.
**FRs covered:** FR9, FR10, FR11, FR12
**UX-DRs covered:** UX-DR4

### Epic 3: Daily Planning — Core Loop
Users build a daily todo list with effort sizing (XS–XL), goal tagging, and wellness category tagging. They can complete, edit, reorder, and carry forward todos.
**FRs covered:** FR13, FR14, FR15, FR16, FR17, FR18, FR19, FR20, FR21
**UX-DRs covered:** UX-DR3, UX-DR5, UX-DR6, UX-DR7, UX-DR9, UX-DR15, UX-DR16, UX-DR17

### Epic 4: Capacity Intelligence
The app tracks daily effort history, displays a transparent learning state, and delivers a personalized capacity estimate. Users receive contextual feedback when their plan is over or under capacity and can view their history.
**FRs covered:** FR22, FR23, FR24, FR25, FR26, FR27
**UX-DRs covered:** UX-DR2, UX-DR11

### Epic 5: Reinforcement & Goal Streaks
Users receive warm, specific positive reinforcement on todo completion and strong days. The app tracks goal activity and surfaces nudges when a goal goes untouched for 3+ days. An optional day wrap-up ritual provides a satisfying close.
**FRs covered:** FR28, FR29, FR30, FR31
**UX-DRs covered:** UX-DR8, UX-DR10

### Epic 6: Subscription & Billing
New users receive a 30-day free trial with full feature access. Stripe handles payment collection at trial end. Users can view status, cancel, and reactivate their subscription.
**FRs covered:** FR8, FR41, FR42, FR43, FR44, FR45

### Epic 7: Push Notifications
Users receive configurable push notifications for morning planning reminders, capacity alerts, goal inactivity nudges, and strong-day reinforcement.
**FRs covered:** FR32, FR33, FR34, FR35, FR36

### Epic 8: Mobile App & Real-Time Sync
Users have full access to Tend on iOS and Android with offline-first support. Changes sync across platforms within 5 seconds via SSE. The mobile app mirrors all web features with a touch-optimized interface.
**FRs covered:** FR37, FR38, FR39, FR40
**UX-DRs covered:** UX-DR18

---

## Epic 1: Foundation, Authentication & Landing Page

Users can discover Tend on the web, create an account, log in, and access the app as a returning user. The monorepo is initialized as the first story.

### Story 1.1: Monorepo Scaffold & CI Pipeline

As a **developer**,
I want a Turborepo monorepo initialized with Go API, React web app, Expo mobile app, and shared types package,
So that all three apps can be developed, type-checked, and tested in a unified workspace.

**Acceptance Criteria:**

**Given** the repository is cloned,
**When** `pnpm install` is run,
**Then** all JS dependencies install without errors and the workspace structure matches: `apps/api`, `apps/web`, `apps/mobile`, `packages/types`

**Given** the monorepo is set up,
**When** `pnpm dev` is run,
**Then** the web Vite dev server (localhost:5173), Go API (localhost:8080), and Expo mobile bundler (localhost:8081) all start in parallel via Turborepo

**Given** a PR is opened on GitHub,
**When** the CI pipeline runs,
**Then** `go vet ./...` and `go test ./...` pass, TypeScript type-checks pass for web and mobile, and ESLint passes for both TS apps

**Given** `pnpm run generate:types` is run after Go model changes,
**When** the pipeline executes swaggo/swag then openapi-typescript,
**Then** `packages/types/generated/api.ts` is regenerated from `apps/api/docs/swagger.json`

**Given** the CI pipeline runs on a PR,
**When** `packages/types` is out of sync with the Go OpenAPI spec,
**Then** the pipeline fails with a clear message indicating types need regeneration

---

### Story 1.2: Deployment Infrastructure

As a **developer**,
I want the web app deployed to Vercel and the API + PostgreSQL deployed to Railway with development, staging, and production environments,
So that changes can be tested in deployed environments before reaching production.

**Acceptance Criteria:**

**Given** code is merged to `main`,
**When** the CI deploy job runs,
**Then** the web app auto-deploys to the Vercel staging URL and the API auto-deploys to Railway staging

**Given** a PR is opened,
**When** Vercel processes it,
**Then** a unique preview URL is generated for the web app

**Given** the Railway project is configured with staging and production environments,
**When** the Go API starts in either environment,
**Then** `GET /health` returns `200 OK` at the deployed URL

**Given** Railway staging and production are provisioned,
**When** the API connects,
**Then** it can read from and write to the managed PostgreSQL add-on in both environments

**Given** secrets (Clerk keys, Stripe keys, DB URL),
**When** configured in Railway and Vercel environment dashboards,
**Then** they are available to the apps at runtime and are not present in the git repository

---

### Story 1.3: Design System Foundation

As a **user**,
I want a consistent, warm visual experience across every screen,
So that the app feels calm and trustworthy from the first interaction.

**Acceptance Criteria:**

**Given** the web app loads any page,
**When** the page renders,
**Then** Plus Jakarta Sans (variable weight, loaded from Google Fonts) is used as the typeface with the 6-role scale (display/heading/subheading/body/label/caption) applied correctly

**Given** any page is displayed,
**When** the user views it,
**Then** the app background is warm cream (`#FDFAF7`), body text is warm charcoal (`#3D3230`), and the 14 semantic color tokens from UX-DR1 are accessible via Tailwind utility classes

**Given** a viewport width of 600px or wider,
**When** any app page renders,
**Then** the main content column is centered with `max-w-[600px]` and `mx-auto`, with no content extending beyond 600px

**Given** a viewport width of 599px or less,
**When** the page renders,
**Then** the column fills full width with 16px horizontal padding (Tailwind `px-4`)

**Given** any interactive element (button, input, chip),
**When** focused via keyboard Tab key,
**Then** a visible 2px blush rose focus ring is displayed and no `:focus-visible` outline is suppressed

**Given** any page,
**When** the DOM is inspected,
**Then** a skip-to-content link is present, visually hidden by default, and becomes visible when focused via keyboard

---

### Story 1.4: Email/Password Sign-Up & Login

As a **new user**,
I want to create a Tend account with my email and password,
So that I can access my personal planning data securely.

**Acceptance Criteria:**

**Given** a new visitor submits the sign-up form with a valid email and password,
**When** Clerk processes the registration,
**Then** a Clerk account is created, a user record is inserted into the local `users` table via the Clerk webhook, and the user is redirected to onboarding

**Given** a Clerk `user.created` webhook fires,
**When** the API processes `POST /api/v1/auth/webhook`,
**Then** a row is inserted into `users` with the Clerk `user_id` and no duplicate is created if the webhook fires more than once (idempotent)

**Given** a returning user submits valid login credentials,
**When** Clerk authenticates the session,
**Then** the user is redirected to today's view without seeing the login screen again

**Given** an authenticated user navigates to the app root or landing page,
**When** the Clerk session is valid,
**Then** they are automatically redirected to the today view (FR5)

**Given** an incorrect password is submitted,
**When** Clerk rejects the attempt,
**Then** a clear, non-alarming error message is displayed and no session is created

**Given** any protected API route receives a request,
**When** no valid Clerk JWT is present in the request,
**Then** the API returns `401 Unauthorized`

---

### Story 1.5: Social Authentication — Google & Apple Sign-In

As a **new or returning user**,
I want to sign up or log in with my Google or Apple account,
So that I can access Tend without managing a separate password.

**Acceptance Criteria:**

**Given** a user taps "Continue with Google" on the sign-up or login screen,
**When** they complete the Google OAuth flow,
**Then** they are authenticated via Clerk, a user record is created if new, and they are redirected to onboarding (new) or today's view (returning)

**Given** a user taps "Continue with Apple" on the sign-up or login screen,
**When** they complete the Apple Sign-In flow,
**Then** they are authenticated via Clerk and redirected correctly (same as Google flow above)

**Given** a social sign-in completes for a new user,
**When** the Clerk webhook fires,
**Then** a user record is created in the local `users` table

**Given** a user who previously signed up with email uses Google Sign-In with the same email address,
**When** Clerk processes the sign-in,
**Then** no duplicate user record is created in the local `users` table

**Given** the API receives any request with a Clerk token from a social sign-in,
**When** the Clerk middleware validates it,
**Then** the token is accepted and the request proceeds as authenticated

---

### Story 1.6: Public Landing Page

As a **potential user**,
I want to learn about Tend on a public landing page and create an account or log in,
So that I can start using the app or return to my existing account.

**Acceptance Criteria:**

**Given** a search engine crawls the landing page,
**When** it parses the HTML,
**Then** the page has a descriptive `<title>`, `<meta name="description">`, Open Graph tags, and semantic HTML structure (`<main>`, `<header>`, `<section>`)

**Given** a new visitor lands on the root URL,
**When** the page loads on desktop,
**Then** the value proposition, sign-up CTA ("Get started"), and login link are all visible above the fold without scrolling

**Given** a new visitor clicks "Get started",
**When** redirected,
**Then** they arrive at the Clerk sign-up flow

**Given** a returning visitor clicks "Log in",
**When** redirected,
**Then** they arrive at the Clerk sign-in flow

**Given** an authenticated user navigates to the root URL,
**When** the Clerk session is valid,
**Then** they are redirected to the today view within one render cycle (no flash of landing content)

**Given** any viewport width ≥ 320px,
**When** the landing page renders,
**Then** all content is fully readable and usable without horizontal scrolling

---

### Story 1.7: Log Out & Account Deletion

As a **user**,
I want to log out of my account or permanently delete it,
So that I have full control over my session and personal data.

**Acceptance Criteria:**

**Given** a logged-in user selects "Log out" from the account menu,
**When** the action is confirmed,
**Then** the Clerk session is cleared, local state is reset, and the user is redirected to the landing page

**Given** a user navigates to account settings and taps "Delete account",
**When** they confirm via the destructive confirmation dialog,
**Then** `DELETE /api/v1/account` deletes all user data (todos, goals, daily summaries, notification preferences, user record) in a single database transaction and revokes the Clerk session

**Given** account deletion completes,
**When** the former user attempts to log in with the same credentials,
**Then** login fails and no user data is accessible

**Given** the account deletion confirmation dialog is shown,
**When** the user reads it,
**Then** the "Delete" button uses the destructive button style (outlined, `color-warning` text) and the dialog makes clear the action is permanent

**Given** the deletion API endpoint is called,
**When** the transaction fails partway through,
**Then** the transaction rolls back and no partial data deletion occurs

---

## Epic 2: Goal Setup & Onboarding

New users establish up to 3 long-term goals through guided discovery or direct entry. Returning users can edit or delete goals at any time.

### Story 2.1: Goal Setup — Direct Entry Flow

As a **new user who knows their goals**,
I want to type up to 3 long-term goals during onboarding and have them saved to my account,
So that I can start tagging my daily todos to what actually matters to me.

**Acceptance Criteria:**

**Given** a new user completes sign-up,
**When** they are routed to onboarding,
**Then** they see a goal setup screen with up to 3 text fields and a prompt ("What do you want to make progress on?")

**Given** a user types a goal name and moves to the next field,
**When** the field loses focus,
**Then** the goal is auto-saved to `POST /api/v1/goals` and assigned one of the 3 fixed colors (lavender/sage/peach) based on slot order

**Given** a user submits the onboarding form with at least 1 goal,
**When** the goals are saved,
**Then** they are redirected to today's view and their goals are visible in the goal context area

**Given** a user submits with fewer than 3 goals,
**When** onboarding completes,
**Then** only the filled slots are saved; empty slots are not created

**Given** the `goals` table,
**When** any goal is created or queried,
**Then** all queries are scoped to the authenticated user's `user_id` (never accessible to other users)

**Given** a user reaches the onboarding goal screen,
**When** they are not ready to set goals,
**Then** a "Skip for now" option is visible and functional, routing them directly to today's view

---

### Story 2.2: Guided Goal Discovery Flow

As a **new user who isn't sure what their goals are**,
I want to answer reflective prompts that help me identify what I want to work toward,
So that I can set meaningful goals without feeling stuck or pressured.

**Acceptance Criteria:**

**Given** a user on the onboarding goal screen taps "Help me think",
**When** the discovery flow begins,
**Then** the first reflective prompt is shown alone on screen: "What would you work on if you had a free Saturday?"

**Given** a user submits an answer to a prompt,
**When** the next prompt appears,
**Then** only one prompt is visible at a time (never a multi-field form)

**Given** three prompts are answered ("What would you work on if you had a free Saturday?", "What do you wish you made more time for?", "What would feel meaningful in 6 months?"),
**When** all are submitted,
**Then** suggested goal names derived from the answers are presented for the user to confirm or edit

**Given** the suggested goals are shown,
**When** the user taps "These look right",
**Then** the goals are saved and the user is routed to today's view

**Given** the user wants to edit a suggested goal,
**When** they tap a suggested goal,
**Then** it becomes editable inline before saving

**Given** a user is at any point in the discovery flow,
**When** they tap "Skip",
**Then** they exit the flow and land on the direct goal entry screen (never forced back to discovery)

---

### Story 2.3: Goal Management — Edit & Delete

As a **returning user**,
I want to rename or delete my long-term goals,
So that my goals stay current as my life and priorities change.

**Acceptance Criteria:**

**Given** a user navigates to the Goals page,
**When** the page loads,
**Then** all of their current goals are displayed with their assigned colors (lavender/sage/peach by slot)

**Given** a user taps the edit action on a goal,
**When** the name field becomes editable,
**Then** saving the updated name calls `PATCH /api/v1/goals/:id` and the updated name appears immediately

**Given** a user taps "Delete" on a goal,
**When** they confirm the destructive dialog,
**Then** `DELETE /api/v1/goals/:id` removes the goal and all todos previously tagged to it display as untagged (`goal_id` set to null on those todos)

**Given** a user deletes a goal,
**When** the deletion completes,
**Then** the remaining goals retain their original color assignments (goal slot colors do not shift)

**Given** a user has fewer than 3 goals,
**When** they view the Goals page,
**Then** an "Add a goal" option is available to fill the remaining slot(s)

**Given** any goal API endpoint (`GET`, `PATCH`, `DELETE /api/v1/goals`),
**When** called with a valid auth token,
**Then** only goals belonging to the authenticated user are returned or modified

---

### Story 2.4: GoalChip Component

As a **user**,
I want to see my goals represented as colored chips throughout the app,
So that I can instantly identify which goal a task connects to.

**Acceptance Criteria:**

**Given** the GoalChip is rendered in `pill` variant (day header),
**When** displayed,
**Then** it shows a colored dot (`color-goal-1/2/3`) alongside the goal name text — color is never the sole identifier

**Given** the GoalChip is rendered in `badge` variant (todo row),
**When** displayed on a completed todo,
**Then** the badge remains visible with the goal color and name

**Given** the GoalChip is rendered in `selector` variant (todo creation row),
**When** a user taps a goal chip,
**Then** it shows a selected state (filled background) and only one goal chip can be selected at a time; tapping "None" deselects any active selection

**Given** any GoalChip variant,
**When** rendered in the DOM,
**Then** it has `aria-label="[Goal name] tag"` for screen reader identification

**Given** a user has fewer than 3 goals set,
**When** the selector variant renders,
**Then** only the user's existing goals are shown (no empty slots displayed as selectable)

---

## Epic 3: Daily Planning — Core Loop

Users build a daily todo list with effort sizing (XS–XL), goal tagging, and wellness category tagging. They can complete, edit, reorder, and carry forward todos.

### Story 3.1: Today View Shell & Todo List

As a **user**,
I want to see my todo list for the current day when I open the app,
So that I have an immediate, clear view of what I'm working on today.

**Acceptance Criteria:**

**Given** an authenticated user opens the app,
**When** the today view loads,
**Then** `GET /api/v1/todos?date=today` is called scoped to their `user_id` and their todos for today are displayed in the Direction A Airy Column layout

**Given** the today view loads with todos,
**When** rendered,
**Then** each todo row shows the title and is displayed as a soft surface card (`#FFF8F5`, `rounded-2xl`, 14px padding, 8px gap between rows)

**Given** a user has no todos for today,
**When** the today view renders,
**Then** the empty state message "What do you want to get done today?" is shown with the dashed add row below it (no illustration)

**Given** the today view is loading todos,
**When** the initial fetch is in progress,
**Then** 3 skeleton placeholder rows are shown (shimmer in `color-border`); no full-page spinner

**Given** the page renders on any supported viewport,
**When** the column width is measured,
**Then** the content column does not exceed 600px and is centered on viewports ≥ 600px

---

### Story 3.2: Todo Creation — Inline Flow

As a **user**,
I want to add a todo with a title, effort size, goal tag, and optional wellness category — all inline — in under 10 seconds,
So that planning my day stays fast and frictionless.

**Acceptance Criteria:**

**Given** a user taps the dashed "Add a todo…" row,
**When** the creation row expands,
**Then** it slides open over 200ms (eased, not bouncy) and the title input is immediately focused

**Given** the creation row is open,
**When** the user types a title,
**Then** SizeChip (XS/S/M/L/XL) appears inline with M pre-selected, followed by GoalChip selector (user's goals + None, None pre-selected), followed by WellnessIcon selector (exercise/fun-hobby/rest, none pre-selected)

**Given** the user taps a size chip,
**When** the selection changes,
**Then** the CapacityBar updates immediately to reflect the new planned points (before the todo is saved)

**Given** the user taps Add or presses Return,
**When** the todo is submitted,
**Then** `POST /api/v1/todos` is called with `title`, `effort_size` (XS–XL), `goal_id` (nullable), `wellness_category` (nullable), `date` (today); the todo fades into the list over 150ms; the creation row collapses

**Given** the user presses Escape while the creation row is open,
**When** Escape is detected,
**Then** the row collapses without saving and no API call is made

**Given** the user submits without a title,
**When** the form is validated,
**Then** the submission is blocked and the title input is highlighted; no error modal is shown

---

### Story 3.3: SizeChip & WellnessIcon Components

As a **user**,
I want effort sizes and wellness categories displayed as tappable chips with clear visual weight,
So that I can assign them in a single tap without any extra friction.

**Acceptance Criteria:**

**Given** the SizeChip component renders,
**When** displayed,
**Then** 5 chips (XS/S/M/L/XL) appear in a row with opacity scale on `color-primary`: XS at 30%, S at 50%, M at 70%, L at 85%, XL at 100%

**Given** a SizeChip row is rendered,
**When** inspected for accessibility,
**Then** the group has `role="radiogroup"` and each chip has `role="radio"` and `aria-label="Effort: [size]"`; arrow keys navigate between chips

**Given** the WellnessIcon component renders,
**When** displayed,
**Then** 3 icon buttons (🏃 Exercise / 🎨 Fun / 😴 Rest) are shown, each with its text label alongside the icon

**Given** a WellnessIcon is selected,
**When** tapped again,
**Then** it deselects (toggles off); only one wellness category can be selected at a time

**Given** any chip or icon button,
**When** measured for touch target size,
**Then** the tappable area is at least 44×44px

**Given** `prefers-reduced-motion: reduce` is set in the OS,
**When** any animated transition on these components would normally play,
**Then** the transition falls back to an opacity-only change with no movement

---

### Story 3.4: Todo Completion & Micro-interactions

As a **user**,
I want to check off a todo and see immediate, satisfying feedback,
So that completing tasks feels rewarding without being performative.

**Acceptance Criteria:**

**Given** a user taps the checkbox on an active todo,
**When** the optimistic update fires,
**Then** the checkmark fills `color-success` over 200ms and the todo title fades + strikethrough over 300ms — no loading indicator shown

**Given** the optimistic update fires,
**When** `PATCH /api/v1/todos/:id` resolves with success,
**Then** the completed state is confirmed; if the API call fails, the todo reverts to active state

**Given** a user taps a completed todo's checkbox,
**When** the toggle fires,
**Then** the todo returns to active state (strikethrough removed, opacity restored) and `PATCH /api/v1/todos/:id` is called with `status: "active"`

**Given** `prefers-reduced-motion: reduce` is set,
**When** a todo is completed,
**Then** the visual change happens as an immediate opacity shift with no animated checkmark or strikethrough motion

**Given** the TodoRow renders in `complete` state,
**When** displayed,
**Then** the entire row (including goal badge and size chip) is shown at 55% opacity with strikethrough on the title only

---

### Story 3.5: Todo Edit & Delete

As a **user**,
I want to edit a todo's title, size, goal tag, or wellness category, or delete it entirely,
So that I can keep my day plan accurate as things change.

**Acceptance Criteria:**

**Given** a user taps the edit action on a todo,
**When** the edit mode opens,
**Then** the same inline creation row expands pre-populated with the todo's existing title, size, goal, and wellness values

**Given** a user updates any field and taps Save,
**When** `PATCH /api/v1/todos/:id` is called,
**Then** the todo row updates immediately with the new values and the CapacityBar adjusts if the effort size changed

**Given** a user taps delete on a todo,
**When** they confirm via the destructive dialog,
**Then** `DELETE /api/v1/todos/:id` is called, the todo is removed from the list, and the CapacityBar adjusts

**Given** a user cancels an edit (Escape or Cancel),
**When** the edit row closes,
**Then** the original todo values are restored and no API call is made

**Given** any todo edit or delete API call,
**When** processed by the API,
**Then** the operation is scoped to the authenticated user's `user_id`; a user cannot edit or delete another user's todo

---

### Story 3.6: Todo Reorder

As a **user**,
I want to reorder my todos by dragging (desktop) or long-pressing (mobile),
So that I can prioritize my day visually.

**Acceptance Criteria:**

**Given** a user hovers over a todo row on desktop,
**When** the drag handle appears,
**Then** dragging the handle reorders the todo within the list using optimistic UI

**Given** a user long-presses a todo row on mobile,
**When** the drag mode activates,
**Then** the todo can be dragged to a new position in the list

**Given** a reorder action completes,
**When** the user releases the todo,
**Then** `PATCH /api/v1/todos/:id` is called with the updated `sort_order` value; if the call fails, the list reverts to its previous order

**Given** the TodoRow is in `dragging` state,
**When** rendered,
**Then** it shows a visual elevation (subtle shadow) to indicate it is being moved

---

### Story 3.7: Wellness Requirement & Day Completion State

As a **user**,
I want the app to require at least one wellness todo before my day plan is considered complete,
So that I'm reminded to protect time for exercise, fun, or rest every day.

**Acceptance Criteria:**

**Given** a user's today list has no todos tagged with a wellness category,
**When** they view the today view,
**Then** a soft inline nudge is shown: "Add something for yourself today — exercise, fun, or rest"

**Given** a user adds a todo with a wellness category,
**When** it is saved,
**Then** the wellness nudge is no longer shown for the day

**Given** a user attempts to explicitly wrap up their day (tap "Wrap up today"),
**When** no wellness todo exists in the list,
**Then** the wrap-up is blocked with a gentle message: "Add at least one wellness task before wrapping up"

**Given** the wellness nudge is shown,
**When** rendered,
**Then** it uses `color-warning-bg` background and `color-warning` text — not red, not alarming in tone

---

### Story 3.8: Carry-Forward Prompt

As a **user**,
I want to be shown yesterday's incomplete todos on my first login of the day and choose which to carry forward,
So that nothing slips through the cracks without me consciously deciding to skip it.

**Acceptance Criteria:**

**Given** a user logs in for the first time on a given day,
**When** incomplete todos from the previous day exist,
**Then** the CarryForwardPrompt appears at the top of today's view showing the first incomplete todo with "Keep" and "Skip" options

**Given** the CarryForwardPrompt is shown,
**When** the user taps "Keep",
**Then** the todo is added to today's list via `POST /api/v1/todos` (new record for today's date) and the next incomplete todo from yesterday is shown

**Given** the user taps "Skip",
**When** processed,
**Then** the todo is dismissed without being added to today and the next one is shown

**Given** all yesterday's incomplete todos have been handled,
**When** the last is kept or skipped,
**Then** the CarryForwardPrompt dismisses and today's view is shown normally

**Given** a user dismisses the CarryForwardPrompt partway through,
**When** they re-open the app the same day,
**Then** the prompt does not reappear (handled state is stored)

**Given** the CarryForwardPrompt is visible,
**When** inspected for accessibility,
**Then** focus is trapped within the prompt and keyboard users can navigate Keep/Skip without leaving the component

---

## Epic 4: Capacity Intelligence

The app tracks daily effort history, displays a transparent learning state, and delivers a personalized capacity estimate. Users receive contextual feedback when their plan is over or under capacity and can view their history.

### Story 4.1: Daily Effort Tracking & Learning State

As a **new user**,
I want the app to start learning my daily capacity from day one — and to be transparent that it's doing so,
So that I trust the system before it gives me any estimates.

**Acceptance Criteria:**

**Given** a user completes or uncompletes a todo,
**When** the API processes the change,
**Then** a `daily_summaries` record for today is created or updated with the total `points_completed` for that user

**Given** a user views the today view during their first 1-10 days,
**When** the CapacityBar renders,
**Then** it displays in the `learning` state with the message: "Learning your pace — check back in a few days" and no numeric estimate is shown

**Given** fewer than 5 days of completion data exist for a user,
**When** `GET /api/v1/capacity` is called,
**Then** the API returns a `learning: true` flag and no capacity estimate value

**Given** the `daily_summaries` table,
**When** any summary is written or read,
**Then** all queries are scoped to the authenticated user's `user_id`

**Given** a day passes without the user explicitly wrapping up,
**When** the next day begins,
**Then** the daily summary for the previous day is finalized with whatever `points_completed` value was recorded (no user action required)

---

### Story 4.2: Capacity Estimate & CapacityBar

As a **user who has used Tend for at least 5 days**,
I want the app to generate a personalized daily capacity estimate,
So that I can plan a realistic day based on what I actually accomplish.

**Acceptance Criteria:**

**Given** a user has at least 5 days of `daily_summaries` data,
**When** `GET /api/v1/capacity` is called,
**Then** the API returns a capacity estimate (rolling average of recent `points_completed`) in under 1 second

**Given** a capacity estimate exists,
**When** the CapacityBar renders,
**Then** it displays in `on-track` state showing today's planned points vs. the estimate (e.g., "14 / 18 pts") with a blush gradient fill proportional to usage

**Given** a user adds or removes a todo,
**When** the total planned points change,
**Then** the CapacityBar fill animates smoothly to the new width over 400ms ease-in-out

**Given** the capacity estimate is available,
**When** the CapacityBar is inspected for accessibility,
**Then** it has `role="progressbar"`, `aria-valuenow` (planned points), and `aria-valuemax` (capacity estimate)

**Given** the capacity model runs server-side,
**When** the estimate is calculated,
**Then** it is computed in `capacity_service.go` only — no client-side calculation is performed

---

### Story 4.3: Capacity Reveal Moment

As a **user reaching their first capacity estimate**,
I want a special moment when the app first reveals my personalized capacity,
So that the transition from "learning" to "I know your pace" feels meaningful, not invisible.

**Acceptance Criteria:**

**Given** a user's capacity estimate becomes available (≥5 days of data),
**When** they open the app for the first time after the estimate is ready,
**Then** the CapacityRevealCard is shown: the capacity bar animates in and a specific message is displayed ("We've been paying attention. You tend to accomplish around X points on a good day.")

**Given** the CapacityRevealCard is shown,
**When** the user dismisses it,
**Then** a `capacity_reveal_seen: true` flag is stored for the user and the card never appears again

**Given** the user has already seen the reveal,
**When** they open the app on any subsequent day,
**Then** the CapacityRevealCard is not shown

**Given** the CapacityRevealCard message,
**When** rendered,
**Then** the copy uses "You tend to accomplish around X points" — observational language, never "Your limit is X" or any prescriptive framing

**Given** `prefers-reduced-motion: reduce` is set,
**When** the CapacityRevealCard appears,
**Then** the bar entrance is an opacity fade-in with no animated width expansion

---

### Story 4.4: Over-Capacity & Under-Capacity Feedback

As a **user planning their day**,
I want the app to gently flag when I've planned too much or have room for more,
So that I can make realistic adjustments without feeling pressured.

**Acceptance Criteria:**

**Given** a user's planned points exceed their capacity estimate,
**When** the CapacityBar updates,
**Then** the bar shifts to `over-capacity` state: amber fill (`color-warning`) and an inline note below the bar: "You've planned more than your recent average — that's okay"

**Given** a user's planned points are significantly below their capacity estimate (less than 50%),
**When** the CapacityBar renders,
**Then** a soft encouragement note appears: "You have room for more if you want it"

**Given** the capacity warning or encouragement message,
**When** rendered,
**Then** no red color is used anywhere; amber (`color-warning`) is the maximum visual intensity for warnings

**Given** the over-capacity state is active,
**When** a user removes a todo bringing planned points back under capacity,
**Then** the CapacityBar returns to `on-track` state and the warning message disappears

**Given** the capacity feedback messages,
**When** announced by a screen reader via ARIA live region,
**Then** the spoken text is identical to the displayed text and uses the same observational, permissive tone

---

### Story 4.5: History View

As a **user**,
I want to view my historical daily point completion data,
So that I can see my patterns over time and feel a sense of progress.

**Acceptance Criteria:**

**Given** a user navigates to the History view,
**When** the page loads,
**Then** `GET /api/v1/daily-summaries` returns their historical records and a chart or list shows `points_completed` per day

**Given** the history view renders,
**When** the data is displayed,
**Then** it is presented as a personal record — not as performance data or a comparison to targets

**Given** fewer than 3 days of history exist,
**When** the history view renders,
**Then** a warm empty state message is shown: "Keep going — your history will appear here after your first few days"

**Given** the history data,
**When** queried from the API,
**Then** all records are scoped to the authenticated user's `user_id`; no other user's data is accessible

---

## Epic 5: Reinforcement & Goal Streaks

Users receive warm, specific positive reinforcement on todo completion and strong days. The app tracks goal activity and surfaces nudges when a goal goes untouched. An optional day wrap-up ritual provides a satisfying close.

### Story 5.1: Todo Completion Reinforcement

As a **user**,
I want to receive a warm, specific acknowledgment when I complete a goal-tagged todo,
So that daily progress feels connected to what I actually care about.

**Acceptance Criteria:**

**Given** a user completes a todo that is tagged to one of their goals,
**When** the completion is saved,
**Then** a GoalNudgeMessage appears inline beneath that todo row with a specific message referencing the goal name (e.g., "One more step toward your bakery goal")

**Given** the GoalNudgeMessage appears,
**When** 4 seconds pass or the user takes any next action,
**Then** it fades out over 300ms automatically

**Given** a user completes a todo with no goal tag,
**When** the completion is saved,
**Then** no GoalNudgeMessage is shown (silence is intentional)

**Given** the GoalNudgeMessage,
**When** rendered in the DOM,
**Then** it has `role="status"` and `aria-live="polite"` so screen readers announce it without interrupting the user's current focus

**Given** the reinforcement message copy,
**When** generated,
**Then** it always names the specific goal — never generic text like "Great job!" alone

---

### Story 5.2: Strong-Day Reinforcement

As a **user**,
I want to receive positive reinforcement when I meet or exceed my daily capacity target,
So that completing a full day feels genuinely acknowledged.

**Acceptance Criteria:**

**Given** a user completes their last remaining todo for the day,
**When** their completed points meet or exceed their capacity estimate,
**Then** a warm inline message is shown: "You had a solid day. Everything you planned, done."

**Given** a user completes their todos but completed points are below capacity,
**When** the last todo is checked off,
**Then** no message is shown — partial days are presented as neutral data, not failure

**Given** the strong-day message,
**When** rendered,
**Then** it uses `color-success` left border and warm `color-surface` background — no exclamation points, no fanfare

**Given** a user is still in the learning state (no capacity estimate yet),
**When** they complete all their todos,
**Then** a simple warm message is shown: "You got through your list today" (no capacity comparison made)

---

### Story 5.3: Goal Inactivity Nudge

As a **user**,
I want to be gently reminded when I haven't worked toward a goal in 3 or more days,
So that my long-term goals don't quietly slip out of my daily routine.

**Acceptance Criteria:**

**Given** a user has a goal with no completed goal-tagged todos in the last 3 calendar days,
**When** the goal inactivity check runs in `goal_service.go`,
**Then** an in-app nudge is surfaced on the today view: "[Goal name] hasn't come up lately — want to add something today?"

**Given** the nudge is shown,
**When** the user adds a todo tagged to that goal,
**Then** the nudge disappears for that goal

**Given** the nudge is shown,
**When** the user dismisses it without adding a todo,
**Then** it does not reappear for that goal until another 3 days pass without activity

**Given** the nudge message copy,
**When** rendered,
**Then** it is framed as an observation and invitation — never a guilt statement or instruction (wellness language rule)

**Given** multiple goals are inactive,
**When** the nudge logic runs,
**Then** only one nudge is shown at a time (the most overdue goal); multiple simultaneous nudges are not shown

---

### Story 5.4: Goal-Activity Streaks

As a **user**,
I want to see a streak showing consecutive days I've worked toward each goal,
So that I can feel the momentum of consistent effort without being punished for missing a day.

**Acceptance Criteria:**

**Given** a user has completed at least one goal-tagged todo for a goal on consecutive calendar days,
**When** the streak is calculated in `goal_service.go`,
**Then** `GET /api/v1/goals` returns a `current_streak` value per goal representing consecutive days of activity

**Given** the streak value is returned,
**When** displayed on the Goals page or today view,
**Then** it is shown as an encouraging label (e.g., "3 days in a row") — never as a counter that resets dramatically

**Given** a user misses a day for a goal,
**When** the streak is displayed the next day,
**Then** the streak shows 0 or restarts quietly — no alarming reset animation, no guilt messaging

**Given** a user has a streak of 0 for a goal,
**When** the streak is displayed,
**Then** no streak indicator is shown (absence is neutral, not a failure state)

**Given** streak data,
**When** queried from the API,
**Then** all calculations are scoped to the authenticated user's `user_id`

---

### Story 5.5: Day Wrap-Up (Optional Ritual)

As a **user who wants to close out their day intentionally**,
I want an optional "Wrap up today" action that gives me a warm summary of what I accomplished,
So that I can end my planning day with a sense of closure and satisfaction.

**Acceptance Criteria:**

**Given** a user taps "Wrap up today",
**When** the DaySummaryScreen opens,
**Then** it shows: total points completed, which goals were touched today, and a warm specific sign-off message referencing the day's actual data

**Given** the DaySummaryScreen sign-off message,
**When** generated,
**Then** it is specific to the day (e.g., "You made progress on your bakery goal and got your run in") — never a generic template like "Great work today!"

**Given** the user closes the DaySummaryScreen,
**When** dismissed,
**Then** the today view returns to its normal state; the day is not "locked" — todos can still be added or completed

**Given** the day has no wellness todo when "Wrap up today" is tapped,
**When** the attempt is made,
**Then** the wrap-up is blocked with the gentle message from Story 3.7 ("Add at least one wellness task before wrapping up")

**Given** a user does not tap "Wrap up today",
**When** midnight passes,
**Then** the day closes automatically with no penalty, no missed-summary notification, and the data is recorded identically

---

## Epic 6: Subscription & Billing

New users receive a 30-day free trial with full feature access. Stripe handles payment collection at trial end. Users can view status, cancel, and reactivate their subscription.

### Story 6.1: 30-Day Free Trial

As a **new user**,
I want to access all of Tend's features free for 30 days,
So that I can build a real habit before deciding to subscribe.

**Acceptance Criteria:**

**Given** a new user completes sign-up,
**When** their account is created,
**Then** a trial subscription is created in Stripe via `stripe_service.go` and the local `users` record is updated with `subscription_status: "trialing"` and a `trial_ends_at` timestamp 30 days from now

**Given** a user is in trial,
**When** they access any feature,
**Then** all features are fully accessible — no trial limitations or upsell interruptions during the trial period

**Given** a user views their account or subscription page,
**When** in trial,
**Then** the number of days remaining in the trial is displayed (e.g., "18 days left in your free trial")

**Given** the subscription middleware runs on any protected API route,
**When** the user's `trial_ends_at` is in the future,
**Then** the request proceeds normally (trial counts as active)

---

### Story 6.2: Payment Collection & Subscription Activation

As a **user whose trial is ending**,
I want to be prompted to enter payment details and continue with a paid subscription,
So that I don't lose access to my data and planning history.

**Acceptance Criteria:**

**Given** a user's trial ends,
**When** the Stripe trial-end webhook fires,
**Then** the API processes `POST /api/v1/stripe/webhook` idempotently and updates the user's `subscription_status` to `"past_due"` if no payment method is on file

**Given** a user with `subscription_status: "past_due"` opens the app,
**When** any page loads,
**Then** a non-blocking prompt is shown: "Your trial has ended — add a payment method to keep access" with a CTA to the subscription page

**Given** a user enters their payment details on the subscription page,
**When** Stripe processes the first charge ($5/month),
**Then** the Stripe webhook fires, the API updates `subscription_status` to `"active"`, and the user's access continues uninterrupted

**Given** payment is successfully collected,
**When** the webhook is processed,
**Then** the idempotent handler does not create a duplicate subscription record if the webhook fires more than once

**Given** payment data,
**When** submitted by the user,
**Then** it is handled entirely by Stripe's hosted elements — no raw card data passes through Tend's API servers (NFR10)

---

### Story 6.3: Subscription Management

As a **subscriber**,
I want to view my subscription status and cancel or reactivate my subscription,
So that I stay in full control of my billing relationship with Tend.

**Acceptance Criteria:**

**Given** a subscribed user navigates to the subscription page,
**When** the page loads,
**Then** `GET /api/v1/subscription` returns their current status (`trialing`, `active`, `canceled`, `past_due`) and next billing date, displayed clearly

**Given** an active subscriber taps "Cancel subscription",
**When** they confirm via the destructive dialog,
**Then** `POST /api/v1/subscription/cancel` calls Stripe to cancel at period end and the UI updates to show "Access until [date]"

**Given** a canceled subscriber still within their paid period,
**When** they access the app,
**Then** all features remain accessible until the period end date

**Given** a canceled subscriber taps "Reactivate subscription",
**When** confirmed,
**Then** `POST /api/v1/subscription/reactivate` calls Stripe to resume the subscription and `subscription_status` returns to `"active"`

**Given** the subscription page,
**When** rendered,
**Then** the current plan ($5/month), status, and next billing date are all visible; no upsell or upgrade options are shown (single plan)

---

### Story 6.4: Billing Notifications & Failed Payments

As a **subscriber**,
I want to be notified before I'm charged and when a payment fails,
So that I'm never surprised by a charge or locked out without warning.

**Acceptance Criteria:**

**Given** a user's billing date is 3 days away,
**When** the Stripe upcoming-invoice webhook fires,
**Then** the API logs the event and an in-app notification is queued: "Your $5 renewal is coming up on [date]"

**Given** a payment attempt fails,
**When** the Stripe `invoice.payment_failed` webhook fires,
**Then** the API updates `subscription_status` to `"past_due"` and the in-app prompt from Story 6.2 is shown

**Given** all Stripe webhook events,
**When** processed by `POST /api/v1/stripe/webhook`,
**Then** each event is validated against the Stripe signature header before processing; invalid signatures return `400`

**Given** a Stripe webhook fires more than once for the same event,
**When** the handler processes it,
**Then** the second and subsequent calls are no-ops (idempotent processing via event ID deduplication)

**Given** a user's subscription lapses (payment repeatedly fails),
**When** Stripe marks the subscription as `canceled`,
**Then** the Stripe webhook updates `subscription_status` to `"canceled"` and the user sees a prompt to resubscribe on next login

---

## Epic 7: Push Notifications

Users receive configurable push notifications for morning planning reminders, capacity alerts, goal inactivity nudges, and strong-day reinforcement.

### Story 7.1: Notification Preferences

As a **user**,
I want to configure which push notifications I receive and when,
So that Tend's nudges feel helpful rather than intrusive.

**Acceptance Criteria:**

**Given** a user navigates to notification settings,
**When** the page loads,
**Then** `GET /api/v1/notification-preferences` returns their current preferences and each notification type is shown with an on/off toggle and (where applicable) a time picker

**Given** a user toggles a notification type on or off,
**When** the switch is changed,
**Then** `PATCH /api/v1/notification-preferences` is called immediately and the change persists across sessions

**Given** a user sets a morning reminder time,
**When** the time picker is used,
**Then** the selected time is saved and the morning notification will fire at that time in the user's local timezone

**Given** a user disables all notifications,
**When** they return to the settings page,
**Then** all toggles are shown in the off state and no push notifications are sent for that user

**Given** the notification preferences,
**When** stored,
**Then** they are scoped to the authenticated user's `user_id`

---

### Story 7.2: Morning Planning Reminder

As a **user**,
I want to receive a push notification in the morning reminding me to plan my day,
So that building my daily plan becomes a consistent habit.

**Acceptance Criteria:**

**Given** a user has the morning reminder enabled with a configured time,
**When** that time arrives in the user's timezone,
**Then** `notification_service.go` dispatches a push notification via APNs (iOS) or FCM (Android) with the message: "Time to plan your day — what do you want to get done?"

**Given** the push notification is dispatched,
**When** the APNs/FCM call is made,
**Then** it runs in a goroutine and any delivery failure is logged via `slog.Error` without blocking the API response or any other operation (NFR20)

**Given** a user taps the morning reminder notification,
**When** the app opens,
**Then** they are taken directly to the today view

**Given** push notification copy,
**When** written,
**Then** it is meaningful without visual context — readable as plain text on the lock screen (NFR17)

---

### Story 7.3: Capacity & Goal Nudge Notifications

As a **user**,
I want push notifications when my plan exceeds capacity or a goal goes untouched,
So that I'm prompted to adjust my day even when I'm not actively in the app.

**Acceptance Criteria:**

**Given** a user's planned points exceed their capacity estimate,
**When** the over-capacity state is detected after a todo is added,
**Then** a push notification is dispatched (if the preference is enabled): "You've planned more than usual today — feel free to remove something"

**Given** a goal has had no completed goal-tagged todos in 3+ days,
**When** the daily inactivity check runs in `goal_service.go`,
**Then** a push notification is dispatched (if enabled): "[Goal name] hasn't come up lately — want to add something today?"

**Given** either notification type,
**When** the user has that notification type disabled in their preferences,
**Then** no notification is sent regardless of the trigger condition

**Given** a goal inactivity notification is sent,
**When** the user taps it,
**Then** the app opens to the today view with that goal's pill subtly highlighted

**Given** the capacity notification,
**When** dispatched,
**Then** the copy uses permissive, observational language — never alarming or prescriptive (wellness language rule, NFR17)

---

### Story 7.4: Strong-Day Reinforcement Notification

As a **user**,
I want a push notification when I complete a strong day,
So that positive reinforcement reaches me even after I've closed the app.

**Acceptance Criteria:**

**Given** a user's completed points meet or exceed their capacity estimate for the day,
**When** the final todo is completed,
**Then** a push notification is dispatched (if enabled): "You had a solid day — everything you planned, done"

**Given** a user is still in the learning state (no capacity estimate),
**When** they complete all their todos,
**Then** no strong-day push notification is sent (no estimate means no comparison can be made)

**Given** the strong-day notification,
**When** dispatched,
**Then** it runs in a goroutine and delivery failure is logged without blocking the completion API response

**Given** a user taps the strong-day notification,
**When** the app opens,
**Then** they are taken to the today view (no special screen — the moment already passed)

---

## Epic 8: Mobile App & Real-Time Sync

Users have full access to Tend on iOS and Android with offline-first support. Changes sync across platforms within 5 seconds via SSE. The mobile app mirrors all web features with a touch-optimized interface.

### Story 8.1: React Native App Foundation

As a **user**,
I want to access Tend as a native iOS and Android app,
So that I can plan and check off todos from my phone with a native feel.

**Acceptance Criteria:**

**Given** the Expo app is opened on iOS or Android,
**When** it launches,
**Then** the app starts within 2 seconds on mid-range devices (NFR2) and displays the Tend visual identity (blush palette, Plus Jakarta Sans via NativeWind)

**Given** a user opens the mobile app for the first time,
**When** no Clerk session exists,
**Then** they are routed to the sign-in/sign-up screen with email/password and social auth options (same auth methods as web)

**Given** an authenticated user opens the mobile app,
**When** the app loads,
**Then** they are routed to today's view and their todos are fetched from `GET /api/v1/todos?date=today`

**Given** the mobile app layout,
**When** rendered on any screen,
**Then** the bottom tab bar (Today / Goals / History / Settings) is always visible and the active tab uses `color-primary`

**Given** the React Native app,
**When** any interactive element is rendered,
**Then** it has `accessibilityLabel`, `accessibilityRole`, and meets the 44pt minimum touch target size (NFR16, UX-DR18)

---

### Story 8.2: Full Feature Parity on Mobile

As a **mobile user**,
I want access to all the same features I have on the web — goal management, daily planning, capacity feedback, reinforcement, and settings,
So that I can use Tend fully from my phone without needing to switch to a browser.

**Acceptance Criteria:**

**Given** a mobile user navigates to the Today tab,
**When** the view renders,
**Then** the full daily planning flow works identically to web: todo creation (with inline SizeChip, GoalChip, WellnessIcon), completion, edit, delete, reorder via long-press drag, and carry-forward prompt on first daily open

**Given** a mobile user navigates to the Goals tab,
**When** the view renders,
**Then** they can view, edit, and delete goals with the same behavior as web

**Given** a mobile user navigates to the History tab,
**When** the view renders,
**Then** their daily point completion history is displayed

**Given** a mobile user navigates to Settings,
**When** the view renders,
**Then** they can manage notification preferences, view subscription status, and access account settings (log out, delete account)

**Given** the CapacityBar, GoalNudgeMessage, CarryForwardPrompt, DaySummaryScreen, and CapacityRevealCard components,
**When** rendered in the React Native app,
**Then** they use NativeWind equivalents of the web Tailwind styles and behave identically to their web counterparts

---

### Story 8.3: Offline-First Todo Access

As a **mobile user without network access**,
I want to view and check off todos while offline,
So that my planning routine isn't interrupted by poor connectivity.

**Acceptance Criteria:**

**Given** a user has previously loaded today's todos while online,
**When** they lose network connectivity,
**Then** today's todos and goals remain fully visible and interactive using TanStack Query's `persistQueryClient` + AsyncStorage cache

**Given** a user completes, adds, or edits a todo while offline,
**When** the mutation is attempted,
**Then** the change is applied optimistically to the local cache and queued in `offline-queue-store.ts` for later sync

**Given** the device regains network connectivity,
**When** the app detects the reconnect,
**Then** all queued mutations are flushed to the API in order and the cache is invalidated to pull fresh server state

**Given** an offline mutation conflicts with a server change (same todo edited on both platforms),
**When** the sync resolves the conflict,
**Then** last-write-wins is applied for todo status; capacity model data defers to server-authoritative values (NFR21)

**Given** the app is offline,
**When** the user views the today view,
**Then** a small pulsing indicator in the top bar signals offline mode — no alarming banner, no blocking UI

---

### Story 8.4: Real-Time Cross-Platform Sync via SSE

As a **user on multiple devices**,
I want changes I make on one platform to appear on all my other connected devices within 5 seconds,
So that my web and mobile apps always show the same state.

**Acceptance Criteria:**

**Given** an authenticated user is connected on both web and mobile,
**When** a todo is created, updated, or completed on one platform,
**Then** the SSE stream (`GET /api/v1/stream`) pushes a `todo.updated` or `todo.created` event to all other connected clients for that user within 5 seconds (NFR4)

**Given** an SSE event is received by a client,
**When** the event is processed,
**Then** TanStack Query invalidates `['todos', userId, date]` and refetches; the updated list appears without a full page reload

**Given** the SSE connection drops,
**When** the `EventSource` detects the disconnect,
**Then** it attempts automatic reconnection; the user sees the small pulsing offline indicator but is not shown an error

**Given** 3 or more consecutive SSE reconnection failures,
**When** the client cannot re-establish the stream,
**Then** the client falls back to polling on the next user action (no user-visible error)

**Given** a capacity recalculation occurs on the server,
**When** the `capacity.recalculated` SSE event is pushed,
**Then** the CapacityBar updates on all connected clients without requiring a page refresh

---

### Story 8.5: EAS Build & App Store Submission

As a **developer**,
I want the Expo app built via EAS Build and submitted to the App Store and Google Play,
So that users can download Tend from their platform's app store.

**Acceptance Criteria:**

**Given** the Expo app is configured with `app.json` (bundle ID, version, app name "Tend"),
**When** `eas build --platform all` is run,
**Then** EAS produces an `.ipa` (iOS) and `.aab` (Android) build without errors

**Given** the iOS build,
**When** submitted to App Store Connect,
**Then** it passes Apple's automated pre-review checks (icon, splash screen, required permissions declared)

**Given** the Android build,
**When** submitted to Google Play Console,
**Then** it passes Google's automated review requirements (target SDK, permissions, signing)

**Given** the app store listings,
**When** published,
**Then** the app name is "Tend", the description communicates the core value proposition, and screenshots represent the actual app UI
