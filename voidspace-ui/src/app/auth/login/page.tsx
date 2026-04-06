"use client";

import { useEffect } from "react";
import { useRouter } from "next/navigation";
import { useAuthStore } from "@/store/useAuthStore";
import AuthLayout from "@/components/auth/AuthLayout";
import LoginForm from "@/components/auth/LoginForm";

export default function LoginPage() {
  const router = useRouter();
  const { isLoggedIn, _hasHydrated } = useAuthStore();

  useEffect(() => {
    if (_hasHydrated && isLoggedIn) {
      router.replace("/");
    }
  }, [isLoggedIn, _hasHydrated, router]);

  if (!_hasHydrated) return null;
  if (isLoggedIn) return null;

  return (
    <AuthLayout>
      <LoginForm />
    </AuthLayout>
  );
}
