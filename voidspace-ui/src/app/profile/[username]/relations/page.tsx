"use client";

import { useState, useEffect } from "react";
import { useParams, useSearchParams, useRouter } from "next/navigation";
import DashboardLayout from "@/components/layout/DashboardLayout";
import { motion, AnimatePresence } from "framer-motion";
import { ChevronLeft, Users, UserPlus, Loader2 } from "lucide-react";
import Link from "next/link";
import { useAuthStore } from "@/store/useAuthStore";
import { useUser } from "@/hooks/useUser";
import { UserBanner as UserBannerType } from "@/types";
import UserBanner from "@/components/profile/UserBanner";

const tabs = [
  { id: "following", label: "Following", icon: UserPlus },
  { id: "followers", label: "Followers", icon: Users },
];

export default function RelationsPage() {
  const params = useParams();
  const searchParams = useSearchParams();
  const router = useRouter();
  const { user: currentUser, _hasHydrated } = useAuthStore();
  const { getFollowers, getFollowing } = useUser();
  
  const username = params.username as string;
  const initialTab = searchParams.get("tab") || "following";
  const [activeTab, setActiveTab] = useState(initialTab);
  const [users, setUsers] = useState<UserBannerType[]>([]);
  const [isLoading, setIsLoading] = useState(true);

  useEffect(() => {
    setActiveTab(initialTab);
  }, [initialTab]);

  useEffect(() => {
    const fetchData = async () => {
      if (!_hasHydrated) return;
      
      setIsLoading(true);
      try {
        const response = activeTab === "followers" 
          ? await getFollowers(username)
          : await getFollowing(username);
          
        if (response.success && response.data) {
          setUsers(response.data);
        } else {
          setUsers([]);
        }
      } catch (err) {
        console.error("Failed to fetch relations:", err);
        setUsers([]);
      } finally {
        setIsLoading(false);
      }
    };

    fetchData();
  }, [activeTab, username, _hasHydrated]);

  const handleTabChange = (tabId: string) => {
    setActiveTab(tabId);
    router.push(`/profile/${username}/relations?tab=${tabId}`);
  };

  if (!_hasHydrated) return null;

  return (
    <DashboardLayout fullWidth={true}>
      <div className="flex flex-col min-h-screen">
        {/* Header */}
        <div className="sticky top-0 z-30 bg-background/80 backdrop-blur-md border-b border-foreground/10 px-6 py-4 flex items-center gap-4">
          <Link 
            href={`/profile/${username}`}
            className="p-2 hover:bg-foreground/5 rounded-full transition-colors text-foreground/60 hover:text-foreground"
          >
            <ChevronLeft size={20} />
          </Link>
          <div>
            <h1 className="font-space-grotesk text-xl font-bold tracking-tight leading-none">
              {username}
            </h1>
            <p className="text-[10px] text-foreground/40 uppercase tracking-[1px] mt-1">
              Network Connections
            </p>
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
                onClick={() => handleTabChange(tab.id)}
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
                    layoutId="activeTabRelations"
                    className="absolute bottom-0 left-0 right-0 h-[2px] bg-foreground"
                    transition={{ type: "spring", bounce: 0.2, duration: 0.6 }}
                  />
                )}
              </button>
            );
          })}
        </div>

        {/* Content Area */}
        <div className="flex-1 p-6 max-w-4xl mx-auto w-full">
          {isLoading ? (
            <div className="flex items-center justify-center py-20">
              <Loader2 className="size-8 text-foreground/20 animate-spin" />
            </div>
          ) : (
            <AnimatePresence mode="wait">
              <motion.div
                key={activeTab}
                initial={{ opacity: 0, y: 10 }}
                animate={{ opacity: 1, y: 0 }}
                exit={{ opacity: 0, y: -10 }}
                transition={{ duration: 0.2 }}
                className="space-y-4"
              >
                {users.length > 0 ? (
                  users.map((user) => (
                    <UserBanner key={user.id} user={user} />
                  ))
                ) : (
                  <div className="flex flex-col items-center justify-center py-20 text-center">
                    <div className="size-16 rounded-full bg-foreground/5 flex items-center justify-center mb-4">
                      {(() => {
                        const ActiveIcon = tabs.find((t) => t.id === activeTab)?.icon;
                        return ActiveIcon ? <ActiveIcon size={24} className="text-foreground/20" /> : null;
                      })()}
                    </div>
                    <h2 className="font-space-grotesk text-lg font-bold uppercase tracking-tight mb-2">
                      No {activeTab} yet
                    </h2>
                    <p className="text-foreground/40 text-sm max-w-xs mx-auto">
                      {activeTab === "followers" 
                        ? "This user doesn't have any followers in the void yet." 
                        : "This user isn't following anyone in the void yet."}
                    </p>
                  </div>
                )}
              </motion.div>
            </AnimatePresence>
          )}
        </div>
      </div>
    </DashboardLayout>
  );
}
