---
stepsCompleted:
  - step-01-init
  - step-02-context
  - step-03-starter
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
