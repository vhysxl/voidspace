import { apiFetch } from "@/lib/api";
import { ApiResponse, User, UpdateProfileRequest } from "@/types";
import { useAuthStore } from "@/store/useAuthStore";

export const useUser = () => {
  const { updateUser } = useAuthStore();

  const getUser = async (username: string) => {
    return await apiFetch<ApiResponse<User>>(`/user/${username}`);
  };

  const updateProfile = async (data: UpdateProfileRequest) => {
    const response = await apiFetch<ApiResponse<User>>("/user/me", {
      method: "PUT",
      body: JSON.stringify(data),
    });

    if (response.success && response.data) {
      updateUser(response.data);
    }

    return response;
  };

  const deleteUser = async () => {
    return await apiFetch<ApiResponse<null>>("/user/me", {
      method: "DELETE",
    });
  };

  return { getUser, updateProfile, deleteUser };
};
