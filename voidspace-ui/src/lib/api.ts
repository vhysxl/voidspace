import { useAuthStore } from "@/store/useAuthStore";

const BASE_URL = process.env.NEXT_PUBLIC_API_URL || "http://localhost:5000/api/v2";
const API_KEY = process.env.NEXT_PUBLIC_API_KEY;

interface FetchOptions extends RequestInit {
  params?: Record<string, string>;
}

export async function apiFetch<T>(endpoint: string, options: FetchOptions = {}): Promise<T> {
  if (!API_KEY) {
    throw new Error("NEXT_PUBLIC_API_KEY is not defined");
  }

  const { accessToken } = useAuthStore.getState();

  const headers = new Headers(options.headers);
  headers.set("x-api-key", API_KEY);
  headers.set("Content-Type", "application/json");

  if (accessToken) {
    headers.set("Authorization", `Bearer ${accessToken}`);
  }

  const url = new URL(`${BASE_URL}${endpoint.startsWith("/") ? endpoint : `/${endpoint}`}`);
  if (options.params) {
    Object.entries(options.params).forEach(([key, value]) => url.searchParams.append(key, value));
  }

  const response = await fetch(url.toString(), {
    ...options,
    headers,
    credentials: "include",
  });

  const data = await response.json();

  if (!response.ok) {
    throw new Error(data.detail || "An error occurred");
  }

  return data as T;
}
