import { ApiResponse, Post } from "@/types";

export default defineEventHandler(async (event) => {
  const config = useRuntimeConfig();
  const postId = getRouterParam(event, "id");

  try {
    const response = await $fetch<ApiResponse<Post>>(
      `${config.apiUrl}/comment/post/${postId}`,
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
      statusMessage: error.data?.detail || "Get Post Comments Failed",
    });
  }
});
