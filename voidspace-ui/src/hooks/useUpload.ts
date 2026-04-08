import { apiFetch } from "@/lib/api";
import { ApiResponse, SignedURLRequest, SignedURLResponse, UploadFolder } from "@/types";

export const useUpload = () => {
  const getSignedURL = async (folder: UploadFolder, contentType: string) => {
    const data: SignedURLRequest = { folder, contentType };
    return await apiFetch<ApiResponse<SignedURLResponse>>("/upload/signed-url", {
      method: "POST",
      body: JSON.stringify(data),
    });
  };

  const uploadFile = async (file: File, signedUrl: string) => {
    const response = await fetch(signedUrl, {
      method: "PUT",
      body: file,
      headers: {
        "Content-Type": file.type,
      },
    });

    if (!response.ok) {
      throw new Error("Failed to upload file to storage");
    }

    return true;
  };

  const uploadImage = async (file: File, folder: UploadFolder) => {
    // 1. Get signed URL
    const signedRes = await getSignedURL(folder, file.type);
    if (!signedRes.success || !signedRes.data) {
      throw new Error(signedRes.detail || "Failed to get signed URL");
    }

    const { signedUrl, publicUrl } = signedRes.data;

    // 2. Upload to GCS
    await uploadFile(file, signedUrl);

    // 3. Return public URL and image dimensions (for posts)
    return new Promise<{ url: string; width: number; height: number }>((resolve, reject) => {
      const img = new Image();
      img.onload = () => {
        resolve({
          url: publicUrl,
          width: img.width,
          height: img.height,
        });
      };
      img.onerror = () => {
        // Fallback if image loading fails for some reason
        resolve({ url: publicUrl, width: 0, height: 0 });
      };
      img.src = URL.createObjectURL(file);
    });
  };

  return { uploadImage };
};
