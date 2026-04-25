import * as v from "valibot";

export const loginSchema = v.object({
  usernameoremail: v.pipe(
    v.string(),
    v.nonEmpty("Username or email is required")
  ),
  password: v.pipe(
    v.string(),
    v.nonEmpty("Password is required")
  ),
});

export const registerSchema = v.object({
  username: v.pipe(
    v.string(),
    v.nonEmpty("Username is required"),
    v.minLength(3, "Username must be at least 3 characters"),
    v.maxLength(30, "Username must be at most 30 characters"),
    v.regex(/^[a-zA-Z0-9]+$/, "Username must be alphanumeric")
  ),
  email: v.pipe(
    v.string(),
    v.nonEmpty("Email is required"),
    v.email("Invalid email address")
  ),
  password: v.pipe(
    v.string(),
    v.nonEmpty("Password is required"),
    v.minLength(6, "Password must be at least 6 characters")
  ),
});

export type LoginInput = v.InferOutput<typeof loginSchema>;
export type RegisterInput = v.InferOutput<typeof registerSchema>;
