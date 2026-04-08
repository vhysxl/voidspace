"use client";

import { useEffect, useState, useCallback, useRef } from "react";
import DashboardLayout from "@/components/layout/DashboardLayout";
import PostCard from "@/components/posts/PostCard";
import { useAuthStore } from "@/store/useAuthStore";
import { useFeed } from "@/hooks/useFeed";
import { Post } from "@/types";
import { motion } from "framer-motion";
import { Loader2 } from "lucide-react";

export default function HomePage() {
  const { _hasHydrated, isLoggedIn } = useAuthStore();
  const { getGlobalFeed, getFollowingFeed } = useFeed();
  const [activeTab, setActiveTab] = useState<"global" | "following">("global");
  
  const [posts, setPosts] = useState<Post[]>([]);
  const [isLoading, setIsLoading] = useState(false);
  const [isFetchingMore, setIsFetchingMore] = useState(false);
  const [hasMore, setHasMore] = useState(true);
  const [error, setError] = useState<string | null>(null);
  
  // Ref to track if we've already done the initial fetch for current tab
  const initialFetchDone = useRef<string | null>(null);
  const observerTarget = useRef<HTMLDivElement>(null);

  const fetchFeed = useCallback(async (isMore = false) => {
    if (!_hasHydrated) return;
    if (isLoading || (isMore && !hasMore) || (isMore && isFetchingMore)) return;

    if (isMore) {
      setIsFetchingMore(true);
    } else {
      setIsLoading(true);
      setPosts([]); // Clear posts on tab switch
    }
    
    setError(null);

    try {
      const lastPost = isMore && posts.length > 0 ? posts[posts.length - 1] : null;
      const cursor = lastPost?.created_at;
      const cursorid = lastPost?.id;

      const response = activeTab === "following" && isLoggedIn
        ? await getFollowingFeed(cursor, cursorid) 
        : await getGlobalFeed(cursor, cursorid);
      
      if (response.success && response.data) {
        const newPosts = response.data.posts || [];
        setPosts(prev => isMore ? [...prev, ...newPosts] : newPosts);
        setHasMore(response.data.has_more);
      } else {
        setError(response.detail || "Failed to fetch feed");
      }
    } catch (err: any) {
      setError(err.message || "An error occurred while fetching the feed");
    } finally {
      setIsLoading(false);
      setIsFetchingMore(false);
    }
  }, [_hasHydrated, isLoggedIn, activeTab, getGlobalFeed, getFollowingFeed, posts, hasMore, isLoading, isFetchingMore]);

  // Initial fetch effect
  useEffect(() => {
    if (_hasHydrated && initialFetchDone.current !== activeTab) {
      initialFetchDone.current = activeTab;
      fetchFeed(false);
    }
  }, [_hasHydrated, activeTab, fetchFeed]);

  // Reset tracker on auth change
  useEffect(() => {
    initialFetchDone.current = null;
  }, [isLoggedIn]);

  // Intersection Observer for Infinite Scroll
  useEffect(() => {
    const observer = new IntersectionObserver(
      entries => {
        if (entries[0].isIntersecting && hasMore && !isLoading && !isFetchingMore && posts.length > 0) {
          fetchFeed(true);
        }
      },
      { threshold: 1.0 }
    );

    if (observerTarget.current) {
      observer.observe(observerTarget.current);
    }

    return () => observer.disconnect();
  }, [hasMore, isLoading, isFetchingMore, posts.length, fetchFeed]);

  if (!_hasHydrated) return null;

  return (
    <DashboardLayout fullWidth={true}>
      <div className="flex flex-col min-h-screen">
        {/* Header & Tabs */}
        <div className="sticky top-0 z-30 bg-background/80 backdrop-blur-md border-b border-foreground/10">
          <div className="px-6 py-4">
            <h1 className="font-space-grotesk text-xl font-bold tracking-tight uppercase">
              Feed
            </h1>
          </div>
          
          <div className="flex px-2">
            <button
              onClick={() => setActiveTab("global")}
              className={`flex-1 relative py-4 text-[11px] font-bold uppercase tracking-[2px] transition-colors ${
                activeTab === "global" ? "text-foreground" : "text-foreground/40 hover:text-foreground/60"
              }`}
            >
              Global
              {activeTab === "global" && (
                <motion.div
                  layoutId="activeFeedTab"
                  className="absolute bottom-0 left-0 right-0 h-[2px] bg-foreground"
                  transition={{ type: "spring", bounce: 0.2, duration: 0.6 }}
                />
              )}
            </button>
            
            {isLoggedIn && (
              <button
                onClick={() => setActiveTab("following")}
                className={`flex-1 relative py-4 text-[11px] font-bold uppercase tracking-[2px] transition-colors ${
                  activeTab === "following" ? "text-foreground" : "text-foreground/40 hover:text-foreground/60"
                }`}
              >
                Following
                {activeTab === "following" && (
                  <motion.div
                    layoutId="activeFeedTab"
                    className="absolute bottom-0 left-0 right-0 h-[2px] bg-foreground"
                    transition={{ type: "spring", bounce: 0.2, duration: 0.6 }}
                  />
                )}
              </button>
            )}
          </div>
        </div>

        {/* Feed Content */}
        <div className="flex-1 divide-y divide-foreground/10">
          {isLoading ? (
            <div className="flex items-center justify-center py-20">
              <Loader2 className="size-8 text-foreground/20 animate-spin" />
            </div>
          ) : error ? (
            <div className="p-10 text-center space-y-4">
              <p className="text-red-500 uppercase text-xs tracking-widest font-bold">{error}</p>
              <button 
                onClick={() => fetchFeed(false)}
                className="px-6 py-2 border border-foreground/20 text-[10px] font-bold uppercase tracking-[2px] hover:bg-foreground hover:text-background transition-all"
              >
                Retry
              </button>
            </div>
          ) : posts.length > 0 ? (
            <>
              {posts.map((post) => (
                <PostCard key={post.id} post={post} />
              ))}
              
              {/* Infinite Scroll Trigger */}
              <div ref={observerTarget} className="py-10 flex items-center justify-center">
                {isFetchingMore && <Loader2 className="size-6 text-foreground/20 animate-spin" />}
                {!hasMore && (
                  <p className="text-foreground/20 text-[10px] uppercase tracking-[2px]">
                    End of the void reached.
                  </p>
                )}
              </div>
            </>
          ) : (
            <div className="p-20 text-center">
              <p className="text-foreground/40 uppercase text-xs tracking-widest">
                {activeTab === "global" 
                  ? "The void is silent. No transmissions found." 
                  : "You aren't following anyone in the void yet."}
              </p>
            </div>
          )}
        </div>

        {/* Space for bottom nav on mobile */}
        <div className="h-20" />
      </div>
    </DashboardLayout>
  );
}
