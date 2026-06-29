"use client";

import { useEffect, useMemo, useState } from "react";
import type { Difficulty, Question } from "@/lib/types";

const DIFF_ORDER: Record<Difficulty, number> = { Easy: 0, Medium: 1, Hard: 2 };

const DIFF_BADGE: Record<Difficulty, string> = {
  Easy: "bg-green-100 text-green-700",
  Medium: "bg-amber-100 text-amber-700",
  Hard: "bg-red-100 text-red-700",
};

type SortKey = "topic" | "difficulty" | "title" | "status";

export default function Home() {
  const [questions, setQuestions] = useState<Question[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  // filters
  const [topic, setTopic] = useState("All");
  const [difficulty, setDifficulty] = useState("All");
  const [status, setStatus] = useState("All"); // All | Done | Todo
  const [microsoftOnly, setMicrosoftOnly] = useState(false);
  const [search, setSearch] = useState("");

  // sorting
  const [sortKey, setSortKey] = useState<SortKey>("topic");
  const [sortDir, setSortDir] = useState<"asc" | "desc">("asc");

  useEffect(() => {
    fetch("/api/questions")
      .then((r) => {
        if (!r.ok) throw new Error("Failed to load questions");
        return r.json();
      })
      .then((data: Question[]) => setQuestions(data))
      .catch((e: Error) => setError(e.message))
      .finally(() => setLoading(false));
  }, []);

  const topics = useMemo(
    () => ["All", ...Array.from(new Set(questions.map((q) => q.topic)))],
    [questions]
  );

  async function toggle(id: number) {
    const before = questions;
    // optimistic update
    setQuestions((prev) =>
      prev.map((q) =>
        q.id === id
          ? {
              ...q,
              is_complete: !q.is_complete,
              completed_at: !q.is_complete ? new Date().toISOString() : null,
            }
          : q
      )
    );
    try {
      const res = await fetch(`/api/questions/${id}`, { method: "PATCH" });
      if (!res.ok) throw new Error();
      const updated = await res.json();
      setQuestions((prev) =>
        prev.map((q) => (q.id === id ? { ...q, ...updated } : q))
      );
    } catch {
      setQuestions(before); // revert on failure
    }
  }

  function setSort(key: SortKey) {
    if (key === sortKey) {
      setSortDir((d) => (d === "asc" ? "desc" : "asc"));
    } else {
      setSortKey(key);
      setSortDir("asc");
    }
  }

  const filtered = useMemo(() => {
    const list = questions.filter((q) => {
      if (topic !== "All" && q.topic !== topic) return false;
      if (difficulty !== "All" && q.difficulty !== difficulty) return false;
      if (status === "Done" && !q.is_complete) return false;
      if (status === "Todo" && q.is_complete) return false;
      if (microsoftOnly && !q.is_microsoft) return false;
      if (search && !q.title.toLowerCase().includes(search.toLowerCase()))
        return false;
      return true;
    });

    list.sort((a, b) => {
      let cmp = 0;
      switch (sortKey) {
        case "difficulty":
          cmp = DIFF_ORDER[a.difficulty] - DIFF_ORDER[b.difficulty];
          break;
        case "title":
          cmp = a.title.localeCompare(b.title);
          break;
        case "status":
          cmp = Number(a.is_complete) - Number(b.is_complete);
          break;
        case "topic":
        default:
          cmp =
            a.topic.localeCompare(b.topic) ||
            (a.leetcode_id ?? 99999) - (b.leetcode_id ?? 99999);
      }
      return sortDir === "asc" ? cmp : -cmp;
    });

    return list;
  }, [questions, topic, difficulty, status, microsoftOnly, search, sortKey, sortDir]);

  const doneCount = questions.filter((q) => q.is_complete).length;
  const total = questions.length;
  const pct = total ? Math.round((doneCount / total) * 100) : 0;

  const arrow = (key: SortKey) =>
    sortKey === key ? (sortDir === "asc" ? " ▲" : " ▼") : "";

  return (
    <main className="mx-auto max-w-6xl px-4 py-8">
      <header className="mb-6">
        <h1 className="text-2xl font-bold">DSA Tracker</h1>
        <p className="text-sm text-slate-500">Microsoft SWE II prep</p>

        <div className="mt-4">
          <div className="mb-1 flex justify-between text-sm">
            <span className="font-medium">
              {doneCount} / {total} solved
            </span>
            <span className="text-slate-500">{pct}%</span>
          </div>
          <div className="h-2 w-full overflow-hidden rounded bg-slate-200">
            <div
              className="h-full bg-emerald-500 transition-all"
              style={{ width: `${pct}%` }}
            />
          </div>
        </div>
      </header>

      {/* Filters */}
      <div className="mb-4 flex flex-wrap items-center gap-3">
        <select
          className="rounded border border-slate-300 bg-white px-3 py-1.5 text-sm"
          value={topic}
          onChange={(e) => setTopic(e.target.value)}
        >
          {topics.map((t) => (
            <option key={t} value={t}>
              {t === "All" ? "All topics" : t}
            </option>
          ))}
        </select>

        <select
          className="rounded border border-slate-300 bg-white px-3 py-1.5 text-sm"
          value={difficulty}
          onChange={(e) => setDifficulty(e.target.value)}
        >
          {["All", "Easy", "Medium", "Hard"].map((d) => (
            <option key={d} value={d}>
              {d === "All" ? "All difficulties" : d}
            </option>
          ))}
        </select>

        <select
          className="rounded border border-slate-300 bg-white px-3 py-1.5 text-sm"
          value={status}
          onChange={(e) => setStatus(e.target.value)}
        >
          <option value="All">All statuses</option>
          <option value="Todo">Not done</option>
          <option value="Done">Done</option>
        </select>

        <label className="flex items-center gap-1.5 text-sm">
          <input
            type="checkbox"
            checked={microsoftOnly}
            onChange={(e) => setMicrosoftOnly(e.target.checked)}
          />
          ⭐ Microsoft only
        </label>

        <input
          className="ml-auto rounded border border-slate-300 bg-white px-3 py-1.5 text-sm"
          placeholder="Search title…"
          value={search}
          onChange={(e) => setSearch(e.target.value)}
        />
      </div>

      {loading && <p className="text-slate-500">Loading…</p>}
      {error && (
        <p className="rounded bg-red-50 px-4 py-3 text-red-700">
          {error}. Is the database running and seeded? See the README.
        </p>
      )}

      {!loading && !error && (
        <>
          <p className="mb-2 text-sm text-slate-500">
            Showing {filtered.length} question{filtered.length === 1 ? "" : "s"}
          </p>
          <div className="overflow-hidden rounded-lg border border-slate-200 bg-white">
            <table className="w-full text-left text-sm">
              <thead className="border-b border-slate-200 bg-slate-50 text-slate-600">
                <tr>
                  <th className="w-12 px-4 py-2">Done</th>
                  <th className="w-16 px-2 py-2">#</th>
                  <th
                    className="cursor-pointer px-2 py-2 hover:text-slate-900"
                    onClick={() => setSort("title")}
                  >
                    Title{arrow("title")}
                  </th>
                  <th
                    className="cursor-pointer px-2 py-2 hover:text-slate-900"
                    onClick={() => setSort("topic")}
                  >
                    Topic{arrow("topic")}
                  </th>
                  <th
                    className="cursor-pointer px-2 py-2 hover:text-slate-900"
                    onClick={() => setSort("difficulty")}
                  >
                    Difficulty{arrow("difficulty")}
                  </th>
                </tr>
              </thead>
              <tbody>
                {filtered.map((q) => (
                  <tr
                    key={q.id}
                    className="border-b border-slate-100 last:border-0 hover:bg-slate-50"
                  >
                    <td className="px-4 py-2">
                      <input
                        type="checkbox"
                        className="h-4 w-4 cursor-pointer accent-emerald-500"
                        checked={q.is_complete}
                        onChange={() => toggle(q.id)}
                      />
                    </td>
                    <td className="px-2 py-2 text-slate-400">
                      {q.leetcode_id ?? "—"}
                    </td>
                    <td className="px-2 py-2">
                      <a
                        href={q.url}
                        target="_blank"
                        rel="noopener noreferrer"
                        className={`hover:underline ${
                          q.is_complete
                            ? "text-slate-400 line-through"
                            : "text-slate-900"
                        }`}
                      >
                        {q.title}
                      </a>
                      {q.is_microsoft && <span title="Microsoft-frequent"> ⭐</span>}
                    </td>
                    <td className="px-2 py-2 text-slate-600">{q.topic}</td>
                    <td className="px-2 py-2">
                      <span
                        className={`rounded-full px-2 py-0.5 text-xs font-medium ${
                          DIFF_BADGE[q.difficulty]
                        }`}
                      >
                        {q.difficulty}
                      </span>
                    </td>
                  </tr>
                ))}
              </tbody>
            </table>
          </div>
        </>
      )}
    </main>
  );
}
