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
  const auth = useAuthStore();
  const { fetchWithAuth } = useApi();

  const getGlobalFeed = async (): Promise<FeedResponse> => {
    try {
      let response;

      if (auth.isLoggedIn) {
        response = await fetchWithAuth(`/api/feed/global`, {
          method: "GET",
          headers: { "Content-Type": "application/json" },
        });
      } else {
        response = await $fetch<FeedResponse>(`/api/feed/`, {
          method: "GET",
          headers: { "Content-Type": "application/json" },
        });
      }

      return response as FeedResponse;
    } catch (error: any) {
      throw new Error(error.message || "Failed to get feed");
    }
  };

  //follow feed

  return { getGlobalFeed };
};
