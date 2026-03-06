import { ApiResponse } from "@/types";
import { authData } from "@/utils/authParser";

export default defineEventHandler(async (event) => {
  const config = useRuntimeConfig();
  const authCookie = getCookie(event, "auth");
  const commentId = getRouterParam(event, "id");
  let token;

  //decode cookie
  if (authCookie) {
    const auth = authData(authCookie);
    token = auth.accessToken;
  } else {
    throw createError({ statusCode: 401, statusMessage: "Unauthorized" });
  }

  try {
    return await $fetch<ApiResponse>(`${config.apiUrl}/comment/${commentId}`, {
      method: "DELETE",
      headers: {
        "Content-Type": "application/json",
        "x-api-key": config.apiSecret,
        ...(token && { Authorization: `Bearer ${token}` }),
      },
      credentials: "include",
    });
  } catch (error: any) {
    throw createError({
      statusCode: error.status || 500,
      statusMessage: error.data?.detail || "Delete Comment Failed",
    });
  }
});
