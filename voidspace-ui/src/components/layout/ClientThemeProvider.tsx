"use client";

import { useThemeStore } from "@/store/useThemeStore";
import { useEffect, ReactNode, useState } from "react";

export default function ClientThemeProvider({ children }: { children: ReactNode }) {
  const { theme, _hasHydrated } = useThemeStore();
  const [mounted, setMounted] = useState(false);

  useEffect(() => {
    setMounted(true);
  }, []);

  useEffect(() => {
    if (_hasHydrated && mounted) {
      const root = window.document.documentElement;
      root.classList.remove("light", "dark");
      root.classList.add(theme);
      
      // Force update metadata or other things if needed
      if (theme === "dark") {
        root.style.colorScheme = "dark";
      } else {
        root.style.colorScheme = "light";
      }
    }
  }, [theme, _hasHydrated, mounted]);

  // Avoid hydration mismatch by not rendering anything theme-dependent on the server
  // Since we're using suppressHydrationWarning on html tag, it should be fine.
  
  return <>{children}</>;
}
