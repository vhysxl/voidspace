"use client";

import React, { useState } from "react";
import Link from "next/link";
import { useAuth } from "@/hooks/useAuth";
import Input from "@/components/ui/Input";
import Button from "@/components/ui/Button";
import AuthHeader from "./AuthHeader";
import AuthFooter from "./AuthFooter";
import { MoveRight } from "lucide-react";

export default function LoginForm() {
  const { login } = useAuth();
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);
  const [formData, setFormData] = useState({
    usernameoremail: "",
    password: "",
  });

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setIsLoading(true);
    setError(null);

    try {
      await login(formData.usernameoremail, formData.password);
    } catch (err: any) {
      setError(err.message || "Login failed. Please check your credentials.");
    } finally {
      setIsLoading(false);
    }
  };

  return (
    <div className="w-full">
      <AuthHeader 
        title="Welcome Back" 
        subtitle="Log in to your hour account" 
      />

      <form onSubmit={handleSubmit} className="space-y-8">
        <div className="space-y-8">
          <Input
            label="Email Address"
            placeholder="Email"
            required
            value={formData.usernameoremail}
            onChange={(e) => setFormData({ ...formData, usernameoremail: e.target.value })}
          />

          <Input
            label="Password"
            placeholder="Password"
            type="password"
            required
            action={
              <Link href="/auth/forgot" className="text-[10px] text-[#666] tracking-[1px] uppercase hover:text-white transition-colors">
                Forgot Password?
              </Link>
            }
            value={formData.password}
            onChange={(e) => setFormData({ ...formData, password: e.target.value })}
          />
        </div>

        {error && (
          <p className="text-red-500/80 text-[12px] tracking-[0.6px] uppercase text-center font-manrope">
            {error}
          </p>
        )}

        <div className="space-y-8 pt-2">
          <Button type="submit" isLoading={isLoading}>
            Sign In {!isLoading && <MoveRight size={16} />}
          </Button>

          <div className="text-center text-[12px] tracking-[0.6px] uppercase font-manrope">
            <span className="text-[#666]">New to Voidspace? </span>
            <Link href="/auth/register" className="text-white font-bold hover:underline">
              Create account
            </Link>
          </div>
        </div>
      </form>

      <AuthFooter />
    </div>
  );
}
