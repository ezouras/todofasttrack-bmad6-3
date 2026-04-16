---
stepsCompleted:
  - step-01-init
  - step-02-discovery
  - step-03-core-experience
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
