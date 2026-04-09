"use client";

import { useState, useRef, useEffect } from "react";
import { Comment } from "@/types";
import Link from "next/link";
import { motion, AnimatePresence } from "framer-motion";
import { MoreHorizontal, Trash2 } from "lucide-react";
import { useAuthStore } from "@/store/useAuthStore";
import { useComments } from "@/hooks/useComments";
import ConfirmationModal from "@/components/ui/ConfirmationModal";

interface CommentItemProps {
  comment: Comment;
}

export default function CommentItem({ comment }: CommentItemProps) {
  const { author, content, created_at, comment_id } = comment;
  const { user } = useAuthStore();
  const { deleteComment } = useComments();

  const [showOptions, setShowOptions] = useState(false);
  const [showDeleteConfirm, setShowDeleteConfirm] = useState(false);
  const [isDeleting, setIsDeleting] = useState(false);
  const optionsRef = useRef<HTMLDivElement>(null);

  const isOwner = user?.id && author?.id ? user.id === author.id : false;

  const formattedDate = new Date(created_at).toLocaleDateString('en-US', {
    month: 'short',
    day: 'numeric'
  });

  useEffect(() => {
    const handleClickOutside = (event: MouseEvent) => {
      if (optionsRef.current && !optionsRef.current.contains(event.target as Node)) {
        setShowOptions(false);
      }
    };
    document.addEventListener("mousedown", handleClickOutside);
    return () => document.removeEventListener("mousedown", handleClickOutside);
  }, []);

  const handleDelete = async () => {
    setIsDeleting(true);
    try {
      const res = await deleteComment(comment_id);
      if (res.success) {
        setShowDeleteConfirm(false);
        window.location.reload(); // Refresh to reflect change
      } else {
        alert(res.detail || "Failed to delete comment");
      }
    } catch (err: any) {
      alert(err.message || "An error occurred");
    } finally {
      setIsDeleting(false);
    }
  };

  return (
    <>
      <motion.div 
        initial={{ opacity: 0, y: 10 }}
        animate={{ opacity: 1, y: 0 }}
        className="p-4 md:p-6 border-b border-foreground/10 hover:bg-foreground/[0.02] transition-colors group"
      >
        <div className="flex gap-4">
          {/* Avatar */}
          <Link 
            href={author ? `/profile/${author.username}` : "#"}
            className={`shrink-0 pt-1 ${!author ? "pointer-events-none" : ""}`}
          >
            <div className="size-8 md:size-10 rounded-full bg-void-gray border border-foreground/5 overflow-hidden flex items-center justify-center">
              {author?.profile?.avatar_url ? (
                <img src={author.profile.avatar_url} alt="" className="w-full h-full object-cover" />
              ) : (
                <span className="text-[10px] font-bold uppercase text-foreground/20">
                  {author ? author.username.slice(0, 2) : "??"}
                </span>
              )}
            </div>
          </Link>

          {/* Content Area */}
          <div className="flex-1 space-y-2 min-w-0">
            {/* Header */}
            <div className="flex justify-between items-start">
              <div className="flex items-center gap-1.5 min-w-0">
                {author ? (
                  <>
                    <Link 
                      href={`/profile/${author.username}`}
                      className="font-bold text-sm text-foreground tracking-tight truncate hover:underline"
                    >
                      {author.profile.display_name || author.username}
                    </Link>
                    <span className="text-foreground/40 text-[10px] md:text-xs tracking-wider truncate">
                      @{author.username}
                    </span>
                  </>
                ) : (
                  <span className="font-bold text-sm text-foreground/40 italic">
                    Unknown User
                  </span>
                )}
                <span className="text-foreground/20 text-xs">•</span>
                <span className="text-foreground/40 text-[10px] md:text-xs tracking-wider shrink-0">
                  {formattedDate}
                </span>
              </div>
              
              <div className="relative" ref={optionsRef}>
                <button 
                  onClick={() => setShowOptions(!showOptions)}
                  className="text-foreground/20 hover:text-foreground p-1 transition-colors cursor-pointer"
                >
                  <MoreHorizontal size={16} />
                </button>

                <AnimatePresence>
                  {showOptions && (
                    <motion.div
                      initial={{ opacity: 0, scale: 0.95, y: -10 }}
                      animate={{ opacity: 1, scale: 1, y: 0 }}
                      exit={{ opacity: 0, scale: 0.95, y: -10 }}
                      className="absolute right-0 top-full mt-1 w-40 bg-void-dark border border-foreground/10 rounded-sm shadow-2xl z-50 overflow-hidden"
                    >
                      {isOwner ? (
                        <button
                          onClick={() => {
                            setShowOptions(false);
                            setShowDeleteConfirm(true);
                          }}
                          className="w-full px-4 py-3 flex items-center gap-3 text-sm font-bold uppercase tracking-wider text-red-500/80 hover:bg-red-500/5 hover:text-red-500 transition-colors text-left cursor-pointer"
                        >
                          <Trash2 size={14} />
                          Delete
                        </button>
                      ) : (
                        <div className="px-4 py-3 text-[10px] font-bold uppercase tracking-wider text-foreground/40 italic">
                          No Actions
                        </div>
                      )}
                    </motion.div>
                  )}
                </AnimatePresence>
              </div>
            </div>

            {/* Text Content */}
            <p className="text-foreground/80 text-sm md:text-[15px] leading-relaxed break-words font-manrope">
              {content}
            </p>
          </div>
        </div>
      </motion.div>

      <ConfirmationModal
        isOpen={showDeleteConfirm}
        onClose={() => setShowDeleteConfirm(false)}
        onConfirm={handleDelete}
        isLoading={isDeleting}
        title="Delete Resonance"
        description="Are you sure you want to silence this transmission? This action is irreversible."
        confirmText="Erase Comment"
        variant="danger"
      />
    </>
  );
}
