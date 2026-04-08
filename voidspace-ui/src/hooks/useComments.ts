import { apiFetch } from "@/lib/api";
import { ApiResponse, Comment, CreateCommentRequest } from "@/types";

export const useComments = () => {
  const createComment = async (post_id: number, content: string) => {
    const data: CreateCommentRequest = { post_id, content };
    return await apiFetch<ApiResponse<Comment>>("/comments", {
      method: "POST",
      body: JSON.stringify(data),
    });
  };

  const getPostComments = async (post_id: number | string) => {
    return await apiFetch<ApiResponse<Comment[]>>(`/comments/post/${post_id}`);
  };

  const getUserComments = async (username: string) => {
    return await apiFetch<ApiResponse<Comment[]>>(`/comments/user/${username}`);
  };

  const deleteComment = async (id: number | string) => {
    return await apiFetch<ApiResponse<null>>(`/comments/${id}`, {
      method: "DELETE",
    });
  };

  return {
    createComment,
    getPostComments,
    getUserComments,
    deleteComment,
  };
};
