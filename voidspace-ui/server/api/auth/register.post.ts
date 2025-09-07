import type { AuthResponse, ApiResponse } from "@/types/index";

export default defineEventHandler(async (event) => {
  type RegisterResponse = ApiResponse<AuthResponse>;
  const config = useRuntimeConfig();
  const body = await readBody(event);

  try {
    const response = await fetch(`${config.apiUrl}/auth/register`, {
      method: "POST",
      body: JSON.stringify(body),
      headers: {
        "Content-Type": "application/json",
        "x-api-key": config.apiSecret,
      },
    });

    const data: RegisterResponse = await response.json();

    const setCookieHeader = response.headers.get("set-cookie");
    if (setCookieHeader) {
      setHeader(event, "set-cookie", setCookieHeader);
    }

    if (!response.ok || !data.success) {
      throw createError({
        statusCode: response.status,
        statusMessage: data.detail || "Register failed",
      });
    }

    const successData = data as RegisterResponse;
    return successData;
  } catch (error: any) {
    throw createError({
      statusCode: error.status || 500,
      statusMessage: error.detail || error.message || "Register failed",
    });
  }
});
