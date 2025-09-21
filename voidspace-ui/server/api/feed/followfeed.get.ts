import { authData } from "@/utils/authParser";

export default defineEventHandler(async (event) => {
  const config = useRuntimeConfig();

  const authCookie = getCookie(event, "auth");
  let token: string | undefined;

  if (authCookie) {
    const auth = authData(authCookie);
    token = auth.accessToken;
  } else {
    throw createError({ statusCode: 401, statusMessage: "Unauthorized" });
  }

  const query = getQuery(event);
  const queryString = new URLSearchParams(
    query as Record<string, string>
  ).toString();
  const url = `${config.apiUrl}/feed/following/${
    queryString ? `?${queryString}` : ""
  }`;

  try {
    const response = await $fetch(url, {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
        "x-api-key": config.apiSecret,
        ...(token ? { Authorization: `Bearer ${token}` } : {}),
      },
      credentials: "include",
    });

    return response;
  } catch (error: any) {
    throw createError({
      statusCode: error.status || 500,
      statusMessage: error.data?.detail || "Failed to get feed",
    });
  }
});
