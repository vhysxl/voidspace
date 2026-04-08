"use client";

import { create } from "zustand";
import { Post } from "@/types";

interface UIState {
  isNewPostModalOpen: boolean;
  openNewPostModal: () => void;
  closeNewPostModal: () => void;
  
  isEditProfileModalOpen: boolean;
  openEditProfileModal: () => void;
  closeEditProfileModal: () => void;

  editingPost: Post | null;
  openEditPostModal: (post: Post) => void;
  closeEditPostModal: () => void;
}

export const useUIStore = create<UIState>((set) => ({
  isNewPostModalOpen: false,
  openNewPostModal: () => set({ isNewPostModalOpen: true }),
  closeNewPostModal: () => set({ isNewPostModalOpen: false }),

  isEditProfileModalOpen: false,
  openEditProfileModal: () => set({ isEditProfileModalOpen: true }),
  closeEditProfileModal: () => set({ isEditProfileModalOpen: false }),

  editingPost: null,
  openEditPostModal: (post) => set({ editingPost: post }),
  closeEditPostModal: () => set({ editingPost: null }),
}));
