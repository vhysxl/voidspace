"use client";

import {
  Home,
  Search,
  User,
  Plus,
  Settings
} from "lucide-react";
import Link from "next/link";
import { usePathname } from "next/navigation";
import { clsx, type ClassValue } from "clsx";
import { twMerge } from "tailwind-merge";

import { useUIStore } from "@/store/useUIStore";

function cn(...inputs: ClassValue[]) {
  return twMerge(clsx(inputs));
}

const navItems = [
  { href: "/", icon: Home, label: "Home" },
  { href: "/explore", icon: Search, label: "Explore" },
  { href: "/profile", icon: User, label: "Profile" },
  { href: "/settings", icon: Settings, label: "Settings" },
];

export default function MobileNav() {
  const pathname = usePathname();
  const { openNewPostModal } = useUIStore();

  return (
    <div className="md:hidden fixed bottom-0 left-0 right-0 h-16 bg-background/80 backdrop-blur-lg border-t border-foreground/10 flex items-center justify-around px-2 z-50">
      {navItems.slice(0, 2).map((item) => {
        const isActive = pathname === item.href;
        return (
          <Link
            key={item.href}
            href={item.href}
            className={cn(
              "flex flex-col items-center justify-center size-12 transition-colors",
              isActive ? "text-foreground" : "text-foreground/40"
            )}
          >
            <item.icon size={22} strokeWidth={isActive ? 2.5 : 2} />
          </Link>
        );
      })}

      {/* Floating Action Button for New Post */}
      <button 
        onClick={openNewPostModal}
        className="size-12 bg-foreground text-background rounded-full flex items-center justify-center -mt-8 shadow-lg active:scale-95 transition-transform"
      >
        <Plus size={24} strokeWidth={3} />
      </button>

      {navItems.slice(2).map((item) => {
        const isActive = pathname === item.href;
        return (
          <Link
            key={item.href}
            href={item.href}
            className={cn(
              "flex flex-col items-center justify-center size-12 transition-colors",
              isActive ? "text-foreground" : "text-foreground/40"
            )}
          >
            <item.icon size={22} strokeWidth={isActive ? 2.5 : 2} />
          </Link>
        );
      })}
    </div>
  );
}
