export interface UploadResponse {
  data: {
    publicUrl: string;
    signedUrl: string;
  };
  publicUrl: string;
  detail: string;
  success: boolean;
}

export const useUpload = () => {
  const config = useRuntimeConfig();
  const { fetchWithAuth } = useApi();
  const apiUrl = config.public.apiUrl;

  const uploadFile = async (file: File): Promise<string> => {
    try {
      // Get signed URL
      const response = (await fetchWithAuth(`${apiUrl}/upload/signed-url`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          contentType: file.type,
        }),
      })) as UploadResponse;

      // Upload file
      await fetch(response.data.signedUrl, {
        method: "PUT",
        headers: {
          "Content-Type": file.type,
        },
        body: file,
      });

      return response.data.publicUrl;
    } catch (error: any) {
      throw new Error(error.data?.detail || "Failed to upload image");
    }
  };

  return { uploadFile };
};
