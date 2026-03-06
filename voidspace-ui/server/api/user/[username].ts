import { authData } from "@/utils/authParser";

export default defineEventHandler(async (event) => {
  const config = useRuntimeConfig();
  const username = event.context.params?.username;

  const authCookie = getCookie(event, "auth");
  let token;

  if (!username) {
    throw createError({ statusCode: 400, statusMessage: "Username required" });
  }

  //decode cookie
  if (authCookie) {
    const auth = authData(authCookie);
    {
    }
    token = auth.accessToken;
  } else {
    throw createError({ statusCode: 401, statusMessage: "Unauthorized" });
  }

  try {
    const response = await $fetch(`${config.apiUrl}/users/${username}`, {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
        "x-api-key": config.apiSecret,
        ...(token && { Authorization: `Bearer ${token}` }),
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
