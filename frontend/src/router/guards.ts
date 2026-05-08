import { getToken } from "../utils/token";

export function requireAuth(_to: any, _from: any, next: any) {
  if (!getToken()) {
    next("/login");
  } else {
    next();
  }
}