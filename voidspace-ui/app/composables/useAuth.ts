import type { ApiResponse, AuthResponse } from "@/types/index";

export const useAuth = () => {
  const apiUrl = "/api";
  type useAuthResponse = ApiResponse<AuthResponse>;

  const login = async (
    usernameoremail: string,
    password: string
  ): Promise<useAuthResponse> => {
    try {
      return await $fetch<useAuthResponse>(`${apiUrl}/auth/login`, {
        method: "POST",
        body: { usernameoremail, password },
        headers: { "Content-Type": "application/json" },
        credentials: "include",
      });
    } catch (error: any) {
      throw new Error(error.statusMessage || "Login failed");
    }
  };

  const register = async (
    username: string,
    email: string,
    password: string
  ): Promise<useAuthResponse> => {
    try {
      return await $fetch<useAuthResponse>(`${apiUrl}/auth/register`, {
        method: "POST",
        body: { username, email, password },
        headers: { "Content-Type": "application/json" },
        credentials: "include",
      });
    } catch (error: any) {
      throw new Error(error.statusMessage || "Register failed");
    }
  };

  const logout = async () => {
    try {
      return await $fetch(`${apiUrl}/auth/logout`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        credentials: "include",
      });
    } catch (error: any) {
      throw new Error(error.statusMessage || "Logout failed");
    }
  };

  const refresh = async (): Promise<useAuthResponse> => {
    try {
      return await $fetch<useAuthResponse>(`${apiUrl}/auth/refresh`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        credentials: "include",
      });
    } catch (error: any) {
      throw new Error(error.statusMessage || "Refresh failed");
    }
  };

  return { login, register, logout, refresh };
};
