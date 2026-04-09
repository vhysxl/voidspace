"use client";

import { motion, AnimatePresence } from "framer-motion";
import { X, ZoomIn } from "lucide-react";

interface ImageModalProps {
  isOpen: boolean;
  onClose: () => void;
  src?: string;
  alt?: string;
}

export default function ImageModal({ isOpen, onClose, src, alt }: ImageModalProps) {
  return (
    <AnimatePresence>
      {isOpen && (
        <div 
          className="fixed inset-0 z-[200] flex items-center justify-center p-4 md:p-10 bg-black/95 backdrop-blur-md cursor-zoom-out"
          onClick={onClose}
        >
          <motion.div
            initial={{ opacity: 0, scale: 0.9 }}
            animate={{ opacity: 1, scale: 1 }}
            exit={{ opacity: 0, scale: 0.9 }}
            transition={{ type: "spring", damping: 25, stiffness: 300 }}
            className="relative max-w-full max-h-full flex items-center justify-center"
            onClick={(e) => e.stopPropagation()}
          >
            <button
              onClick={onClose}
              className="absolute -top-12 right-0 md:-right-12 text-white/60 hover:text-white transition-colors p-2"
            >
              <X size={32} />
            </button>
            
            {src && (
              <img 
                src={src} 
                alt={alt || "Full resolution"} 
                className="max-w-full max-h-[85vh] object-contain shadow-2xl rounded-sm border border-white/10"
              />
            )}
            
            {alt && (
              <div className="absolute -bottom-12 left-0 right-0 text-center">
                <p className="text-white/40 text-[10px] uppercase tracking-[2px] font-bold">
                  {alt}
                </p>
              </div>
            )}
          </motion.div>
        </div>
      )}
    </AnimatePresence>
  );
}
