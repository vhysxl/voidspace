import type { Feed } from "@/types";
import type { ApiResponse } from "@/types/index";
import { handleApiError } from "@/utils/apiErrorHandler";
import { defaultOptions, defaultAuthOptions } from "@/utils/apiDefaults";

/**
 * Feed service hooks
 * Provides GET operations for feed with cursor-based pagination
 */

type FeedResponse = ApiResponse<Feed>;

export const useFeed = () => {
  const auth = useAuthStore();
  const { fetchWithAuth } = useApi();

  const getGlobalFeed = async (
    params?: Record<string, string | number>
  ): Promise<FeedResponse> => {
    try {
      const query = params ? `?${new URLSearchParams(params as any)}` : "";
      const url = `/api/feed/global${query}`;

      if (auth.isLoggedIn) {
        return await fetchWithAuth<FeedResponse>(url, {
          ...defaultAuthOptions,
          method: "GET",
        });
      } else {
        return await $fetch<FeedResponse>(url, {
          ...defaultOptions,
          method: "GET",
        });
      }
    } catch (error: unknown) {
      handleApiError(error, "Failed to get feed");
    }
  };

  const getFollowFeed = async (
    params?: Record<string, string | number>
  ): Promise<FeedResponse> => {
    try {
      const query = params ? `?${new URLSearchParams(params as any)}` : "";
      const url = `/api/feed/followfeed${query}`;

      return await fetchWithAuth<FeedResponse>(url, {
        method: "GET",
        headers: { "Content-Type": "application/json" },
      });
    } catch (error: unknown) {
      handleApiError(error, "Failed to get feed");
    }
  };

  return { getFollowFeed, getGlobalFeed };
};
