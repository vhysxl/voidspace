"use client";

import { useState } from "react";
import { useUIStore } from "@/store/useUIStore";
import { useAuthStore } from "@/store/useAuthStore";
import { X, Image as ImageIcon, Sparkles } from "lucide-react";
import { motion, AnimatePresence } from "framer-motion";

export default function NewPostModal() {
  const { isNewPostModalOpen, closeNewPostModal } = useUIStore();
  const { user } = useAuthStore();
  const [content, setContent] = useState("");

  const handlePost = () => {
    if (content.trim()) {
      alert("Post created! (Simulated)");
      setContent("");
      closeNewPostModal();
    }
  };

  return (
    <AnimatePresence>
      {isNewPostModalOpen && (
        <div className="fixed inset-0 z-[100] flex items-center justify-center p-4 bg-background/80 backdrop-blur-sm">
          <motion.div
            initial={{ opacity: 0, scale: 0.95, y: 20 }}
            animate={{ opacity: 1, scale: 1, y: 0 }}
            exit={{ opacity: 0, scale: 0.95, y: 20 }}
            className="w-full max-w-xl bg-background border border-foreground/20 rounded-sm shadow-2xl flex flex-col"
          >
            {/* Header */}
            <div className="flex justify-between items-center p-4 border-b border-foreground/10">
              <h3 className="font-space-grotesk text-sm font-bold uppercase tracking-[1px] flex items-center gap-2">
                <Sparkles size={16} />
                New Broadcast
              </h3>
              <button
                onClick={closeNewPostModal}
                className="text-foreground/40 hover:text-foreground transition-colors"
              >
                <X size={20} />
              </button>
            </div>

            {/* Body */}
            <div className="p-6 flex gap-4">
              <div className="size-12 rounded-full bg-foreground/5 flex-shrink-0 flex items-center justify-center text-xs font-bold text-foreground/20 border border-foreground/10">
                {user?.username?.slice(0, 2).toUpperCase() || "VO"}
              </div>
              <div className="flex-1 space-y-4">
                <textarea
                  autoFocus
                  value={content}
                  onChange={(e) => setContent(e.target.value)}
                  placeholder="What's resonating in the void?"
                  className="w-full min-h-[150px] bg-transparent resize-none text-lg outline-none placeholder:text-foreground/20 leading-relaxed"
                />
                
                {/* Media Attachment (Placeholder) */}
                <div className="flex items-center gap-4 pt-4 border-t border-foreground/5">
                  <button className="text-foreground/40 hover:text-foreground transition-colors p-2 hover:bg-foreground/5 rounded-full">
                    <ImageIcon size={20} />
                  </button>
                </div>
              </div>
            </div>

            {/* Footer */}
            <div className="p-4 border-t border-foreground/10 flex justify-between items-center">
              <div className="text-[10px] text-foreground/30 uppercase tracking-[1px]">
                {content.length} / 280 characters
              </div>
              <button
                disabled={!content.trim()}
                onClick={handlePost}
                className={`px-8 py-3 rounded-sm font-bold text-[11px] uppercase tracking-[2px] transition-all ${
                  content.trim()
                    ? "bg-foreground text-background hover:opacity-90 active:scale-[0.98]"
                    : "bg-foreground/5 text-foreground/20 cursor-not-allowed"
                }`}
              >
                Launch Post
              </button>
            </div>
          </motion.div>
        </div>
      )}
    </AnimatePresence>
  );
}
