export interface AuthResponse {
  success: boolean;
  detail: string;
  data: {
    access_token: string;
    expires_in: number;
  };
}

export const useAuth = () => {
  const apiUrl = "/api";

  const login = async (
    usernameoremail: string,
    password: string
  ): Promise<AuthResponse> => {
    try {
      return await $fetch<AuthResponse>(`${apiUrl}/auth/login`, {
        method: "POST",
        body: { usernameoremail, password },
        headers: { "Content-Type": "application/json" },
        credentials: "include",
      });
    } catch (error: any) {
      console.log(error);
      throw new Error(error.statusMessage || "Login failed");
    }
  };

  const register = async (
    username: string,
    email: string,
    password: string
  ): Promise<AuthResponse> => {
    try {
      return await $fetch<AuthResponse>(`${apiUrl}/auth/register`, {
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
      throw new Error(
        error.data?.message || error.data?.detail || "Logout failed"
      );
    }
  };

  const refresh = async (): Promise<AuthResponse> => {
    try {
      return await $fetch(`${apiUrl}/auth/refresh`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        credentials: "include",
      });
    } catch (error: any) {
      throw new Error(
        error.data?.message || error.data?.detail || "Refresh failed"
      );
    }
  };

  return { login, register, logout, refresh };
};
