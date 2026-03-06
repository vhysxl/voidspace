import type { AuthResponse, ApiResponse } from "@/types/index";

export default defineEventHandler(async (event) => {
  type LoginResponse = ApiResponse<AuthResponse>;
  const config = useRuntimeConfig();
  const body = await readBody(event);

  try {
    const response = await fetch(`${config.apiUrl}/auth/login`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        "x-api-key": config.apiSecret,
      },
      body: JSON.stringify(body),
    });

    const data: LoginResponse = await response.json();

    // Forward Set-Cookie header dari backend
    const setCookieHeader = response.headers.get("set-cookie");
    if (setCookieHeader) {
      setHeader(event, "set-cookie", setCookieHeader);
    }

    if (!response.ok || !data.success) {
      throw createError({
        statusCode: response.status,
        statusMessage: data.detail || "Login failed",
      });
    }

    return data;
  } catch (error: any) {
    throw createError({
      statusCode: error.statusCode || 500,
      statusMessage: error.statusMessage || error.message || "Login failed",
    });
  }
});
