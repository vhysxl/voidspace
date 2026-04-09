"use client";

import { useState, useEffect, useRef } from "react";
import { useUIStore } from "@/store/useUIStore";
import { useAuthStore } from "@/store/useAuthStore";
import { useUser } from "@/hooks/useUser";
import { useUpload } from "@/hooks/useUpload";
import { X, Camera, MapPin, Loader2 } from "lucide-react";
import { motion, AnimatePresence } from "framer-motion";
import Input from "@/components/ui/Input";

export default function EditProfileModal() {
  const { isEditProfileModalOpen, closeEditProfileModal } = useUIStore();
  const { user } = useAuthStore();
  const { updateProfile } = useUser();
  const { uploadImage } = useUpload();
  
  const [formData, setFormData] = useState({
    display_name: "",
    bio: "",
    location: "",
    avatar_url: "",
    banner_url: "",
  });
  const [isSubmitting, setIsSubmitting] = useState(false);
  const [error, setError] = useState<string | null>(null);

  const MAX_FILE_SIZE = 10 * 1024 * 1024; // 10MB

  const avatarInputRef = useRef<HTMLInputElement>(null);
  const bannerInputRef = useRef<HTMLInputElement>(null);

  useEffect(() => {
    if (user) {
      setFormData({
        display_name: user.profile.display_name || "",
        bio: user.profile.bio || "",
        location: user.profile.location || "",
        avatar_url: user.profile.avatar_url || "",
        banner_url: user.profile.banner_url || "",
      });
    }
  }, [user, isEditProfileModalOpen]);

  const handleFileChange = async (e: React.ChangeEvent<HTMLInputElement>, type: "avatars" | "banners") => {
    const file = e.target.files?.[0];
    if (!file) return;

    if (file.size > MAX_FILE_SIZE) {
      setError("File size must be less than 10MB");
      return;
    }

    setIsSubmitting(true);
    setError(null);
    try {
      const result = await uploadImage(file, type);
      setFormData(prev => ({
        ...prev,
        [type === "avatars" ? "avatar_url" : "banner_url"]: result.url
      }));
    } catch (err: any) {
      setError(`Failed to upload ${type === "avatars" ? "avatar" : "banner"}`);
    } finally {
      setIsSubmitting(false);
    }
  };

  const handleSave = async () => {
    setIsSubmitting(true);
    setError(null);
    try {
      const response = await updateProfile({
        display_name: formData.display_name,
        bio: formData.bio,
        location: formData.location,
        avatar_url: formData.avatar_url,
        banner_url: formData.banner_url,
      });

      if (response.success) {
        closeEditProfileModal();
      } else {
        setError(response.detail || "Failed to update profile");
      }
    } catch (err: any) {
      setError(err.message || "An error occurred while updating profile");
    } finally {
      setIsSubmitting(false);
    }
  };

  return (
    <AnimatePresence>
      {isEditProfileModalOpen && (
        <div className="fixed inset-0 z-[100] flex items-center justify-center p-4 bg-background/80 backdrop-blur-sm">
          <motion.div
            initial={{ opacity: 0, scale: 0.95, y: 20 }}
            animate={{ opacity: 1, scale: 1, y: 0 }}
            exit={{ opacity: 0, scale: 0.95, y: 20 }}
            className="w-full max-w-xl bg-background border border-foreground/20 rounded-sm shadow-2xl flex flex-col max-h-[90vh] overflow-hidden"
          >
            {/* Header */}
            <div className="flex justify-between items-center p-4 border-b border-foreground/10 sticky top-0 bg-background z-10">
              <div className="flex items-center gap-4">
                <button
                  onClick={closeEditProfileModal}
                  disabled={isSubmitting}
                  className="text-foreground/40 hover:text-foreground transition-colors disabled:opacity-50"
                >
                  <X size={20} />
                </button>
                <h3 className="font-space-grotesk text-sm font-bold uppercase tracking-[1px]">
                  Edit Profile
                </h3>
              </div>
              <button
                onClick={handleSave}
                disabled={isSubmitting}
                className="px-6 py-2 bg-foreground text-background rounded-sm font-bold text-[11px] uppercase tracking-[2px] hover:opacity-90 active:scale-[0.98] transition-all flex items-center gap-2 disabled:opacity-50"
              >
                {isSubmitting ? (
                  <>
                    <Loader2 size={14} className="animate-spin" />
                    Saving...
                  </>
                ) : (
                  "Save"
                )}
              </button>
            </div>

            {/* Body */}
            <div className="flex-1 overflow-y-auto">
              {/* Banner Upload Area */}
              <div 
                onClick={() => bannerInputRef.current?.click()}
                className="relative h-32 md:h-44 bg-void-dark group overflow-hidden cursor-pointer"
              >
                {formData.banner_url ? (
                  <img src={formData.banner_url} alt="" className="w-full h-full object-cover opacity-50" />
                ) : (
                  <div className="w-full h-full bg-gradient-to-br from-void-gray to-background opacity-30" />
                )}
                <div className="absolute inset-0 flex items-center justify-center">
                  <div className="size-10 rounded-full bg-background/40 backdrop-blur-md flex items-center justify-center text-white group-hover:scale-110 transition-transform">
                    <Camera size={20} />
                  </div>
                </div>
                <input 
                  type="file" 
                  ref={bannerInputRef} 
                  className="hidden" 
                  accept="image/*"
                  onChange={(e) => handleFileChange(e, "banners")}
                />
              </div>

              {/* Avatar Upload Area */}
              <div className="px-6 relative">
                <div 
                  onClick={() => avatarInputRef.current?.click()}
                  className="size-24 md:size-32 rounded-full border-4 border-background bg-void-gray -mt-12 mb-4 relative overflow-hidden shadow-xl cursor-pointer group"
                >
                  {formData.avatar_url ? (
                    <img src={formData.avatar_url} alt="" className="w-full h-full object-cover opacity-50" />
                  ) : (
                    <div className="w-full h-full flex items-center justify-center bg-foreground/5 text-foreground/20 text-2xl font-bold uppercase">
                      {user?.username?.slice(0, 2)}
                    </div>
                  )}
                  <div className="absolute inset-0 flex items-center justify-center">
                    <div className="size-8 rounded-full bg-background/40 backdrop-blur-md flex items-center justify-center text-white group-hover:scale-110 transition-transform">
                      <Camera size={16} />
                    </div>
                  </div>
                  <input 
                    type="file" 
                    ref={avatarInputRef} 
                    className="hidden" 
                    accept="image/*"
                    onChange={(e) => handleFileChange(e, "avatars")}
                  />
                </div>
              </div>

              {/* Form Fields */}
              <div className="p-6 pt-2 space-y-8">
                {error && (
                  <div className="bg-red-500/5 border border-red-500/20 py-3 px-4 text-red-500 text-xs tracking-widest uppercase font-bold text-center">
                    {error}
                  </div>
                )}

                <Input
                  label="Display Name"
                  placeholder="Enter your display name"
                  value={formData.display_name}
                  onChange={(e) => setFormData({ ...formData, display_name: e.target.value })}
                  disabled={isSubmitting}
                />

                <div className="space-y-3 w-full">
                  <label className="text-[12px] text-foreground/40 tracking-[2.4px] uppercase block font-manrope px-1">
                    Bio
                  </label>
                  <div className="relative">
                    <textarea
                      value={formData.bio}
                      onChange={(e) => setFormData({ ...formData, bio: e.target.value })}
                      placeholder="Tell the void about yourself"
                      disabled={isSubmitting}
                      className="w-full bg-foreground/5 border border-foreground/10 px-4 py-3.5 text-foreground text-[16px] tracking-[0.32px] focus:outline-none focus:border-foreground/30 transition-all placeholder:text-foreground/10 font-manrope min-h-[120px] resize-none disabled:opacity-50"
                    />
                  </div>
                </div>

                <Input
                  label="Location"
                  placeholder="Where in the void are you?"
                  value={formData.location}
                  onChange={(e) => setFormData({ ...formData, location: e.target.value })}
                  disabled={isSubmitting}
                  action={<MapPin size={14} className="text-foreground/20" />}
                />
              </div>
            </div>
          </motion.div>
        </div>
      )}
    </AnimatePresence>
  );
}
