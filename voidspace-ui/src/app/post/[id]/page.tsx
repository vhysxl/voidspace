"use client";

import React, { useEffect, useState, useCallback, useRef } from "react";
import { useParams, useRouter } from "next/navigation";
import DashboardLayout from "@/components/layout/DashboardLayout";
import PostCard from "@/components/posts/PostCard";
import CommentItem from "@/components/posts/CommentItem";
import CommentInput from "@/components/posts/CommentInput";
import { ChevronLeft, Loader2 } from "lucide-react";
import { Post, Comment } from "@/types";
import { usePosts } from "@/hooks/usePosts";
import { useComments } from "@/hooks/useComments";
import { useAuthStore } from "@/store/useAuthStore";
import { usePostStore } from "@/store/usePostStore";

export default function PostDetailPage() {
  const params = useParams();
  const router = useRouter();
  const { getPost } = usePosts();
  const { getPostComments, createComment } = useComments();
  const { _hasHydrated } = useAuthStore();
  const { activePost, setActivePost } = usePostStore();

  const postId = params.id as string;
  const [post, setPost] = useState<Post | null>(activePost);
  const [comments, setComments] = useState<Comment[]>([]);
  const [isLoading, setIsLoading] = useState(!activePost);
  const [error, setError] = useState<string | null>(null);

  // Ref to ensure we only fetch once per postId
  const initialFetchDone = useRef<string | null>(null);

  // Ref for the comment input to scroll to it
  const commentInputRef = useRef<HTMLDivElement>(null);

  const fetchPostData = useCallback(async () => {
    if (!postId || !_hasHydrated) return;

    if (!post) setIsLoading(true);
    setError(null);
    try {
      const [postRes, commentsRes] = await Promise.all([
        getPost(postId),
        getPostComments(postId)
      ]);

      if (postRes.success && postRes.data) {
        setPost(postRes.data);
      } else if (!post) {
        setError(postRes.detail || "Post not found");
      }

      if (commentsRes.success && commentsRes.data) {
        setComments(commentsRes.data);
      }
    } catch (err: any) {
      if (!post) setError(err.message || "Failed to load transmission");
    } finally {
      setIsLoading(false);
    }
  }, [postId, getPost, getPostComments, _hasHydrated, post]);


  const handleCommentCreate = async (content: string) => {
    if (!postId || !content.trim()) return;

    try {
      const response = await createComment(Number(postId), content);
      if (response.success) {
        if (response.data) {
          handleCommentAdded(response.data);
        } else {
          // If backend doesn't return the comment object, refresh the list
          fetchPostData();
        }
      } else {
        throw new Error(response.detail || "Failed to post comment");
      }
    } catch (err: any) {
      throw err;
    }
  };

  const handleCommentAdded = (newComment: Comment) => {
    setComments(prev => [newComment, ...prev]);
    if (post) {
      setPost({
        ...post,
        comments_count: (post.comments_count || 0) + 1
      });
    }
  };

  const scrollToCommentInput = () => {
    commentInputRef.current?.scrollIntoView({ behavior: "smooth" });
    // Focus the textarea inside CommentInput if possible
    const textarea = commentInputRef.current?.querySelector("textarea");
    if (textarea) {
      textarea.focus();
    }
  };


  useEffect(() => {
    if (_hasHydrated && postId && initialFetchDone.current !== postId) {
      initialFetchDone.current = postId;
      fetchPostData();
    }
  }, [_hasHydrated, postId, fetchPostData]);

  // Cleanup on unmount
  useEffect(() => {
    return () => {
      setActivePost(null);
    };
  }, [setActivePost]);

  if (!_hasHydrated) return null;

  return (
    <DashboardLayout fullWidth={true}>
      <div className="flex flex-col min-h-screen">
        {/* Header */}
        <div className="sticky top-0 z-30 bg-background/80 backdrop-blur-md border-b border-foreground/10 px-4 py-4 flex items-center gap-4">
          <button
            onClick={() => router.back()}
            className="p-2 hover:bg-foreground/5 rounded-full transition-colors text-foreground/60 hover:text-foreground"
          >
            <ChevronLeft size={20} />
          </button>
          <h1 className="font-space-grotesk text-xl font-bold tracking-tight uppercase">
            Post
          </h1>
        </div>

        <div className="flex-1">
          {isLoading ? (
            <div className="flex items-center justify-center py-20">
              <Loader2 className="size-8 text-foreground/20 animate-spin" />
            </div>
          ) : error ? (
            <div className="p-20 text-center space-y-4">
              <p className="text-red-500 uppercase text-xs tracking-widest font-bold">{error}</p>
              <button
                onClick={() => {
                  initialFetchDone.current = null;
                  fetchPostData();
                }}
                className="px-8 py-3 bg-foreground text-background font-bold uppercase text-[11px] tracking-[2px]"
              >
                Retry
              </button>
            </div>
          ) : post ? (
            <>
              {/* The Post */}
              <PostCard 
                post={post} 
                isDetailed={true} 
                onCommentClick={scrollToCommentInput}
              />

              {/* Comment Input */}
              <div ref={commentInputRef}>
                <CommentInput 
                  postId={post.id} 
                  onSubmit={handleCommentCreate}
                />
              </div>

              {/* Comments Section */}
              <div className="border-b border-foreground/10 bg-foreground/[0.01] px-6 py-4">
                <h2 className="text-[11px] font-bold uppercase tracking-[2px] text-foreground/40">
                  Comments ({post.comments_count})
                </h2>
              </div>

              <div className="divide-y divide-foreground/10">
                {comments.length > 0 ? (
                  comments.map((comment) => (
                    <CommentItem key={comment.comment_id} comment={comment} />
                  ))
                ) : (
                  <div className="p-10 text-center">
                    <p className="text-foreground/20 text-xs uppercase tracking-widest">No responses in this frequency.</p>
                  </div>
                )}
              </div>
            </>
          ) : null}

          {/* Empty space for scrolling */}
          <div className="h-40" />
        </div>
      </div>
    </DashboardLayout>
  );
}
