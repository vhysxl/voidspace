"use client";

import React, { useState } from "react";
import Link from "next/link";
import { useAuth } from "@/hooks/useAuth";
import Input from "@/components/ui/Input";
import Button from "@/components/ui/Button";
import AuthHeader from "./AuthHeader";
import AuthFooter from "./AuthFooter";
import { MoveRight } from "lucide-react";

export default function RegisterForm() {
  const { register } = useAuth();
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);
  const [formData, setFormData] = useState({
    username: "",
    email: "",
    password: "",
  });

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setIsLoading(true);
    setError(null);

    try {
      await register(formData.username, formData.email, formData.password);
    } catch (err: any) {
      setError(err.message || "Registration failed. Please try again.");
    } finally {
      setIsLoading(false);
    }
  };

  return (
    <div className="w-full">
      <AuthHeader 
        title="Join Voidspace" 
        subtitle="Create your hour account to get started" 
      />

      <form onSubmit={handleSubmit} className="space-y-8">
        <div className="space-y-8">
          <Input
            label="Username"
            placeholder="Username"
            required
            value={formData.username}
            onChange={(e) => setFormData({ ...formData, username: e.target.value })}
          />

          <Input
            label="Email Address"
            placeholder="Email"
            type="email"
            required
            value={formData.email}
            onChange={(e) => setFormData({ ...formData, email: e.target.value })}
          />

          <Input
            label="Password"
            placeholder="Password"
            type="password"
            required
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
            Register {!isLoading && <MoveRight size={16} />}
          </Button>

          <div className="text-center text-[12px] tracking-[0.6px] uppercase font-manrope">
            <span className="text-[#666]">Already have an account? </span>
            <Link href="/auth/login" className="text-white font-bold hover:underline">
              Log In
            </Link>
          </div>
        </div>
      </form>

      <AuthFooter />
    </div>
  );
}
