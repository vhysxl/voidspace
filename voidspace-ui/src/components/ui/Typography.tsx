import React from "react";
import { clsx, type ClassValue } from "clsx";
import { twMerge } from "tailwind-merge";

function cn(...inputs: ClassValue[]) {
  return twMerge(clsx(inputs));
}

interface TypographyProps {
  children: React.ReactNode;
  className?: string;
  as?: React.ElementType;
}

export function Heading({ children, className, as: Component = "h1" }: TypographyProps) {
  return (
    <Component
      className={cn(
        "font-space-grotesk font-bold text-[36px] text-white tracking-[-1.08px] uppercase",
        className
      )}
    >
      {children}
    </Component>
  );
}

export function Subtext({ children, className, as: Component = "p" }: TypographyProps) {
  return (
    <Component
      className={cn(
        "font-manrope text-[14px] text-[#666] tracking-[0.7px] uppercase",
        className
      )}
    >
      {children}
    </Component>
  );
}

export function Label({ children, className, as: Component = "span" }: TypographyProps) {
  return (
    <Component
      className={cn(
        "font-manrope text-[12px] text-[#666] tracking-[2.4px] uppercase",
        className
      )}
    >
      {children}
    </Component>
  );
}
