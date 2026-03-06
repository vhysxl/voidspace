export const authData = (authCookie: string) => {
  try {
    const data = JSON.parse(decodeURIComponent(authCookie));
    return data;
  } catch (error) {
    console.error("Error parsing auth cookie:", error);
  }
};
