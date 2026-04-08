"use client";

import { useEffect, useState } from "react";
import { useParams, useRouter } from "next/navigation";
import { useAuthStore } from "@/store/useAuthStore";
import DashboardLayout from "@/components/layout/DashboardLayout";
import ProfileHeader from "@/components/profile/ProfileHeader";
import PostCard from "@/components/posts/PostCard";
import { User, Post } from "@/types";

export default function UserProfilePage() {
  const params = useParams();
  const router = useRouter();
  const { user: currentUser, _hasHydrated } = useAuthStore();
  const [user, setUser] = useState<User | null>(null);
  const [isLoading, setIsLoading] = useState(true);
  const [userPosts, setUserPosts] = useState<Post[]>([]);

  const username = params.username as string;

  useEffect(() => {
    if (_hasHydrated) {
      if (currentUser && currentUser.username === username) {
        setUser(currentUser);
        setIsLoading(false);
      } else {
        // Here we would typically fetch the user from the API
        // For now, we set loading to false and let the empty state handle it
        setIsLoading(false);
      }
    }
  }, [_hasHydrated, username, currentUser]);

  if (!_hasHydrated || isLoading) {
    return (
      <DashboardLayout fullWidth={true}>
        <div className="flex items-center justify-center min-h-[60vh]">
          <div className="size-8 border-2 border-foreground/20 border-t-foreground rounded-full animate-spin" />
        </div>
      </DashboardLayout>
    );
  }

  if (!user) {
    return (
      <DashboardLayout fullWidth={true}>
        <div className="flex flex-col items-center justify-center min-h-[60vh] space-y-4">
          <h1 className="text-2xl font-bold uppercase tracking-tighter font-space-grotesk text-foreground">Voyager Not Found</h1>
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
        <ProfileHeader user={user} isOwnProfile={isOwnProfile} />

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
