import type { User } from "@/types";

export const useAccountDeletion = (
  userData: Ref<User | null>,
  isSubmitting: Ref<boolean>
) => {
  const toast = useToast();
  const auth = useAuthStore();
  const { deleteUser } = useUser();

  const deleteConfirmation = ref("");

  const deleteAccount = async () => {
    if (deleteConfirmation.value !== userData.value?.username) return;
    if (isSubmitting.value) return;

    isSubmitting.value = true;

    try {
      await deleteUser();

      toast.add({
        title: "Account Deleted",
        description: "Your account has been permanently deleted",
        color: "error",
      });

      auth.user = null;
      auth.accessToken = "";
      auth.expiresIn = 0;

      await navigateTo("/auth/login");
    } catch (error: any) {
      toast.add({
        title: "Delete Failed",
        description: error.message || "Failed to delete account",
        color: "error",
      });
    } finally {
      isSubmitting.value = false;
    }
  };

  const canDelete = computed(
    () => deleteConfirmation.value === userData.value?.username
  );

  return {
    deleteConfirmation,
    deleteAccount,
    canDelete,
  };
};
