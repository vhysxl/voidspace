import { useApi } from "./useApi";
import type { ApiResponse, UpdateProfileRequest, User } from "@/types/index";
import { handleApiError } from "@/utils/apiErrorHandler";
import { defaultOptions, defaultAuthOptions } from "@/utils/apiDefaults";

/**
 * User service hooks
 * Provides user-related operations (profile management, user data).
 */
type UserApiResponse = ApiResponse<User>;

export const useUser = () => {
  const { fetchWithAuth } = useApi();
  const apiUrl = "/api/user";

  const getUser = async (username: string): Promise<UserApiResponse> => {
    try {
      return await $fetch<UserApiResponse>(`${apiUrl}/${username}`, {
        ...defaultOptions,
        method: "GET",
      });
    } catch (error: unknown) {
      return handleApiError(error, "Failed to get user");
    }
  };

  const getCurrentUser = async (): Promise<UserApiResponse> => {
    try {
      return await fetchWithAuth<UserApiResponse>(`${apiUrl}/me`, {
        ...defaultAuthOptions,
        method: "GET",
      });
    } catch (error: unknown) {
      return handleApiError(error, "Failed to get current user");
    }
  };

  const updateProfile = async (
    profileData: UpdateProfileRequest
  ): Promise<ApiResponse> => {
    try {
      return await fetchWithAuth<ApiResponse>(`${apiUrl}/me`, {
        ...defaultAuthOptions,
        method: "PUT",
        body: profileData,
      });
    } catch (error: unknown) {
      return handleApiError(error, "Failed to update profile");
    }
  };

  const deleteUser = async (): Promise<ApiResponse> => {
    try {
      const response = await fetchWithAuth<ApiResponse>(`${apiUrl}/me`, {
        ...defaultAuthOptions,
        method: "DELETE",
      });

      // Clear auth state after successful deletion
      const authCookie = useCookie("auth");
      authCookie.value = null;

      return response;
    } catch (error: unknown) {
      return handleApiError(error, "Failed to delete user");
    }
  };

  return {
    getUser,
    getCurrentUser,
    updateProfile,
    deleteUser,
  };
};
