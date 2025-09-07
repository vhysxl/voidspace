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

    //langsung masuk block catch pake custom biar bisa baca astoge
    if (!response.ok || !data.success) {
      throw createError({
        statusCode: response.status,
        statusMessage: data.detail || "Login failed",
      });
    }

    const successData = data as LoginResponse;
    return successData;
  } catch (error: any) {
    throw createError({
      statusCode: error.status || 500,
      statusMessage: error.detail || "Login failed",
    });
  }
});
