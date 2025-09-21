import type { Post } from "@/types";

export const usePostState = () => {
  const { getPost } = usePost();
  const toast = useToast();

  const state = reactive({
    post: null as Post | null,
    pending: false,
    error: null as string | null,
  });

  const fetchPost = async (id: string) => {
    state.pending = true;
    state.error = null;

    try {
      const res = await getPost(id);
      state.post = res.data ?? null;
    } catch (err: any) {
      state.error = err.message;
      toast?.add?.({
        title: "Failed to Load Post",
        description: err.message || "Unable to load post",
        color: "error",
      });
      if (err.statusCode === 404)
        throw createError({ statusCode: 404, statusMessage: "Post not found" });
    } finally {
      state.pending = false;
    }
  };

  return { state, fetchPost };
};
