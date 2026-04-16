---
stepsCompleted:
  - step-01-init
  - step-02-discovery
  - step-02b-vision
  - step-02c-executive-summary
  - step-03-success
  - step-04-journeys
  - step-05-domain
  - step-06-innovation
  - step-07-project-type
  - step-08-scoping
  - step-09-functional
  - step-10-nonfunctional
  - step-11-polish
  - step-12-complete
inputDocuments: []
workflowType: 'prd'
classification:
  projectType: full-stack (web app + mobile app + API backend + database)
  domain: Productivity & Wellness
  complexity: medium
  projectContext: greenfield
---

# Product Requirements Document - toDoFastTrack

**Author:** Evelynzouras
**Date:** 2026-04-15

## Executive Summary

toDoFastTrack is a cross-platform productivity and wellness application (web + mobile) that bridges the gap between long-term aspirations and daily task management. It targets individuals who pursue meaningful goals but chronically overcommit, burn out, or lose sight of their bigger ambitions in the noise of daily obligations. The product solves a self-worth problem masquerading as a productivity problem: users feel perpetually behind not because they lack discipline, but because their task lists are divorced from reality and from what genuinely matters to them.

The app anchors every user experience around three long-term personal goals, uses an adaptive capacity model to learn realistic daily output limits, and protects time for exercise, rest, and enjoyment — treating these not as luxuries but as requirements for sustainable performance.

**Project Classification:** Full-stack (Web App + Mobile App + API Backend + Database) · Productivity & Wellness · Medium Complexity · Greenfield

### What Makes This Special

Most productivity tools amplify the anxiety spiral: more features, more tasks, more ways to feel behind. toDoFastTrack inverts this. The core differentiator is **honest capacity management** — the app learns how many effort points a user can realistically accomplish in a day and pushes back when the list exceeds that ceiling. Users receive permission, backed by data, to do less and feel good about it.

A hybrid effort-scoring system (user-assigned + app-suggested based on historical patterns) removes estimation friction while improving accuracy over time. Goal-tagging on individual todos creates a visible thread between daily actions and long-term ambitions, making progress feel meaningful rather than mechanical. Phase 2 will add social media and screen time awareness to surface where time quietly disappears before the day begins.

The delight moment: weeks in, users realize they're hitting their goals *and* enjoying their lives — not despite the constraints the app imposes, but because of them.

## Success Criteria

### User Success

- Within 2-3 days of use, users receive their first capacity warning or encouragement — sufficient data to begin reflecting realistic daily output
- Users actively adjust their daily list based on app feedback (removing overloaded tasks, adding items when under capacity)
- Point budgets include non-work categories (exercise, hobbies, rest) in every day's plan — treated as non-negotiable allocations, not afterthoughts
- Users receive positive reinforcement when completing todos and reaching capacity targets, reinforcing habit formation
- Within 2-4 weeks, users report feeling more balanced — daily todos consistently include progress toward at least one long-term goal alongside personal wellness tasks
- Goal-alignment gaps are surfaced: if a user's todos haven't touched a long-term goal in 3+ days, the app flags it

### Business Success

- **3 months:** 100 paying subscribers ($500 MRR) — product-market fit signal
- **6 months:** 400 paying subscribers ($2,000 MRR) — growth trajectory confirmed
- **12 months:** 1,000 paying subscribers ($5,000 MRR / $60,000 ARR)
- Monthly churn below 5%
- Subscription model: $5/month flat rate, 30-day free trial

### Technical Success

- Seamless data sync across web and mobile — changes appear on all connected platforms within 5 seconds
- Capacity model accuracy improves measurably in weeks 3-4 vs. week 1
- 99.9% uptime — downtime breaks daily habit routines and erodes trust
- All user data encrypted at rest and in transit

### Measurable Outcomes

- Daily active usage rate >60% among paying subscribers
- >70% of daily todo lists include at least one goal-tagged task
- >80% of daily lists include at least one wellness/fun/exercise point allocation
- Overexertion alert acted-upon rate >50% (users remove tasks after warning)

## User Journeys

### Journey 1: Laura — The Developer with a Dream

**Persona:** Laura, 45, senior software developer. Her real dream is opening a vegan bakery — she's been "working toward it" for three years, but her todo lists are dominated by work tasks and the bakery keeps slipping to "someday."

**Opening Scene:** Laura downloads toDoFastTrack on a Sunday night, skeptical — she's tried every productivity app and they all make her feel more behind. During onboarding, the app asks her to name three long-term goals. She types: *Open a vegan bakery. Learn bread scoring techniques. Get to a healthy weight.* For the first time, her goals are written down somewhere that connects to her daily life.

**Rising Action:** Monday morning she adds 14 todos. The app flags: *"That's 38 points — you've only completed an average of 22 points on past Mondays. Consider removing some tasks."* Annoyed, she removes three work tasks and adds *"Watch one bread scoring video (2pts)"* — her first bakery task in weeks. A wellness allocation sits unfilled; the app nudges her to add something for exercise or fun.

**Climax:** By Thursday of week two, Laura notices she's been adding at least one bakery task every day. Not because she forced it — because the app asks *"does this move you toward your goals?"* when she creates todos. The question itself is the intervention. She completes her fourth bakery task of the week and the app says: *"You're on a streak — 4 days working toward 'Open a vegan bakery.' Keep going, Laura."* She screenshots it and sends it to her sister.

**Resolution:** Six weeks in, Laura has a realistic picture of her daily capacity (20-24 points on workdays, 30+ on weekends). Her bakery goal gets touched 4-5 days a week. She hasn't opened the bakery yet — but for the first time, she believes she actually will. She's also exercising three times a week because the app wouldn't let her ignore it.

*Capabilities revealed: goal onboarding, goal-tagging, capacity model, overload warnings, wellness allocation, streak tracking, positive reinforcement.*

---

### Journey 2: Julie — The Reporter Racing the Clock

**Persona:** Julie, 30, journalist at a fast-paced news outlet. Two ambitious goals: write a novel and run a marathon. Her current todo app has 47 items. Her manuscript is at chapter three, where it's been for four months.

**Opening Scene:** Julie hears about toDoFastTrack from a podcast. During onboarding she types: *Write my novel. Run a marathon. Be present in my relationships.* She doesn't need the guided goal discovery — she's known her goals for years. The problem is the gap between knowing and doing.

**Rising Action:** Her first daily list has 22 items. The app flags overload immediately. She ignores it and marks 9 things done — zero novel tasks, zero running. Day two: *"Yesterday you completed 18 points. Today you have 34 — want to trim it down?"* She removes overflow tasks and adds *"Write 200 words of novel (3pts)"* and *"20 min run (4pts)."* Both get done. First time in months both happened in the same day.

**Climax:** Week three. A major story breaks and her work todos explode. The app flags that she's on track to go three days without touching her book or running. It doesn't lecture her — it surfaces the gap: *"You haven't worked toward 'Write my novel' in 4 days."* She carves out 20 minutes at 6am. 200 words. Enough.

**Resolution:** Three months in, Julie is at chapter seven and has a half-marathon on the calendar. She still has brutal news weeks — but she no longer lets them consume the whole person. The app taught her that something is always better than nothing.

*Capabilities revealed: goal-alignment nudges, capacity warnings, overload alerts, cross-goal balancing, progress tracking.*

---

### Journey 3: Alex — "I Don't Know My Goals" (Edge Case)

**Persona:** Alex, 38, feeling stuck and burned out. Knows something needs to change but can't articulate what they actually want.

**Journey:** Alex opens toDoFastTrack and stares at the goal-setting prompt. Blank. The app offers guided discovery — *"What would you regret not doing in 5 years? What do you do that makes you lose track of time? What do you wish you had more of?"* Alex lands on: *Get healthier. Reconnect with photography. Spend more time offline.* None of these felt like "goals" before — but seeing them written down makes them real.

*Capabilities revealed: guided goal discovery flow, reflection prompts, onboarding flexibility.*

---

### Journey 4: Platform Admin (Operator)

**Persona:** Evelynzouras, the founder managing the product.

**Journey:** Logs in to monitor subscription health (MRR, churn, new signups), review support issues, and manage user accounts. Views aggregate usage patterns — are users engaging with wellness allocation? Is the capacity model being used? Individual user data is never surfaced.

*Capabilities revealed: admin dashboard, subscription management, aggregate analytics, account management.*

---

### Journey Requirements Summary

| Capability | Revealed By |
|---|---|
| Goal onboarding (known + guided discovery) | Laura, Alex |
| Daily todo creation with point assignment | Laura, Julie |
| Goal-tagging per todo | Laura, Julie |
| Wellness point allocation (exercise, fun, rest) | Laura |
| Capacity model + overload warnings | Laura, Julie |
| Positive reinforcement / compliments | Laura |
| Goal-alignment nudges ("X days since you worked toward Y") | Julie, Laura |
| Streak tracking | Laura |
| Cross-platform sync (web + mobile) | Laura, Julie |
| Admin dashboard + subscription management | Admin |

## Domain-Specific Requirements

### Compliance & Regulatory

- **Global privacy (GDPR + CCPA baseline):** Users have the right to export and delete all their data. Privacy policy required at launch. Consent obtained at account creation.
- **Data retention:** User data retained for subscription duration + 90 days post-cancellation (reactivation window). Deleted on user request within 30 days.
- **Payment compliance:** Payments processed via Stripe. No raw card data stored or transmitted through application servers. PCI-DSS compliance delegated to payment processor.

### Technical Constraints

- **Data privacy:** Personal goals, task history, and capacity data encrypted at rest and in transit. No third-party data sharing.
- **Wellness positioning:** All app messaging is observational and factual — never clinical, diagnostic, or prescriptive. Example: *"You completed X points over 5 days — your list of Y points may be optimistic."*
- **Account deletion:** Full data purge, user-initiated, completed within 30 days.

### Pre-Launch Requirements

- Privacy policy and terms of service required before launch
- Wellness language guidelines (internal copy standards) established before launch
- Stripe integration handles recurring billing and failed payment retries

## Innovation & Novel Patterns

### Detected Innovation Areas

**1. Adaptive Capacity Learning**
toDoFastTrack learns a user's *actual* daily output over time — not what they plan to do, but what they consistently complete — and uses that as a personalization baseline. The app shifts from a passive list to an active reality-check engine. This behavioral feedback loop applied to personal capacity management is uncommon in consumer productivity tools.

**2. The Permission Inversion**
Productivity apps typically make users feel behind. toDoFastTrack is designed to do the opposite: it tells users when they're *over-planning* and gives data-backed permission to do less. The app's job is to reduce the list, not grow it.

**3. Wellness as a Non-Negotiable Budget**
Exercise, rest, and fun are required point allocations in every day's plan — not optional. The app treats life balance as a system constraint, not a nice-to-have.

**4. Goal-Reality Thread**
Long-term goals are woven into every daily task via tagging. The app actively monitors when goals go untouched and surfaces the gap in real time.

### Market Context & Competitive Landscape

Existing productivity tools split into simple task lists (Todoist, Things, Apple Reminders) and complex project management (Notion, Asana, Linear). Neither addresses personal capacity learning or goal-reality alignment. The closest analogues — Habitica (gamification), Finch (wellness), time-blocking tools — each solve one piece but not the whole. toDoFastTrack's combination of adaptive capacity + goal threading + wellness budgeting occupies a distinct position.

### Innovation Validation Metrics

- Capacity model accuracy: does point estimate precision improve week-over-week?
- Goal touch rate: do users who tag goals retain at higher rates?
- Wellness allocation compliance: do users fill wellness categories daily, and does it correlate with retention?
- Alert response rate: do users remove tasks after overexertion warnings?

## Full-Stack Specific Requirements

### Architecture Overview

toDoFastTrack consists of three layers sharing a single data model: a React SPA (web), a React Native mobile app (iOS + Android), and a shared REST API backend with PostgreSQL database. User account, goals, todos, and capacity history are consistent across all surfaces.

### Web Application

- **Architecture:** Single-Page Application (React) — no full page reloads after initial load
- **Landing page:** SEO-optimized, publicly indexed — clear value proposition, sign-up CTA, login option
- **Session detection:** Returning authenticated users redirected to app automatically; valid unexpired tokens trigger auto-login
- **Authenticated app:** Full todo/goal interface behind authentication — not publicly indexed
- **Browser support:** Chrome, Firefox, Safari, Edge (modern versions) — no IE
- **Responsive design:** Fully usable on mobile browsers
- **Performance:** Initial load under 3 seconds; SPA transitions under 300ms

### Mobile Application

- **Platform:** React Native — single codebase, iOS + Android
- **Store compliance:** Apple App Store and Google Play Store requirements met at launch
- **Push notifications:** Morning planning reminder (configurable), overexertion alert, goal nudge, positive reinforcement
- **Offline:** Offline-first — today's todos and goals stored locally (AsyncStorage/SQLite). Syncs automatically on reconnect. Conflict resolution: last-write-wins for todo status; server authoritative for capacity model data.

### Backend & API

- **Architecture:** Stateless REST API serving web and mobile clients
- **Authentication:** Email/password (bcrypt), Google Sign-In, Apple Sign-In; JWT with refresh token rotation; cookie-based sessions for web, token-based for mobile
- **Database:** PostgreSQL — users, goals, todos, daily point history, capacity model snapshots
- **Sync:** Endpoints support delta sync (changes since last sync timestamp) for efficient mobile updates
- **Subscriptions:** Stripe webhooks for trial start, payment success/failure, cancellation — idempotent processing

### Implementation Considerations

- **Monorepo:** Shared types, utilities, and API client between web and React Native
- **Capacity model:** Server-side only — consistent across devices, survives app reinstalls
- **Environments:** Development, staging, and production required before launch

## Project Scoping & Phased Development

### MVP Strategy

**Approach:** Experience MVP — users feel core value (realistic planning + goal progress) within their first week. 30-day free trial converts to $5/month.

**Constraint:** Solo developer — scope discipline is critical. Web launches first; mobile follows 4-6 weeks later.

### Phase 1 — MVP Feature Set

**Must-Have Capabilities:**

- **Onboarding:** Account creation (email/password + Google/Apple Sign-In), guided goal-setting for up to 3 long-term goals, optional guided discovery flow for undecided users
- **Daily Planning:** Todo creation with user-set effort points, goal-tagging per todo, wellness category allocation (exercise, fun/hobby, rest) required before day plan is complete
- **Capacity Model:** Tracks completed points daily; "learning" state for first 7-10 days with transparent messaging; capacity estimates and overload warnings from day 5+ onward
- **Feedback & Reinforcement:** Overexertion alerts, positive reinforcement on completion and goal streaks, goal-alignment nudges after 3+ days without touching a goal
- **Cross-Platform:** React SPA (web) + React Native (iOS + Android) + shared REST API + PostgreSQL
- **Auth & Sessions:** JWT tokens, cookie sessions, auto-login for returning web users
- **Push Notifications:** Morning planning reminder, overexertion alert, goal nudge, positive reinforcement
- **Offline Support:** Local storage of today's plan, auto-sync on reconnect
- **Subscription:** 30-day free trial, $5/month via Stripe, failed payment handling
- **SEO Landing Page:** Public marketing page, sign-up and login

**Excluded from MVP:**
- App-suggested point estimates (Phase 2 — requires history data)
- Social media / screen time tracking (Phase 2)
- Admin analytics dashboard (Stripe dashboard sufficient at early scale)

### Phase 2 — Growth

- App-suggested effort point estimates based on historical patterns
- Social media / screen time awareness integration
- Goal progress visualization (weekly/monthly breakdown by goal)
- Streak tracking and milestone celebrations
- Admin dashboard (DAU, retention, capacity model accuracy)

### Phase 3 — Expansion

- AI-driven daily plan suggestions
- Calendar integration
- Community / accountability features (optional goal-sharing)
- Coaching-style insights: patterns, burnout risk trends, personalized recommendations

### Risk Mitigation

**Technical:**
- *Offline sync complexity (solo dev):* Web-first launch; mobile to TestFlight/Play internal testing in parallel; public mobile release 4-6 weeks post web launch
- *Capacity cold start:* "Learning" state shown for first 10 days; no warnings until 5 days of data exist
- *App Store review delays:* Submit 2-3 weeks before planned launch date

**Market:**
- *Habit abandonment:* Morning planning push notification is MVP-required, not a growth feature; 30-day trial provides full habit loop before payment
- *Competitor response:* Wellness-as-constraint philosophy and goal-reality threading are the moat — ship fast, build retention

**Resource:**
- *Solo velocity:* Web-first launch, mobile follows; admin tooling deferred to Phase 2

## Functional Requirements

### User Account Management

- FR1: Users can create an account with email and password
- FR2: Users can create an account using Google Sign-In
- FR3: Users can create an account using Apple Sign-In
- FR4: Users can log in to an existing account
- FR5: Returning authenticated users are automatically redirected to the app from the landing page
- FR6: Users can log out of their account
- FR7: Users can delete their account and all associated data
- FR8: Users can manage their subscription (view status, cancel, reactivate)

### Onboarding & Goal Setup

- FR9: New users are guided through a goal-setting flow to establish up to 3 long-term goals
- FR10: Users who don't know their goals can access a guided discovery flow that prompts reflection questions to help identify them
- FR11: Users can name, edit, and delete their long-term goals at any time
- FR12: Users can skip the guided discovery flow and set goals manually

### Daily Planning

- FR13: Users can create todos for the current day
- FR14: Users can assign an effort point value to each todo
- FR15: Users can tag each todo with one of their long-term goals (or mark it as untagged)
- FR16: Users can assign each todo to a wellness category (exercise, fun/hobby, rest)
- FR17: Users can mark todos as complete
- FR18: Users can edit or delete todos
- FR19: Users can reorder todos within their daily list
- FR20: Users must include at least one wellness-category todo before their day plan is considered complete
- FR21: Users can carry incomplete todos forward to the next day

### Capacity Management

- FR22: The system tracks effort points completed by a user each day over time
- FR23: The system displays a "learning" state for the first 7-10 days with messaging explaining it is building a capacity baseline
- FR24: The system generates a daily capacity estimate once sufficient history exists (minimum 5 days)
- FR25: The system alerts users when planned daily points exceed their capacity estimate
- FR26: The system encourages users to add tasks when planned total is significantly below capacity estimate
- FR27: Users can view their historical daily point completion data

### Feedback & Reinforcement

- FR28: The system delivers positive reinforcement messages when a user completes a todo
- FR29: The system delivers positive reinforcement when a user meets or exceeds their daily capacity target
- FR30: The system notifies users when a long-term goal has not been touched in 3 or more days
- FR31: The system tracks and displays goal-activity streaks (consecutive days with at least one goal-tagged todo completed)

### Notifications

- FR32: Users receive a configurable morning push notification to plan their day
- FR33: Users receive a push notification when their planned list exceeds their capacity estimate
- FR34: Users receive a push notification when a long-term goal has not been worked on for 3+ days
- FR35: Users receive a push notification with positive reinforcement when they complete a strong day
- FR36: Users can configure notification preferences (enable/disable each type, set notification time)

### Cross-Platform & Sync

- FR37: Users can access their account, goals, and todos from both the web app and the mobile app
- FR38: Changes made on one platform appear on other connected platforms within 5 seconds
- FR39: Users can view and interact with their todo list while offline on mobile
- FR40: Changes made offline sync automatically to the server when connectivity is restored

### Subscription & Billing

- FR41: New users receive a 30-day free trial with full feature access
- FR42: Users are prompted to enter payment details at trial end
- FR43: The system charges users $5/month via Stripe upon trial completion
- FR44: Users receive notification of upcoming billing and failed payment attempts
- FR45: Users can cancel their subscription and retain access until the end of the billing period

### Landing Page & Discoverability

- FR46: The public landing page is optimized for search engine indexing
- FR47: The landing page allows new visitors to sign up
- FR48: The landing page allows existing users to log in
- FR49: The landing page detects authenticated returning users and redirects them to the app

## Non-Functional Requirements

### Performance

- Initial web app load: under 3 seconds on standard broadband
- React Native app launch: under 2 seconds on mid-range devices
- UI response to user actions (todo creation, completion, point updates): under 300ms
- Cross-platform sync latency: changes appear on connected platforms within 5 seconds
- Capacity model calculation: server response under 1 second after day plan submission

### Security

- All user data encrypted at rest (AES-256) and in transit (TLS 1.2+)
- Passwords stored using bcrypt — plaintext passwords never stored or logged
- JWT tokens expire after 24 hours; refresh tokens rotate on use
- User data is private and isolated — no cross-user data access possible
- Payment data never touches application servers — fully delegated to Stripe
- Account deletion purges all user data within 30 days of request

### Scalability

- System supports up to 10,000 users without architectural changes (10x runway beyond 12-month target)
- Database queries optimized for per-user access patterns — no global table scans in hot paths
- Stateless API — horizontally scalable post-MVP if needed

### Accessibility

- Web app meets WCAG 2.1 AA — keyboard navigable, screen reader compatible, sufficient color contrast
- Mobile app follows iOS Human Interface Guidelines and Android Material accessibility guidelines
- Push notification content meaningful without visual context

### Integration

- **Stripe:** Subscription lifecycle webhooks (trial start, payment success/failure, cancellation) processed idempotently
- **Google Sign-In / Apple Sign-In:** OAuth 2.0 tokens validated server-side
- **Push notifications (APNs / FCM):** Delivery failures logged; non-blocking to core app function
- **Offline sync:** Last-write-wins for todo status; capacity model data server-authoritative on reconnect
