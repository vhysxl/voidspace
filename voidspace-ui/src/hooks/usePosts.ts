import { apiFetch } from "@/lib/api";
import { ApiResponse, Post, CreatePostRequest, UpdatePostRequest, PostImage } from "@/types";

export const usePosts = () => {
  const createPost = async (content: string, post_images?: PostImage[]) => {
    const data: CreatePostRequest = { content, post_images };
    return await apiFetch<ApiResponse<Post>>("/posts", {
      method: "POST",
      body: JSON.stringify(data),
    });
  };

  const getPost = async (id: number | string) => {
    return await apiFetch<ApiResponse<Post>>(`/posts/${id}`);
  };

  const getUserPosts = async (username: string) => {
    return await apiFetch<ApiResponse<Post[]>>(`/posts/user/${username}`);
  };

  const getLikedPosts = async (username: string) => {
    return await apiFetch<ApiResponse<Post[]>>(`/posts/liked/${username}`);
  };

  const updatePost = async (id: number | string, content: string) => {
    const data: UpdatePostRequest = { content };
    return await apiFetch<ApiResponse<Post>>(`/posts/${id}`, {
      method: "PUT",
      body: JSON.stringify(data),
    });
  };

  const deletePost = async (id: number | string) => {
    return await apiFetch<ApiResponse<null>>(`/posts/${id}`, {
      method: "DELETE",
    });
  };

  const likePost = async (id: number | string) => {
    return await apiFetch<ApiResponse<null>>(`/posts/${id}/like`, {
      method: "POST",
    });
  };

  const unlikePost = async (id: number | string) => {
    return await apiFetch<ApiResponse<null>>(`/posts/${id}/like`, {
      method: "DELETE",
    });
  };

  return {
    createPost,
    getPost,
    getUserPosts,
    getLikedPosts,
    updatePost,
    deletePost,
    likePost,
    unlikePost,
  };
};
