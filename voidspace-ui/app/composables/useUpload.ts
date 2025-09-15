import * as nsfwjs from "nsfwjs";

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
  const { fetchWithAuth } = useApi();

  let nsfwModel: any = null;

  const loadNSFWModel = async () => {
    if (!nsfwModel) {
      nsfwModel = await nsfwjs.load();
    }
    return nsfwModel;
  };

  const checkNSFW = async (file: File): Promise<boolean> => {
    try {
      const model = await loadNSFWModel();

      // createImageBitmap
      const imageBitmap = await createImageBitmap(file);

      // start predictions
      const predictions = await model.classify(imageBitmap);

      const nsfwScore =
        predictions.find(
          (p: any) =>
            p.className === "Porn" ||
            p.className === "Sexy" ||
            p.className === "Hentai"
        )?.probability || 0;

      return nsfwScore > 0.7;
    } catch (error) {
      console.error("Failed to check NSFW:", error);
      return false;
    }
  };

  const uploadFile = async (file: File): Promise<string> => {
    try {
      const res = await checkNSFW(file);
      if (res) {
        throw new Error("Image not allowed. Upload something else!");
      }
      // Get signed URL
      const response = (await fetchWithAuth(`/api/upload/upload`, {
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
      throw new Error(
        error.data?.detail || error.message || "Failed to upload image"
      );
    }
  };

  return { uploadFile };
};
