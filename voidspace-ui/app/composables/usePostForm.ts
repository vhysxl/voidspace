import { reactive, ref, computed } from "vue";
import type { CreatePostReq } from "@/types";
import { createPostSchema } from "@/validations";
import * as v from "valibot";

export const usePostForm = () => {
  const toast = useToast();
  const { uploadFile } = useUpload();
  const { createPost } = usePost();
  const router = useRouter();
  const auth = useAuthStore();

  const maxLength = 240;

  const state = reactive({
    content: "",
    postImages: undefined as File[] | undefined,
  });

  const remainingChars = computed(() => maxLength - state.content.length);
  const isSubmitting = ref(false);

  const resetForm = () => {
    state.content = "";
    state.postImages = undefined;
  };

  const submitPost = async (): Promise<void> => {
    if (isSubmitting.value) return;
    if (!auth.user) {
      await router.push("/auth/login");
      return;
    }

    if (state.postImages && state.postImages.length > 5) {
      toast.add({
        title: "Too many images",
        description: "You can only upload up to 5 images",
        color: "error",
      });
      return;
    }

    try {
      isSubmitting.value = true;
      await v.parseAsync(createPostSchema, state);

      const uploadedUrls: string[] = [];

      if (state.postImages) {
        const results = await Promise.all(
          state.postImages.map((file) => uploadFile(file))
        );
        uploadedUrls.push(...results);
      }

      const payload: CreatePostReq = {
        content: state.content || "",
        post_images: uploadedUrls,
      };

      await createPost(payload);

      toast.add({
        title: "Post Created",
        description: "Your post has been created successfully",
        color: "neutral",
      });

      resetForm();

      window.location.reload();
    } catch (error: any) {
      toast.add({
        title: "Creation Failed",
        description:
          error.message || "Failed to create post, please try again later",
        color: "error",
      });
    } finally {
      isSubmitting.value = false;
    }
  };

  return {
    state,
    remainingChars,
    isSubmitting,
    resetForm,
    submitPost,
  };
};
