import { http } from "./http";

export async function executeAlgorithm(data: number[]) {
   return http('/api/execute', {
    method: "POST",
    body: JSON.stringify({
      algorithm: "bubble_sort",
      data: data,
    })
   })
}

export function getExecution(id: string) {
  return http(`/api/execution/${id}`);
}