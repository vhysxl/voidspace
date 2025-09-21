export const defaultOptions = {
  headers: { "Content-Type": "application/json" },
};

export const defaultAuthOptions = {
  ...defaultOptions,
  credentials: "include" as const,
};
