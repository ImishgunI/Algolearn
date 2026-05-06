import { ref } from "vue";
import { me } from "../api/user";

export const user = ref<any>(null);

export async function fetchUser() {
  try {
    user.value = await me();
  } catch {
    user.value = null;
  }
}

export function isAdmin() {
  return user.value?.role === "Admin";
}