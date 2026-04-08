"use client";

import { Comment } from "@/types";
import Link from "next/link";
import { motion } from "framer-motion";
import { MoreHorizontal } from "lucide-react";

interface CommentItemProps {
  comment: Comment;
}

export default function CommentItem({ comment }: CommentItemProps) {
  const { author, content, created_at } = comment;

  const formattedDate = new Date(created_at).toLocaleDateString('en-US', {
    month: 'short',
    day: 'numeric'
  });

  return (
    <motion.div 
      initial={{ opacity: 0, y: 10 }}
      animate={{ opacity: 1, y: 0 }}
      className="p-4 md:p-6 border-b border-foreground/10 hover:bg-foreground/[0.02] transition-colors group"
    >
      <div className="flex gap-4">
        {/* Avatar */}
        <Link 
          href={`/profile/${author.username}`}
          className="shrink-0 pt-1"
        >
          <div className="size-8 md:size-10 rounded-full bg-void-gray border border-foreground/5 overflow-hidden flex items-center justify-center">
            {author.profile.avatar_url ? (
              <img src={author.profile.avatar_url} alt="" className="w-full h-full object-cover" />
            ) : (
              <span className="text-[10px] font-bold uppercase text-foreground/20">
                {author.username.slice(0, 2)}
              </span>
            )}
          </div>
        </Link>

        {/* Content Area */}
        <div className="flex-1 space-y-2 min-w-0">
          {/* Header */}
          <div className="flex justify-between items-start">
            <div className="flex items-center gap-1.5 min-w-0">
              <Link 
                href={`/profile/${author.username}`}
                className="font-bold text-sm text-foreground uppercase tracking-tight truncate hover:underline"
              >
                {author.profile.display_name || author.username}
              </Link>
              <span className="text-foreground/40 text-[10px] md:text-xs tracking-wider uppercase truncate">
                @{author.username}
              </span>
              <span className="text-foreground/20 text-xs">•</span>
              <span className="text-foreground/40 text-[10px] md:text-xs tracking-wider uppercase shrink-0">
                {formattedDate}
              </span>
            </div>
            <button className="text-foreground/20 hover:text-foreground p-1 transition-colors">
              <MoreHorizontal size={16} />
            </button>
          </div>

          {/* Text Content */}
          <p className="text-foreground/80 text-sm md:text-[15px] leading-relaxed break-words font-manrope">
            {content}
          </p>
        </div>
      </div>
    </motion.div>
  );
}
