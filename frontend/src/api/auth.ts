import { http } from "./http";

export function login(data: any) {
  return http("/api/login", {
    method: "POST",
    body: JSON.stringify(data),
  });
}

export function register(data: any) {
  return http("/api/register", {
    method: "POST",
    body: JSON.stringify(data),
  });
}