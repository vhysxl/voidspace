import { authData } from "@/utils/authParser";

export default defineEventHandler(async (event) => {
  const config = useRuntimeConfig();
  const authCookie = getCookie(event, "auth");
  const username = getRouterParam(event, "username");
  let token;

  //decode cookie
  if (authCookie) {
    const auth = authData(authCookie);
    token = auth.accessToken;
  } else {
    throw createError({ statusCode: 401, statusMessage: "Unauthorized" });
  }

  try {
    const response = await $fetch(`${config.apiUrl}/follow/${username}`, {
      method: "POST",
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
        error.data?.detail || "Follow action failed, Please try again later",
    });
  }
});
