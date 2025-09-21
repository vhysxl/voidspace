export default defineEventHandler(async (event) => {
  const config = useRuntimeConfig();

  try {
    return await $fetch(`${config.apiUrl}/auth/logout`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        "x-api-key": config.apiSecret,
      },
      credentials: "include",
    });
  } catch (error: any) {
    throw createError({
      statusCode: error.status || 500,
      statusMessage: error.data?.detail || "Logout failed",
    });
  }
});
