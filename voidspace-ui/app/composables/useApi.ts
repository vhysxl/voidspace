import { useAuth } from "./useAuth";

export const useApi = () => {
  const auth = useAuth();
  const authStore = useAuthStore();

  const isTokenExpired = (token: string) => {
    try {
      const payload = JSON.parse(atob(token.split(".")[1]!));
      const currentTime = Math.floor(Date.now() / 1000);
      return payload.exp < currentTime;
    } catch {
      return true;
    }
  };

  const fetchWithAuth = async (url: string, options: any = {}) => {
    let token = authStore.accessToken;

    // Check token expiry BEFORE making request
    if (token && isTokenExpired(token)) {
      try {
        const refreshResponse = await auth.refresh();
        token = refreshResponse.data.access_token;
        authStore.accessToken = token;
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
      if (error.status === 401 && token && !isTokenExpired(token)) {
        try {
          const refreshResponse = await auth.refresh();
          authStore.accessToken = refreshResponse.data.access_token;

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
