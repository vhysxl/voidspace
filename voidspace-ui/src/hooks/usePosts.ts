import { useCallback } from "react";
import { useMemo } from "react";
import { apiFetch } from "@/lib/api";
import { ApiResponse, Post, CreatePostRequest, UpdatePostRequest, PostImage } from "@/types";

export const usePosts = () => {
  return useMemo(() => ({
    createPost: async (content: string, post_images?: PostImage[]) => {
      const data: CreatePostRequest = { content, post_images };
      return await apiFetch<ApiResponse<Post>>("/posts", {
        method: "POST",
        body: JSON.stringify(data),
      });
    },

    getPost: async (id: number | string) => {
      return await apiFetch<ApiResponse<Post>>(`/posts/${id}`);
    },

    getUserPosts: async (username: string) => {
      return await apiFetch<ApiResponse<Post[]>>(`/posts/user/${username}`);
    },

    getLikedPosts: async (username: string) => {
      return await apiFetch<ApiResponse<Post[]>>(`/posts/liked/${username}`);
    },

    updatePost: async (id: number | string, content: string, post_images?: PostImage[]) => {
      const data: UpdatePostRequest = { content, post_images };
      return await apiFetch<ApiResponse<Post>>(`/posts/${id}`, {
        method: "PUT",
        body: JSON.stringify(data),
      });
    },

    deletePost: async (id: number | string) => {
      return await apiFetch<ApiResponse<null>>(`/posts/${id}`, {
        method: "DELETE",
      });
    },

    likePost: async (id: number | string) => {
      return await apiFetch<ApiResponse<null>>(`/posts/${id}/like`, {
        method: "POST",
      });
    },

    unlikePost: async (id: number | string) => {
      return await apiFetch<ApiResponse<null>>(`/posts/${id}/like`, {
        method: "DELETE",
      });
    },
  }), []);
};
