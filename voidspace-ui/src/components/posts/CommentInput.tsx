"use client";

import { useState } from "react";
import { useAuthStore } from "@/store/useAuthStore";
import { useComments } from "@/hooks/useComments";
import { Loader2 } from "lucide-react";
import Button from "@/components/ui/Button";
import { Comment } from "@/types";

interface CommentInputProps {
  postId: number;
  onCommentAdded?: (comment: Comment) => void;
  onSubmit?: (content: string) => Promise<void>;
}

export default function CommentInput({ postId, onCommentAdded, onSubmit }: CommentInputProps) {
  const { user, isLoggedIn } = useAuthStore();
  const { createComment } = useComments();
  const [content, setContent] = useState("");
  const [isSubmitting, setIsSubmitting] = useState(false);

  const handleSubmit = async () => {
    if (!content.trim() || isSubmitting) return;

    setIsSubmitting(true);
    try {
      if (onSubmit) {
        await onSubmit(content);
        setContent("");
      } else {
        const response = await createComment(postId, content);
        if (response.success) {
          setContent("");
          if (onCommentAdded && response.data) {
            onCommentAdded(response.data);
          }
        } else {
          alert(response.detail || "Failed to post comment");
        }
      }
    } catch (err: any) {
      console.error("Failed to add comment:", err);
      alert(err.message || "An error occurred");
    } finally {
      setIsSubmitting(false);
    }
  };

  if (!isLoggedIn) return null;

  return (
    <div className="p-4 md:p-6 border-b border-foreground/10 flex gap-4 bg-foreground/[0.02]">
      <div className="size-10 rounded-full bg-foreground/5 shrink-0 flex items-center justify-center text-xs font-bold text-foreground/20 border border-foreground/10 overflow-hidden">
        {user?.profile.avatar_url ? (
          <img src={user.profile.avatar_url} alt="" className="w-full h-full object-cover" />
        ) : (
          user?.username?.slice(0, 2).toUpperCase()
        )}
      </div>
      <div className="flex-1 space-y-3">
        <textarea
          value={content}
          onChange={(e) => setContent(e.target.value)}
          placeholder="Add a comment..."
          className="w-full bg-transparent resize-none text-sm md:text-base outline-none placeholder:text-foreground/20 leading-relaxed min-h-[80px]"
        />
        <div className="flex justify-end">
          <Button
            onClick={handleSubmit}
            disabled={!content.trim() || isSubmitting}
            isLoading={isSubmitting}
            className="w-auto px-8 h-10 text-[11px]"
          >
            Post Comment
          </Button>
        </div>
      </div>
    </div>
  );
}
