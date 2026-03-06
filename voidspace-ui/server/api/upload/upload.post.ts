import { authData } from "@/utils/authParser";

export default defineEventHandler(async (event) => {
  const config = useRuntimeConfig();
  const body = await readBody(event);
  const authCookie = getCookie(event, "auth");
  let token = null;

  //decode cookie
  if (authCookie) {
    const auth = authData(authCookie);
    token = auth.accessToken;
  } else {
    throw createError({ statusCode: 401, statusMessage: "Unaothired" });
  }

  try {
    const response = await $fetch(`${config.apiUrl}/upload/signed-url`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        "x-api-key": config.apiSecret,
        ...(token && { Authorization: `Bearer ${token}` }),
      },
      body,
      credentials: "include",
    });

    return response;
  } catch (error: any) {
    throw createError({
      statusCode: error.status || 500,
      statusMessage: error.data?.detail || "Upload failed",
    });
  }
});
