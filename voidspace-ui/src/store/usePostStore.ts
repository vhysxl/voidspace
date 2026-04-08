"use client";

import { create } from "zustand";
import { Post } from "@/types";

interface PostState {
  activePost: Post | null;
  setActivePost: (post: Post | null) => void;
}

export const usePostStore = create<PostState>((set) => ({
  activePost: null,
  setActivePost: (post) => set({ activePost: post }),
}));
