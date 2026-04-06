"use client";

import { useState } from "react";
import DashboardLayout from "@/components/layout/DashboardLayout";
import { useThemeStore } from "@/store/useThemeStore";
import { useAuthStore } from "@/store/useAuthStore";
import { Sun, Moon, Trash2, AlertTriangle, X } from "lucide-react";
import { motion, AnimatePresence } from "framer-motion";

export default function SettingsPage() {
  const { theme, toggleTheme, _hasHydrated } = useThemeStore();
  const { user } = useAuthStore();
  const [showDeleteModal, setShowDeleteModal] = useState(false);
  const [confirmUsername, setConfirmUsername] = useState("");

  const isConfirmed = confirmUsername === user?.username;

  if (!_hasHydrated) {
    return (
      <DashboardLayout fullWidth={true}>
        <div className="flex flex-col min-h-screen animate-pulse p-6">
          <div className="h-8 bg-foreground/5 w-32 mb-8" />
          <div className="h-20 bg-foreground/5 w-full rounded-sm" />
        </div>
      </DashboardLayout>
    );
  }

  const handleDeleteAccount = () => {
    if (isConfirmed) {
      alert("Account deletion process started. (Simulated)");
      // Add actual deletion logic here
      setShowDeleteModal(false);
    }
  };

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
          {/* Theme Section */}
          <section className="space-y-4">
            <h2 className="text-[11px] font-bold uppercase tracking-[2px] text-foreground/40">
              Appearance
            </h2>
            <div className="flex items-center justify-between p-4 border border-foreground/10 rounded-sm bg-foreground/5">
              <div className="flex items-center gap-3">
                {theme === "dark" ? <Moon size={20} /> : <Sun size={20} />}
                <div>
                  <p className="text-sm font-bold uppercase tracking-tight">
                    {theme === "dark" ? "Dark Mode" : "Light Mode"}
                  </p>
                  <p className="text-[11px] text-foreground/40 uppercase tracking-[1px]">
                    Switch between celestial modes
                  </p>
                </div>
              </div>
              <button
                onClick={toggleTheme}
                className="relative inline-flex h-6 w-11 items-center rounded-full bg-foreground/10 transition-colors focus:outline-none"
              >
                <span
                  className={`${
                    theme === "dark" ? "translate-x-6 bg-foreground" : "translate-x-1 bg-foreground/40"
                  } inline-block h-4 w-4 transform rounded-full transition-transform`}
                />
              </button>
            </div>
          </section>

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
                  {user?.name || "GUEST VOYAGER"}
                </p>
                <p className="text-[11px] text-foreground/40 tracking-[1px]">
                  @{user?.username || "guest"}
                </p>
              </div>

              <button
                onClick={() => setShowDeleteModal(true)}
                className="w-full p-4 flex items-center gap-3 text-red-500 hover:bg-red-500/5 transition-colors text-left"
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
        </div>
      </div>

      {/* Delete Confirmation Modal */}
      <AnimatePresence>
        {showDeleteModal && (
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
                  className="text-foreground/40 hover:text-foreground transition-colors"
                >
                  <X size={20} />
                </button>
              </div>

              <div className="space-y-2">
                <h3 className="font-space-grotesk text-xl font-bold uppercase tracking-tight text-red-500">
                  Are you absolutely sure?
                </h3>
                <p className="text-sm text-foreground/60 leading-relaxed">
                  This action is irreversible. It will permanently delete your posts, comments, and profile from VOIDSPACE.
                </p>
              </div>

              <div className="space-y-4">
                <div className="p-4 bg-foreground/5 border border-foreground/10 rounded-sm">
                  <p className="text-xs text-foreground/40 uppercase tracking-[1px] mb-2">
                    To confirm, please type your username:
                    <span className="block mt-1 font-bold text-foreground">
                      {user?.username || "guest"}
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
                        ? "bg-red-500 text-white hover:bg-red-600 active:scale-[0.98]"
                        : "bg-foreground/5 text-foreground/20 cursor-not-allowed"
                    }`}
                  >
                    Permanently Delete
                  </button>
                  <button
                    onClick={() => setShowDeleteModal(false)}
                    className="w-full py-4 rounded-sm font-bold text-[11px] uppercase tracking-[2px] border border-foreground/10 hover:bg-foreground/5 transition-all"
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
