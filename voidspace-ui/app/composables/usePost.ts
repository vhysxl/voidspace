import type { ApiResponse, CreatePostReq, Post } from "@/types";
import { handleApiError } from "@/utils/apiErrorHandler";
import { defaultOptions, defaultAuthOptions } from "@/utils/apiDefaults";

/**
 * Post service hooks
 * Provides CRUD operations for posts (create, read, update, delete).
 */

type PostApiResponse = ApiResponse<Post>;
type PostsApiResponse = ApiResponse<Post[]>;

export const usePost = () => {
  const { fetchWithAuth } = useApi();
  const apiUrl = "/api/post";

  const createPost = async (req: CreatePostReq): Promise<PostApiResponse> => {
    try {
      return await fetchWithAuth<PostApiResponse>(`${apiUrl}/post`, {
        ...defaultAuthOptions,
        method: "POST",
        body: req,
      });
    } catch (error: unknown) {
      return handleApiError(error, "Failed to create post");
    }
  };

  const getPost = async (postId: string): Promise<PostApiResponse> => {
    try {
      return await $fetch<PostApiResponse>(`${apiUrl}/${postId}`, {
        ...defaultOptions,
        method: "GET",
      });
    } catch (error: unknown) {
      return handleApiError(error, "Failed to get post");
    }
  };

  const getUserPosts = async (username: string): Promise<PostsApiResponse> => {
    try {
      return await $fetch<PostsApiResponse>(`${apiUrl}/user/${username}`, {
        ...defaultOptions,
        method: "GET",
      });
    } catch (error: unknown) {
      return handleApiError(error, "Failed to get user posts");
    }
  };

  const updatePost = async (
    req: CreatePostReq,
    postId: string
  ): Promise<ApiResponse> => {
    try {
      return await fetchWithAuth<ApiResponse>(`${apiUrl}/${postId}`, {
        ...defaultAuthOptions,
        method: "PUT",
        body: req,
      });
    } catch (error: unknown) {
      return handleApiError(error, "Failed to update post");
    }
  };

  const deletePost = async (postId: string): Promise<ApiResponse> => {
    try {
      return await fetchWithAuth<ApiResponse>(`${apiUrl}/${postId}`, {
        ...defaultAuthOptions,
        method: "DELETE",
      });
    } catch (error: unknown) {
      return handleApiError(error, "Failed to delete post");
    }
  };

  return { createPost, getPost, getUserPosts, updatePost, deletePost };
};
