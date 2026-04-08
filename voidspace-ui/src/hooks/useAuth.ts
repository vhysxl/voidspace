import { apiFetch } from "@/lib/api";
import { useAuthStore } from "@/store/useAuthStore";
import { ApiResponse, AuthResponse, User } from "@/types";
import { useRouter } from "next/navigation";

export const useAuth = () => {
  const { setToken, setUser, logout: clearAuth } = useAuthStore();
  const router = useRouter();

  const login = async (usernameoremail: string, password: string) => {
    const response = await apiFetch<ApiResponse<AuthResponse>>("/auth/login", {
      method: "POST",
      body: JSON.stringify({ usernameoremail, password }),
    });

    if (response.success && response.data) {
      // Set the token immediately so subsequent call to /user/me uses it
      setToken(response.data.access_token, response.data.expires_in);

      // After login, we need to fetch user profile to populate store
      const userProfile = await apiFetch<ApiResponse<User>>("/user/me");

      if (userProfile.success && userProfile.data) {
        setUser(userProfile.data);
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
      setToken(response.data.access_token, response.data.expires_in);

      const userProfile = await apiFetch<ApiResponse<User>>("/user/me");

      if (userProfile.success && userProfile.data) {
        setUser(userProfile.data);
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
      router.push("/");
    }
  };

  return { login, register, logout };
};
