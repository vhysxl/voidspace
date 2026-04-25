"use client";

import Link from "next/link";
import { usePathname } from "next/navigation";
import { clsx, type ClassValue } from "clsx";
import { twMerge } from "tailwind-merge";
import { LucideIcon } from "lucide-react";

function cn(...inputs: ClassValue[]) {
  return twMerge(clsx(inputs));
}

interface SidebarItemProps {
  href: string;
  icon: LucideIcon;
  label: string;
  collapsed?: boolean;
}

export default function SidebarItem({
  href,
  icon: Icon,
  label,
  collapsed = false,
}: SidebarItemProps) {
  const pathname = usePathname();
  const isActive = pathname === href;

  return (
    <Link
      href={href}
      className={cn(
        "flex items-center gap-4 px-4 py-3 transition-all duration-300 group relative",
        isActive
          ? "text-foreground bg-foreground/5"
          : "text-foreground/40 hover:text-foreground hover:bg-foreground/5"
      )}
    >
      {/* Active Indicator */}
      {isActive && (
        <div className="absolute left-0 top-0 bottom-0 w-[2px] bg-foreground shadow-[0_0_8px_rgba(var(--foreground),0.5)]" />
      )}

      <Icon
        size={20}
        className={cn(
          "transition-transform duration-300 group-hover:scale-110",
          isActive && "text-foreground"
        )}
      />

      {!collapsed && (
        <span className="font-manrope text-[13px] tracking-[2px] uppercase font-medium">
          {label}
        </span>
      )}
    </Link>
  );
}
