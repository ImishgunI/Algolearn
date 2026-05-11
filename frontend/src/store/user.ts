import { ref } from "vue";
import { me } from "../api/user";
import { clearToken } from "../utils/token";
import router from "../router";

export const user = ref<any>(null);

export async function fetchUser() {
  try {
    user.value = await me();
  } catch (err) {
    console.error("Failed to fetch user:", err);
    user.value = null;
  }
}

export function logout() {
  clearToken();
  user.value = null;
  router.push("/login");
}

export function isAdmin() {
  return user.value?.role === "admin" || user.value?.user_role === "admin";
}