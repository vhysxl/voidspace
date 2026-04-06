"use client";

import { create } from "zustand";

interface UIState {
  isNewPostModalOpen: boolean;
  openNewPostModal: () => void;
  closeNewPostModal: () => void;
}

export const useUIStore = create<UIState>((set) => ({
  isNewPostModalOpen: false,
  openNewPostModal: () => set({ isNewPostModalOpen: true }),
  closeNewPostModal: () => set({ isNewPostModalOpen: false }),
}));
