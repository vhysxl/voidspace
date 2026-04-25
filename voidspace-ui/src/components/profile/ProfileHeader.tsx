"use client";

import { useState, useEffect } from "react";
import { User } from "@/types";
import { Calendar, MapPin, Link as LinkIcon, Search } from "lucide-react";
import Button from "@/components/ui/Button";
import Link from "next/link";
import { motion } from "framer-motion";
import { useUIStore } from "@/store/useUIStore";
import ImageModal from "@/components/ui/ImageModal";
import { useUser } from "@/hooks/useUser";

interface ProfileHeaderProps {
  user: User;
  isOwnProfile?: boolean;
  activeTab?: string;
  onTabChange?: (tab: string) => void;
}

export default function ProfileHeader({ 
  user, 
  isOwnProfile = true,
  activeTab = "Posts",
  onTabChange
}: ProfileHeaderProps) {
  const { profile, username, created_at, is_followed } = user;
  const { openEditProfileModal } = useUIStore();
  const { followUser, unfollowUser, me } = useUser();
  const [previewImage, setPreviewImage] = useState<{ src: string; alt: string } | null>(null);
  
  const [isFollowed, setIsFollowed] = useState(is_followed);
  const [followersCount, setFollowersCount] = useState(profile.followers);
  const [isActionLoading, setIsActionLoading] = useState(false);

  // Sync state if user prop changes
  useEffect(() => {
    setIsFollowed(is_followed);
    setFollowersCount(profile.followers);
  }, [is_followed, profile.followers]);

  const handleFollowToggle = async () => {
    if (isActionLoading) return;
    
    setIsActionLoading(true);
    
    // Optimistic update
    const previousFollowed = isFollowed;
    const previousCount = followersCount;
    
    setIsFollowed(!previousFollowed);
    setFollowersCount(prev => previousFollowed ? prev - 1 : prev + 1);

    try {
      const response = previousFollowed 
        ? await unfollowUser(username)
        : await followUser(username);

      if (response.success) {
        // Refresh current user stats to update following count globally
        me();
      } else {
        // Revert on failure
        setIsFollowed(previousFollowed);
        setFollowersCount(previousCount);
      }
    } catch (error) {
      console.error("Follow action failed:", error);
      // Revert on error
      setIsFollowed(previousFollowed);
      setFollowersCount(previousCount);
    } finally {
      setIsActionLoading(false);
    }
  };

  const joinedDate = new Date(created_at).toLocaleDateString('en-US', {
    month: 'long',
    year: 'numeric'
  });

  const tabs = ["Posts", "Comments"];

  return (
    <div className="border-b border-foreground/10">
      {/* Banner */}
      <div 
        className="h-48 md:h-64 bg-void-dark relative overflow-hidden cursor-zoom-in group"
        onClick={() => profile.banner_url && setPreviewImage({ src: profile.banner_url, alt: "Profile Banner" })}
      >
        {profile.banner_url ? (
          <img
            src={profile.banner_url}
            alt="Banner"
            className="w-full h-full object-cover transition-transform duration-700 group-hover:scale-105"
          />
        ) : (
          <div className="w-full h-full bg-linear-to-br from-void-gray to-background opacity-50" />
        )}
        {profile.banner_url && (
          <div className="absolute inset-0 bg-black/20 opacity-0 group-hover:opacity-100 transition-opacity flex items-center justify-center">
            <Search className="text-white/50" size={24} />
          </div>
        )}
      </div>

      {/* Profile Info Area */}
      <div className="max-w-4xl mx-auto px-4 md:px-8 pb-6 relative">
        {/* Avatar - Overlapping Banner */}
        <div className="relative -mt-16 mb-4 flex justify-between items-end">
          <div 
            className="size-32 md:size-40 rounded-full border-4 border-background bg-void-gray overflow-hidden shrink-0 cursor-zoom-in group/avatar relative shadow-2xl"
            onClick={() => profile.avatar_url && setPreviewImage({ src: profile.avatar_url, alt: "Profile Avatar" })}
          >
            {profile.avatar_url ? (
              <img
                src={profile.avatar_url}
                alt={username}
                className="w-full h-full object-cover transition-transform duration-500 group-hover/avatar:scale-110"
              />
            ) : (
              <div className="w-full h-full flex items-center justify-center uppercase bg-void-gray text-foreground/20 text-4xl font-bold">
                {username.slice(0, 2)}
              </div>
            )}
            {profile.avatar_url && (
              <div className="absolute inset-0 bg-black/20 opacity-0 group-hover/avatar:opacity-100 transition-opacity flex items-center justify-center">
                <Search className="text-white/50" size={20} />
              </div>
            )}
          </div>

          <div className="mb-2">
            {isOwnProfile ? (
              <Button 
                variant="secondary" 
                className="w-auto px-6 h-10 text-[11px]"
                onClick={openEditProfileModal}
              >
                Edit Profile
              </Button>
            ) : (
              <Button 
                variant={isFollowed ? "secondary" : "primary"}
                className="w-auto px-8 h-10 text-[11px]"
                onClick={handleFollowToggle}
                isLoading={isActionLoading}
              >
                {isFollowed ? "Unfollow" : "Follow"}
              </Button>
            )}
          </div>
        </div>

        {/* User Details */}
        <div className="space-y-1">
          <h1 className="text-2xl md:text-3xl font-bold text-foreground tracking-tight font-space-grotesk">
            {profile.display_name || username}
          </h1>
          <p className="text-foreground/40 text-sm md:text-base tracking-[1px]">
            @{username}
          </p>
        </div>

        {profile.bio && (
          <p className="mt-4 text-foreground/90 text-[15px] leading-relaxed max-w-2xl font-manrope">
            {profile.bio}
          </p>
        )}

        {/* Metadata */}
        <div className="flex flex-wrap gap-x-6 gap-y-2 mt-4 text-foreground/40 text-sm tracking-wide font-manrope">
          {profile.location && (
            <div className="flex items-center gap-1.5">
              <MapPin size={16} />
              <span>{profile.location}</span>
            </div>
          )}
          <div className="flex items-center gap-1.5">
            <Calendar size={16} />
            <span>Joined {joinedDate}</span>
          </div>
        </div>

        {/* Stats */}
        <div className="flex gap-6 mt-6 font-manrope">
          <Link 
            href={`/profile/${username}/relations?tab=following`}
            className="flex gap-1.5 items-baseline hover:underline decoration-foreground/20 underline-offset-4 transition-all"
          >
            <span className="text-foreground font-bold">{profile.following}</span>
            <span className="text-foreground/40 text-sm uppercase tracking-wider">Following</span>
          </Link>
          <Link 
            href={`/profile/${username}/relations?tab=followers`}
            className="flex gap-1.5 items-baseline hover:underline decoration-foreground/20 underline-offset-4 transition-all"
          >
            <span className="text-foreground font-bold">{followersCount}</span>
            <span className="text-foreground/40 text-sm uppercase tracking-wider">Followers</span>
          </Link>
        </div>
      </div>

      {/* Tabs */}
      <div className="border-t border-foreground/5">
        <div className="max-w-4xl mx-auto flex px-4 md:px-8">
          {tabs.map((tab) => {
            const isActive = activeTab === tab;
            return (
              <div
                key={tab}
                onClick={() => onTabChange?.(tab)}
                className={`px-6 py-4 text-sm font-bold uppercase tracking-[2px] cursor-pointer transition-colors relative group ${isActive ? "text-foreground" : "text-foreground/40 hover:text-foreground"}`}
              >
                {tab}
                {isActive && (
                  <motion.div 
                    layoutId="profileActiveTab"
                    className="absolute bottom-0 left-0 right-0 h-1 bg-foreground shadow-[0_0_8px_var(--color-foreground)]" 
                  />
                )}
              </div>
            );
          })}
        </div>
      </div>

      {/* Image Preview Modal */}
      <ImageModal 
        isOpen={!!previewImage} 
        onClose={() => setPreviewImage(null)} 
        src={previewImage?.src || undefined} 
        alt={previewImage?.alt}
      />
    </div>
  );
}
