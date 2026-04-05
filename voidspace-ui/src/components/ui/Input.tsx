import React from "react";
import { clsx, type ClassValue } from "clsx";
import { twMerge } from "tailwind-merge";

function cn(...inputs: ClassValue[]) {
  return twMerge(clsx(inputs));
}

interface InputProps extends React.InputHTMLAttributes<HTMLInputElement> {
  label?: string;
  error?: string;
  action?: React.ReactNode;
}

export default function Input({
  className,
  label,
  error,
  action,
  ...props
}: InputProps) {
  return (
    <div className="space-y-3 w-full">
      {(label || action) && (
        <div className="flex justify-between items-baseline px-1">
          {label && (
            <label className="text-[12px] text-[#666] tracking-[2.4px] uppercase block font-manrope">
              {label}
            </label>
          )}
          {action && <div className="font-manrope">{action}</div>}
        </div>
      )}
      <div className="relative">
        <input
          className={cn(
            "w-full bg-[#121212] border border-white/10 px-4 py-3.5 text-white text-[16px] tracking-[0.32px] focus:outline-none focus:border-white/30 transition-all placeholder:text-white/10 font-manrope",
            error && "border-red-500/50 focus:border-red-500",
            className
          )}
          {...props}
        />
      </div>
      {error && (
        <p className="text-red-500 text-[10px] tracking-[1px] uppercase px-1 font-manrope">
          {error}
        </p>
      )}
    </div>
  );
}
