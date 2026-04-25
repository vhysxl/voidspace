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
    primary: "bg-foreground text-background hover:opacity-90 font-bold text-[14px] tracking-[2.4px] uppercase h-[56px] w-full disabled:opacity-50",
    secondary: "border border-foreground/20 bg-transparent text-foreground/60 hover:text-foreground hover:border-foreground/40 disabled:opacity-50",
    ghost: "text-foreground/40 hover:text-foreground transition-colors disabled:opacity-50",
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
