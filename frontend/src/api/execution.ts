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

export async function getExecution(id: string) {
  const res = await fetch(`/api/execution/${id}`);
  return res.json();
}