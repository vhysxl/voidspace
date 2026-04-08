"use client";

import { User } from "@/types";
import { Calendar, MapPin, Link as LinkIcon } from "lucide-react";
import Button from "@/components/ui/Button";
import Link from "next/link";
import { motion } from "framer-motion";
import { useUIStore } from "@/store/useUIStore";

interface ProfileHeaderProps {
  user: User;
  isOwnProfile?: boolean;
}

export default function ProfileHeader({ user, isOwnProfile = true }: ProfileHeaderProps) {
  const { profile, username, created_at } = user;
  const { openEditProfileModal } = useUIStore();

  const joinedDate = new Date(created_at).toLocaleDateString('en-US', {
    month: 'long',
    year: 'numeric'
  });

  return (
    <div className="border-b border-foreground/10">
      {/* Banner */}
      <div className="h-48 md:h-64 bg-void-dark relative overflow-hidden">
        {profile.banner_url ? (
          <img
            src={profile.banner_url}
            alt="Banner"
            className="w-full h-full object-cover"
          />
        ) : (
          <div className="w-full h-full bg-gradient-to-br from-void-gray to-background opacity-50" />
        )}
      </div>

      {/* Profile Info Area */}
      <div className="max-w-4xl mx-auto px-4 md:px-8 pb-6 relative">
        {/* Avatar - Overlapping Banner */}
        <div className="relative -mt-16 mb-4 flex justify-between items-end">
          <div className="size-32 md:size-40 rounded-full border-4 border-background bg-void-gray overflow-hidden shrink-0">
            {profile.avatar_url ? (
              <img
                src={profile.avatar_url}
                alt={username}
                className="w-full h-full object-cover"
              />
            ) : (
              <div className="w-full h-full flex items-center justify-center bg-void-gray text-foreground/20 text-4xl font-bold uppercase">
                {username.slice(0, 2)}
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
              <Button className="w-auto px-8 h-10 text-[11px]">
                Follow
              </Button>
            )}
          </div>
        </div>

        {/* User Details */}
        <div className="space-y-1">
          <h1 className="text-2xl md:text-3xl font-bold text-foreground tracking-tight uppercase font-space-grotesk">
            {profile.display_name || username}
          </h1>
          <p className="text-foreground/40 text-sm md:text-base uppercase tracking-[1px]">
            @{username}
          </p>
        </div>

        {profile.bio && (
          <p className="mt-4 text-foreground/90 text-[15px] leading-relaxed max-w-2xl font-manrope">
            {profile.bio}
          </p>
        )}

        {/* Metadata */}
        <div className="flex flex-wrap gap-x-6 gap-y-2 mt-4 text-foreground/40 text-sm uppercase tracking-wide font-manrope">
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
            <span className="text-foreground font-bold">{profile.followers}</span>
            <span className="text-foreground/40 text-sm uppercase tracking-wider">Followers</span>
          </Link>
        </div>
      </div>

      {/* Tabs Placeholder */}
      <div className="border-t border-foreground/5">
        <div className="max-w-4xl mx-auto flex px-4 md:px-8">
          {["Posts", "Comments", "Media", "Likes"].map((tab, i) => (
            <div
              key={tab}
              className={`px-6 py-4 text-sm font-bold uppercase tracking-[2px] cursor-pointer transition-colors relative group ${i === 0 ? "text-foreground" : "text-foreground/40 hover:text-foreground"}`}
            >
              {tab}
              {i === 0 && (
                <motion.div 
                  layoutId="profileActiveTab"
                  className="absolute bottom-0 left-0 right-0 h-1 bg-foreground shadow-[0_0_8px_var(--color-foreground)]" 
                />
              )}
            </div>
          ))}
        </div>
      </div>
    </div>
  );
}
