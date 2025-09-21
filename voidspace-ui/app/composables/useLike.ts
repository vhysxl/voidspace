import type { ApiResponse } from "@/types";
import { handleApiError } from "@/utils/apiErrorHandler";
import { defaultAuthOptions } from "@/utils/apiDefaults";

/**
 * like service hooks
 * Provides operations for like and unlike.
 */

export const useLike = () => {
  const { fetchWithAuth } = useApi();
  const apiUrl = "/api/like";

  const like = async (postId: string): Promise<ApiResponse> => {
    try {
      return await fetchWithAuth<ApiResponse>(`${apiUrl}/${postId}`, {
        ...defaultAuthOptions,
        method: "POST",
      });
    } catch (error: unknown) {
      handleApiError(error, "Failed to like post");
    }
  };

  const unlike = async (postId: string): Promise<ApiResponse> => {
    try {
      return await fetchWithAuth<ApiResponse>(`${apiUrl}/${postId}`, {
        ...defaultAuthOptions,
        method: "DELETE",
      });
    } catch (error: unknown) {
      handleApiError(error, "Failed to unlike post");
    }
  };

  return { like, unlike };
};
