import { authFetch } from "./http";

export async function me() {
  const res = await authFetch("/api/me");
  return res.json();
}