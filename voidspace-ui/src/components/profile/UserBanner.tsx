"use client";

import { UserBanner as UserBannerType } from "@/types";
import Link from "next/link";

interface UserBannerProps {
  user: UserBannerType;
}

export default function UserBanner({ user }: UserBannerProps) {
  const { username, display_name, avatar_url } = user;

  return (
    <Link 
      href={`/profile/${username}`}
      className="flex items-center justify-between p-4 border border-foreground/10 bg-foreground/[0.02] rounded-sm group hover:border-foreground/20 transition-all hover:bg-foreground/[0.04]"
    >
      <div className="flex items-center gap-3">
        <div className="size-12 rounded-full bg-foreground/10 border border-foreground/5 overflow-hidden flex items-center justify-center">
          {avatar_url ? (
            <img src={avatar_url} alt={username} className="w-full h-full object-cover" />
          ) : (
            <span className="text-xs font-bold uppercase text-foreground/20">
              {username.slice(0, 2)}
            </span>
          )}
        </div>
        <div>
          <p className="text-sm font-bold tracking-tight text-foreground group-hover:underline">
            {display_name || username}
          </p>
          <p className="text-[10px] text-foreground/40 tracking-[1px]">
            @{username}
          </p>
        </div>
      </div>
      
      {/* Follow button could go here, but UserBanner type doesn't include is_followed status yet */}
      <div className="px-4 py-2 border border-foreground/10 text-[9px] font-bold uppercase tracking-[2px] text-foreground/20 group-hover:text-foreground/60 transition-colors">
        View Profile
      </div>
    </Link>
  );
}
