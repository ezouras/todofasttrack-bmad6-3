---
stepsCompleted:
  - step-01-init
  - step-02-discovery
  - step-03-core-experience
  - step-04-emotional-response
  - step-05-inspiration
  - step-06-design-system
  - step-07-defining-experience
inputDocuments:
  - _bmad-output/planning-artifacts/prd.md
  - _bmad-output/planning-artifacts/architecture.md
workflowType: 'ux-design'
project_name: 'toDoFastTrack'
user_name: 'Evelynzouras'
date: '2026-04-15'
---

# UX Design Specification — toDoFastTrack

**Author:** Evelynzouras
**Date:** 2026-04-15

---

<!-- UX design content will be appended sequentially through collaborative workflow steps -->

## Executive Summary

### Project Vision

toDoFastTrack is a daily planning app built around 3 long-term personal goals. Users assign effort points to each todo (like Jira story points), tag them to goals and wellness categories (exercise, fun/hobby, rest), and the app learns their personal daily capacity over time. The core promise: realistic planning that protects users from burnout while keeping their long-term ambitions visible in daily life. Available as a web app and mobile app ($5/month, 30-day free trial).

### Target Users

**Laura** — 45, software developer pursuing a long-term goal of owning a vegan bakery. Driven and high-achieving, but at risk of systematically over-scheduling. Needs a planning system that respects her real bandwidth, not an idealized version of it.

**Julie** — 30, journalist juggling a demanding job while pursuing two long-term goals: writing a book and running a marathon. Needs help keeping her ambitious side-goals visible and actionable without overwhelming her primary work responsibilities.

**The shared pattern:** Goal-oriented people who suspect their to-do lists are working against them — packed with tasks that don't connect to what they actually care about, and scaled to a version of themselves that doesn't have a job, family, or need for rest.

### Key Design Challenges

1. **The capacity learning state (days 1-10)** — The app cannot produce a capacity estimate until it has observed at least 5 days of usage. This is a high-risk UX moment: users expect immediate value. The "still learning" experience must feel intentional and trust-building, not like a broken feature.

2. **Effort point assignment friction** — If tagging a todo with a point value feels like extra work, users will skip it and the entire capacity model breaks down. Point assignment must be fast enough to feel like part of creating the todo, not a separate step.

3. **Wellness category tagging** — Same friction risk. The PRD explicitly requires a one-tap interaction. A modal form or multi-step flow here will kill daily adoption.

4. **Onboarding goal discovery** — Many users won't have 3 clear long-term goals articulated and ready to type. The guided discovery flow must feel reflective and supportive — more journal prompt than registration form.

5. **Overload alerting tone** — When the app warns that the day's plan exceeds capacity, it must read as a helpful, observational nudge. Any hint of scolding, alarm, or prescriptive instruction violates the wellness language constraint and risks emotional harm to users with productivity anxiety.

### Design Opportunities

1. **The learning period as emotional investment** — Framing days 1-10 as "you're teaching the app about yourself" (rather than a broken state) could create genuine anticipation for the capacity estimate reveal — turning a product limitation into a feature.

2. **Goal tagging as identity, not taxonomy** — Seeing daily todos visually mapped to real long-term aspirations could be genuinely motivating. Goal badges and streaks should feel expressive, not administrative.

3. **Daily completion as a ritual** — A satisfying, warm end-of-day moment — not just a counter — could be the emotional hook that makes toDoFastTrack a habit rather than just another productivity tool.

## Core User Experience

### Defining Experience

The core loop is **daily planning as a morning ritual**. Users open the app and land on a view that shows their capacity status and goal context first — then their todo list for the day beneath it. The most frequent action is adding a todo, which happens multiple times a day and must be frictionless enough to do in 10 seconds.

### Platform Strategy

- **Web (desktop) is primary** — where users plan and think; the morning ritual happens here
- **Mobile is the companion** — for checking off todos throughout the day and quick additions on the go; launched 4-6 weeks after web
- **Both platforms share the same information architecture** — mobile is not a simplified version, just a touch-optimized one
- **Offline-first on mobile** — todo list must be usable without network; sync on reconnect

### Effortless Interactions

- **Effort sizing via XS/S/M/L/XL** — five tappable chips, no number entry, no free text. Should feel like choosing a coffee size, not filling out a form. Default to M if untouched.
- **Goal tagging** — tap one of up to 3 goal chips (or "none"). Single tap, no modal.
- **Wellness category** — tap one of 3 icons (exercise / fun-hobby / rest), or skip. Optional unless the day plan is missing one entirely.
- **Todo creation** — title + sizing + goal tag + wellness category, all inline. No separate "add details" flow.
- **Morning carry-forward prompt** — on first login of the day (any platform), app surfaces yesterday's incomplete todos one at a time: "Add this to today? Keep / Skip." User is always in control; the app never silently rolls items forward.
- **Optional day completion** — users can explicitly close out their day via a single "Wrap up today" action, which triggers the end-of-day reinforcement moment. If they don't, the day closes automatically at midnight. No penalty either way; the app treats the data identically.

### Critical Success Moments

1. **First todo added** — user experiences how fast and low-friction the full flow is (title → size → goal → done in under 10 seconds)
2. **First capacity estimate (around day 5-7)** — the moment the app becomes predictive; needs to feel like a reveal, not just a number appearing
3. **First overload warning** — the first time the app says "you've planned more than you typically accomplish" must feel helpful and kind, never alarming
4. **Completing a goal-tagged todo** — the moment of reinforcement that ties daily work to long-term identity
5. **Morning carry-forward prompt** — first time the app "remembers" yesterday and checks in; must feel like a thoughtful assistant, not a guilt trip
6. **Explicit day completion (optional)** — for users who choose to close out the day intentionally, this becomes a satisfying ritual: a summary of what they accomplished, a note on goal progress, a warm sign-off. Not a report card — more like putting a book down with a sense of satisfaction.

### Experience Principles

1. **The app observes, the user decides** — every nudge, warning, and suggestion is framed as information, never instruction. The app never tells users what to do.
2. **Speed is a feature** — any interaction that recurs daily must be completable without thinking. Friction compounds across 365 days.
3. **Goals stay visible, not buried** — long-term goals aren't in a settings page; they're woven into the daily planning view. Every todo is a chance to connect to something bigger.
4. **The day ends on your terms** — users can explicitly wrap up their day to get the reinforcement moment and a clean close-out. If they don't, the day closes automatically at midnight with no friction. The choice is always theirs; the app works either way.
5. **Warmth over gamification** — positive reinforcement is human and observational ("You've made real progress toward your bakery goal this week") not points, badges, or streaks for their own sake.

## Desired Emotional Response

### Primary Emotional Goals

**Calm and grounded** is the north star. toDoFastTrack should feel like a deep breath, not a productivity dashboard. Users — especially those prone to anxiety about their to-do lists — should consistently leave the app feeling more settled than when they opened it, not more pressured.

**Connected to purpose** is the secondary goal. Not motivational-poster inspiration, but the quiet, real feeling of seeing daily work connecting to something that actually matters. The difference between "I replied to 14 emails" and "I spent an hour on my book today."

**Trusted, not judged** is the foundation. On bad days especially — when 2 of 8 todos got done — the app must meet the user with neutrality and warmth, not disappointment.

### Emotional Journey Mapping

| Moment | Target Feeling | Design Implication |
|---|---|---|
| First opens app | Hopeful, not overwhelmed | Onboarding is reflective and unhurried |
| Daily planning | Clear-headed, settled | Clean layout; no clutter or urgency signals |
| Adding a todo | Effortless, natural | Fast inline flow; no modal forms |
| Capacity warning | Informed, not alarmed | Soft tone; observational language, no red alerts |
| Completing a goal-tagged todo | Quiet satisfaction | Warm, specific acknowledgment — not a fanfare |
| Great day, explicit wrap-up | Genuinely acknowledged | Warm summary; specific to what they did, not generic praise |
| Hard day (low completion) | Neutral, unjudged | No negative framing; just tomorrow is a fresh start |
| Returning next morning | Welcomed, familiar | Like opening a notebook you trust |

### Micro-Emotions

- **Trust over skepticism** — especially during the 10-day learning period; the app needs to feel credible before it has data
- **Confidence over confusion** — every screen has one clear thing to do next
- **Quiet satisfaction over excitement** — reinforcement is warm and human, not performative
- **Acceptance over frustration** — incomplete days are presented as neutral data, not failure

### Design Implications

- **Visual language:** Soft, warm palette — inspired by Calm's blues and earth tones, not bright productivity-app primaries. No red for warnings; amber or soft orange at most.
- **Reinforcement copy is specific, not generic** — "You worked on your bakery goal three days this week" lands differently than "Great job!" Specificity makes it feel human.
- **No alarm iconography** — capacity warnings use informational icons (info circle, soft indicator), never warning triangles or red badges
- **No punishing streaks** — goal streaks are displayed as encouragement, never as something that resets dramatically or shames the user for a missed day
- **Micro-animations are quiet** — a gentle checkmark animation on completion, not a confetti explosion. The emotional register is a warm nod, not a parade.
- **Silence is okay** — not every moment needs a nudge, a tip, or a message. The app is present without being chatty.

### Emotional Design Principles

1. **The app is a calm companion, not a coach** — it observes and reflects; it never pushes or prescribes
2. **Specificity is warmth** — generic praise feels hollow; specific acknowledgments ("you finished something toward your marathon goal") feel real
3. **Bad days are neutral data** — the app treats low-completion days as information, not failure; no negative framing, ever
4. **Quiet > loud** — every moment of celebration or reinforcement should feel like a warm nod from someone who noticed, not a notification asking for attention
5. **The emotional tone of Calm and Insight Timer, applied to planning** — unhurried, warm, on your side

## UX Pattern Analysis & Inspiration

### Inspiring Products Analysis

**Calm**
- Breathing room in layouts — generous whitespace, one focal point per screen
- Soft layered visual hierarchy — type and color carry the weight; no harsh boxes or dividers
- Unhurried transitions — screens ease rather than snap; the pacing itself creates calm
- Progressive disclosure — reveals options when needed, not all at once

**Insight Timer**
- History as encouragement, not tracking — your record feels like self-care, not performance
- Onboarding feels like a conversation — asks what you need, not who you are
- Warm, specific acknowledgments — "47 people meditating right now" feels human, not gamified
- Post-session screen is quiet — simple, warm, done. No upsell, no share prompt.

### Transferable UX Patterns

**Layout & Visual:**
- Single focal point per screen — Tend never competes with itself for attention
- Generous whitespace as a design element — breathing room signals calm, not emptiness
- Soft transitions between states — ease-in/out, never snap or flash

**Interaction:**
- Progressive disclosure — todo creation reveals size/goal/wellness fields naturally, not all at once in a form
- Conversational onboarding — goal discovery flow feels like thoughtful questions, not a registration form

**Tone & Copy:**
- Specific acknowledgments over generic praise — mirrors Insight Timer's human warmth
- Post-completion quietness — the wrap-up screen is settled and warm, not an upsell moment

### Anti-Patterns to Avoid

- **Ads or promotional content on open** — destroys trust in the first 2 seconds; Tend will never show ads (subscription model eliminates this entirely)
- **Gamification pressure** — leaderboards, aggressive streak counters, badges that shame on a missed day
- **Feature overload on first launch** — showing everything at once signals complexity, not capability
- **Alarm-style notifications and warnings** — red badges, warning triangles, urgent language for capacity nudges
- **Generic praise** — "Great job!" without specificity feels hollow and automated

### Design Inspiration Strategy

**Adopt from Calm:**
- Layout breathing room and whitespace philosophy
- Soft, warm color palette (blues, earth tones, muted greens) — no productivity-app primaries
- Unhurried micro-animations

**Adopt from Insight Timer:**
- Conversational onboarding structure
- Quiet post-completion/wrap-up screen design
- History presented as personal record, not performance data

**Tend's unique direction:**
- Minimalist with happy colors — warmer and more expressive than Calm's muted palette, but still clean
- Goals as visible identity — goal chips and tags are visually expressive, not just labels
- The capacity reveal (day 5-7) as a designed moment — Calm and Insight Timer don't have an equivalent; this is Tend's own signature interaction to design

## Design System Foundation

### Design System Choice

**Tailwind CSS v4 + shadcn/ui** (web) and **NativeWind v4** (mobile), customized with Tend's own design tokens. shadcn/ui provides unstyled, accessible component primitives — Tend's visual identity is built on top of that foundation.

### Rationale for Selection

- Already established in architecture — no new decisions needed
- shadcn/ui's unstyled approach means nothing fights against Tend's warm visual identity
- Shared Tailwind tokens between web and mobile ensure both platforms feel like the same product
- Solo developer: shadcn/ui's copy-paste component model is practical and fast

### Tend's Visual Identity

**Color Palette:**
- **Primary:** Warm blush rose (`#E8A0A0` range) — the emotional anchor of the brand. Soft, warm, not aggressive.
- **Accent:** Dusty coral/peach (`#F0B090` range) — for CTAs, highlights, and positive reinforcement moments
- **Background:** Warm cream (`#FDFAF7`) — not pure white; softer and warmer
- **Surface:** Slightly warmer white (`#FFF8F5`) for cards and panels
- **Text:** Warm charcoal (`#3D3230`) — not pure black; reads warmer
- **Warning/capacity:** Soft amber (`#D4A843`) — informational, never alarming. No red.
- **Success/completion:** Muted warm green (`#7BAF8A`) — growth, not clinical

**Goal Colors (each of a user's 3 goals gets a distinct warm tone):**
- Goal slot 1: Soft lavender (`#B8A0D4`) — calm, purposeful
- Goal slot 2: Warm sage green (`#8EB89A`) — grounded, growth
- Goal slot 3: Warm peach (`#E8B49A`) — energetic, creative
- These are harmonious as a set — distinct enough to differentiate at a glance, warm enough to coexist on screen without clashing

**Typography:**
- **Typeface:** Plus Jakarta Sans — rounded, humanist, approachable without being playful. Works well at both small (mobile) and large (desktop) sizes.
- **Scale:** Generous line height (1.6–1.7) to support the breathing-room principle
- **Weight:** Regular for body, Medium for labels, Semibold for headings — never Bold or Black

**Mode:** Light mode only at launch. Dark mode considered post-MVP.

### Customization Strategy

- Design tokens defined once in `packages/config` — primary, accent, goal colors, neutrals, typography scale
- shadcn/ui components reskinned via Tailwind config — no component source modifications
- Goal color assignment: user's first goal gets lavender, second gets sage, third gets peach — fixed assignment so colors feel personal and consistent across sessions
- XS/S/M/L/XL effort size chips use the primary blush palette with varying opacity to signal relative weight
