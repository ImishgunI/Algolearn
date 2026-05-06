import { http } from "./http";

export async function me() {
  const res = await http("/api/me");
  return res.json();
}