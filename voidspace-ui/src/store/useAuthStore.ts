import { create } from "zustand";
import { persist } from "zustand/middleware";
import { User } from "@/types";

interface AuthState {
  user: User | null;
  accessToken: string | null;
  expiresAt: number | null;
  isLoggedIn: boolean;
  _hasHydrated: boolean;

  setToken: (token: string, expires_in: number) => void;
  setUser: (user: User) => void;
  updateUser: (user: User) => void;
  logout: () => void;
  clearAuth: () => void;
  setHasHydrated: (state: boolean) => void;
}

export const useAuthStore = create<AuthState>()(
  persist(
    (set) => ({
      user: null,
      accessToken: null,
      expiresAt: null,
      isLoggedIn: false,
      _hasHydrated: false,

      setToken: (token: string, expires_in: number) => {
        set({
          accessToken: token,
          expiresAt: Date.now() + expires_in * 1000,
        });
      },

      setUser: (user: User) => {
        set({
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

      setHasHydrated: (state) => {
        set({ _hasHydrated: state });
      },
    }),
    {
      name: "voidspace-auth",
      onRehydrateStorage: (state) => {
        return () => state.setHasHydrated(true);
      },
    }
  )
);
