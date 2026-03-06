export const isValidTokenResponse = (
  data: any
): data is { access_token: string; expires_in: number } => {
  return (
    data &&
    typeof data.access_token === "string" &&
    data.access_token.trim() !== "" &&
    typeof data.expires_in === "number" &&
    data.expires_in > 0
  );
};
