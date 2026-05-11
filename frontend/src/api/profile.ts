import { http } from "./http";

export async function updateProfile(data: { user_name: string; user_surname: string; email: string }) {
  return http("/api/profile", {
    method: "PUT",
    body: JSON.stringify(data),
  });
}

export async function updatePassword(data: { old_password: string; new_password: string }) {
  return http("/api/profile/password", {
    method: "PUT",
    body: JSON.stringify(data),
  });
}

export async function getStats() {
  return http("/api/profile/stats");
}

export async function getFavorites() {
  return http("/api/profile/favorites");
}