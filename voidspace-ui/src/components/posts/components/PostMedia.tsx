"use client";

import { useState } from "react";
import { motion, AnimatePresence } from "framer-motion";
import { ChevronLeft, ChevronRight } from "lucide-react";

interface PostMediaProps {
  images: { image_url: string }[];
  isDetailed: boolean;
  onImageClick: (url: string) => void;
}

function DetailCarousel({ images, onImageClick }: { images: { image_url: string }[], onImageClick: (url: string) => void }) {
  const [currentIndex, setCurrentIndex] = useState(0);
  const [direction, setDirection] = useState(0);

  const slideVariants = {
    enter: (direction: number) => ({
      x: direction > 0 ? 300 : -300,
      opacity: 0
    }),
    center: {
      zIndex: 1,
      x: 0,
      opacity: 1
    },
    exit: (direction: number) => ({
      zIndex: 0,
      x: direction < 0 ? 300 : -300,
      opacity: 0
    })
  };

  const paginate = (newDirection: number) => {
    setDirection(newDirection);
    setCurrentIndex((prevIndex) => (prevIndex + newDirection + images.length) % images.length);
  };

  return (
    <div className="relative mt-4 rounded-sm overflow-hidden border border-foreground/10 bg-black/20 aspect-video group/carousel shadow-inner">
      <AnimatePresence initial={false} custom={direction} mode="popLayout">
        <motion.div
          key={currentIndex}
          custom={direction}
          variants={slideVariants}
          initial="enter"
          animate="center"
          exit="exit"
          transition={{
            x: { type: "spring", stiffness: 300, damping: 30 },
            opacity: { duration: 0.2 }
          }}
          className="w-full h-full flex items-center justify-center cursor-zoom-in"
          onClick={() => onImageClick(images[currentIndex].image_url)}
        >
          <img
            src={images[currentIndex].image_url}
            alt={`Transmission data ${currentIndex + 1}`}
            className="w-full h-full object-contain"
          />
        </motion.div>
      </AnimatePresence>

      <div className="absolute top-4 right-4 bg-black/50 backdrop-blur-md px-2.5 py-1 rounded-sm text-[10px] font-bold text-white uppercase tracking-[2px] z-20 border border-white/10">
        {currentIndex + 1} / {images.length}
      </div>

      <button
        onClick={(e) => { e.stopPropagation(); paginate(-1); }}
        className="absolute left-4 top-1/2 -translate-y-1/2 z-30 size-10 rounded-full bg-black/40 backdrop-blur-md text-white flex items-center justify-center opacity-0 group-hover/carousel:opacity-100 transition-all hover:bg-black/60 active:scale-95 border border-white/5"
      >
        <ChevronLeft size={20} />
      </button>
      <button
        onClick={(e) => { e.stopPropagation(); paginate(1); }}
        className="absolute right-4 top-1/2 -translate-y-1/2 z-30 size-10 rounded-full bg-black/40 backdrop-blur-md text-white flex items-center justify-center opacity-0 group-hover/carousel:opacity-100 transition-all hover:bg-black/60 active:scale-95 border border-white/5"
      >
        <ChevronRight size={20} />
      </button>

      <div className="absolute bottom-4 left-0 right-0 z-30 flex justify-center gap-1.5 px-4">
        {images.map((_, idx) => (
          <button
            key={idx}
            onClick={(e) => {
              e.stopPropagation();
              setDirection(idx > currentIndex ? 1 : -1);
              setCurrentIndex(idx);
            }}
            className={`h-1 rounded-full transition-all duration-300 ${idx === currentIndex ? "bg-white w-8 shadow-[0_0_10px_rgba(255,255,255,0.5)]" : "bg-white/20 w-4 hover:bg-white/40"}`}
          />
        ))}
      </div>
    </div>
  );
}

export default function PostMedia({ images, isDetailed, onImageClick }: PostMediaProps) {
  const validImages = images?.filter(img => img.image_url) || [];
  if (validImages.length === 0) return null;

  if (isDetailed && validImages.length > 1) {
    return <DetailCarousel images={validImages} onImageClick={onImageClick} />;
  }

  return (
    <div className={`mt-4 rounded-sm overflow-hidden border max-w-[800px] border-foreground/10 grid gap-1 ${validImages.length === 1 ? "grid-cols-1" : "grid-cols-2"}`}>
      {(isDetailed ? validImages : validImages.slice(0, 2)).map((img, idx) => {
        const isLastVisible = !isDetailed && idx === 1 && validImages.length > 2;
        return (
          <div
            key={idx}
            className={`relative overflow-hidden group/img cursor-zoom-in ${validImages.length === 3 && idx === 0 && isDetailed ? "row-span-2" : ""} ${validImages.length === 1
              ? (isDetailed ? "w-full min-h-[200px] max-h-[70vh] flex items-center justify-center bg-black/20" : "aspect-video")
              : "aspect-square"}`}
            onClick={(e) => {
              e.stopPropagation();
              onImageClick(img.image_url);
            }}
          >
            <img
              src={img.image_url}
              alt={`Post content ${idx + 1}`}
              className={`w-full h-full transition-transform duration-500 group-hover/img:scale-[1.03] ${validImages.length === 1 && isDetailed ? "object-contain max-h-[80vh]" : "object-cover"}`}
            />
            {isLastVisible && (
              <div className="absolute inset-0 bg-black/60 flex items-center justify-center backdrop-blur-[2px] transition-colors group-hover/img:bg-black/50">
                <span className="text-white font-space-grotesk text-2xl font-bold tracking-widest">+{validImages.length - 1}</span>
              </div>
            )}
          </div>
        );
      })}
    </div>
  );
}
