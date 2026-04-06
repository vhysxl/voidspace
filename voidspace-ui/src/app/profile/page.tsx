"use client";

import { useEffect } from "react";
import { useRouter } from "next/navigation";
import { useAuthStore } from "@/store/useAuthStore";
import DashboardLayout from "@/components/layout/DashboardLayout";
import ProfileHeader from "@/components/profile/ProfileHeader";
import { motion } from "framer-motion";

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
      <DashboardLayout fullWidth>
        <div className="flex items-center justify-center min-h-[60vh]">
          <div className="size-8 border-2 border-white/20 border-t-white rounded-full animate-spin" />
        </div>
      </DashboardLayout>
    );
  }

  if (!isLoggedIn || !user) {
    return null;
  }

  return (
    <DashboardLayout fullWidth>
      <div className="w-full">
        <ProfileHeader user={user} isOwnProfile={true} />

        {/* Profile Content / Feed Placeholder */}
        <div className="max-w-4xl mx-auto p-4 md:p-8 space-y-6">
          {[1, 2, 3].map((i) => (
            <motion.div
              key={i}
              initial={{ opacity: 0, y: 20 }}
              animate={{ opacity: 1, y: 0 }}
              transition={{ delay: i * 0.1 }}
              className="p-6 border border-white/10 bg-white/5 rounded-sm space-y-4 hover:border-white/20 transition-colors group"
            >
              <div className="flex items-center gap-3">
                <div className="size-10 rounded-full bg-white/10 border border-white/5 overflow-hidden shrink-0">
                  {user.profile.avatar_url ? (
                    <img src={user.profile.avatar_url} alt="" className="w-full h-full object-cover" />
                  ) : (
                    <div className="w-full h-full flex items-center justify-center text-xs font-bold uppercase text-white/20">
                      {user.username.slice(0, 2)}
                    </div>
                  )}
                </div>
                <div className="flex flex-col">
                  <span className="text-white text-sm font-bold tracking-tight uppercase">
                    {user.profile.display_name || user.username}
                  </span>
                  <span className="text-[#666] text-xs uppercase tracking-wider">
                    @{user.username}
                  </span>
                </div>
              </div>
              <div className="space-y-2">
                <div className="h-4 w-full bg-white/10 rounded-full" />
                <div className="h-4 w-3/4 bg-white/10 rounded-full" />
              </div>
              <div className="flex gap-4 pt-2">
                <div className="h-4 w-12 bg-white/5 rounded-full" />
                <div className="h-4 w-12 bg-white/5 rounded-full" />
                <div className="h-4 w-12 bg-white/5 rounded-full" />
              </div>
            </motion.div>
          ))}
        </div>
      </div>
    </DashboardLayout>
  );
}
