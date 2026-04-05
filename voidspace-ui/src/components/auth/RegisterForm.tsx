"use client";

import React, { useState } from "react";
import Link from "next/link";
import { useAuth } from "@/hooks/useAuth";
import Input from "@/components/ui/Input";
import Button from "@/components/ui/Button";
import AuthHeader from "./AuthHeader";
import AuthFooter from "./AuthFooter";
import { MoveRight } from "lucide-react";
import { safeParse, flatten } from "valibot";
import { registerSchema } from "@/lib/validations";

export default function RegisterForm() {
  const { register } = useAuth();
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);
  const [fieldErrors, setFieldErrors] = useState<Record<string, string>>({});
  const [formData, setFormData] = useState({
    username: "",
    email: "",
    password: "",
  });

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setIsLoading(true);
    setError(null);
    setFieldErrors({});

    const result = safeParse(registerSchema, formData);
    if (!result.success) {
      const flattened = flatten(result.issues);
      const errors: Record<string, string> = {};
      if (flattened.nested) {
        Object.entries(flattened.nested).forEach(([key, value]) => {
          if (value) errors[key] = value[0];
        });
      }
      setFieldErrors(errors);
      setIsLoading(false);
      return;
    }

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
        subtitle="Create your your account to get started" 
      />

      <form onSubmit={handleSubmit} className="space-y-8" noValidate>
        <div className="space-y-8">
          <Input
            label="Username"
            placeholder="Username"
            required
            value={formData.username}
            onChange={(e) => setFormData({ ...formData, username: e.target.value })}
            error={fieldErrors.username}
          />

          <Input
            label="Email Address"
            placeholder="Email"
            type="email"
            required
            value={formData.email}
            onChange={(e) => setFormData({ ...formData, email: e.target.value })}
            error={fieldErrors.email}
          />

          <Input
            label="Password"
            placeholder="Password"
            type="password"
            required
            value={formData.password}
            onChange={(e) => setFormData({ ...formData, password: e.target.value })}
            error={fieldErrors.password}
          />
        </div>

        {error && (
          <div className="bg-red-500/5 border border-red-500/20 py-3 px-4 text-red-500/90 text-[14px] tracking-[0.6px] uppercase text-center font-manrope animate-in fade-in duration-300">
            {error}
          </div>
        )}

        <div className="space-y-8 pt-2">
          <Button type="submit" isLoading={isLoading}>
            Register {!isLoading && <MoveRight size={16} />}
          </Button>

          <div className="text-center text-[15px] tracking-[0.6px] uppercase font-manrope">
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
