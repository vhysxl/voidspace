import { useCallback } from "react";
import { apiFetch } from "@/lib/api";
import { ApiResponse, User, UpdateProfileRequest, UserBanner } from "@/types";
import { useAuthStore } from "@/store/useAuthStore";

export const useUser = () => {
  const { updateUser } = useAuthStore();

  const me = useCallback(async () => {
    const userProfile = await apiFetch<ApiResponse<User>>("/user/me");

    if (userProfile.success && userProfile.data) {
      updateUser(userProfile.data);
    }
  }, [updateUser]);

  const getUser = useCallback(async (username: string) => {
    return await apiFetch<ApiResponse<User>>(`/user/${username}`);
  }, []);

  const updateProfile = useCallback(async (data: UpdateProfileRequest) => {
    const response = await apiFetch<ApiResponse>("/user/me", {
      method: "PUT",
      body: JSON.stringify(data),
    });

    if (response.success) {
      me();
    }

    return response;
  }, [me]);

  const deleteUser = useCallback(async () => {
    return await apiFetch<ApiResponse<null>>("/user/me", {
      method: "DELETE",
    });
  }, []);

  const followUser = useCallback(async (username: string) => {
    return await apiFetch<ApiResponse<null>>(`/follow/${username}`, {
      method: "POST",
    });
  }, []);

  const unfollowUser = useCallback(async (username: string) => {
    return await apiFetch<ApiResponse<null>>(`/follow/${username}`, {
      method: "DELETE",
    });
  }, []);

  const getFollowers = useCallback(async (username: string) => {
    return await apiFetch<ApiResponse<UserBanner[]>>(`/user/${username}/followers`);
  }, []);

  const getFollowing = useCallback(async (username: string) => {
    return await apiFetch<ApiResponse<UserBanner[]>>(`/user/${username}/following`);
  }, []);

  return { getUser, updateProfile, deleteUser, followUser, unfollowUser, me, getFollowers, getFollowing };
};
