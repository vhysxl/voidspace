"use client";

import { useState } from "react";
import DashboardLayout from "@/components/layout/DashboardLayout";
import { motion, AnimatePresence } from "framer-motion";
import { Search, MessageSquare, LayoutGrid, Users } from "lucide-react";

const tabs = [
  { id: "posts", label: "Posts", icon: LayoutGrid },
  { id: "comments", label: "Comments", icon: MessageSquare },
  { id: "users", label: "Users", icon: Users },
];

export default function ExplorePage() {
  const [activeTab, setActiveTab] = useState("posts");

  return (
    <DashboardLayout fullWidth={true}>
      <div className="flex flex-col min-h-screen">
        {/* Header */}
        <div className="sticky top-0 z-30 bg-background/80 backdrop-blur-md border-b border-foreground/10 px-6 py-4">
          <h1 className="font-space-grotesk text-xl font-bold tracking-tight uppercase">
            Explore
          </h1>
        </div>

        {/* Search Bar (Mobile/Internal) */}
        <div className="p-6">
          <div className="relative group">
            <div className="absolute left-4 top-1/2 -translate-y-1/2 text-foreground/40 group-focus-within:text-foreground transition-colors">
              <Search size={18} />
            </div>
            <input
              type="text"
              placeholder="SEARCH VOIDSPACE"
              className="w-full bg-foreground/5 border border-foreground/10 rounded-sm py-4 pl-12 pr-4 text-xs tracking-[2px] uppercase outline-none focus:border-foreground/30 transition-all placeholder:text-foreground/30"
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
                onClick={() => setActiveTab(tab.id)}
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
        <div className="flex-1 p-6">
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
              <h2 className="font-space-grotesk text-lg font-bold uppercase tracking-tight mb-2">
                No {activeTab} found
              </h2>
              <p className="text-foreground/40 text-sm max-w-xs mx-auto">
                The void is currently empty. Try searching for something else or explore different realms.
              </p>
            </motion.div>
          </AnimatePresence>
        </div>
      </div>
    </DashboardLayout>
  );
}
