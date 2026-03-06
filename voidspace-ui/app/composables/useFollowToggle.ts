import type { User } from "@/types";

export const useFollowActions = (
  userData: Ref<User | null>,
  isSubmitting: Ref<boolean>
) => {
  const toast = useToast();
  const { follow, unFollow } = useFollow();

  const handleFollowToggle = async () => {
    if (!userData.value || isSubmitting.value) return;

    isSubmitting.value = true;

    try {
      if (userData.value.is_followed) {
        const res = await unFollow(userData.value.username);
        if (res.success) {
          userData.value.is_followed = false;
          userData.value.profile.followers = Math.max(
            0,
            userData.value.profile.followers - 1
          );
          toast.add({
            title: "Unfollowed",
            description: `You unfollowed ${userData.value.username}`,
            color: "neutral",
          });
        } else {
          throw new Error(res.detail || "Failed to unfollow");
        }
      } else {
        const res = await follow(userData.value.username);
        if (res.success) {
          userData.value.is_followed = true;
          userData.value.profile.followers += 1;
          toast.add({
            title: "Followed",
            description: `You are now following ${userData.value.username}`,
            color: "neutral",
          });
        } else {
          throw new Error(res.detail || "Failed to follow");
        }
      }
    } catch (error: any) {
      toast.add({
        title: "Action Failed",
        description: error.message || "Please try again later",
        color: "error",
      });
    } finally {
      isSubmitting.value = false;
    }
  };

  return {
    handleFollowToggle,
  };
};
