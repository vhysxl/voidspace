import { ApiResponse, User } from "@/types";

export default defineEventHandler(async (event) => {
  const config = useRuntimeConfig();
  const username = getRouterParam(event, "username");

  try {
    const response = await $fetch<ApiResponse<User>>(
      `${config.apiUrl}/comment/user/${username}`,
      {
        method: "GET",
        headers: {
          "Content-Type": "application/json",
          "x-api-key": config.apiSecret,
        },
      }
    );

    return response;
  } catch (error: any) {
    throw createError({
      statusCode: error.status || 500,
      statusMessage: error.data?.detail || "Get User Comments Failed",
    });
  }
});
