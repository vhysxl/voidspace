"use client";

import { useState, useRef, useEffect } from "react";
import { Post } from "@/types";
import { MessageSquare, Heart, Share2, MoreHorizontal, Edit2, Trash2 } from "lucide-react";
import Link from "next/link";
import { motion, AnimatePresence } from "framer-motion";
import { useAuthStore } from "@/store/useAuthStore";
import { useUIStore } from "@/store/useUIStore";
import { usePostStore } from "@/store/usePostStore";
import { usePosts } from "@/hooks/usePosts";
import { useRouter } from "next/navigation";
import ConfirmationModal from "@/components/ui/ConfirmationModal";

interface PostCardProps {
  post: Post;
  isDetailed?: boolean;
}

export default function PostCard({ post, isDetailed = false }: PostCardProps) {
  const { author, content, post_images, created_at, likes_count, comments_count, is_liked } = post;
  const { user } = useAuthStore();
  const { openEditPostModal } = useUIStore();
  const { setActivePost } = usePostStore();
  const { deletePost } = usePosts();
  const router = useRouter();

  const [showOptions, setShowOptions] = useState(false);
  const [showDeleteConfirm, setShowDeleteConfirm] = useState(false);
  const [isDeleting, setIsDeleting] = useState(false);
  const optionsRef = useRef<HTMLDivElement>(null);

  const isOwner = user?.id === author.id;

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
      const res = await deletePost(post.id);
      if (res.success) {
        setShowDeleteConfirm(false);
        if (isDetailed) {
          router.push("/");
        } else {
          window.location.reload();
        }
      } else {
        alert(res.detail || "Failed to delete post");
      }
    } catch (err: any) {
      alert(err.message || "An error occurred");
    } finally {
      setIsDeleting(false);
    }
  };

  const handleCardClick = () => {
    if (!isDetailed) {
      setActivePost(post);
      router.push(`/post/${post.id}`);
    }
  };

  return (
    <>
      <motion.div 
        initial={{ opacity: 0, y: 10 }}
        animate={{ opacity: 1, y: 0 }}
        onClick={handleCardClick}
        className={`p-4 md:p-6 border-b border-foreground/10 hover:bg-foreground/[0.02] transition-colors group cursor-pointer ${isDetailed ? "bg-background border-x !cursor-default" : ""}`}
      >
        <div className="flex gap-4">
          {/* Avatar */}
          <Link 
            href={`/profile/${author.username}`}
            className="shrink-0 pt-1"
            onClick={(e) => e.stopPropagation()}
          >
            <div className="size-10 md:size-12 rounded-full bg-void-gray border border-foreground/5 overflow-hidden flex items-center justify-center">
              {author.profile.avatar_url ? (
                <img src={author.profile.avatar_url} alt="" className="w-full h-full object-cover" />
              ) : (
                <span className="text-xs font-bold uppercase text-foreground/20">
                  {author.username.slice(0, 2)}
                </span>
              )}
            </div>
          </Link>

          {/* Content Area */}
          <div className="flex-1 space-y-3 min-w-0">
            {/* Header */}
            <div className="flex justify-between items-start relative">
              <div className="flex items-center gap-1.5 min-w-0">
                <Link 
                  href={`/profile/${author.username}`}
                  className="font-bold text-sm md:text-[15px] text-foreground uppercase tracking-tight truncate hover:underline"
                  onClick={(e) => e.stopPropagation()}
                >
                  {author.profile.display_name || author.username}
                </Link>
                <span className="text-foreground/40 text-[11px] md:text-xs tracking-wider uppercase truncate">
                  @{author.username}
                </span>
                <span className="text-foreground/20 text-xs">•</span>
                <span className="text-foreground/40 text-[11px] md:text-xs tracking-wider uppercase shrink-0">
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
                              openEditPostModal(post);
                            }}
                            className="w-full px-4 py-3 flex items-center gap-3 text-sm font-bold uppercase tracking-wider text-foreground/60 hover:bg-foreground/5 hover:text-foreground transition-colors text-left"
                          >
                            <Edit2 size={14} />
                            Edit Post
                          </button>
                          <button 
                            onClick={() => {
                              setShowOptions(false);
                              setShowDeleteConfirm(true);
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

            {/* Post Content */}
            <p className="text-foreground/90 text-sm md:text-base leading-relaxed break-words font-manrope">
              {content}
            </p>

            {/* Images Grid */}
            {post_images && post_images.length > 0 && (
              <div className={`mt-4 rounded-sm overflow-hidden border border-foreground/10 bg-void-dark grid gap-1 ${
                post_images.length === 1 ? "grid-cols-1" : "grid-cols-2"
              }`}>
                {post_images.map((img, idx) => (
                  <div 
                    key={idx} 
                    className={`relative overflow-hidden group/img ${
                      post_images.length === 3 && idx === 0 ? "row-span-2" : ""
                    } ${
                      post_images.length === 1 ? "aspect-auto" : "aspect-square md:aspect-video"
                    }`}
                  >
                    <img 
                      src={img.image_url} 
                      alt={`Post content ${idx + 1}`} 
                      className="w-full h-full object-cover transition-transform duration-500 group-hover/img:scale-[1.03]" 
                    />
                  </div>
                ))}
              </div>
            )}

            {/* Footer Actions */}
            <div className="flex items-center gap-8 pt-4">
              <div 
                className="flex items-center gap-2 text-foreground/40 hover:text-blue-500 transition-colors group/action"
              >
                <div className="size-8 rounded-full flex items-center justify-center group-hover/action:bg-blue-500/10 transition-colors">
                  <MessageSquare size={18} />
                </div>
                <span className="text-[11px] font-bold tracking-widest">{comments_count}</span>
              </div>

              <button 
                className={`flex items-center gap-2 transition-colors group/action ${is_liked ? "text-pink-500" : "text-foreground/40 hover:text-pink-500"}`}
                onClick={(e) => {
                  e.stopPropagation();
                  // handle like
                }}
              >
                <div className={`size-8 rounded-full flex items-center justify-center group-hover/action:bg-pink-500/10 transition-colors ${is_liked ? "bg-pink-500/5" : ""}`}>
                  <Heart size={18} fill={is_liked ? "currentColor" : "none"} />
                </div>
                <span className="text-[11px] font-bold tracking-widest">{likes_count}</span>
              </button>

              <button 
                className="flex items-center gap-2 text-foreground/40 hover:text-foreground transition-colors group/action ml-auto"
                onClick={(e) => {
                  e.stopPropagation();
                  // handle share
                }}
              >
                <div className="size-8 rounded-full flex items-center justify-center group-hover/action:bg-foreground/5 transition-colors">
                  <Share2 size={18} />
                </div>
              </button>
            </div>
          </div>
        </div>
      </motion.div>

      <ConfirmationModal 
        isOpen={showDeleteConfirm}
        onClose={() => setShowDeleteConfirm(false)}
        onConfirm={handleDelete}
        isLoading={isDeleting}
        title="Delete Transmission"
        description="Are you absolutely sure you want to erase this data from the void? This action is irreversible and will remove all associated resonance."
        confirmText="Erase Post"
        variant="danger"
      />
    </>
  );
}
