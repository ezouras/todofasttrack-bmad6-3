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

### What Makes This Special

Most productivity tools amplify the anxiety spiral: more features, more tasks, more ways to feel behind. toDoFastTrack inverts this. The core differentiator is **honest capacity management** — the app learns how many effort points a user can realistically accomplish in a day and pushes back when the list exceeds that ceiling. Users receive permission, backed by data, to do less and feel good about it.

A hybrid effort-scoring system (user-assigned + app-suggested based on historical patterns) removes the friction of estimation while improving accuracy over time. Goal-tagging on individual todos creates a visible thread between daily actions and long-term ambitions, making progress feel meaningful rather than mechanical. Optional social media and screen time awareness surfaces where time quietly disappears before the day begins — giving users a complete picture of their actual capacity.

The delight moment: weeks in, users realize they're hitting their goals *and* enjoying their lives — not despite the constraints the app imposes, but because of them.

## Project Classification

- **Project Type:** Full-stack — Web App + Mobile App + API Backend + Database
- **Domain:** Productivity & Wellness
- **Complexity:** Medium
- **Project Context:** Greenfield

## Success Criteria

### User Success

- Within **2-3 days** of use, users receive their first capacity warning or encouragement — the app has enough data to begin reflecting realistic daily output back to them
- Users actively adjust their daily list based on app feedback (removing overloaded tasks, adding items when under capacity)
- Point budgets include non-work categories (exercise, hobbies, rest) in every day's plan — the app treats these as non-negotiable allocations, not afterthoughts
- Users receive positive reinforcement (compliments, encouragement) when completing todos, reinforcing habit formation
- Within **2-4 weeks**, users report feeling more balanced — daily todos consistently include progress toward at least one long-term goal alongside personal wellness tasks
- The app surfaces goal-alignment gaps: if a user's todos haven't touched a long-term goal in several days, it flags it

### Business Success

- **3 months:** 100 paying subscribers ($500 MRR) — product-market fit signal
- **6 months:** 400 paying subscribers ($2,000 MRR) — growth trajectory confirmed
- **12 months:** 1,000 paying subscribers ($5,000 MRR / $60,000 ARR)
- Monthly churn below 5% — users staying indicates the habit is forming
- Subscription model: $5/month flat rate

### Technical Success

- Seamless data sync across web and mobile — a todo added on mobile appears instantly on web and vice versa
- Capacity model improves accuracy over time — app suggestions for point estimates become more reliable the longer a user is active
- 99.9% uptime — todos are a daily habit; downtime breaks the routine and erodes trust
- Secure user data storage — personal goals and task history are sensitive; proper auth and encryption required

### Measurable Outcomes

- Daily active usage rate >60% among paying subscribers (this is a daily-habit app — low DAU = product isn't sticky)
- >70% of daily todo lists include at least one goal-tagged task
- >80% of daily lists include at least one wellness/fun/exercise point allocation
- Average capacity accuracy improves measurably in weeks 3-4 vs week 1

## Product Scope

### MVP — Minimum Viable Product

- User onboarding with long-term goal discovery (guided flow for users who don't know their goals yet)
- Daily todo creation with effort point assignment (user-set)
- Goal tagging per todo (which goal does this move forward?)
- Wellness point categories: exercise, hobby/fun, rest — required allocation per day
- Adaptive capacity model: tracks daily completed points over time, warns when list exceeds realistic capacity
- Positive reinforcement messaging when todos are completed
- Cross-platform: web app + mobile app with shared backend and database
- $5/month subscription with account creation and billing

### Growth Features (Post-MVP)

- App-suggested effort point estimates based on historical task patterns
- Social media / screen time awareness integration — surfaces time lost before the day begins
- Goal progress visualization — how much of your weekly/monthly effort went toward each goal?
- Streak tracking and milestone celebrations
- Smart goal-alignment nudges — "You haven't worked toward Goal X in 3 days"

### Vision (Future)

- AI-driven daily plan suggestions — the app proposes your todo list based on goals, capacity, and wellness balance
- Integration with calendars and external tools
- Community or accountability features (optional sharing of goal progress)
- Coaching-style insights: patterns, trends, personalized recommendations

## User Journeys

### Journey 1: Laura — The Developer with a Dream

**Persona:** Laura, 45, senior software developer. Sharp, disciplined at work, but her real dream is opening a vegan bakery. She's been "working toward it" for three years but her todo lists are always dominated by work tasks, home responsibilities, and the bakery dream keeps slipping to "someday."

**Opening Scene:** Laura downloads toDoFastTrack on a Sunday night, skeptical — she's tried every productivity app and they all make her feel more behind. During onboarding, the app asks her to name three long-term goals. She types: *Open a vegan bakery. Learn bread scoring techniques. Get to a healthy weight.* For the first time, her goals are written down somewhere that connects to her daily life.

**Rising Action:** Monday morning she adds 14 todos. The app flags: *"That's 38 points — you've only completed an average of 22 points on past Mondays. Consider removing some tasks."* She's annoyed but removes three work tasks and adds: *"Watch one bread scoring video (2pts)"* — her first bakery task in weeks. She also has a wellness allocation sitting unfilled: the app nudges her to add something for exercise or fun.

**Climax:** By Thursday of week two, Laura notices something: she's been adding at least one bakery task every single day. Not because she forced it — because the app keeps asking *"does this move you toward your goals?"* when she creates todos. The question itself is the intervention. She completes her fourth bakery task of the week and the app says: *"You're on a streak — 4 days working toward 'Open a vegan bakery.' Keep going, Laura."* She screenshots it and sends it to her sister.

**Resolution:** Six weeks in, Laura has a realistic picture of her actual daily capacity (around 20-24 points on workdays, 30+ on weekends). Her bakery goal gets touched 4-5 days a week. She hasn't opened the bakery yet — but for the first time, she believes she actually will. She's also exercising three times a week because the app wouldn't let her ignore it.

*Capabilities revealed: goal onboarding, goal-tagging on todos, capacity model, overload warnings, wellness point allocation, goal-streak tracking, positive reinforcement messaging.*

---

### Journey 2: Julie — The Reporter Racing the Clock

**Persona:** Julie, 30, journalist at a fast-paced news outlet. Two ambitious goals: write a novel and run a marathon. She's smart, driven, and chronically overcommitted. Her current todo app has 47 items. She hasn't run in two weeks. Her manuscript is at chapter three, where it's been for four months.

**Opening Scene:** Julie hears about toDoFastTrack from a podcast. During goal onboarding she types: *Write my novel. Run a marathon. Be present in my relationships.* The app asks if she needs help defining these goals — she doesn't, she's known them for years. The problem is the gap between knowing and doing.

**Rising Action:** Her first daily list has 22 items. The app flags overload immediately. She ignores the warning and marks 9 things done by end of day — including zero novel tasks and zero running. Day two, the app suggests: *"Yesterday you completed 18 points. Today you have 34 on your list — want to trim it down?"* She removes news-work overflow tasks and adds: *"Write 200 words of novel (3pts)"* and *"20 min run (4pts)."* Both get done. It's the first time in months both happened in the same day.

**Climax:** Week three. Julie has a brutal news week — a major story breaks and her work todos explode. The app flags that she's on track to go three days without touching her book or running. It doesn't lecture her — it just surfaces the gap: *"You haven't worked toward 'Write my novel' in 4 days."* She carves out 20 minutes at 6am. 200 words. Enough.

**Resolution:** Three months in, Julie is at chapter seven. She's run two 10Ks and has a half-marathon on the calendar. She still has crazy news weeks — but she no longer lets them consume the whole person. The app taught her that 200 words counts. That 20 minutes counts. That something is always better than nothing.

*Capabilities revealed: goal-alignment nudges, capacity suggestions, overload warnings, cross-goal balancing, progress tracking, "you haven't worked toward X" alerts.*

---

### Journey 3: The "I Don't Know My Goals" User — Edge Case

**Persona:** Alex, 38, feeling stuck and burned out. Knows something needs to change but can't articulate what they actually want. Opens toDoFastTrack during onboarding and stares at the "What are your 3 long-term goals?" prompt. Blank.

**Journey:** The app offers guided goal discovery — a short conversational flow asking questions like: *"What would you regret not doing in 5 years? What do you do that makes you lose track of time? What do you wish you had more of?"* Alex works through it and lands on: *Get healthier. Reconnect with photography. Spend more time offline.* None of these felt like "goals" before — but seeing them written down makes them real.

*Capabilities revealed: guided goal discovery flow, goal suggestion prompts, onboarding flexibility for undecided users.*

---

### Journey 4: Platform Admin (Operator)

**Persona:** Evelynzouras, the founder managing the product.

**Journey:** Logs into an admin dashboard to monitor subscription health (MRR, churn, new signups), review flagged support issues, and manage user accounts if needed. Can see aggregate usage patterns (not individual user data) — e.g., are users engaging with the wellness allocation feature? Is the capacity model being used?

*Capabilities revealed: admin dashboard, subscription management, aggregate analytics, user account management, support tooling.*

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

- **Global privacy compliance (GDPR + CCPA baseline):** Users have the right to export and delete all their data. Privacy policy required at launch explaining data collection, use, and retention. Consent obtained at account creation.
- **Data retention policy:** User data retained for duration of active subscription + 90 days post-cancellation (allows reactivation). Deleted on user request within 30 days.
- **Payment compliance:** Payments processed via third-party processor (e.g., Stripe). No raw card data stored or transmitted through application servers. PCI-DSS compliance delegated to payment processor.

### Technical Constraints

- **Data privacy:** Personal goals, task history, and capacity data are sensitive — encrypted at rest and in transit. No selling or sharing of user data with third parties.
- **Wellness positioning:** App communicates factually about user behavior (points completed, capacity trends) and never makes clinical, diagnostic, or prescriptive health claims. Language is observational: *"You completed X points over 5 days — your list of Y points may be optimistic."*
- **Account deletion:** Full data purge capability required — user-initiated, completed within 30 days.

### Risk Mitigations

- **Privacy policy + terms of service** required before launch — standard for any subscription product collecting personal data
- **Wellness language guidelines** — establish internal copy standards to keep all app messaging factual and non-prescriptive
- **Stripe integration** for subscriptions — handles payment security, recurring billing, and failed payment retries

## Innovation & Novel Patterns

### Detected Innovation Areas

**1. Adaptive Capacity Learning**
Most productivity apps let you add unlimited tasks. toDoFastTrack learns your *actual* daily output over time — not what you plan to do, but what you consistently complete — and uses that as a personalization baseline. The app gets smarter the longer you use it, shifting from a passive list to an active reality-check engine. This behavioral feedback loop applied to personal capacity management is uncommon in consumer productivity tools.

**2. The Permission Inversion**
Productivity apps typically make you feel behind. toDoFastTrack is architecturally designed to do the opposite: it tells you when you're *over-planning* and gives you data-backed permission to do less. This is a philosophical inversion of the standard productivity app model — the app's job is to reduce your list, not grow it.

**3. Wellness as a Non-Negotiable Budget**
Exercise, rest, and fun are not optional add-ons — they are required point allocations in every day's plan. The app won't consider a day "set up" until wellness categories are represented. This treats life balance as a system constraint rather than a nice-to-have, which is a meaningful architectural decision.

**4. Goal-Reality Thread**
Long-term goals aren't a separate vision board — they're woven into every daily task via tagging, and the app actively monitors when goals go untouched. The connection between a Tuesday todo and a 5-year dream is made explicit and tracked in real time.

### Market Context & Competitive Landscape

Existing productivity tools fall into two camps: simple task lists (Todoist, Things, Apple Reminders) and complex project management systems (Notion, Asana, Linear). Neither camp addresses personal capacity learning or goal-reality alignment at the individual level. The closest analogues — Habitica (gamification), Finch (wellness as self-care), and time-blocking tools — each solve one piece but not the whole. toDoFastTrack's combination of adaptive capacity + goal threading + wellness budgeting is a distinct position.

### Validation Approach

- **Capacity model accuracy**: Track whether app point estimates improve week-over-week vs. user's actual completions
- **Goal touch rate**: Measure whether users who engage with goal-tagging have higher retention than those who don't
- **Wellness allocation compliance**: Are users filling wellness categories daily? Does it correlate with retention?
- **Overexertion alerts acted upon**: Do users actually remove tasks when warned? Track list adjustment rate post-warning

### Risk Mitigation

- **Capacity model cold start**: New users have no history — default to conservative capacity estimates and improve rapidly in the first 2 weeks
- **Wellness nags becoming annoying**: Nudges must be encouraging, not guilt-inducing — copy standards and user testing critical here
- **Goal-tagging friction**: If tagging feels like overhead, users will skip it — keep it lightweight (one tap, not a form)

## Full-Stack Specific Requirements

### Project-Type Overview

toDoFastTrack is a cross-platform consumer application consisting of three layers: a React SPA (web), a React Native mobile app (iOS + Android), and a shared REST API backend with a persistent database. All three layers share a single data model — user account, goals, todos, and capacity history are consistent across surfaces.

### Web Application

- **Architecture:** Single-Page Application (React) — no full page reloads after initial load
- **Public landing page:** SEO-optimized marketing page (discoverable via Google) with clear value proposition, sign-up CTA, and login option
- **Session detection:** Landing page detects returning authenticated users via cookie/session token and redirects them directly to the app (or auto-logs them in if token is valid and unexpired)
- **Authenticated app:** Full todo/goal interface lives behind authentication — not publicly indexed
- **Browser support:** Modern browsers (Chrome, Firefox, Safari, Edge) — no IE support required
- **Responsive design:** Web app must be fully usable on mobile browsers, not just desktop
- **Performance target:** Initial load under 3 seconds; SPA transitions under 300ms

### Mobile Application

- **Platform:** Cross-platform via React Native — single codebase targeting iOS and Android
- **Store compliance:** Must meet Apple App Store and Google Play Store submission requirements at launch
- **Push notifications:**
  - Morning planning reminder (configurable time)
  - Overexertion alert ("your list today is over your typical capacity")
  - Goal nudge ("you haven't worked toward [goal] in X days")
  - Positive reinforcement ("great day — you hit your capacity target!")
- **Offline support:** Offline-first architecture — today's todos and goals stored locally on device (AsyncStorage / SQLite). Changes made offline sync automatically when connection is restored. Conflict resolution: last-write-wins for todo status; server authoritative for capacity model data.

### Backend & API

- **Architecture:** REST API serving both web and mobile clients
- **Authentication:**
  - Email/password (with secure password hashing)
  - OAuth via Google Sign-In and Apple Sign-In
  - JWT session tokens with refresh token rotation
  - Cookie-based session for web; token-based for mobile
- **Database:** Relational database (PostgreSQL recommended) — user accounts, goals, todos, daily point history, capacity model snapshots
- **Data sync:** API designed for efficient mobile sync — endpoints support fetching changes since last sync timestamp
- **Subscription management:** Stripe integration for $5/month subscription — webhook handling for payment events (success, failure, cancellation)

### Implementation Considerations

- **Monorepo recommended:** Shared types, utilities, and API client code between web and React Native to reduce duplication
- **Capacity model:** Server-side calculation — not client-side — so data is consistent across devices and survives app reinstalls
- **Environment separation:** Development, staging, and production environments required before launch

## Project Scoping & Phased Development

### MVP Strategy & Philosophy

**MVP Approach:** Experience MVP — the goal is for users to feel the core value (realistic planning + goal progress) within their first week, and to be converted to paying subscribers after a 30-day free trial.

**Resource Requirements:** Solo developer. Sequencing and scope discipline are critical — every feature added to MVP extends the timeline significantly.

**Free Trial:** 30 days free, then $5/month via Stripe. Full feature access during trial.

### MVP Feature Set (Phase 1)

**Core User Journeys Supported:**
- Laura (goal-pursuer): onboarding → goal setup → daily planning → capacity feedback → goal streaks
- Julie (overcommitter): overload warning → list trimming → goal-alignment nudges
- Alex (undecided): guided goal discovery flow
- Admin: Stripe subscription management + basic account tooling

**Must-Have Capabilities:**

- **Onboarding:** Account creation (email/password + Google/Apple Sign-In), guided goal-setting for up to 3 long-term goals, optional goal discovery flow for users who aren't sure
- **Daily Planning:** Create todos with effort point assignment (user-set), goal-tag each todo (which goal does it serve?), wellness category allocation (exercise, fun, rest) required before day is "set"
- **Capacity Model:** Tracks completed points per day, enters "learning" state for first 7-10 days with transparent messaging ("Still learning your pace — keep completing tasks"), produces capacity estimates and overload warnings from week 2 onward
- **Feedback & Reinforcement:** Overexertion alerts when daily list exceeds capacity estimate, positive reinforcement messages on todo completion and goal streaks, goal-alignment nudges when a goal hasn't been touched in 3+ days
- **Cross-Platform:** React SPA (web) + React Native (iOS + Android), shared REST API + PostgreSQL backend
- **Auth & Sessions:** JWT tokens, cookie-based web sessions, auto-login for returning users on landing page
- **Push Notifications:** Morning planning reminder, overexertion alert, goal nudge, positive reinforcement
- **Offline Support:** Local storage of today's plan, sync on reconnect
- **Subscription:** Stripe integration, 30-day free trial, $5/month recurring billing, failed payment handling
- **SEO Landing Page:** Public marketing page with sign-up and login, optimized for search

**Deliberately excluded from MVP:**
- App-suggested point estimates (requires more history data — Phase 2)
- Social media / screen time tracking (Phase 2)
- Admin analytics dashboard (manual Stripe dashboard sufficient at early scale)

### Post-MVP Features

**Phase 2 — Growth (post product-market fit):**
- App-suggested effort point estimates based on historical task patterns
- Social media / screen time awareness integration
- Goal progress visualization (weekly/monthly effort breakdown by goal)
- Streak tracking and milestone celebrations
- Admin dashboard with aggregate analytics (DAU, retention, capacity model accuracy)

**Phase 3 — Expansion:**
- AI-driven daily plan suggestions
- Calendar integration (block time for high-point todos)
- Community / accountability features (optional goal-sharing)
- Coaching-style insights: patterns, burnout risk trends, personalized recommendations

### Risk Mitigation Strategy

**Technical Risks:**
- *Biggest risk:* Offline sync conflict resolution and cross-platform parity as a solo developer. **Mitigation:** Launch web first, ship mobile to TestFlight/Play internal testing in parallel, release mobile publicly 4-6 weeks after web launch. Reduces surface area at launch.
- *Capacity model cold start:* Display explicit "learning" state for first 10 days with encouraging copy. Don't show capacity warnings until minimum 5 days of data exist.
- *App Store review delays:* Submit to both stores 2-3 weeks before planned launch date to absorb review cycles.

**Market Risks:**
- *Biggest risk:* Users don't maintain the daily planning habit. **Mitigation:** Push notification for morning planning reminder is an MVP requirement, not a growth feature. The 30-day trial gives enough time to experience the full habit loop before payment.
- *Competitor response:* Established apps (Todoist, Notion) could add capacity features. **Mitigation:** The wellness-as-constraint philosophy and goal-reality threading are the moat — ship fast and build retention.

**Resource Risks:**
- Solo developer means a smaller, more focused MVP is essential. **Mitigation:** Web-first launch with mobile following 4-6 weeks later. Admin tooling deferred to Phase 2 — Stripe dashboard handles early subscription management.

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
- FR20: Users are required to include at least one wellness-category todo before their day plan is considered complete
- FR21: Users can carry incomplete todos forward to the next day

### Capacity Management

- FR22: The system tracks the number of effort points a user completes each day over time
- FR23: The system displays a "learning" state during the first 7-10 days of use, with messaging that explains it is building a capacity baseline
- FR24: The system generates a daily capacity estimate for the user once sufficient history exists (minimum 5 days)
- FR25: The system alerts users when their planned daily point total exceeds their capacity estimate
- FR26: The system encourages users to add more tasks when their planned total is significantly below their capacity estimate
- FR27: Users can view their historical daily point completion data

### Feedback & Reinforcement

- FR28: The system delivers positive reinforcement messages when a user completes a todo
- FR29: The system delivers positive reinforcement when a user hits or exceeds their capacity target for the day
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
- FR38: Changes made on one platform are reflected on the other platform in real time when connected
- FR39: Users can view and interact with their todo list while offline on mobile
- FR40: Changes made offline are automatically synced to the server when connectivity is restored

### Subscription & Billing

- FR41: New users receive a 30-day free trial with full access to all features
- FR42: Users are prompted to enter payment details at the end of their free trial
- FR43: The system charges users $5/month via Stripe upon trial completion
- FR44: Users receive notification of upcoming billing and failed payment attempts
- FR45: Users can cancel their subscription and retain access until the end of the billing period

### Landing Page & Discoverability

- FR46: The public landing page is optimized for search engine indexing
- FR47: The landing page allows new visitors to sign up for an account
- FR48: The landing page allows existing users to log in
- FR49: The landing page detects authenticated returning users and redirects them to the app
