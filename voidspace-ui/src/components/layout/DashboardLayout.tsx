"use client";

import { ReactNode } from "react";
import Sidebar from "./Sidebar";
import RightSidebar from "./RightSidebar";
import MobileNav from "./MobileNav";
import NewPostModal from "./NewPostModal";
import EditPostModal from "../posts/EditPostModal";
import EditProfileModal from "../profile/EditProfileModal";

interface DashboardLayoutProps {
  children: ReactNode;
  fullWidth?: boolean;
}

export default function DashboardLayout({ children, fullWidth = true }: DashboardLayoutProps) {
  return (
    <div className="min-h-screen bg-background text-foreground font-manrope">
      {/* Left Sidebar */}
      <Sidebar />

      {/* Main Content Area */}
      <main className="md:pl-[260px] lg:pr-[350px] min-h-screen pb-20 md:pb-0">
        <div className={fullWidth ? "w-full min-h-screen" : "max-w-[600px] mx-auto min-h-screen border-x border-foreground/10"}>
          {children}
        </div>
      </main>

      {/* Right Sidebar */}
      <RightSidebar />

      {/* Mobile Navigation */}
      <MobileNav />

      {/* Modals */}
      <NewPostModal />
      <EditPostModal />
      <EditProfileModal />
    </div>
  );
}
