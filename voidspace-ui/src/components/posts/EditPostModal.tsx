"use client";

import { useState, useEffect } from "react";
import { useUIStore } from "@/store/useUIStore";
import { usePosts } from "@/hooks/usePosts";
import { X, Sparkles, Loader2 } from "lucide-react";
import { motion, AnimatePresence } from "framer-motion";

export default function EditPostModal() {
  const { editingPost, closeEditPostModal } = useUIStore();
  const { updatePost } = usePosts();
  
  const [content, setContent] = useState("");
  const [isSubmitting, setIsSubmitting] = useState(false);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    if (editingPost) {
      setContent(editingPost.content);
    }
  }, [editingPost]);

  const handleUpdate = async () => {
    if (!editingPost || !content.trim()) return;

    setIsSubmitting(true);
    setError(null);
    try {
      const response = await updatePost(editingPost.id, content);
      if (response.success) {
        closeEditPostModal();
        window.location.reload(); 
      } else {
        setError(response.detail || "Failed to update transmission");
      }
    } catch (err: any) {
      setError(err.message || "An error occurred while updating the post");
    } finally {
      setIsSubmitting(false);
    }
  };

  return (
    <AnimatePresence>
      {editingPost && (
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
                Modify Transmission
              </h3>
              <button
                onClick={closeEditPostModal}
                className="text-foreground/40 hover:text-foreground transition-colors"
              >
                <X size={20} />
              </button>
            </div>

            {/* Body */}
            <div className="p-6 flex gap-4">
              <div className="size-12 rounded-full bg-foreground/5 flex-shrink-0 flex items-center justify-center text-xs font-bold text-foreground/20 border border-foreground/10 uppercase">
                {editingPost.author.username.slice(0, 2)}
              </div>
              <div className="flex-1 space-y-4">
                <textarea
                  autoFocus
                  disabled={isSubmitting}
                  value={content}
                  onChange={(e) => setContent(e.target.value)}
                  placeholder="What's resonating in the void?"
                  className="w-full min-h-[150px] bg-transparent resize-none text-lg outline-none placeholder:text-foreground/20 leading-relaxed disabled:opacity-50"
                />
                
                {error && (
                  <p className="text-red-500 text-xs uppercase tracking-wider font-bold">
                    {error}
                  </p>
                )}
              </div>
            </div>

            {/* Footer */}
            <div className="p-4 border-t border-foreground/10 flex justify-between items-center">
              <div className="text-[10px] text-foreground/30 uppercase tracking-[1px]">
                {content.length} / 280 characters
              </div>
              <div className="flex gap-3">
                <button
                  onClick={closeEditPostModal}
                  className="px-6 py-3 rounded-sm font-bold text-[11px] uppercase tracking-[2px] border border-foreground/10 hover:bg-foreground/5 transition-all"
                >
                  Cancel
                </button>
                <button
                  disabled={!content.trim() || isSubmitting || content === editingPost.content}
                  onClick={handleUpdate}
                  className={`px-8 py-3 rounded-sm font-bold text-[11px] uppercase tracking-[2px] transition-all flex items-center gap-2 ${
                    content.trim() && !isSubmitting && content !== editingPost.content
                      ? "bg-foreground text-background hover:opacity-90 active:scale-[0.98]"
                      : "bg-foreground/5 text-foreground/20 cursor-not-allowed"
                  }`}
                >
                  {isSubmitting ? (
                    <>
                      <Loader2 size={14} className="animate-spin" />
                      Updating...
                    </>
                  ) : (
                    "Update Post"
                  )}
                </button>
              </div>
            </div>
          </motion.div>
        </div>
      )}
    </AnimatePresence>
  );
}
