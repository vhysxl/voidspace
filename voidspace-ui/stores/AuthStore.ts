import { defineStore } from "pinia";
import { jwtDecode } from "jwt-decode";
import { useUser } from "~/composables/useUser";
import type { User } from "@/types";

interface JWTClaims {
  ID: string;
  Username: string;
  exp: number;
}

function mapUser(resData: User): User {
  return {
    id: resData.id,
    username: resData.username,
    profile: {
      display_name: resData.profile.display_name,
      bio: resData.profile.bio,
      avatar_url: resData.profile.avatar_url,
      banner_url: resData.profile.banner_url,
      location: resData.profile.location,
      followers: resData.profile.followers,
      following: resData.profile.following,
    },
    created_at: resData.created_at,
    is_followed: false,
  };
}

export const useAuthStore = defineStore(
  "auth",
  () => {
    const user = ref<User | null>(null);
    const accessToken = ref("");
    const expiresIn = ref(0);
    const isLoggedIn = ref(false);
    const authApi = useAuth();
    const usersApi = useUser();

    async function login(token: string, expires: number) {
      accessToken.value = token;
      expiresIn.value = Date.now() + expires * 1000;

      const decoded = jwtDecode<JWTClaims>(token);

      // initially set user with decoded info
      user.value = {
        id: Number(decoded.ID),
        username: decoded.Username,
        is_followed: false,
        created_at: "",
        profile: {
          display_name: "",
          bio: "",
          avatar_url: "",
          location: "",
          banner_url: "",
          followers: 0,
          following: 0,
        },
      };

      isLoggedIn.value = true;

      // fetch full user data from API
      try {
        const res = await usersApi.getCurrentUser();
        if (res.success && res.data) {
          user.value = mapUser(res.data);
        }
      } catch (err) {
        console.error("Failed to fetch current user after login:", err);
      }
    }

    async function refreshUser() {
      if (!accessToken.value) return;

      try {
        const res = await usersApi.getCurrentUser();
        if (res.success && res.data) {
          user.value = mapUser(res.data);
          isLoggedIn.value = true;
        } else {
          logout();
        }
      } catch {
        logout();
      }
    }

    async function logout() {
      try {
        await authApi.logout();
      } catch (err) {
        console.error("Failed to fetch current user after login:", err);
      }
      user.value = null;
      isLoggedIn.value = false;
      accessToken.value = "";
      expiresIn.value = 0;
    }

    return {
      user,
      accessToken,
      expiresIn,
      isLoggedIn,
      login,
      logout,
      refreshUser,
    };
  },
  {
    persist: true,
  }
);
