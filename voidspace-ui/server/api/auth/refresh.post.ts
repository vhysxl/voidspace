import type { AuthResponse, ApiResponse } from "@/types/index";

export default defineEventHandler(async (event) => {
  const config = useRuntimeConfig();
  type RefreshResponse = ApiResponse<AuthResponse>;

  try {
    const clientCookies = getHeader(event, "cookie");

    const response: RefreshResponse = await $fetch(
      `${config.apiUrl}/auth/refresh`,
      {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          "x-api-key": config.apiSecret,
          ...(clientCookies && { cookie: clientCookies }),
        },
        credentials: "include",
      }
    );

    return response;
  } catch (error: any) {
    throw createError({
      statusCode: error.status || 500,
      statusMessage: error.data?.detail || "Refresh failed",
    });
  }
});
