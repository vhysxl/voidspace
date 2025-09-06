export default defineEventHandler(async (event) => {
  const config = useRuntimeConfig();
  const method = event.node.req.method?.toUpperCase();

  const authCookie = getCookie(event, "auth");
  let token = null;

  //decode cookie
  if (authCookie) {
    try {
      const authData = JSON.parse(decodeURIComponent(authCookie));
      token = authData.accessToken;
    } catch (error) {
      console.error("Error parsing auth cookie:", error);
    }
  }

  try {
    if (method === "GET") {
      return await $fetch(`${config.apiUrl}/users/me`, {
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
      const response = await $fetch(`${config.apiUrl}/users/me`, {
        method: "PUT",
        headers: {
          "Content-Type": "application/json",
          "x-api-key": config.apiSecret,
          ...(token && { Authorization: `Bearer ${token}` }),
        },
        body: body,
        credentials: "include",
      });

      return response;
    }

    if (method === "DELETE") {
      return await $fetch(`${config.apiUrl}/users/me`, {
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
      statusMessage:
        error.data?.detail || error.message || "User request failed",
    });
  }
});
