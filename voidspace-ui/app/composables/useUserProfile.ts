import type { User } from "@/types";

export const useUserProfile = () => {
  const userData = ref<User | null>(null);
  const auth = useAuthStore();
  const isLoadingProfile = ref(false);
  const error = ref<string | null>(null);

  const fetchUserProfile = async (username: string) => {
    isLoadingProfile.value = true;
    error.value = null;

    try {
      const { getUser } = useUser();
      const res = await getUser(username);
      if (res.success && res.data) {
        userData.value = res.data;
      } else {
        throw new Error("Failed to fetch user data");
      }
    } catch (err) {
      error.value = err instanceof Error ? err.message : "An error occurred";
    } finally {
      isLoadingProfile.value = false;
    }
  };

  return {
    userData,
    isLoadingProfile,
    error,
    fetchUserProfile,
  };
};
