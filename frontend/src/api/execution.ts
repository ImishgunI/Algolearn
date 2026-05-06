import { authFetch } from "./http";

export async function executeAlgorithm(data: number[]) {
  const res = await fetch("/api/execute", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
      algorithm: "bubble_sort",
      data,
    }),
  });

  return res.json();
}

export function getExecution(id: string) {
  return authFetch(`/api/execution/${id}`);
}