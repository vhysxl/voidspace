import { create } from "zustand";
import { persist } from "zustand/middleware";
import { User } from "@/types";

interface AuthState {
  user: User | null;
  accessToken: string | null;
  expiresAt: number | null;
  isLoggedIn: boolean;
  
  setAuth: (token: string, expires_in: number, user: User) => void;
  updateUser: (user: User) => void;
  logout: () => void;
  clearAuth: () => void;
}

export const useAuthStore = create<AuthState>()(
  persist(
    (set) => ({
      user: null,
      accessToken: null,
      expiresAt: null,
      isLoggedIn: false,

      setAuth: (token, expires_in, user) => {
        set({
          accessToken: token,
          expiresAt: Date.now() + expires_in * 1000,
          user,
          isLoggedIn: true,
        });
      },

      updateUser: (user) => {
        set({ user });
      },

      logout: () => {
        set({
          user: null,
          accessToken: null,
          expiresAt: null,
          isLoggedIn: false,
        });
      },

      clearAuth: () => {
        set({
          user: null,
          accessToken: null,
          expiresAt: null,
          isLoggedIn: false,
        });
      },
    }),
    {
      name: "voidspace-auth",
    }
  )
);
