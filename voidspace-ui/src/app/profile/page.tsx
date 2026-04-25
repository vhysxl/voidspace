"use client";

import { useEffect, useState, useCallback } from "react";
import { useRouter } from "next/navigation";
import { useAuthStore } from "@/store/useAuthStore";
import DashboardLayout from "@/components/layout/DashboardLayout";
import ProfileHeader from "@/components/profile/ProfileHeader";
import PostCard from "@/components/posts/PostCard";
import CommentItem from "@/components/posts/CommentItem";
import { Post, Comment } from "@/types";
import { usePosts } from "@/hooks/usePosts";
import { useComments } from "@/hooks/useComments";
import { Loader2 } from "lucide-react";

type TabType = "Posts" | "Comments" | "Likes";

export default function ProfilePage() {
  const { user, isLoggedIn, _hasHydrated } = useAuthStore();
  const router = useRouter();
  const { getUserPosts, getLikedPosts } = usePosts();
  const { getUserComments } = useComments();

  const [activeTab, setActiveTab] = useState<TabType>("Posts");
  const [contentLoading, setContentLoading] = useState(false);
  const [posts, setPosts] = useState<Post[]>([]);
  const [comments, setComments] = useState<Comment[]>([]);
  const [fetchedTabs, setFetchedTabs] = useState<Set<TabType>>(new Set());

  useEffect(() => {
    if (_hasHydrated && !isLoggedIn) {
      router.replace("/auth/login");
    }
  }, [isLoggedIn, _hasHydrated, router]);

  // Reset state when username changes
  useEffect(() => {
    setPosts([]);
    setComments([]);
    setFetchedTabs(new Set());
  }, [user?.username]);

  const fetchTabContent = useCallback(async (username: string, tab: TabType) => {
    if (fetchedTabs.has(tab)) return;
    
    setContentLoading(true);
    try {
      let success = false;
      if (tab === "Posts") {
        const res = await getUserPosts(username);
        if (res.success && res.data) {
          setPosts(res.data);
          success = true;
        }
      } else if (tab === "Comments") {
        const res = await getUserComments(username);
        if (res.success && res.data) {
          setComments(res.data);
          success = true;
        }
      } else if (tab === "Likes") {
        const res = await getLikedPosts(username);
        if (res.success && res.data) {
          setPosts(res.data);
          success = true;
        }
      }

      if (success) {
        setFetchedTabs(prev => new Set(prev).add(tab));
      }
    } catch (err) {
      console.error(`Failed to fetch ${tab}:`, err);
    } finally {
      setContentLoading(false);
    }
  }, [getUserPosts, getUserComments, getLikedPosts, fetchedTabs]);

  useEffect(() => {
    if (isLoggedIn && user?.username && !fetchedTabs.has(activeTab)) {
      fetchTabContent(user.username, activeTab);
    }
  }, [isLoggedIn, user?.username, activeTab, fetchTabContent, fetchedTabs]);

  if (!_hasHydrated) {
    return (
      <DashboardLayout fullWidth={true}>
        <div className="flex items-center justify-center min-h-[60vh]">
          <Loader2 className="size-8 text-foreground/20 animate-spin" />
        </div>
      </DashboardLayout>
    );
  }

  if (!isLoggedIn || !user) {
    return null;
  }

  return (
    <DashboardLayout fullWidth={true}>
      <div className="w-full">
        <ProfileHeader 
          user={user} 
          isOwnProfile={true} 
          activeTab={activeTab}
          onTabChange={(tab) => setActiveTab(tab as TabType)}
        />

        {/* Tab Content */}
        <div className="flex-1 divide-y divide-foreground/10 border-t border-foreground/5 min-h-[40vh]">
          {contentLoading ? (
            <div className="flex items-center justify-center py-20">
              <Loader2 className="size-6 text-foreground/20 animate-spin" />
            </div>
          ) : (
            <>
              {(activeTab === "Posts" || activeTab === "Likes") && (
                <>
                  {posts.map((post) => (
                    <PostCard key={post.id} post={post} />
                  ))}
                  {posts.length === 0 && (
                    <div className="p-20 text-center">
                      <p className="text-foreground/40 uppercase text-xs tracking-widest">
                        {activeTab === "Posts" ? "No transmissions from this sector yet." : "No resonated matter found."}
                      </p>
                    </div>
                  )}
                </>
              )}

              {activeTab === "Comments" && (
                <>
                  {comments.map((comment) => (
                    <CommentItem key={comment.comment_id} comment={comment} />
                  ))}
                  {comments.length === 0 && (
                    <div className="p-20 text-center">
                      <p className="text-foreground/40 uppercase text-xs tracking-widest">The void is silent here.</p>
                    </div>
                  )}
                </>
              )}
            </>
          )}
        </div>

        {/* Empty space for scrolling */}
        <div className="h-40" />
      </div>
    </DashboardLayout>
  );
}
