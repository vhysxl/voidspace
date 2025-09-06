export default defineEventHandler(async (event) => {
  const config = useRuntimeConfig();
  const body = await readBody(event);

  try {
    const response = await fetch(`${config.apiUrl}/auth/login`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        "x-api-key": config.apiSecret,
      },
      body: JSON.stringify(body),
    });

    const data = await response.json();

    // Forward Set-Cookie header dari backend
    const setCookieHeader = response.headers.get("set-cookie");
    if (setCookieHeader) {
      setHeader(event, "set-cookie", setCookieHeader);
    }

    console.log(data.detail);

    //langsung masuk block catch pake custom biar bisa baca astoge
    if (!response.ok) {
      throw { status: response.status, detail: data.detail || "Login failed" };
    }

    return data;
  } catch (error: any) {
    throw createError({
      statusCode: error.status || 500,
      statusMessage: error.detail || error.message || "Login failed",
    });
  }
});
