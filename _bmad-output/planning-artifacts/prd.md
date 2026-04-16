---
stepsCompleted:
  - step-01-init
  - step-02-discovery
  - step-02b-vision
  - step-02c-executive-summary
  - step-03-success
  - step-04-journeys
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
