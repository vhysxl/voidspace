import { useAuth } from "./useAuth";

export const useApi = () => {
  const auth = useAuth();
  const authStore = useAuthStore();

  const fetchWithAuth = async (url: string, options: any = {}) => {
    let token = authStore.accessToken;
    const isExpired = Date.now() >= authStore.expiresIn;

    // Check token expiry BEFORE making request
    if (token && isExpired) {
      try {
        const refreshResponse = await auth.refresh();
        const token = refreshResponse.data?.access_token;
        const expiredIn = refreshResponse.data?.expires_in;

        if (
          typeof token !== "string" ||
          token.trim() === "" ||
          typeof expiredIn !== "number"
        ) {
          throw new Error("Invalid access token");
        }

        authStore.login(token, expiredIn);
      } catch (refreshError) {
        // Refresh failed, redirect to login
        await auth.logout();
        throw refreshError;
      }
    }

    try {
      return await $fetch(url, {
        ...options,
        headers: {
          ...options.headers,
          ...(token && { Authorization: `Bearer ${token}` }),
        },
      });
    } catch (error: any) {
      // Fallback: jika masih 401, coba refresh sekali lagi
      if (error.status === 401 && token && !isExpired) {
        try {
          const refreshResponse = await auth.refresh();
          const token = refreshResponse.data?.access_token;

          if (typeof token !== "string" || token.trim() === "") {
            throw new Error("Invalid access token");
          }

          authStore.accessToken = token;

          return await $fetch(url, {
            ...options,
            headers: {
              ...options.headers,
              Authorization: `Bearer ${authStore.accessToken}`,
            },
          });
        } catch (refreshError) {
          await auth.logout();
          throw refreshError;
        }
      }
      throw error;
    }
  };

  return { fetchWithAuth };
};
