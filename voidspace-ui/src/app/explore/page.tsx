"use client";

import { useState, useEffect, useCallback, Suspense } from "react";
import DashboardLayout from "@/components/layout/DashboardLayout";
import { motion, AnimatePresence } from "framer-motion";
import { Search, MessageSquare, LayoutGrid, Users, Loader2 } from "lucide-react";
import { useSearch, SearchType } from "@/hooks/useSearch";
import PostCard from "@/components/posts/PostCard";
import CommentItem from "@/components/posts/CommentItem";
import Link from "next/link";
import { useSearchParams } from "next/navigation";
import { Heading, Subtext } from "@/components/ui/Typography";

const tabs = [
  { id: "post", label: "Posts", icon: LayoutGrid },
  { id: "comment", label: "Comments", icon: MessageSquare },
  { id: "user", label: "Users", icon: Users },
];

function ExploreContent() {
  const searchParams = useSearchParams();
  const urlQuery = searchParams.get("q") || "";

  const [activeTab, setActiveTab] = useState<SearchType>("post");
  const [query, setQuery] = useState(urlQuery);
  const [results, setResults] = useState<any[]>([]);
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

  const { search } = useSearch();

  const handleSearch = useCallback(async (searchQuery: string, type: SearchType) => {
    const trimmedQuery = searchQuery.trim();
    if (!trimmedQuery) {
      setResults([]);
      return;
    }

    setIsLoading(true);
    setError(null);
    try {
      const res = await search(trimmedQuery, type);
      if (res.success) {
        const data = res.data;
        if (Array.isArray(data)) {
          setResults(data);
        } else if (data && typeof data === 'object') {
          const key = type === 'post' ? 'posts' : type === 'user' ? 'users' : 'comments';
          setResults(data[key] || []);
        } else {
          setResults([]);
        }
      } else {
        setError(res.detail || "Search failed");
      }
    } catch (err: any) {
      setError(err.message || "An error occurred");
    } finally {
      setIsLoading(false);
    }
  }, [search]);

  // Sync internal query with URL query
  useEffect(() => {
    if (urlQuery) {
      setQuery(urlQuery);
    }
  }, [urlQuery]);

  // Combined effect to trigger search when query OR tab changes
  // But only if we have a valid query
  useEffect(() => {
    const trimmedQuery = query.trim();
    if (trimmedQuery) {
      handleSearch(trimmedQuery, activeTab);
    }
  }, [activeTab, handleSearch]); // Note: query is NOT here to avoid debounce/loop issues
  // The 'query' state is manually handled by handleKeyDown/handleSearch

  // Separate effect to handle the initial load or URL change
  useEffect(() => {
    if (urlQuery.trim()) {
      handleSearch(urlQuery.trim(), activeTab);
    }
  }, [urlQuery, handleSearch]);

  const handleKeyDown = (e: React.KeyboardEvent) => {
    if (e.key === "Enter") {
      handleSearch(query, activeTab);
    }
  };

  return (
    <div className="flex flex-col min-h-screen">
      {/* Header */}
      <div className="sticky top-0 z-30 bg-background/80 backdrop-blur-md border-b border-foreground/10 px-6 py-4">
        <Heading as="h1" className="text-xl md:text-xl tracking-tight">
          Explore
        </Heading>
      </div>

      {/* Search Bar */}
      <div className="p-6">
        <div className="relative group">
          <button 
            onClick={() => handleSearch(query, activeTab)}
            className="absolute left-4 top-1/2 -translate-y-1/2 text-foreground/40 group-focus-within:text-foreground transition-colors cursor-pointer hover:text-foreground"
          >
            <Search size={18} />
          </button>
          <input
            type="text"
            value={query}
            onChange={(e) => setQuery(e.target.value)}
            onKeyDown={handleKeyDown}
            placeholder="Search voidspace (Press Enter)"
            className="w-full bg-foreground/5 border border-foreground/10 rounded-sm py-4 pl-12 pr-4 text-xs tracking-[2px] outline-none focus:border-foreground/30 transition-all placeholder:text-foreground/30"
          />
        </div>
      </div>

      {/* Tabs */}
      <div className="flex border-b border-foreground/10 px-2">
        {tabs.map((tab) => {
          const Icon = tab.icon;
          const isActive = activeTab === tab.id;
          return (
            <button
              key={tab.id}
              onClick={() => setActiveTab(tab.id as SearchType)}
              className="flex-1 relative py-4 flex items-center justify-center gap-2 group transition-colors"
            >
              <Icon
                size={16}
                className={isActive ? "text-foreground" : "text-foreground/40 group-hover:text-foreground/60"}
              />
              <span
                className={`text-[11px] font-bold uppercase tracking-[2px] ${
                  isActive ? "text-foreground" : "text-foreground/40 group-hover:text-foreground/60"
                }`}
              >
                {tab.label}
              </span>
              {isActive && (
                <motion.div
                  layoutId="activeTab"
                  className="absolute bottom-0 left-0 right-0 h-[2px] bg-foreground"
                  transition={{ type: "spring", bounce: 0.2, duration: 0.6 }}
                />
              )}
            </button>
          );
        })}
      </div>

        {/* Content Area */}
        <div className="flex-1 divide-y divide-foreground/10">
          {isLoading ? (
            <div className="flex flex-col items-center justify-center py-20">
              <Loader2 className="size-8 text-foreground/20 animate-spin" />
              <Subtext className="mt-4 text-[10px] tracking-[2px] text-foreground/20">
                Scanning Void...
              </Subtext>
            </div>
          ) : results.length > 0 ? (
            <div className="flex flex-col">
              {results.map((item, idx) => {
                if (activeTab === "post") {
                  return <PostCard key={`post-${item.id}-${idx}`} post={item} />;
                }
                if (activeTab === "comment") {
                  return <CommentItem key={`comment-${item.comment_id}-${idx}`} comment={item} />;
                }
                if (activeTab === "user") {
                  return (
                    <Link 
                      key={`user-${item.id}-${idx}`} 
                      href={item.username ? `/profile/${item.username}` : "#"}
                      className={`p-4 md:p-6 flex items-center gap-4 hover:bg-foreground/5 transition-colors border-b border-foreground/5 last:border-0 ${!item.username ? "pointer-events-none" : ""}`}
                    >
                      <div className="size-12 rounded-full bg-void-gray border border-foreground/5 overflow-hidden">
                        {item.avatar_url ? (
                          <img src={item.avatar_url} alt="" className="w-full h-full object-cover" />
                        ) : (
                          <div className="w-full h-full flex items-center justify-center text-xs font-bold text-foreground/20 uppercase">
                            {item.username?.slice(0, 2) || "??"}
                          </div>
                        )}
                      </div>
                      <div className="flex-1 min-w-0">
                        <p className="font-bold text-sm tracking-tight truncate">
                          {item.display_name || item.username || "Unknown User"}
                        </p>
                        <Subtext className="text-xs truncate tracking-wider">
                          {item.username ? `@${item.username}` : "anonymous"}
                        </Subtext>
                      </div>
                    </Link>
                  );
                }
                return null;
              })}
            </div>
          ) : (
          <AnimatePresence mode="wait">
            <motion.div
              key={activeTab}
              initial={{ opacity: 0, y: 10 }}
              animate={{ opacity: 1, y: 0 }}
              exit={{ opacity: 0, y: -10 }}
              transition={{ duration: 0.2 }}
              className="flex flex-col items-center justify-center py-20 text-center"
            >
              <div className="size-16 rounded-full bg-foreground/5 flex items-center justify-center mb-4">
                {(() => {
                  const ActiveIcon = tabs.find((t) => t.id === activeTab)?.icon;
                  return ActiveIcon ? <ActiveIcon size={24} className="text-foreground/20" /> : null;
                })()}
              </div>
              <Heading as="h2" className="text-lg md:text-lg tracking-tight mb-2">
                {query ? `No ${activeTab}s found for "${query}"` : `No ${activeTab}s yet`}
              </Heading>
              <Subtext className="max-w-xs mx-auto text-sm leading-loose">
                {query ? "Try adjusting your search query or filters." : "Start typing to explore the void."}
              </Subtext>
            </motion.div>
          </AnimatePresence>
        )}
      </div>
    </div>
  );
}

export default function ExplorePage() {
  return (
    <DashboardLayout fullWidth={true}>
      <Suspense fallback={
        <div className="flex items-center justify-center min-h-screen">
          <Loader2 className="size-8 text-foreground/20 animate-spin" />
        </div>
      }>
        <ExploreContent />
      </Suspense>
    </DashboardLayout>
  );
}
