import { apiFetch } from "@/lib/api";
import { useAuthStore } from "@/store/useAuthStore";
import { ApiResponse, AuthResponse, User } from "@/types";
import { useRouter } from "next/navigation";

export const useAuth = () => {
  const { setAuth, logout: clearAuth } = useAuthStore();
  const router = useRouter();

  const login = async (usernameoremail: string, password: string) => {
    const response = await apiFetch<ApiResponse<AuthResponse>>("/auth/login", {
      method: "POST",
      body: JSON.stringify({ usernameoremail, password }),
    });

    if (response.success && response.data) {
      // After login, we need to fetch user profile to populate store
      const userProfile = await apiFetch<ApiResponse<User>>("/user/me", {
        headers: {
          Authorization: `Bearer ${response.data.access_token}`,
        },
      });

      if (userProfile.success && userProfile.data) {
        setAuth(response.data.access_token, response.data.expires_in, userProfile.data);
        router.push("/");
      }
    }
    return response;
  };

  const register = async (username: string, email: string, password: string) => {
    const response = await apiFetch<ApiResponse<AuthResponse>>("/auth/register", {
      method: "POST",
      body: JSON.stringify({ username, email, password }),
    });

    if (response.success && response.data) {
      const userProfile = await apiFetch<ApiResponse<User>>("/user/me", {
        headers: {
          Authorization: `Bearer ${response.data.access_token}`,
        },
      });

      if (userProfile.success && userProfile.data) {
        setAuth(response.data.access_token, response.data.expires_in, userProfile.data);
        router.push("/");
      }
    }
    return response;
  };

  const logout = async () => {
    try {
      await apiFetch("/auth/logout", { method: "POST" });
    } finally {
      clearAuth();
      router.push("/auth/login");
    }
  };

  return { login, register, logout };
};
