export function handleApiError(error: any, fallbackMessage: string): never {
  const message =
    error?.statusMessage || error?.message || error?.detail || fallbackMessage;
  throw new Error(message);
}
