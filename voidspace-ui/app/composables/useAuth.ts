import type { ApiResponse, AuthResponse } from "@/types/index";
import { handleApiError } from "@/utils/apiErrorHandler";
import { defaultOptions } from "@/utils/apiDefaults";

/**
 * Authentication service hooks
 * Wraps API calls related to authentication (login, register, logout, refresh).
 * Each method returns a typed ApiResponse or throws an error handled by handleApiError.
 */

type UseAuthResponse = ApiResponse<AuthResponse>;

export const useAuth = () => {
  const apiUrl = "/api/auth";

  const login = async (
    usernameoremail: string,
    password: string
  ): Promise<UseAuthResponse> => {
    try {
      return await $fetch<UseAuthResponse>(`${apiUrl}/login`, {
        ...defaultOptions,
        method: "POST",
        body: { usernameoremail, password },
      });
    } catch (error: unknown) {
      handleApiError(error, "Login failed");
    }
  };

  const register = async (
    username: string,
    email: string,
    password: string
  ): Promise<UseAuthResponse> => {
    try {
      return await $fetch<UseAuthResponse>(`${apiUrl}/register`, {
        ...defaultOptions,
        method: "POST",
        body: { username, email, password },
      });
    } catch (error: unknown) {
      handleApiError(error, "Register failed");
    }
  };

  const logout = async (): Promise<void> => {
    try {
      await $fetch(`${apiUrl}/logout`, {
        ...defaultOptions,
        method: "POST",
      });
    } catch (error: unknown) {
      handleApiError(error, "Logout failed");
    }
  };

  const refresh = async (): Promise<UseAuthResponse> => {
    try {
      return await $fetch<UseAuthResponse>(`${apiUrl}/refresh`, {
        ...defaultOptions,
        method: "POST",
      });
    } catch (error: unknown) {
      handleApiError(error, "Refresh token failed");
    }
  };

  return { login, register, logout, refresh };
};
