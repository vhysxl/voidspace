"use client";

import { useState, useRef, useEffect } from "react";
import Link from "next/link";
import { MoreHorizontal, Edit2, Trash2, Share2 } from "lucide-react";
import { motion, AnimatePresence } from "framer-motion";

interface PostHeaderProps {
  author: {
    username: string;
    profile: {
      display_name: string;
    };
  };
  formattedDate: string;
  isOwner: boolean;
  onEdit: () => void;
  onDelete: () => void;
}

export default function PostHeader({ author, formattedDate, isOwner, onEdit, onDelete }: PostHeaderProps) {
  const [showOptions, setShowOptions] = useState(false);
  const optionsRef = useRef<HTMLDivElement>(null);

  useEffect(() => {
    const handleClickOutside = (event: MouseEvent) => {
      if (optionsRef.current && !optionsRef.current.contains(event.target as Node)) {
        setShowOptions(false);
      }
    };
    document.addEventListener("mousedown", handleClickOutside);
    return () => document.removeEventListener("mousedown", handleClickOutside);
  }, []);

  return (
    <div className="flex justify-between items-start relative">
      <div className="flex items-center gap-1.5 min-w-0">
        <Link
          href={`/profile/${author.username}`}
          className="font-bold text-sm md:text-[15px] text-foreground tracking-tight truncate hover:underline"
          onClick={(e) => e.stopPropagation()}
        >
          {author.profile.display_name || author.username}
        </Link>
        <span className="text-foreground/40 text-[11px] md:text-xs tracking-wider truncate">
          @{author.username}
        </span>
        <span className="text-foreground/20 text-xs">•</span>
        <span className="text-foreground/40 text-[11px] md:text-xs tracking-wider shrink-0">
          {formattedDate}
        </span>
      </div>

      <div className="relative" ref={optionsRef}>
        <button
          className="text-foreground/20 hover:text-foreground p-1 transition-colors"
          onClick={(e) => {
            e.stopPropagation();
            setShowOptions(!showOptions);
          }}
        >
          <MoreHorizontal size={18} />
        </button>

        <AnimatePresence>
          {showOptions && (
            <motion.div
              initial={{ opacity: 0, scale: 0.95, y: -10 }}
              animate={{ opacity: 1, scale: 1, y: 0 }}
              exit={{ opacity: 0, scale: 0.95, y: -10 }}
              className="absolute right-0 top-full mt-2 w-48 bg-void-dark border border-foreground/10 rounded-sm shadow-2xl z-50 overflow-hidden"
              onClick={(e) => e.stopPropagation()}
            >
              {isOwner ? (
                <>
                  <button
                    onClick={() => {
                      setShowOptions(false);
                      onEdit();
                    }}
                    className="w-full px-4 py-3 flex items-center gap-3 text-sm font-bold uppercase tracking-wider text-foreground/60 hover:bg-foreground/5 hover:text-foreground transition-colors text-left"
                  >
                    <Edit2 size={14} />
                    Edit Post
                  </button>
                  <button
                    onClick={() => {
                      setShowOptions(false);
                      onDelete();
                    }}
                    className="w-full px-4 py-3 flex items-center gap-3 text-sm font-bold uppercase tracking-wider text-red-500/80 hover:bg-red-500/5 hover:text-red-500 transition-colors text-left"
                  >
                    <Trash2 size={14} />
                    Delete
                  </button>
                </>
              ) : (
                <button
                  className="w-full px-4 py-3 flex items-center gap-3 text-sm font-bold uppercase tracking-wider text-foreground/60 hover:bg-foreground/5 hover:text-foreground transition-colors text-left"
                >
                  <Share2 size={14} />
                  Share Post
                </button>
              )}
            </motion.div>
          )}
        </AnimatePresence>
      </div>
    </div>
  );
}
