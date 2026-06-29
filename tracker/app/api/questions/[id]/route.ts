import { NextResponse } from "next/server";
import { pool } from "@/lib/db";

export const dynamic = "force-dynamic";

// Toggle completion. Flips is_complete and sets/clears completed_at atomically.
export async function PATCH(
  _req: Request,
  { params }: { params: { id: string } }
) {
  const id = Number(params.id);
  if (!Number.isInteger(id)) {
    return NextResponse.json({ error: "invalid id" }, { status: 400 });
  }

  const { rows } = await pool.query(
    `UPDATE questions
     SET is_complete  = NOT is_complete,
         completed_at = CASE WHEN is_complete THEN NULL ELSE now() END
     WHERE id = $1
     RETURNING id, is_complete, completed_at`,
    [id]
  );

  if (rows.length === 0) {
    return NextResponse.json({ error: "not found" }, { status: 404 });
  }
  return NextResponse.json(rows[0]);
}
