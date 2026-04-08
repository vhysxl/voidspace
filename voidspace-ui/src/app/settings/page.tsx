"use client";

import { useState } from "react";
import DashboardLayout from "@/components/layout/DashboardLayout";
import { useAuthStore } from "@/store/useAuthStore";
import { useAuth } from "@/hooks/useAuth";
import { Trash2, AlertTriangle, X, LogOut, Loader2, LogIn } from "lucide-react";
import { motion, AnimatePresence } from "framer-motion";
import Link from "next/link";

export default function SettingsPage() {
  const { user, isLoggedIn, _hasHydrated } = useAuthStore();
  const { logout } = useAuth();
  const [showDeleteModal, setShowDeleteModal] = useState(false);
  const [confirmUsername, setConfirmUsername] = useState("");

  const isConfirmed = confirmUsername === user?.username;

  const handleDeleteAccount = () => {
    if (isConfirmed) {
      alert("Account deletion process started. (Simulated)");
      // Add actual deletion logic here
      setShowDeleteModal(false);
    }
  };

  if (!_hasHydrated) {
    return (
      <DashboardLayout fullWidth={true}>
        <div className="flex items-center justify-center min-h-[60vh]">
          <Loader2 className="size-8 text-foreground/20 animate-spin" />
        </div>
      </DashboardLayout>
    );
  }

  return (
    <DashboardLayout fullWidth={true}>
      <div className="flex flex-col min-h-screen">
        {/* Header */}
        <div className="sticky top-0 z-30 bg-background/80 backdrop-blur-md border-b border-foreground/10 px-6 py-4">
          <h1 className="font-space-grotesk text-xl font-bold tracking-tight uppercase">
            Settings
          </h1>
        </div>

        <div className="p-6 space-y-10">
          {!isLoggedIn ? (
            <div className="flex flex-col items-center justify-center py-20 text-center space-y-6">
              <div className="size-16 rounded-full bg-foreground/5 flex items-center justify-center">
                <LogIn size={24} className="text-foreground/20" />
              </div>
              <div className="space-y-2">
                <h2 className="font-space-grotesk text-xl font-bold uppercase tracking-tight">
                  Authentication Required
                </h2>
                <p className="text-foreground/40 text-sm max-w-xs mx-auto uppercase tracking-widest leading-loose">
                  You need to sign in to access and manage your settings.
                </p>
              </div>
              <Link 
                href="/auth/login"
                className="px-8 py-3 bg-foreground text-background font-bold uppercase text-[11px] tracking-[2px] hover:opacity-90 transition-all"
              >
                Sign In
              </Link>
            </div>
          ) : (
            <>
              {/* Account Section */}
              <section className="space-y-4">
                <h2 className="text-[11px] font-bold uppercase tracking-[2px] text-foreground/40">
                  Account
                </h2>
                <div className="border border-foreground/10 rounded-sm divide-y divide-foreground/10">
                  <div className="p-4 flex flex-col gap-1">
                    <p className="text-[11px] text-foreground/40 uppercase tracking-[1px]">
                      Logged in as
                    </p>
                    <p className="text-sm font-bold uppercase tracking-tight">
                      {user?.profile.display_name || user?.username || "VOYAGER"}
                    </p>
                    <p className="text-[11px] text-foreground/40 tracking-[1px]">
                      @{user?.username}
                    </p>
                  </div>

                  <button
                    onClick={() => logout()}
                    className="w-full p-4 flex items-center gap-3 text-foreground/60 hover:bg-foreground/5 hover:text-foreground transition-colors text-left cursor-pointer"
                  >
                    <LogOut size={20} />
                    <div>
                      <p className="text-sm font-bold uppercase tracking-tight">
                        Log Out
                      </p>
                      <p className="text-[11px] opacity-60 uppercase tracking-[1px]">
                        End your current session
                      </p>
                    </div>
                  </button>

                  <button
                    onClick={() => setShowDeleteModal(true)}
                    className="w-full p-4 flex items-center gap-3 text-red-500 hover:bg-red-500/5 transition-colors text-left cursor-pointer"
                  >
                    <Trash2 size={20} />
                    <div>
                      <p className="text-sm font-bold uppercase tracking-tight">
                        Delete Account
                      </p>
                      <p className="text-[11px] opacity-60 uppercase tracking-[1px]">
                        Permanently erase your data from the void
                      </p>
                    </div>
                  </button>
                </div>
              </section>
            </>
          )}
        </div>
      </div>

      {/* Delete Confirmation Modal */}
      <AnimatePresence>
        {showDeleteModal && isLoggedIn && (
          <div className="fixed inset-0 z-50 flex items-center justify-center p-4 bg-background/80 backdrop-blur-sm">
            <motion.div
              initial={{ opacity: 0, scale: 0.95 }}
              animate={{ opacity: 1, scale: 1 }}
              exit={{ opacity: 0, scale: 0.95 }}
              className="w-full max-w-md bg-background border border-foreground/20 rounded-sm p-8 space-y-6 shadow-2xl"
            >
              <div className="flex justify-between items-start">
                <div className="size-12 rounded-full bg-red-500/10 flex items-center justify-center text-red-500">
                  <AlertTriangle size={24} />
                </div>
                <button
                  onClick={() => setShowDeleteModal(false)}
                  className="text-foreground/40 hover:text-foreground transition-colors cursor-pointer"
                >
                  <X size={20} />
                </button>
              </div>

              <div className="space-y-2">
                <h3 className="font-space-grotesk text-xl font-bold uppercase tracking-tight text-red-500">
                  Are you absolutely sure?
                </h3>
                <p className="text-sm text-foreground/60 leading-relaxed uppercase tracking-widest">
                  This action is irreversible. It will permanently delete your posts, comments, and profile from VOIDSPACE.
                </p>
              </div>

              <div className="space-y-4">
                <div className="p-4 bg-foreground/5 border border-foreground/10 rounded-sm">
                  <p className="text-xs text-foreground/40 uppercase tracking-[1px] mb-2">
                    To confirm, please type your username:
                    <span className="block mt-1 font-bold text-foreground">
                      {user?.username}
                    </span>
                  </p>
                  <input
                    type="text"
                    value={confirmUsername}
                    onChange={(e) => setConfirmUsername(e.target.value)}
                    placeholder="ENTER USERNAME"
                    className="w-full bg-transparent border-b border-foreground/20 py-2 text-sm tracking-[2px] uppercase outline-none focus:border-red-500 transition-all placeholder:text-foreground/20"
                  />
                </div>

                <div className="flex flex-col gap-3">
                  <button
                    disabled={!isConfirmed}
                    onClick={handleDeleteAccount}
                    className={`w-full py-4 rounded-sm font-bold text-[11px] uppercase tracking-[2px] transition-all ${
                      isConfirmed
                        ? "bg-red-500 text-white hover:bg-red-600 active:scale-[0.98] cursor-pointer"
                        : "bg-foreground/5 text-foreground/20 cursor-not-allowed"
                    }`}
                  >
                    Permanently Delete
                  </button>
                  <button
                    onClick={() => setShowDeleteModal(false)}
                    className="w-full py-4 rounded-sm font-bold text-[11px] uppercase tracking-[2px] border border-foreground/10 hover:bg-foreground/5 transition-all cursor-pointer"
                  >
                    Cancel
                  </button>
                </div>
              </div>
            </motion.div>
          </div>
        )}
      </AnimatePresence>
    </DashboardLayout>
  );
}
