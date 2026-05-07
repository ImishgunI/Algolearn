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

  if (res.status === 204 || res.status === 304) {
    return null;
  }

  const contentType = res.headers.get("content-type") || "";

  if (contentType.includes("application/json")) {
    const text = await res.text();
    return text ? JSON.parse(text) : null;
  }

  return res.text();
}