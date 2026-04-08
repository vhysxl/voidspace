import { apiFetch } from "@/lib/api";
import { ApiResponse, FeedResponse } from "@/types";

export const useFeed = () => {
  const getGlobalFeed = async (cursor?: string, cursorid?: number) => {
    const params: Record<string, string> = {};
    if (cursor) params.cursor = cursor;
    if (cursorid) params.cursorid = cursorid.toString();
    
    return await apiFetch<ApiResponse<FeedResponse>>("/feed", { params });
  };

  const getFollowingFeed = async (cursor?: string, cursorid?: number) => {
    const params: Record<string, string> = {};
    if (cursor) params.cursor = cursor;
    if (cursorid) params.cursorid = cursorid.toString();

    return await apiFetch<ApiResponse<FeedResponse>>("/feed/following", { params });
  };

  return {
    getGlobalFeed,
    getFollowingFeed,
  };
};
