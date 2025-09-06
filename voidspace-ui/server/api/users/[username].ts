export default defineEventHandler(async (event) => {
  const config = useRuntimeConfig();
  const username = event.context.params?.username;

  if (!username) {
    throw createError({ statusCode: 400, statusMessage: "Username required" });
  }

  try {
    const response = await $fetch(`${config.apiUrl}/users/${username}`, {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
        "x-api-key": config.apiSecret,
      },
      credentials: "include",
    });

    return response;
  } catch (error: any) {
    throw createError({
      statusCode: error.status || 500,
      statusMessage:
        error.data?.detail || error.message || "Failed to get user",
    });
  }
});
