import { getToken } from "../utils/token";

export async function authFetch(url: string, options: any = {}) {
  return fetch(url, {
    ...options,
    headers: {
      ...options.headers,
      Authorization: `Bearer ${getToken()}`,
    },
  });
}