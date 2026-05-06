import { clearToken } from "../utils/token";
import router from "../router";

export async function register(data: any) {
  return fetch("/api/register", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(data),
  });
}

export async function login(data: any) {
  const res = await fetch("/api/login", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(data),
  });

  return res.json();
}

export function logout() {
  clearToken();
  router.push("/login");
}