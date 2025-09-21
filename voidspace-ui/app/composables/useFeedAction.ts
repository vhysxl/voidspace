import type { Post } from "@/types";

export const useFeedAction = (posts: Ref<Post[] | null>) => {
  const { getGlobalFeed, getFollowFeed } = useFeed();
  const route = useRoute();

  const feedState = reactive({
    initialLoading: true,
    nextCursor: null as string | null,
    nextCursorId: null as string | null,
    hasMorePosts: true,
    isEndOfFeed: false,
    loadingMore: false,
    scrollContainer: ref<HTMLElement>(),
  });

  const activeTab = computed(() =>
    route.query.tab === "following" ? "following" : "for-you"
  );

  const updateCursorFromResponse = (response: any) => {
    if (response.data.hasmore) {
      const lastPost = posts.value?.[posts.value.length - 1];
      if (lastPost) {
        feedState.nextCursor = lastPost.created_at;
        feedState.nextCursorId = lastPost.id.toString();
        feedState.hasMorePosts = true;
      }
    } else {
      feedState.nextCursor = null;
      feedState.nextCursorId = null;
      feedState.hasMorePosts = false;
    }
  };

  const loadInitialPosts = async () => {
    try {
      feedState.initialLoading = true;
      feedState.nextCursor = null;
      feedState.nextCursorId = null;

      const params: Record<string, string> = {};
      let response;

      if (activeTab.value === "following") {
        response = await getFollowFeed(params);
      } else {
        response = await getGlobalFeed(params);
      }

      posts.value = response.data?.posts || [];
      updateCursorFromResponse(response);
      feedState.isEndOfFeed =
        posts.value.length === 0 || !feedState.hasMorePosts;
    } catch (error) {
      console.error("Error loading initial posts:", error);
    } finally {
      feedState.initialLoading = false;
    }
  };

  const loadMorePosts = async () => {
    if (
      feedState.loadingMore ||
      !feedState.hasMorePosts ||
      !feedState.nextCursor ||
      !feedState.nextCursorId
    ) {
      return;
    }

    try {
      feedState.loadingMore = true;

      const params: Record<string, string> = {
        cursor: feedState.nextCursor,
        cursorid: feedState.nextCursorId,
      };

      const response = await getGlobalFeed(params);
      const newPosts = response.data?.posts || [];

      if (newPosts.length > 0) {
        posts.value = [...(posts.value || []), ...newPosts];
        updateCursorFromResponse(response);
      } else {
        feedState.hasMorePosts = false;
      }

      feedState.isEndOfFeed = !feedState.hasMorePosts;
    } catch (error) {
      console.error("Error loading more posts:", error);
    } finally {
      feedState.loadingMore = false;
    }
  };

  const refreshFeed = () => {
    posts.value = [];
    feedState.hasMorePosts = true;
    feedState.isEndOfFeed = false;
    feedState.nextCursor = null;
    feedState.nextCursorId = null;
    loadInitialPosts();
  };

  const setupInfiniteScroll = (container: Ref<HTMLElement | null>) => {
    useInfiniteScroll(
      container,
      () => {
        if (
          feedState.hasMorePosts &&
          !feedState.initialLoading &&
          !feedState.loadingMore
        ) {
          loadMorePosts();
        }
      },
      { distance: 10, direction: "bottom", interval: 100 }
    );
  };

  return {
    feedState,
    activeTab,
    loadInitialPosts,
    loadMorePosts,
    refreshFeed,
    setupInfiniteScroll,
  };
};
