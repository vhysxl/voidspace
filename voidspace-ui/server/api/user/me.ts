import { authData } from "@/utils/authParser";

export default defineEventHandler(async (event) => {
  const config = useRuntimeConfig();
  const method = event.node.req.method?.toUpperCase();

  const authCookie = getCookie(event, "auth");
  let token;

  //decode cookie
  if (authCookie) {
    const auth = authData(authCookie);
    token = auth.accessToken;
  } else {
    throw createError({ statusCode: 401, statusMessage: "Unaothired" });
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
      const response = await $fetch(`${config.apiUrl}/users/me`, {
        method: "DELETE",
        headers: {
          "Content-Type": "application/json",
          "x-api-key": config.apiSecret,
          ...(token && { Authorization: `Bearer ${token}` }),
        },
        credentials: "include",
      });

      setCookie(event, "refresh_token", "", {
        expires: new Date(0),
        path: "/",
        httpOnly: true,
        secure: true,
        sameSite: "strict",
      });

      return response;
    }

    throw createError({ statusCode: 405, statusMessage: "Method not allowed" });
  } catch (error: any) {
    throw createError({
      statusCode: error.status || 500,
      statusMessage: error.data?.detail || "User request failed",
    });
  }
});
