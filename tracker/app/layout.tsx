import type { Metadata } from "next";
import "./globals.css";

export const metadata: Metadata = {
  title: "DSA Tracker — Microsoft SWE II Prep",
  description: "Track LeetCode questions by topic and difficulty.",
};

export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <html lang="en">
      <body>{children}</body>
    </html>
  );
}
