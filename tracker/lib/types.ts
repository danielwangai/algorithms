export type Difficulty = "Easy" | "Medium" | "Hard";

export interface Question {
  id: number;
  leetcode_id: number | null;
  title: string;
  url: string;
  topic: string;
  difficulty: Difficulty;
  is_microsoft: boolean;
  is_complete: boolean;
  completed_at: string | null;
}
