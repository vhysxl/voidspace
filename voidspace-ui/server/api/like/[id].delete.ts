import { authData } from "@/utils/authParser";

export default defineEventHandler(async (event) => {
  const config = useRuntimeConfig();
  const postId = getRouterParam(event, "id");
  const authCookie = getCookie(event, "auth");

  if (!authCookie) {
    throw createError({ statusCode: 401, statusMessage: "Unauthorized" });
  }

  const { accessToken } = authData(authCookie);

  try {
    return await $fetch(`${config.apiUrl}/likes/${postId}`, {
      method: "DELETE",
      headers: {
        "Content-Type": "application/json",
        "x-api-key": config.apiSecret,
        Authorization: `Bearer ${accessToken}`,
      },
      credentials: "include",
    });
  } catch (error: any) {
    throw createError({
      statusCode: error.status || 500,
      statusMessage: error.data?.detail || "Failed to unlike post",
    });
  }
});
