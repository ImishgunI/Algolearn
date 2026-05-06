import { getToken } from "../utils/token";

export function requireAuth(next: any) {
  if (!getToken()) {
    next("/login");
  } else {
    next();
  }
}