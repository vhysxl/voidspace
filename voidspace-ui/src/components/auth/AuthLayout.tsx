"use client";

import { ReactNode } from "react";
import Link from "next/link";
import { ChevronLeft } from "lucide-react";
import { motion } from "framer-motion";

interface AuthLayoutProps {
  children: ReactNode;
}

export default function AuthLayout({ children }: AuthLayoutProps) {
  return (
    <div className="min-h-screen bg-black flex items-center justify-center p-6 relative overflow-hidden font-inter selection:bg-white/20 selection:text-white">
      {/* Background Glow */}
      <div className="absolute inset-0 pointer-events-none">
        <div className="absolute inset-0 opacity-[0.03] bg-[radial-gradient(circle_at_center,_rgba(255,255,255,0.1)_0%,_rgba(255,255,255,0)_70%)]"></div>
      </div>

      {/* Back Button */}
      <Link 
        href="/"
        className="absolute left-4 top-4 md:left-8 md:top-8 size-10 border border-white/20 flex items-center justify-center text-white/60 hover:text-white hover:border-white/40 transition-all rounded-sm z-50 group bg-black/50 backdrop-blur-sm"
      >
        <ChevronLeft size={16} className="group-hover:-translate-x-0.5 transition-transform" />
      </Link>

      <motion.div 
        initial={{ opacity: 0, y: 20 }}
        animate={{ opacity: 1, y: 0 }}
        transition={{ duration: 0.8, ease: [0.16, 1, 0.3, 1] }}
        className="w-full max-w-[440px] relative z-10 py-12 md:py-0"
      >
        {children}
      </motion.div>
    </div>
  );
}
