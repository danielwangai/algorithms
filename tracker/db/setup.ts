/**
 * Creates the schema and seeds questions by parsing ../STUDY_PLAN.md.
 * The markdown plan is the single source of truth. Re-running is idempotent:
 * existing rows (matched by url) are left untouched, so your progress is preserved.
 *
 * Run: npm run db:setup
 */
import { config } from "dotenv";
config({ path: ".env.local" });

import { readFileSync } from "fs";
import { join } from "path";
import { Pool } from "pg";

interface SeedRow {
  leetcode_id: number | null;
  title: string;
  url: string;
  topic: string;
  difficulty: "Easy" | "Medium" | "Hard";
  is_microsoft: boolean;
  is_complete: boolean;
}

function parsePlan(markdown: string): SeedRow[] {
  const lines = markdown.split("\n");
  const rows: SeedRow[] = [];
  const seen = new Set<string>();

  // "## 12. Bit Manipulation & Math (...)" -> "Bit Manipulation & Math"
  const topicRe = /^##\s+\d+\.\s+(.+)$/;
  // "- [ ] [Title](https://...) — Easy · #1 ⭐"
  const itemRe = /^- \[( |x)\]\s+\[(.+?)\]\((https?:\/\/[^)]+)\)/;

  let topic = "";
  for (const line of lines) {
    const topicMatch = line.match(topicRe);
    if (topicMatch) {
      topic = topicMatch[1].replace(/\s*\(.*$/, "").trim();
      continue;
    }

    const item = line.match(itemRe);
    if (!item || !topic) continue;

    const url = item[3];
    if (seen.has(url)) continue; // keep first (primary) topic for cross-listed problems
    seen.add(url);

    const diffMatch = line.match(/\b(Easy|Medium|Hard)\b/);
    const idMatch = line.match(/#(\d+)/);

    rows.push({
      leetcode_id: idMatch ? parseInt(idMatch[1], 10) : null,
      title: item[2],
      url,
      topic,
      difficulty: (diffMatch?.[1] as SeedRow["difficulty"]) ?? "Medium",
      is_microsoft: line.includes("⭐"),
      is_complete: item[1] === "x",
    });
  }
  return rows;
}

async function main() {
  if (!process.env.DATABASE_URL) {
    throw new Error("DATABASE_URL is not set. Copy .env.local.example to .env.local first.");
  }

  const pool = new Pool({ connectionString: process.env.DATABASE_URL });

  const schema = readFileSync(join(process.cwd(), "db", "schema.sql"), "utf8");
  await pool.query(schema);

  const planPath = join(process.cwd(), "..", "STUDY_PLAN.md");
  const rows = parsePlan(readFileSync(planPath, "utf8"));

  let inserted = 0;
  for (const r of rows) {
    const res = await pool.query(
      `INSERT INTO questions
         (leetcode_id, title, url, topic, difficulty, is_microsoft, is_complete, completed_at)
       VALUES ($1, $2, $3, $4, $5, $6, $7, CASE WHEN $7 THEN now() ELSE NULL END)
       ON CONFLICT (url) DO NOTHING`,
      [r.leetcode_id, r.title, r.url, r.topic, r.difficulty, r.is_microsoft, r.is_complete]
    );
    inserted += res.rowCount ?? 0;
  }

  console.log(
    `Parsed ${rows.length} questions from STUDY_PLAN.md; inserted ${inserted} new (existing rows untouched).`
  );
  await pool.end();
}

main().catch((err) => {
  console.error(err);
  process.exit(1);
});
