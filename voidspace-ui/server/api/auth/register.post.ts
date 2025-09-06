export default defineEventHandler(async (event) => {
  const config = useRuntimeConfig();
  const body = await readBody(event);

  try {
    const response = await fetch(`${config.apiUrl}/auth/register`, {
      method: "POST",
      body: JSON.stringify(body),
      headers: {
        "Content-Type": "application/json",
        "x-api-key": config.apiSecret,
      },
    });

    const data = await response.json();

    const setCookieHeader = response.headers.get("set-cookie");
    if (setCookieHeader) {
      setHeader(event, "set-cookie", setCookieHeader);
    }

    if (!response.ok) {
      throw {
        status: response.status,
        detail: data.detail || "Register failed",
      };
    }

    return data;
  } catch (error: any) {
    throw createError({
      statusCode: error.status || 500,
      statusMessage: error.detail || error.message || "Register failed",
    });
  }
});
