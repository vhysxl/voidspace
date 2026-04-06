"use client";

import Link from "next/link";
import { useAuthStore } from "@/store/useAuthStore";
import { LogOut } from "lucide-react";

export default function UserMiniProfile() {
  const { user, isLoggedIn, logout, _hasHydrated } = useAuthStore();

  if (!_hasHydrated) {
    return (
      <div className="p-6 border-t border-foreground/5 animate-pulse">
        <div className="h-10 bg-foreground/5 rounded-full" />
      </div>
    );
  }

  if (!isLoggedIn || !user) {
    return (
      <div className="p-6 border-t border-foreground/5">
        <Link
          href="/auth/login"
          className="w-full border border-foreground/20 text-foreground/40 h-[44px] rounded-sm font-bold text-[11px] tracking-[2px] uppercase flex items-center justify-center hover:text-foreground hover:border-foreground/40 transition-all"
        >
          Sign In
        </Link>
      </div>
    );
  }

  return (
    <div className="mx-2 mb-4 flex items-center justify-between group">
      <Link 
        href="/profile"
        className="flex-1 p-3 hover:bg-foreground/5 transition-colors rounded-full flex items-center gap-3 cursor-pointer overflow-hidden"
      >
        {user.profile.avatar_url ? (
          <img 
            src={user.profile.avatar_url} 
            alt={user.username} 
            className="size-10 rounded-full object-cover border border-foreground/10"
          />
        ) : (
          <div className="size-10 rounded-full bg-foreground/10 border border-foreground/5 flex items-center justify-center text-foreground/20 font-bold text-xs uppercase shrink-0">
            {user.username.slice(0, 2)}
          </div>
        )}
        <div className="flex flex-col truncate">
          <span className="text-foreground text-[13px] font-bold tracking-[0.5px] truncate">
            {user.profile.display_name || user.username}
          </span>
          <span className="text-foreground/40 text-[11px] uppercase tracking-[1px] truncate">
            @{user.username}
          </span>
        </div>
      </Link>

      <button 
        onClick={(e) => {
          e.preventDefault();
          logout();
        }}
        className="text-foreground/40 hover:text-red-500 transition-colors p-3 mr-1"
        title="Log Out"
      >
        <LogOut size={16} />
      </button>
    </div>
  );
}
