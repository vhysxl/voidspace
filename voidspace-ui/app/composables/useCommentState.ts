import type { CommentType, CreateCommentReq } from "@/types";

export const useCommentState = () => {
  const { getCommentsByPost, createComment } = useComment();
  const toast = useToast();

  const state = reactive({
    comments: [] as CommentType[],
    pending: false,
    error: null as string | null,
    newComment: "",
    isSubmittingComment: false,
  });

  const fetchComments = async (postId: string) => {
    state.pending = true;
    state.error = null;

    try {
      const res = await getCommentsByPost(postId);
      state.comments = Array.isArray(res.data) ? res.data : [];
    } catch (err: any) {
      state.comments = [];
      state.error = err.message;
      toast?.add?.({
        title: "Failed to Load Comments",
        description: err.message || "Unable to load comments",
        color: "error",
      });
    } finally {
      state.pending = false;
    }
  };

  const submitComment = async (req: CreateCommentReq) => {
    if (!req.content.trim()) return;
    state.isSubmittingComment = true;
    try {
      await createComment(req);
      state.newComment = "";
      await fetchComments(req.post_id.toString());
      toast?.add?.({
        title: "Comment Posted",
        description: "Your comment has been posted",
        color: "neutral",
      });
    } catch (err: any) {
      toast?.add?.({
        title: "Create Comment Failed",
        description: err.message || "Failed to create comment",
        color: "error",
      });
    } finally {
      state.isSubmittingComment = false;
    }
  };

  return { state, fetchComments, submitComment };
};
