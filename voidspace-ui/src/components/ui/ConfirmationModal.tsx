"use client";

import { motion, AnimatePresence } from "framer-motion";
import { X, AlertTriangle, Loader2 } from "lucide-react";

interface ConfirmationModalProps {
  isOpen: boolean;
  onClose: () => void;
  onConfirm: () => void;
  title: string;
  description: string;
  confirmText?: string;
  cancelText?: string;
  isLoading?: boolean;
  variant?: "danger" | "warning" | "info";
}

export default function ConfirmationModal({
  isOpen,
  onClose,
  onConfirm,
  title,
  description,
  confirmText = "Confirm",
  cancelText = "Cancel",
  isLoading = false,
  variant = "danger"
}: ConfirmationModalProps) {
  
  const variantStyles = {
    danger: "bg-red-500/10 text-red-500",
    warning: "bg-yellow-500/10 text-yellow-500",
    info: "bg-blue-500/10 text-blue-500",
  };

  const confirmButtonStyles = {
    danger: "bg-red-500 hover:bg-red-600 text-white",
    warning: "bg-yellow-500 hover:bg-yellow-600 text-black",
    info: "bg-foreground hover:opacity-90 text-background",
  };

  return (
    <AnimatePresence>
      {isOpen && (
        <div className="fixed inset-0 z-[110] flex items-center justify-center p-4 bg-background/80 backdrop-blur-sm">
          <motion.div
            initial={{ opacity: 0, scale: 0.95, y: 10 }}
            animate={{ opacity: 1, scale: 1, y: 0 }}
            exit={{ opacity: 0, scale: 0.95, y: 10 }}
            className="w-full max-w-md bg-background border border-foreground/20 rounded-sm shadow-2xl p-8 space-y-6"
          >
            <div className="flex justify-between items-start">
              <div className={`size-12 rounded-full flex items-center justify-center ${variantStyles[variant]}`}>
                <AlertTriangle size={24} />
              </div>
              <button
                onClick={onClose}
                className="text-foreground/40 hover:text-foreground transition-colors"
              >
                <X size={20} />
              </button>
            </div>

            <div className="space-y-2">
              <h3 className={`font-space-grotesk text-xl font-bold uppercase tracking-tight ${variant === 'danger' ? 'text-red-500' : 'text-foreground'}`}>
                {title}
              </h3>
              <p className="text-sm text-foreground/60 leading-relaxed">
                {description}
              </p>
            </div>

            <div className="flex flex-col gap-3 pt-2">
              <button
                disabled={isLoading}
                onClick={onConfirm}
                className={`w-full py-4 rounded-sm font-bold text-[11px] uppercase tracking-[2px] transition-all flex items-center justify-center gap-2 ${confirmButtonStyles[variant]} disabled:opacity-50 disabled:cursor-not-allowed`}
              >
                {isLoading ? (
                  <>
                    <Loader2 size={14} className="animate-spin" />
                    Processing...
                  </>
                ) : confirmText}
              </button>
              <button
                disabled={isLoading}
                onClick={onClose}
                className="w-full py-4 rounded-sm font-bold text-[11px] uppercase tracking-[2px] border border-foreground/10 hover:bg-foreground/5 transition-all text-foreground/60 hover:text-foreground disabled:opacity-50"
              >
                {cancelText}
              </button>
            </div>
          </motion.div>
        </div>
      )}
    </AnimatePresence>
  );
}
