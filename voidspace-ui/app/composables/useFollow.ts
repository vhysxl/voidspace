import type { ApiResponse } from "@/types";
import { handleApiError } from "@/utils/apiErrorHandler";
import { defaultAuthOptions } from "@/utils/apiDefaults";

/**
 * Follow service hooks
 * Provides operations for follow and unfollow.
 */

export const useFollow = () => {
  const { fetchWithAuth } = useApi();
  const apiUrl = "/api/follow";

  const follow = async (username: string): Promise<ApiResponse> => {
    try {
      return await fetchWithAuth<ApiResponse>(`${apiUrl}/${username}`, {
        ...defaultAuthOptions,
        method: "POST",
      });
    } catch (error: unknown) {
      return handleApiError(error, "Failed to follow user");
    }
  };

  const unFollow = async (username: string): Promise<ApiResponse> => {
    try {
      return await fetchWithAuth(`${apiUrl}/${username}`, {
        ...defaultAuthOptions,
        method: "DELETE",
      });
    } catch (error: unknown) {
      return handleApiError(error, "Failed to Unfollow user");
    }
  };

  return { follow, unFollow };
};
