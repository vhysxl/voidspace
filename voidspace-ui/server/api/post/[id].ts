import { authData } from "@/utils/authParser";

export default defineEventHandler(async (event) => {
  const config = useRuntimeConfig();
  const method = event.node.req.method?.toUpperCase();
  const authCookie = getCookie(event, "auth");
  const postId = event.context.params?.id;
  let token;

  //decode cookie
  if (authCookie) {
    const auth = authData(authCookie);
    token = auth.accessToken;
  }

  try {
    if (method === "GET") {
      return await $fetch(`${config.apiUrl}/posts/${postId}`, {
        method: "GET",
        headers: {
          "Content-Type": "application/json",
          "x-api-key": config.apiSecret,
          ...(token && { Authorization: `Bearer ${token}` }),
        },
        credentials: "include",
      });
    }

    if (method === "PUT") {
      const body = await readBody(event);
      return await $fetch(`${config.apiUrl}/posts/${postId}`, {
        method: "PUT",
        body: body,
        headers: {
          "Content-Type": "application/json",
          "x-api-key": config.apiSecret,
          ...(token && { Authorization: `Bearer ${token}` }),
        },
        credentials: "include",
      });
    }

    if (method === "DELETE") {
      return await $fetch(`${config.apiUrl}/posts/${postId}`, {
        method: "DELETE",
        headers: {
          "Content-Type": "application/json",
          "x-api-key": config.apiSecret,
          ...(token && { Authorization: `Bearer ${token}` }),
        },
        credentials: "include",
      });
    }

    throw createError({ statusCode: 405, statusMessage: "Method not allowed" });
  } catch (error: any) {
    throw createError({
      statusCode: error.status || 500,
      statusMessage: error.data?.detail || "Posts request failed",
    });
  }
});
