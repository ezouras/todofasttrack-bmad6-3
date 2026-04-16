---
stepsCompleted:
  - step-01-init
  - step-02-discovery
  - step-03-core-experience
  - step-04-emotional-response
  - step-05-inspiration
  - step-06-design-system
  - step-07-defining-experience
  - step-08-visual-foundation
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

## Defining Core Experience

### The Defining Experience

Tend's signature interaction: **intentional todo creation**. The moment a user types a task, taps a size (XS→XL), taps a goal chip, and — if relevant — taps a wellness icon, they've done something no plain to-do app lets them do: told themselves what this task costs and what it's for. That micro-ritual, done every day, is what makes Tend different from a notes app or a plain checklist.

**Tend's headline interaction:** "Plan your day around what matters — in under two minutes."

### User Mental Model

Users come to Tend from one of two places:
- **The overwhelmed list-keeper** — they've been writing everything down in Notes or Todoist with no sense of what's achievable. Mental model: a to-do list is a wishlist. Tend reframes it as a plan.
- **The goal-aspirant who loses the thread** — they know they want to write a book or run a marathon but their daily todos never connect to that. Mental model: goals live separately from daily work. Tend collapses that gap.

Both users understand the concept of a checkbox. The novel element is the **size + goal tag** pairing — two extra taps that carry significant meaning. This needs to feel like a natural extension of creating a task, not a form to fill out.

### Success Criteria

The todo creation flow succeeds when:
- It takes **under 10 seconds** from tap to saved
- The user **never has to open a separate screen** — everything inline
- Size defaults to **M** so users can skip it when in a hurry
- Goal tag defaults to **none** so it's always optional
- After saving, the capacity bar updates **instantly** — the user immediately sees the impact of what they just added

### Novel UX Patterns

Tend combines familiar patterns in a novel way:
- **Familiar:** text input for task title, tap-to-select chips, swipe to complete
- **Novel:** the **XS/S/M/L/XL sizing vocabulary applied to daily tasks** — close enough to t-shirt sizing that education is minimal
- **Novel:** the **capacity bar as a live consequence** — seeing the bar move as you add todos is new behavior that needs to feel satisfying, not stressful

Teaching moment: the first time a user adds a todo and sees the capacity bar respond, a one-line tooltip ("This shows how your day is filling up") is enough. No tutorial required.

### Experience Mechanics

**1. Initiation:**
- Prominent "+" button, always visible at bottom of the daily view (web: bottom-right; mobile: bottom-center)
- Tapping slides up an inline creation row — no modal, no new screen

**2. Interaction:**
- Type title → cursor moves naturally to size chips
- Size chips (XS / S / M / L / XL) displayed inline, M pre-selected
- Below size: goal chips (user's 3 goals by name + color, plus "None") — single tap
- Below goal: wellness icons (exercise / fun-hobby / rest) — optional, single tap, can skip
- Capacity bar updates live as size is selected

**3. Feedback:**
- Capacity bar animates smoothly on size change — fills from left, warm blush color
- If selection would push over capacity: bar shifts to amber, soft inline message appears ("You've planned more than your recent average — that's okay")
- No blocking, no error — purely informational

**4. Completion:**
- "Add" tap (or Return on desktop) saves the todo and collapses the creation row
- New todo appears at bottom of list with a gentle fade-in
- Capacity bar settles at new level

## Visual Design Foundation

### Color System

**Semantic Color Tokens:**

| Token | Value | Usage |
|---|---|---|
| `color-primary` | `#E8A0A0` | Primary actions, active states, capacity bar fill |
| `color-primary-dark` | `#C97070` | Hover states, pressed states |
| `color-primary-light` | `#F5D0D0` | Subtle highlights, selected chip backgrounds |
| `color-accent` | `#F0B090` | CTAs, positive reinforcement moments |
| `color-bg` | `#FDFAF7` | App background |
| `color-surface` | `#FFF8F5` | Cards, panels, todo rows |
| `color-surface-raised` | `#FFFFFF` | Modals, popovers |
| `color-text-primary` | `#3D3230` | Body text, headings |
| `color-text-secondary` | `#7A6865` | Labels, metadata, secondary info |
| `color-text-muted` | `#B0A09E` | Placeholder text, disabled states |
| `color-warning` | `#D4A843` | Capacity over-limit indicator |
| `color-warning-bg` | `#FDF3DC` | Warning message background |
| `color-success` | `#7BAF8A` | Completion checkmark, day wrap-up |
| `color-border` | `#EDE5E3` | Subtle dividers, input borders |

**Goal Colors (per-slot, consistent per user):**

| Token | Value | Usage |
|---|---|---|
| `color-goal-1` | `#B8A0D4` | Soft lavender — goal slot 1 chip + badge |
| `color-goal-2` | `#8EB89A` | Warm sage — goal slot 2 chip + badge |
| `color-goal-3` | `#E8B49A` | Warm peach — goal slot 3 chip + badge |

**Effort Size Chip Colors (opacity scale on primary):**

| Size | Opacity | Visual signal |
|---|---|---|
| XS | 30% | Lightest — quick task |
| S | 50% | Light |
| M | 70% | Default |
| L | 85% | Heavy |
| XL | 100% | Heaviest — significant commitment |

**Accessibility:** All text/background pairings target WCAG AA minimum (4.5:1 contrast). `color-text-primary` on `color-bg` achieves ~8:1. Warning amber checked against white background. Goal colors used only for non-text UI elements (chips, badges) where AA large-text threshold (3:1) applies.

### Typography System

**Typeface:** Plus Jakarta Sans (Google Fonts, variable weight)

| Role | Size | Weight | Line Height | Usage |
|---|---|---|---|---|
| `text-display` | 28px | Semibold (600) | 1.3 | Page titles, day header |
| `text-heading` | 20px | Semibold (600) | 1.4 | Section headings |
| `text-subheading` | 16px | Medium (500) | 1.5 | Card titles, goal names |
| `text-body` | 15px | Regular (400) | 1.65 | Todo titles, body copy |
| `text-label` | 13px | Medium (500) | 1.5 | Chips, tags, metadata labels |
| `text-caption` | 12px | Regular (400) | 1.6 | Timestamps, helper text |

**Mobile:** Base sizes reduced by 1px across scale; minimum body size 14px.

### Spacing & Layout Foundation

**Base unit:** 4px. All spacing values are multiples of 4.

| Token | Value | Usage |
|---|---|---|
| `space-1` | 4px | Tight internal padding (chip icon gap) |
| `space-2` | 8px | Chip padding, small gaps |
| `space-3` | 12px | Input padding, list item internal spacing |
| `space-4` | 16px | Section padding, card padding |
| `space-6` | 24px | Between sections |
| `space-8` | 32px | Page margins (mobile) |
| `space-10` | 40px | Page margins (desktop) |
| `space-16` | 64px | Large breathing room between major sections |

**Layout:**
- Web: max content width 720px, centered — feels like a focused document, not a dashboard
- Mobile: full-width with 16px horizontal margin
- Todo list rows: 56px min height — comfortable touch target on mobile, not oversized on desktop
- Generous vertical rhythm — 24px between todo items; never cramped

**Border radius:**
- Chips and tags: `rounded-full` (fully rounded)
- Cards and panels: `rounded-2xl` (16px) — soft, not boxy
- Inputs: `rounded-xl` (12px)
- Buttons: `rounded-full` for primary; `rounded-xl` for secondary

### Accessibility Considerations

- All interactive elements minimum 44×44px touch target (mobile)
- Focus rings visible and high-contrast (2px `color-primary-dark` outline)
- Color never the sole indicator of meaning — size chips labeled XS/S/M/L/XL in addition to opacity variation
- Goal colors supplemented by goal name text on all tagged todos — color-blind users can still identify goals
- Capacity warning uses both color (amber) and icon + text — never color alone
- Plus Jakarta Sans has good legibility at small sizes; minimum rendered size 12px
