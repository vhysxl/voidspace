"use client";

import { useState, useRef } from "react";
import { useUIStore } from "@/store/useUIStore";
import { useAuthStore } from "@/store/useAuthStore";
import { usePosts } from "@/hooks/usePosts";
import { useUpload } from "@/hooks/useUpload";
import { X, Image as ImageIcon, Sparkles, Loader2, Plus } from "lucide-react";
import { motion, AnimatePresence } from "framer-motion";
import { PostImage } from "@/types";

export default function NewPostModal() {
  const { isNewPostModalOpen, closeNewPostModal } = useUIStore();
  const { user } = useAuthStore();
  const { createPost } = usePosts();
  const { uploadImage } = useUpload();
  
  const [content, setContent] = useState("");
  const [images, setImages] = useState<{ file: File; preview: string }[]>([]);
  const [isSubmitting, setIsSubmitting] = useState(false);
  const [error, setError] = useState<string | null>(null);
  const fileInputRef = useRef<HTMLInputElement>(null);

  const handleFileChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const files = Array.from(e.target.files || []);
    if (files.length + images.length > 5) {
      setError("Maximum 5 images allowed");
      return;
    }

    const newImages = files.map(file => ({
      file,
      preview: URL.createObjectURL(file)
    }));

    setImages(prev => [...prev, ...newImages]);
    setError(null);
  };

  const removeImage = (index: number) => {
    setImages(prev => {
      const updated = [...prev];
      URL.revokeObjectURL(updated[index].preview);
      updated.splice(index, 1);
      return updated;
    });
  };

  const handlePost = async () => {
    if (!content.trim() && images.length === 0) return;

    setIsSubmitting(true);
    setError(null);
    try {
      // 1. Upload all images first
      const uploadedImages: PostImage[] = [];
      
      for (let i = 0; i < images.length; i++) {
        const { file } = images[i];
        const result = await uploadImage(file, "posts");
        uploadedImages.push({
          image_url: result.url,
          order: i + 1,
          width: result.width,
          height: result.height
        });
      }

      // 2. Create post with image data
      const response = await createPost(content, uploadedImages);
      if (response.success) {
        setContent("");
        setImages([]);
        closeNewPostModal();
        window.location.reload(); 
      } else {
        setError(response.detail || "Failed to launch post");
      }
    } catch (err: any) {
      setError(err.message || "An error occurred while creating the post");
    } finally {
      setIsSubmitting(false);
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
            className="w-full max-w-xl bg-background border border-foreground/20 rounded-sm shadow-2xl flex flex-col max-h-[90vh]"
          >
            {/* Header */}
            <div className="flex justify-between items-center p-4 border-b border-foreground/10">
              <h3 className="font-space-grotesk text-sm font-bold uppercase tracking-[1px] flex items-center gap-2">
                <Sparkles size={16} />
                New Broadcast
              </h3>
              <button
                onClick={closeNewPostModal}
                disabled={isSubmitting}
                className="text-foreground/40 hover:text-foreground transition-colors disabled:opacity-50"
              >
                <X size={20} />
              </button>
            </div>

            {/* Body */}
            <div className="p-6 flex gap-4 overflow-y-auto">
              <div className="size-12 rounded-full bg-foreground/5 flex-shrink-0 flex items-center justify-center text-xs font-bold text-foreground/20 border border-foreground/10">
                {user?.username?.slice(0, 2).toUpperCase() || "VO"}
              </div>
              <div className="flex-1 space-y-4">
                <textarea
                  autoFocus
                  disabled={isSubmitting}
                  value={content}
                  onChange={(e) => setContent(e.target.value)}
                  placeholder="What's resonating in the void?"
                  className="w-full min-h-[120px] bg-transparent resize-none text-lg outline-none placeholder:text-foreground/20 leading-relaxed disabled:opacity-50"
                />
                
                {/* Image Previews */}
                {images.length > 0 && (
                  <div className="grid grid-cols-2 gap-2 mt-4">
                    {images.map((img, idx) => (
                      <div key={idx} className="relative aspect-video rounded-sm overflow-hidden border border-foreground/10 group">
                        <img src={img.preview} alt="Preview" className="w-full h-full object-cover" />
                        <button
                          onClick={() => removeImage(idx)}
                          disabled={isSubmitting}
                          className="absolute top-2 right-2 size-8 bg-background/80 backdrop-blur-md rounded-full flex items-center justify-center text-foreground hover:bg-foreground hover:text-background transition-all opacity-0 group-hover:opacity-100 disabled:opacity-50"
                        >
                          <X size={14} />
                        </button>
                      </div>
                    ))}
                    {images.length < 5 && (
                      <button
                        onClick={() => fileInputRef.current?.click()}
                        disabled={isSubmitting}
                        className="aspect-video rounded-sm border border-dashed border-foreground/20 flex flex-col items-center justify-center text-foreground/20 hover:border-foreground/40 hover:text-foreground/40 transition-all disabled:opacity-50"
                      >
                        <Plus size={24} />
                        <span className="text-[10px] font-bold uppercase tracking-widest mt-2">Add Image</span>
                      </button>
                    )}
                  </div>
                )}

                {error && (
                  <p className="text-red-500 text-xs uppercase tracking-wider font-bold">
                    {error}
                  </p>
                )}

                {/* Hidden File Input */}
                <input
                  type="file"
                  ref={fileInputRef}
                  onChange={handleFileChange}
                  accept="image/*"
                  multiple
                  className="hidden"
                />

                {/* Media Attachment Action */}
                <div className="flex items-center gap-4 pt-4 border-t border-foreground/5">
                  <button 
                    onClick={() => fileInputRef.current?.click()}
                    disabled={isSubmitting || images.length >= 5}
                    className="text-foreground/40 hover:text-foreground transition-colors p-2 hover:bg-foreground/5 rounded-full disabled:opacity-20"
                  >
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
                disabled={(!content.trim() && images.length === 0) || isSubmitting}
                onClick={handlePost}
                className={`px-8 py-3 rounded-sm font-bold text-[11px] uppercase tracking-[2px] transition-all flex items-center gap-2 ${
                  (content.trim() || images.length > 0) && !isSubmitting
                    ? "bg-foreground text-background hover:opacity-90 active:scale-[0.98]"
                    : "bg-foreground/5 text-foreground/20 cursor-not-allowed"
                }`}
              >
                {isSubmitting ? (
                  <>
                    <Loader2 size={14} className="animate-spin" />
                    Launching...
                  </>
                ) : (
                  "Launch Post"
                )}
              </button>
            </div>
          </motion.div>
        </div>
      )}
    </AnimatePresence>
  );
}
