import { getToken } from "../utils/token";

export async function http(url: string, options: any = {}) {
  const res = await fetch(url, {
    ...options,
    headers: {
      "Content-Type": "application/json",
      ...(options.headers || {}),
      ...(getToken() && {
        Authorization: `Bearer ${getToken()}`
      }),
    },
  });

  if (!res.ok) {
    throw new Error("API error");
  }

  return res.json();
}