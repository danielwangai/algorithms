-- Single-table schema for the DSA question tracker.
CREATE TABLE IF NOT EXISTS questions (
  id            SERIAL PRIMARY KEY,
  leetcode_id   INTEGER,
  title         TEXT        NOT NULL,
  url           TEXT        NOT NULL UNIQUE,
  topic         TEXT        NOT NULL,
  difficulty    TEXT        NOT NULL CHECK (difficulty IN ('Easy', 'Medium', 'Hard')),
  is_microsoft  BOOLEAN     NOT NULL DEFAULT false,
  is_complete   BOOLEAN     NOT NULL DEFAULT false,
  completed_at  TIMESTAMPTZ,
  created_at    TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX IF NOT EXISTS idx_questions_topic      ON questions (topic);
CREATE INDEX IF NOT EXISTS idx_questions_difficulty ON questions (difficulty);
