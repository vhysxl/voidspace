import { authData } from "@/utils/authParser";

export default defineEventHandler(async (event) => {
  const config = useRuntimeConfig();

  const authCookie = getCookie(event, "auth");
  let token;

  //decode cookie
  if (authCookie) {
    const auth = authData(authCookie);
    token = auth.accessToken;
  }

  try {
    const response = await $fetch(`${config.apiUrl}/feed/`, {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
        "x-api-key": config.apiSecret,
        ...(token ? { Authorization: `Bearer ${token}` } : {}),
      },
      credentials: "include",
    });

    return response;
  } catch (error: any) {
    throw createError({
      statusCode: error.status || 500,
      statusMessage: error.data?.detail || "Failed to get feed",
    });
  }
});
