export type Post = {
  id: number;
  content: string;
  user_id: number;
  post_images: string[] | null;
  likes_count: number;
  created_at: string;
  updated_at: string;
  author: User;
};

type FeedResponse = {
  success: boolean;
  detail: string;
  data: {
    posts: Post[];
    hasmore: boolean;
  };
};

export const useFeed = () => {
  const { fetchWithAuth } = useApi();
  const config = useRuntimeConfig();
  const apiUrl = config.public.apiUrl;

  const getVanillaFeed = async (): Promise<FeedResponse> => {
    try {
      const response = await $fetch<FeedResponse>(`${apiUrl}/feed/`, {
        method: "GET",
        headers: { "Content-Type": "application/json" },
      });

      return response;
    } catch (error: any) {
      throw new Error(error.data?.detail || "Failed to get feed");
    }
  };

  //authenticatedfeed
  //follow feed

  return { getVanillaFeed };
};
