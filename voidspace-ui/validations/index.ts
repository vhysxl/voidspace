import * as v from "valibot";

export const loginSchema = v.object({
  credential: v.pipe(v.string(), v.nonEmpty("Username or Email is required")),
  password: v.pipe(v.string(), v.nonEmpty("Password is required")),
});

export const registerSchema = v.pipe(
  v.object({
    username: v.pipe(v.string(), v.nonEmpty("Username is required")),
    email: v.pipe(
      v.string(),
      v.nonEmpty("Please enter your email."),
      v.email("The email address is not valid.")
    ),
    password: v.pipe(
      v.string(),
      v.nonEmpty("Please enter your password."),
      v.minLength(8, "Your password must have 8 characters or more.")
    ),
    confirmPassword: v.pipe(
      v.string(),
      v.nonEmpty("Please confirm your password.")
    ),
  }),
  v.forward(
    v.partialCheck(
      [["password"], ["confirmPassword"]],
      (input) => input.password === input.confirmPassword,
      "The two passwords do not match."
    ),
    ["confirmPassword"]
  )
);

export const editProfileSchema = v.object({
  displayName: v.optional(
    v.pipe(
      v.string(),
      v.maxLength(20, "Display name must be 20 characters or less")
    )
  ),
  bio: v.optional(
    v.pipe(v.string(), v.maxLength(160, "Bio must be 160 characters or less"))
  ),
  location: v.optional(
    v.pipe(
      v.string(),
      v.maxLength(50, "Location must be 50 characters or less")
    )
  ),
  avatar: v.optional(
    v.pipe(
      v.file(),
      v.mimeType(
        ["image/jpeg", "image/png", "image/webp"],
        "Format must be JPEG, PNG, or WebP"
      ),
      v.maxSize(2 * 1024 * 1024, "Avatar must be maximum 2MB")
    )
  ),
  banner: v.optional(
    v.pipe(
      v.file(),
      v.mimeType(
        ["image/jpeg", "image/png", "image/webp"],
        "Format must be JPEG, PNG, or WebP"
      ),
      v.maxSize(5 * 1024 * 1024, "Banner must be maximum 5MB")
    )
  ),
});
