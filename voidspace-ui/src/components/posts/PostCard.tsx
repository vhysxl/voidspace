"use client";

import { useState } from "react";
import { PostCardProps } from "@/types";
import { motion } from "framer-motion";
import { useAuthStore } from "@/store/useAuthStore";
import { useUIStore } from "@/store/useUIStore";
import { usePostStore } from "@/store/usePostStore";
import { usePosts } from "@/hooks/usePosts";
import { useRouter } from "next/navigation";

// UI Components
import ConfirmationModal from "@/components/ui/ConfirmationModal";
import ImageModal from "@/components/ui/ImageModal";

// Atomic Components
import PostAvatar from "./components/PostAvatar";
import PostHeader from "./components/PostHeader";
import PostMedia from "./components/PostMedia";
import PostActions from "./components/PostActions";

export default function PostCard({ post, isDetailed = false, onCommentClick }: PostCardProps) {
  const { author, content, post_images, created_at } = post;
  const { user } = useAuthStore();
  const { openEditPostModal } = useUIStore();
  const { setActivePost } = usePostStore();
  const { deletePost } = usePosts();
  const router = useRouter();

  const [showDeleteConfirm, setShowDeleteConfirm] = useState(false);
  const [isDeleting, setIsDeleting] = useState(false);
  const [selectedImage, setSelectedImage] = useState<string | undefined>(undefined);

  const isOwner = user?.id && author?.id ? user.id === author.id : false;

  const formattedDate = new Date(created_at).toLocaleDateString('en-US', {
    month: 'short',
    day: 'numeric'
  });

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
          <PostAvatar
            username={author?.username || "unknown"}
            avatarUrl={author?.profile?.avatar_url}
          />

          <div className="flex-1 space-y-3 min-w-0">
            <PostHeader
              author={author || { username: "unknown", profile: { display_name: "Unknown User" } } as any}
              formattedDate={formattedDate}
              isOwner={isOwner}
              onEdit={() => openEditPostModal(post)}
              onDelete={() => setShowDeleteConfirm(true)}
            />

            <p className="text-foreground/90 text-sm md:text-base leading-relaxed break-words font-manrope">
              {content}
            </p>

            <PostMedia
              images={post_images}
              isDetailed={isDetailed}
              onImageClick={setSelectedImage}
            />

            <PostActions
              post={post}
              onCommentClick={onCommentClick}
            />
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

      <ImageModal
        isOpen={!!selectedImage}
        onClose={() => setSelectedImage(undefined)}
        src={selectedImage}
        alt="Post content"
      />
    </>
  );
}

