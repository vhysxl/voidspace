// import * as nsfwjs from "nsfwjs";
import type { UploadResponse } from "@/types";
import { handleApiError } from "@/utils/apiErrorHandler";

/**
 * Upload service hook
 * Handles file uploads with NSFW content detection
 */
export const useUpload = () => {
  const { fetchWithAuth } = useApi();

  // Store NSFW model instance globally within the composable
  // This prevents re-loading the model on every function call
  let nsfwModel: any = null;

  /**
   * Loads the NSFW detection model (client-side only)
   * Uses dynamic import to avoid SSR issues with Node.js dependencies
   */
  const loadNSFWModel = async () => {
    // Double check: model not loaded AND we're on client-side
    if (!nsfwModel && import.meta.client) {
      // DYNAMIC IMPORT - This is the key solution!
      // Why dynamic import?
      // 1. Static imports are resolved by bundler at build time, before code execution
      // 2. nsfwjs requires Node.js built-in modules (buffer) that don't exist in browsers
      // 3. Dynamic import only executes when this code actually runs
      // 4. With import.meta.client, this only runs in browser, not on server
      const nsfwjs = await import("nsfwjs");

      // Load AI model for NSFW detection
      nsfwModel = await nsfwjs.load();
    }
    return nsfwModel;
  };

  /**
   * Checks if an image file contains NSFW content using AI
   * @param file - Image file to analyze
   * @returns Promise<boolean> - true if NSFW content detected
   */
  const checkNSFW = async (file: File): Promise<boolean> => {
    // Safety check: skip all processing if on server-side
    // import.meta.client = true when code runs in browser
    // import.meta.client = false when code runs on server (SSR)
    if (!import.meta.client) return false;

    try {
      const model = await loadNSFWModel();

      // If model failed to load (network issues, etc), skip detection
      if (!model) return false;

      // Convert file to bitmap that can be processed by AI
      const imageBitmap = await createImageBitmap(file);

      // AI classification - returns array of predictions with confidence scores
      const predictions = await model.classify(imageBitmap);

      // Find confidence score for NSFW categories
      // Take the highest from Porn, Sexy, or Hentai
      const nsfwScore =
        predictions.find(
          (p: any) =>
            p.className === "Porn" ||
            p.className === "Sexy" ||
            p.className === "Hentai"
        )?.probability || 0;

      // Threshold 0.7 = 70% confidence
      // If AI is 70% confident this is NSFW, block upload
      return nsfwScore > 0.7;
    } catch (error) {
      console.error("Failed to check NSFW:", error);
      // On error, default to allow upload (graceful failure)
      return false;
    }
  };

  /**
   * Uploads a file to cloud storage with NSFW content filtering
   * @param file - File to upload
   * @returns Promise<string> - Public URL of uploaded file
   * @throws Error if NSFW content detected or upload fails
   */
  const uploadFile = async (file: File): Promise<string> => {
    try {
      // Check for NSFW content before uploading
      const isNSFW = await checkNSFW(file);
      if (isNSFW) {
        throw new Error(
          "Image contains inappropriate content. Please upload a different image."
        );
      }

      // Get signed URL for direct upload to cloud storage
      const response = await fetchWithAuth<UploadResponse>(
        `/api/upload/upload`,
        {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify({
            contentType: file.type,
          }),
        }
      );

      // Upload file directly to cloud storage using signed URL
      const uploadResponse = await fetch(response.data.signedUrl, {
        method: "PUT",
        headers: {
          "Content-Type": file.type,
        },
        body: file,
      });

      // Check if direct upload to cloud storage failed
      if (!uploadResponse.ok) {
        handleApiError(
          uploadResponse,
          "Failed to upload image, please try again later"
        );
      }

      return response.data.publicUrl;
    } catch (error: unknown) {
      handleApiError(error, "Failed to upload image. Please try again.");
    }
  };

  return { uploadFile };
};
