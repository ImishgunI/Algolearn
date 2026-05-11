import { http } from "./http";

export async function executeAlgorithm(algorithm_name: string, data: number[]) {
   return http('/api/execute', {
    method: "POST",
    body: JSON.stringify({
      algorithm: algorithm_name,
      data: data,
    })
   })
}

export function getExecution(id: string) {
  return http(`/api/execution/${id}`);
}