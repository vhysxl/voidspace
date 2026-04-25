"use client";

import {
  Home,
  Search,
  User,
  Settings,
  Plus
} from "lucide-react";
import SidebarItem from "./SidebarItem";
import Link from "next/link";
import UserMiniProfile from "./UserMiniProfile";

import { useUIStore } from "@/store/useUIStore";

const navItems = [
  { href: "/", icon: Home, label: "Home" },
  { href: "/explore", icon: Search, label: "Explore" },
  { href: "/profile", icon: User, label: "Profile" },
  { href: "/settings", icon: Settings, label: "Settings" },
];

export default function Sidebar() {
  const { openNewPostModal } = useUIStore();

  return (
    <aside className="hidden md:flex flex-col w-[260px] h-screen fixed left-0 top-0 border-r border-foreground/10 bg-background z-40">
      {/* Logo */}
      <div className="p-8">
        <Link href="/" className="group">
          <h1 className="font-space-grotesk text-2xl font-bold tracking-[-1px] text-foreground transition-all group-hover:tracking-[1px]">
            VOIDSPACE
          </h1>
        </Link>
      </div>

      {/* Navigation */}
      <nav className="flex-1 mt-4">
        {navItems.map((item) => (
          <SidebarItem
            key={item.href}
            href={item.href}
            icon={item.icon}
            label={item.label}
          />
        ))}
      </nav>

      {/* Post Button */}
      <div className="p-6 mt-auto">
        <button 
          onClick={openNewPostModal}
          className="w-full bg-foreground text-background h-[52px] rounded-sm font-bold text-[13px] tracking-[2px] uppercase flex items-center justify-center gap-2 hover:opacity-90 transition-all active:scale-[0.98]"
        >
          <Plus size={18} />
          <span>New Post</span>
        </button>
      </div>

      <UserMiniProfile />
    </aside>
  );
}
