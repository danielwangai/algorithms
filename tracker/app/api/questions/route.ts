import { NextResponse } from "next/server";
import { pool } from "@/lib/db";

// Always read fresh from the DB (don't statically cache).
export const dynamic = "force-dynamic";

export async function GET() {
  const { rows } = await pool.query(
    `SELECT id, leetcode_id, title, url, topic, difficulty,
            is_microsoft, is_complete, completed_at
     FROM questions
     ORDER BY topic, COALESCE(leetcode_id, 99999), title`
  );
  return NextResponse.json(rows);
}
