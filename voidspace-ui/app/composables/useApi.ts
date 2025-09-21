import { useAuth } from "./useAuth";
import { isValidTokenResponse } from "@/utils/validation";

/**
 * Custom hook that provides an authenticated fetch function
 * Handles token expiration, automatic refresh, and fallback retry logic
 */
export const useApi = () => {
  const auth = useAuth();
  const authStore = useAuthStore();

  /**
   * Authenticated fetch wrapper that automatically handles token management
   * @param url - The endpoint URL to fetch from
   * @param options - Fetch options (headers, method, body, etc.)
   * @returns Promise with the fetch response
   */
  const fetchWithAuth = async <T>(url: string, options: any = {}) => {
    // Check if current token is expired before making any request
    const isExpired = Date.now() >= authStore.expiresIn;

    // PROACTIVE TOKEN REFRESH: Refresh token before request if it's expired
    if (authStore.accessToken && isExpired) {
      try {
        // Attempt to refresh the access token
        const refreshResponse = await auth.refresh();

        // Validate the refresh response contains valid token data
        if (!isValidTokenResponse(refreshResponse.data)) {
          throw new Error("Invalid refresh response");
        }

        // Update auth store with new token and expiration time
        authStore.login(
          refreshResponse.data.access_token,
          refreshResponse.data.expires_in
        );
      } catch (error) {
        // Refresh failed - logout user and redirect to login
        await auth.logout();
        throw error;
      }
    }

    // MAIN REQUEST: Attempt the original fetch with current/refreshed token
    try {
      return await $fetch<T>(url, {
        ...options,
        headers: {
          ...options.headers,
          // Add Authorization header only if we have a valid token
          ...(authStore.accessToken
            ? { Authorization: `Bearer ${authStore.accessToken}` }
            : {}),
        },
      });
    } catch (error: any) {
      // FALLBACK RETRY: Handle cases where token was valid but server returned 401
      if (error.status === 401 && authStore.accessToken && !isExpired) {
        try {
          // Attempt one more token refresh
          const refreshResponse = await auth.refresh();

          // Validate the refreshed token
          if (!isValidTokenResponse(refreshResponse.data)) {
            // Invalid refresh response - logout and fail
            await auth.logout();
            throw new Error("Invalid access token");
          }

          // Update token in store
          authStore.accessToken = refreshResponse.data.access_token;

          // Retry the original request with the new token
          return await $fetch<T>(url, {
            ...options,
            headers: {
              ...options.headers,
              Authorization: `Bearer ${authStore.accessToken}`,
            },
          });
        } catch (refreshError) {
          // Final fallback failed - logout user
          await auth.logout();
          throw refreshError;
        }
      }

      // Re-throw any other errors (network issues, 404, 500, etc.)
      throw error;
    }
  };

  return { fetchWithAuth };
};
