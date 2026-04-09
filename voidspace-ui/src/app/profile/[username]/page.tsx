"use client";

import { useEffect, useState, useCallback } from "react";
import { useParams, useRouter } from "next/navigation";
import { useAuthStore } from "@/store/useAuthStore";
import DashboardLayout from "@/components/layout/DashboardLayout";
import ProfileHeader from "@/components/profile/ProfileHeader";
import PostCard from "@/components/posts/PostCard";
import CommentItem from "@/components/posts/CommentItem";
import { User, Post, Comment } from "@/types";
import { useUser } from "@/hooks/useUser";
import { usePosts } from "@/hooks/usePosts";
import { useComments } from "@/hooks/useComments";
import { Loader2 } from "lucide-react";

type TabType = "Posts" | "Comments" | "Likes";

export default function UserProfilePage() {
  const params = useParams();
  const router = useRouter();
  const { getUser } = useUser();
  const { getUserPosts, getLikedPosts } = usePosts();
  const { getUserComments } = useComments();
  const { user: currentUser, _hasHydrated } = useAuthStore();
  
  const [user, setUser] = useState<User | null>(null);
  const [isUserLoading, setIsUserLoading] = useState(true);
  
  const [activeTab, setActiveTab] = useState<TabType>("Posts");
  const [contentLoading, setContentLoading] = useState(false);
  const [posts, setPosts] = useState<Post[]>([]);
  const [comments, setComments] = useState<Comment[]>([]);
  const [fetchedTabs, setFetchedTabs] = useState<Set<TabType>>(new Set());

  const username = params.username as string;

  // Reset state when username changes (navigation between profiles)
  useEffect(() => {
    setUser(null);
    setPosts([]);
    setComments([]);
    setFetchedTabs(new Set());
    setIsUserLoading(true);
  }, [username]);

  const fetchUserData = useCallback(async () => {
    if (!_hasHydrated || user) return;

    if (currentUser && currentUser.username === username) {
      setUser(currentUser);
      setIsUserLoading(false);
      return;
    }

    try {
      setIsUserLoading(true); 
      const response = await getUser(username);
      if (response.success && response.data) {
        setUser(response.data);
      } else {
        setUser(null); 
      }
    } catch (err) {
      console.error("Failed to fetch user:", err);
      setUser(null);
    } finally {
      setIsUserLoading(false);
    }
  }, [_hasHydrated, username, currentUser, getUser, user]);

  const fetchTabContent = useCallback(async (tab: TabType) => {
    if (!username || fetchedTabs.has(tab)) return;
    
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
  }, [username, fetchedTabs, getUserPosts, getUserComments, getLikedPosts]);

  // Initial User Fetch
  useEffect(() => {
    fetchUserData();
  }, [fetchUserData]);

  // Tab Content Fetch (only when tab changes and not already fetched)
  useEffect(() => {
    if (user && !fetchedTabs.has(activeTab)) {
      fetchTabContent(activeTab);
    }
  }, [user, activeTab, fetchTabContent, fetchedTabs]);

  if (!_hasHydrated || isUserLoading) {
    return (
      <DashboardLayout fullWidth={true}>
        <div className="flex items-center justify-center min-h-[60vh]">
          <Loader2 className="size-8 text-foreground/20 animate-spin" />
        </div>
      </DashboardLayout>
    );
  }

  if (!user) {
    return (
      <DashboardLayout fullWidth={true}>
        <div className="flex flex-col items-center justify-center min-h-[60vh] space-y-4">
          <h1 className="text-2xl font-bold uppercase tracking-tighter font-space-grotesk text-foreground">User Not Found</h1>
          <p className="text-foreground/40 uppercase text-xs tracking-widest">They might have drifted too far into the void.</p>
          <button 
            onClick={() => router.back()}
            className="px-8 py-3 bg-foreground text-background font-bold uppercase text-[11px] tracking-[2px]"
          >
            Go Back
          </button>
        </div>
      </DashboardLayout>
    );
  }

  const isOwnProfile = currentUser?.username === username;

  return (
    <DashboardLayout fullWidth={true}>
      <div className="w-full">
        <ProfileHeader 
          user={user} 
          isOwnProfile={isOwnProfile} 
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
                        {activeTab === "Posts" ? "No transmissions from this sector." : "No resonated matter found."}
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
