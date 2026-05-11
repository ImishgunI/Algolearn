import { http } from "./http";

export async function me() {
  return http("/api/me");
}