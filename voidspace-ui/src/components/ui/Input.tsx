import { ReactNode, InputHTMLAttributes } from "react";
import { clsx, type ClassValue } from "clsx";
import { twMerge } from "tailwind-merge";

function cn(...inputs: ClassValue[]) {
  return twMerge(clsx(inputs));
}

interface InputProps extends InputHTMLAttributes<HTMLInputElement> {
  label?: string;
  error?: string;
  action?: ReactNode;
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
            <label className="text-[12px] text-foreground/40 tracking-[2.4px] uppercase block font-manrope">
              {label}
            </label>
          )}
          {action && <div className="font-manrope">{action}</div>}
        </div>
      )}
      <div className="relative">
        <input
          className={cn(
            "w-full bg-foreground/5 border border-foreground/10 px-4 py-3.5 text-foreground text-[16px] tracking-[0.32px] focus:outline-none focus:border-foreground/30 transition-all placeholder:text-foreground/10 font-manrope",
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
