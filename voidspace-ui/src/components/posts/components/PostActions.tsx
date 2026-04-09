"use client";

import { useAuthStore } from "@/store/useAuthStore";
import { usePosts } from "@/hooks/usePosts";
import { useRouter } from "next/navigation";
import { useState, useEffect } from "react";
import { MessageSquare, Heart, Share2 } from "lucide-react";

interface PostActionsProps {
  post: {
    id: number;
    is_liked: boolean;
    likes_count: number;
    comments_count: number;
  };
  onCommentClick?: () => void;
}

export default function PostActions({ post, onCommentClick }: PostActionsProps) {
  const { isLoggedIn } = useAuthStore();
  const { likePost, unlikePost } = usePosts();
  const router = useRouter();

  const [isLiked, setIsLiked] = useState(post.is_liked);
  const [likesCount, setLikesCount] = useState(post.likes_count);
  const [isLoading, setIsLoading] = useState(false);

  useEffect(() => {
    setIsLiked(post.is_liked);
    setLikesCount(post.likes_count);
  }, [post.is_liked, post.likes_count]);

  const handleLikeToggle = async (e: React.MouseEvent) => {
    e.stopPropagation();
    if (!isLoggedIn) {
      router.push("/auth/login");
      return;
    }
    if (isLoading) return;

    setIsLoading(true);
    const previousLiked = isLiked;
    const previousCount = likesCount;

    setIsLiked(!previousLiked);
    setLikesCount((prev) => (previousLiked ? prev - 1 : prev + 1));

    try {
      const res = await (previousLiked ? unlikePost(post.id) : likePost(post.id));
      if (!res.success) {
        setIsLiked(previousLiked);
        setLikesCount(previousCount);
      }
    } catch (err) {
      setIsLiked(previousLiked);
      setLikesCount(previousCount);
    } finally {
      setIsLoading(false);
    }
  };

  return (
    <div className="flex items-center gap-8 pt-4">
      <button
        onClick={(e) => {
          e.stopPropagation();
          onCommentClick?.();
        }}
        className="flex items-center gap-2 text-foreground/40 hover:text-blue-500 transition-colors group/action"
      >
        <div className="size-8 rounded-full flex items-center justify-center group-hover/action:bg-blue-500/10 transition-colors">
          <MessageSquare size={18} />
        </div>
        <span className="text-[11px] font-bold tracking-widest">{post.comments_count}</span>
      </button>

      <button
        className={`flex items-center gap-2 transition-colors group/action ${isLiked ? "text-pink-500" : "text-foreground/40 hover:text-pink-500"}`}
        onClick={handleLikeToggle}
      >
        <div className={`size-8 rounded-full flex items-center justify-center group-hover/action:bg-pink-500/10 transition-colors ${isLiked ? "bg-pink-500/5" : ""}`}>
          <Heart size={18} fill={isLiked ? "currentColor" : "none"} />
        </div>
        <span className="text-[11px] font-bold tracking-widest">{likesCount}</span>
      </button>

      <button
        className="flex items-center gap-2 text-foreground/40 hover:text-foreground transition-colors group/action ml-auto"
        onClick={(e) => {
          e.stopPropagation();
          // Share functionality could go here
        }}
      >
        <div className="size-8 rounded-full flex items-center justify-center group-hover/action:bg-foreground/5 transition-colors">
          <Share2 size={18} />
        </div>
      </button>
    </div>
  );
}
