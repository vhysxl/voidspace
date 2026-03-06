import type { UpdateProfileRequest, User } from "@/types";
import type { FormSubmitEvent } from "@nuxt/ui";
import { safeParse } from "valibot";
import { editProfileSchema } from "@/validations";

export const useProfileForm = (
  userData: Ref<User | null>,
  isSubmitting: Ref<boolean>
) => {
  const toast = useToast();
  const user = useAuthStore();
  const { fetchUserProfile } = useUserProfile();
  const { uploadFile } = useUpload();
  const { updateProfile } = useUser();

  const state = reactive({
    displayName: "",
    bio: "",
    location: "",
    avatar: undefined as File | undefined,
    banner: undefined as File | undefined,
  });

  const resetForm = () => {
    if (!userData.value) return;
    state.displayName = userData.value.profile.display_name || "";
    state.bio = userData.value.profile.bio || "";
    state.location = userData.value.profile.location || "";
    state.avatar = undefined;
    state.banner = undefined;
  };

  const onSubmit = async (event: FormSubmitEvent<any>) => {
    if (isSubmitting.value) return;
    isSubmitting.value = true;

    // Prepare an array of promises, each promise resolves to a [key, value] tuple
    // Example: ["avatar", uploadResult] or ["banner", uploadResult]
    try {
      const validation = safeParse(editProfileSchema, event.data);

      if (!validation.success) {
        const errors = validation.issues
          .map((issue) => issue.message)
          .join(", ");
        toast.add({
          title: "Validation Error",
          description: errors,
          color: "error",
        });
        return;
      }

      const validatedData = validation.output;

      const uploadPromises: Promise<[string, any]>[] = [];

      // If avatar exists, push its upload promise with identifier "avatar"
      if (validatedData.avatar) {
        uploadPromises.push(
          uploadFile(event.data.avatar).then((res) => ["avatar", res])
        );
      }

      // If banner exists, push its upload promise with identifier "banner"
      if (validatedData.banner) {
        uploadPromises.push(
          uploadFile(event.data.banner).then((res) => ["banner", res])
        );
      }

      // Wait for all uploads to finish, results will be an array of [key, value] pairs
      // Example: [["avatar", {...}], ["banner", {...}]]
      const results = await Promise.all(uploadPromises);

      // convert it to an object (hashmap):
      const mapped = Object.fromEntries(results);

      // Use state.displayName if it's not null/undefined,
      // otherwise fall back to user.user?.profile.displayName.
      // Unlike `||`, this allows empty string ("") to be a valid value.
      const profileData: UpdateProfileRequest = {
        display_name:
          validatedData.displayName ?? user.user?.profile.display_name,
        bio: validatedData.bio ?? user.user?.profile.bio,
        avatar_url: mapped.avatar ?? user.user?.profile.avatar_url,
        banner_url: mapped.banner ?? user.user?.profile.banner_url,
        location: validatedData.location ?? user.user?.profile.location,
      };

      await updateProfile(profileData);

      await Promise.all([
        user.refreshUser(),
        fetchUserProfile(user.user!.username!),
      ]);

      toast.add({
        title: "Profile Updated",
        description: "Your profile has been updated successfully",
        color: "neutral",
      });
    } catch (error: any) {
      toast.add({
        title: "Update Failed",
        description:
          error.message || "Failed to update profile, please try again later",
        color: "error",
      });
    } finally {
      isSubmitting.value = false;
    }
  };

  return {
    state,
    isSubmitting,
    resetForm,
    onSubmit,
  };
};
