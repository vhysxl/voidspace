"use client";

import { useState } from "react";
import Link from "next/link";
import { useAuth } from "@/hooks/useAuth";
import Input from "@/components/ui/Input";
import Button from "@/components/ui/Button";
import AuthHeader from "./AuthHeader";
import AuthFooter from "./AuthFooter";
import { MoveRight } from "lucide-react";
import { safeParse, flatten } from "valibot";
import { loginSchema } from "@/lib/validations";

export default function LoginForm() {
  const { login } = useAuth();
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);
  const [fieldErrors, setFieldErrors] = useState<Record<string, string>>({});
  const [formData, setFormData] = useState({
    usernameoremail: "",
    password: "",
  });

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setIsLoading(true);
    setError(null);
    setFieldErrors({});

    const result = safeParse(loginSchema, formData);
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
        subtitle="Log in to your your account"
      />

      <form onSubmit={handleSubmit} className="space-y-6 md:space-y-8" noValidate>
        <div className="space-y-6 md:space-y-8">
          <Input
            label="Username or Email"
            placeholder="Username or Email"
            required
            value={formData.usernameoremail}
            onChange={(e) => setFormData({ ...formData, usernameoremail: e.target.value })}
            error={fieldErrors.usernameoremail}
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

        <div className="space-y-6 md:space-y-8 pt-2">
          <Button type="submit" isLoading={isLoading}>
            Sign In {!isLoading && <MoveRight size={16} />}
          </Button>

          <div className="text-center text-[15px] tracking-[0.6px] uppercase font-manrope">
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
