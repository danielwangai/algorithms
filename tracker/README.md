# DSA Tracker

A tiny Next.js app to track LeetCode questions for Microsoft SWE II prep.
List, filter by topic / difficulty / status, search, and toggle each question done.
Seeded directly from `../STUDY_PLAN.md`.

## Stack

- Next.js (App Router) + Route Handlers as the API — no separate backend
- Postgres accessed directly with `pg` (raw SQL, no ORM)
- One table: `questions`
- Tailwind CSS

## Setup

1. **Install deps**
   ```bash
   npm install
   ```

2. **Create a Postgres database**
   ```bash
   createdb dsa_tracker
   ```

3. **Configure the connection**
   ```bash
   cp .env.local.example .env.local
   # edit DATABASE_URL if your Postgres user/password/port differ
   ```

4. **Create the table and seed from the study plan** (idempotent — keeps your progress)
   ```bash
   npm run db:setup
   ```

5. **Run it**
   ```bash
   npm run dev
   ```
   Open http://localhost:3000

## How it works

- `GET /api/questions` — returns all questions (filtering/sorting happens client-side for snappiness).
- `PATCH /api/questions/:id` — toggles `is_complete` and sets/clears `completed_at` atomically.
- `db/setup.ts` parses `STUDY_PLAN.md`, dedupes cross-listed problems by URL (keeping the first/primary
  topic), and upserts with `ON CONFLICT (url) DO NOTHING` so re-seeding never wipes completion state.

## Re-seeding after editing the plan

Add new problems to `STUDY_PLAN.md` in the same `- [ ] [Title](url) — Difficulty · #N ⭐` format,
then run `npm run db:setup` again. Only new URLs get inserted.
