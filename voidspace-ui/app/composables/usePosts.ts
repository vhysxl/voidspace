export interface CreatePostReq {
  content: string;
  post_images: string[] | null;
}

export interface ApiResponse {
  success: boolean;
  detail: string;
  data: Post | Post[];
}

export const usePosts = () => {
  const { fetchWithAuth } = useApi();
  const config = useRuntimeConfig();
  const apiUrl = config.public.apiUrl;

  const createPost = async (req: CreatePostReq) => {
    try {
      const res = await fetchWithAuth(`${apiUrl}/posts/`, {
        method: "POST",
        body: req,
        headers: { "Content-Type": "application/json" },
        credentials: "include",
      });

      return res as ApiResponse;
    } catch (error: any) {
      throw new Error(error.data?.detail || "Failed to create post");
    }
  };

  const getPost = async (postId: string): Promise<ApiResponse> => {
    try {
      const response = await $fetch<ApiResponse>(`${apiUrl}/posts/${postId}`, {
        method: "GET",
        headers: { "Content-Type": "application/json" },
      });

      return response;
    } catch (error: any) {
      throw new Error(error.data?.detail || "Failed to get post");
    }
  };

  const getUserPosts = async (username: string): Promise<ApiResponse> => {
    try {
      const response = await $fetch<ApiResponse>(
        `${apiUrl}/posts/user/${username}`,
        {
          method: "GET",
          headers: { "Content-Type": "application/json" },
        }
      );

      return response;
    } catch (error: any) {
      throw new Error(error.data?.detail || "Failed to get user posts");
    }
  };

  const updatePost = async (
    req: CreatePostReq,
    postId: string
  ): Promise<ApiResponse> => {
    try {
      const res = await fetchWithAuth(`${apiUrl}/posts/${postId}`, {
        method: "PUT",
        body: req,
        headers: { "Content-Type": "application/json" },
        credentials: "include",
      });

      return res as ApiResponse;
    } catch (error: any) {
      throw new Error(error.data?.detail || "Failed to update post");
    }
  };

  const deletePost = async (postId: string): Promise<ApiResponse> => {
    try {
      const res = await fetchWithAuth(`${apiUrl}/posts/${postId}`, {
        method: "DELETE",
        headers: { "Content-Type": "application/json" },
        credentials: "include",
      });

      return res as ApiResponse;
    } catch (error: any) {
      throw new Error(error.data?.detail || "Failed to delete post");
    }
  };

  return { createPost, getPost, getUserPosts, updatePost, deletePost };
};
