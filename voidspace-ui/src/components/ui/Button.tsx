import React from "react";
import { clsx, type ClassValue } from "clsx";
import { twMerge } from "tailwind-merge";

function cn(...inputs: ClassValue[]) {
  return twMerge(clsx(inputs));
}

interface ButtonProps extends React.ButtonHTMLAttributes<HTMLButtonElement> {
  variant?: "primary" | "secondary" | "ghost";
  isLoading?: boolean;
}

export default function Button({
  className,
  variant = "primary",
  isLoading,
  children,
  ...props
}: ButtonProps) {
  const variants = {
    primary: "bg-white text-black hover:bg-neutral-200 font-bold text-[14px] tracking-[2.4px] uppercase h-[56px] w-full disabled:opacity-50",
    secondary: "border border-white/20 bg-transparent text-white/60 hover:text-white hover:border-white/40 disabled:opacity-50",
    ghost: "text-[#666] hover:text-white transition-colors disabled:opacity-50",
  };

  return (
    <button
      className={cn(
        "flex items-center justify-center gap-2 transition-all active:scale-[0.99]",
        variants[variant],
        className
      )}
      disabled={isLoading || props.disabled}
      {...props}
    >
      {isLoading ? "Processing..." : children}
    </button>
  );
}
