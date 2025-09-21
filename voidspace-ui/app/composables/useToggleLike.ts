import type { Post } from "@/types";

export const useToggleLike = () => {
  const isLikeSubmitting = ref(false);
  const { like, unlike } = useLike();
  const toast = useToast();

  const toggleLike = async (post: Post, event?: Event) => {
    if (event) event.stopPropagation();
    if (isLikeSubmitting.value) return;

    isLikeSubmitting.value = true;

    try {
      if (post.is_liked) {
        const res = await unlike(post.id.toString());
        if (res.success) {
          post.is_liked = false;
          post.likes_count = Math.max(0, post.likes_count - 1);
          toast.add({
            title: "Unliked",
            description: `You unliked post`,
            color: "neutral",
          });
        } else {
          throw new Error(res.detail || "Failed to unlike post");
        }
      } else {
        const res = await like(post.id.toString());
        if (res.success) {
          post.is_liked = true;
          post.likes_count += 1;
          toast.add({
            title: "Liked",
            description: `You liked post`,
            color: "neutral",
          });
        } else {
          throw new Error(res.detail || "Failed to like post");
        }
      }

      return { success: true, post };
    } catch (error: any) {
      toast.add({
        title: "Action failed",
        description: error.message || "Failed to perform action",
        color: "error",
      });
      return { success: false, error: error.message };
    } finally {
      isLikeSubmitting.value = false;
    }
  };

  return {
    toggleLike,
    isSubmitting: readonly(isLikeSubmitting),
  };
};
