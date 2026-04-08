"use client";

import { useEffect } from "react";
import { useRouter } from "next/navigation";
import { useAuthStore } from "@/store/useAuthStore";
import DashboardLayout from "@/components/layout/DashboardLayout";
import ProfileHeader from "@/components/profile/ProfileHeader";
import PostCard from "@/components/posts/PostCard";
import { Post } from "@/types";

export default function ProfilePage() {
  const { user, isLoggedIn, _hasHydrated } = useAuthStore();
  const router = useRouter();

  useEffect(() => {
    if (_hasHydrated && !isLoggedIn) {
      router.replace("/auth/login");
    }
  }, [isLoggedIn, _hasHydrated, router]);

  if (!_hasHydrated) {
    return (
      <DashboardLayout fullWidth={true}>
        <div className="flex items-center justify-center min-h-[60vh]">
          <div className="size-8 border-2 border-foreground/20 border-t-foreground rounded-full animate-spin" />
        </div>
      </DashboardLayout>
    );
  }

  if (!isLoggedIn || !user) {
    return null;
  }

  const userPosts: Post[] = [];

  return (
    <DashboardLayout fullWidth={true}>
      <div className="w-full">
        <ProfileHeader user={user} isOwnProfile={true} />

        {/* User Posts Feed */}
        <div className="flex-1 divide-y divide-foreground/10 border-t border-foreground/5">
          {userPosts.map((post) => (
            <PostCard key={post.id} post={post} />
          ))}
          {userPosts.length === 0 && (
            <div className="p-20 text-center">
              <p className="text-foreground/40 uppercase text-xs tracking-widest">No transmissions from this sector yet.</p>
            </div>
          )}
        </div>

        {/* Empty space for scrolling */}
        <div className="h-40" />
      </div>
    </DashboardLayout>
  );
}
