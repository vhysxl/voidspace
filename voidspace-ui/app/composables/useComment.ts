import type { CommentType, CreateCommentReq } from "@/types";
import type { ApiResponse } from "@/types/index";
import { handleApiError } from "@/utils/apiErrorHandler";
import { defaultOptions, defaultAuthOptions } from "@/utils/apiDefaults";

/**
 * Comment service hooks
 * Provides CRUD operations for comments (create, fetch by post/user, delete).
 */

export type CommentApiResponse = ApiResponse<CommentType | CommentType[]>;

export const useComment = () => {
  const { fetchWithAuth } = useApi();
  const apiUrl = "/api/comment";

  const createComment = async (
    req: CreateCommentReq
  ): Promise<CommentApiResponse> => {
    try {
      const res = await fetchWithAuth<CommentApiResponse>(apiUrl, {
        ...defaultAuthOptions,
        method: "POST",
        body: req,
      });

      return res;
    } catch (error: unknown) {
      return handleApiError(error, "Failed to create comment");
    }
  };

  const getCommentsByPost = async (
    postId: string
  ): Promise<CommentApiResponse> => {
    try {
      return await $fetch<CommentApiResponse>(`${apiUrl}/post/${postId}`, {
        ...defaultOptions,
        method: "GET",
      });
    } catch (error: unknown) {
      return handleApiError(error, "Failed to get post comments");
    }
  };

  const getCommentsByUser = async (
    username: string
  ): Promise<CommentApiResponse> => {
    try {
      return await $fetch<CommentApiResponse>(`${apiUrl}/user/${username}`, {
        ...defaultOptions,
        method: "GET",
      });
    } catch (error: unknown) {
      return handleApiError(error, "Failed to get user comments");
    }
  };

  const deleteComment = async (
    commentId: string
  ): Promise<CommentApiResponse> => {
    try {
      return await fetchWithAuth<CommentApiResponse>(`${apiUrl}/${commentId}`, {
        ...defaultAuthOptions,
        method: "DELETE",
      });
    } catch (error: unknown) {
      return handleApiError(error, "Failed to delete comment");
    }
  };

  return {
    createComment,
    getCommentsByPost,
    getCommentsByUser,
    deleteComment,
  };
};
