import { useCallback } from "react";
import { apiFetch } from "@/lib/api";
import { ApiResponse, Post, User, Comment } from "@/types";

export type SearchType = "user" | "post" | "comment";

export interface SearchResults {
  users?: User[];
  posts?: Post[];
  comments?: Comment[];
}

export const useSearch = () => {
  const search = useCallback(async (query: string, type: SearchType) => {
    const params = {
      q: query,
      type: type,
    };

    return await apiFetch<ApiResponse<any>>("/search", { params });
  }, []);

  return { search };
};
