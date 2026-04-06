"use client";

import DashboardLayout from "@/components/layout/DashboardLayout";
import { Heading, Subtext } from "@/components/ui/Typography";
import { useAuthStore } from "@/store/useAuthStore";
import { motion } from "framer-motion";

export default function HomePage() {
  const { user, isLoggedIn, _hasHydrated } = useAuthStore();

  if (!_hasHydrated) return null;

  return (
    <DashboardLayout>
      <div className="p-4 md:p-8 space-y-8">
        <header className="flex flex-col gap-2">
          <Heading>Home Feed</Heading>
          <Subtext>
            {isLoggedIn ? `Welcome back, ${user?.username}` : "Join the celestial conversation"}
          </Subtext>
        </header>

        {/* Placeholder Feed Content */}
        <div className="space-y-6">
          {[1, 2, 3, 4, 5].map((i) => (
            <motion.div 
              key={i}
              initial={{ opacity: 0, y: 20 }}
              animate={{ opacity: 1, y: 0 }}
              transition={{ delay: i * 0.1 }}
              className="p-6 border border-foreground/10 bg-foreground/5 rounded-sm space-y-4 hover:border-foreground/20 transition-colors group"
            >
              <div className="flex items-center gap-3">
                <div className="size-10 rounded-full bg-foreground/10 border border-foreground/5" />
                <div className="flex flex-col">
                  <div className="h-4 w-24 bg-foreground/10 rounded-full mb-1" />
                  <div className="h-3 w-16 bg-foreground/5 rounded-full" />
                </div>
              </div>
              <div className="space-y-2">
                <div className="h-4 w-full bg-foreground/10 rounded-full" />
                <div className="h-4 w-3/4 bg-foreground/10 rounded-full" />
              </div>
              <div className="flex gap-4 pt-2">
                <div className="h-4 w-12 bg-foreground/5 rounded-full" />
                <div className="h-4 w-12 bg-foreground/5 rounded-full" />
                <div className="h-4 w-12 bg-foreground/5 rounded-full" />
              </div>
            </motion.div>
          ))}
        </div>
      </div>
    </DashboardLayout>
  );
}
