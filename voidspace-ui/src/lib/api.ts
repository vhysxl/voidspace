import { useAuthStore } from "@/store/useAuthStore";

const PROXY_URL = "/api/proxy";
const SKIPPED_ROUTES = ["/auth/login", "/auth/register", "/auth/refresh"];

let isRefreshing = false;
let refreshPromise: Promise<string | null> | null = null;

interface FetchOptions extends RequestInit {
  params?: Record<string, string>;
}

async function refreshToken(): Promise<string | null> {
  try {
    const response = await fetch(`${PROXY_URL}/auth/refresh`, {
      method: "POST",
      credentials: "include", // send http-only cookie
      headers: {
        "Content-Type": "application/json",
      },
    });

    if (!response.ok) return null;

    const data = await response.json();
    if (data.success && data.data?.access_token) {
      useAuthStore.getState().setToken(data.data.access_token, data.data.expires_in);
      return data.data.access_token;
    }
    return null;
  } catch {
    return null;
  }
}

export async function apiFetch<T>(endpoint: string, options: FetchOptions = {}): Promise<T> {
  const isSkipped = SKIPPED_ROUTES.some((route) => endpoint.startsWith(route));
  
  const buildHeaders = (token?: string | null) => {
    const headers = new Headers(options.headers);
    headers.set("Content-Type", "application/json");
    
    // Only set Authorization if not skipped and not manually provided
    if (!isSkipped && token && !headers.has("Authorization")) {
      headers.set("Authorization", `Bearer ${token}`);
    }
    return headers;
  };

  const { accessToken } = useAuthStore.getState();

  const cleanEndpoint = endpoint.startsWith("/") ? endpoint : `/${endpoint}`;
  const url = new URL(`${window.location.origin}${PROXY_URL}${cleanEndpoint}`);
  
  if (options.params) {
    Object.entries(options.params).forEach(([key, value]) => url.searchParams.append(key, value));
  }

  // First attempt
  const response = await fetch(url.toString(), {
    ...options,
    headers: buildHeaders(accessToken),
    credentials: "include",
  });

  const isUnauthorized = response.status === 401 || response.status === 403;
  
  // If not unauthorized, or it's a skipped route (like login), handle normally
  if (!isUnauthorized || isSkipped) {
    const data = await response.json();
    if (!response.ok) {
      throw new Error(data.detail || "An error occurred");
    }
    return data as T;
  }

  // --- Retry with refresh token ---
  if (!isRefreshing) {
    isRefreshing = true;
    refreshPromise = refreshToken().finally(() => {
      isRefreshing = false;
    });
  }

  const newToken = await refreshPromise;

  if (!newToken) {
    useAuthStore.getState().logout();
    throw new Error("Session expired. Please login again.");
  }

  // Retry with new token
  const retryResponse = await fetch(url.toString(), {
    ...options,
    headers: buildHeaders(newToken),
    credentials: "include",
  });

  const retryData = await retryResponse.json();
  if (!retryResponse.ok) {
    throw new Error(retryData.detail || "An error occurred");
  }
  return retryData as T;
}
