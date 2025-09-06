import { useApi } from "./useApi";

export interface User {
  id: number;
  username: string;
  profile: Profile;
  created_at: string;
}

export interface Profile {
  display_name: string;
  bio: string;
  avatar_url: string;
  banner_url: string;
  location: string;
  followers: number;
  following: number;
}

export interface UpdateProfileRequest {
  display_name?: string;
  bio?: string;
  avatar_url?: string;
  banner_url?: string;
  location?: string;
}

export interface ApiResponse<T = any> {
  success: boolean;
  detail: string;
  data?: T;
}

export const useUsers = () => {
  const { fetchWithAuth } = useApi();
  const config = useRuntimeConfig();
  const apiUrl = "/api";

  const getUser = async (username: string): Promise<ApiResponse<User>> => {
    try {
      const response = await $fetch<ApiResponse<User>>(
        `${apiUrl}/users/${username}`,
        {
          method: "GET",
          headers: { "Content-Type": "application/json" },
        }
      );

      return response;
    } catch (error: any) {
      throw new Error(error.data?.detail || "Failed to get user");
    }
  };

  //protected
  const getCurrentUser = async (): Promise<ApiResponse<User>> => {
    try {
      const response = await fetchWithAuth(`${apiUrl}/users/me`, {
        method: "GET",
        headers: { "Content-Type": "application/json" },
      });

      return response as ApiResponse<User>;
    } catch (error: any) {
      throw new Error(error.data?.detail || "Failed to get current user");
    }
  };

  const updateProfile = async (
    profileData: UpdateProfileRequest
  ): Promise<ApiResponse> => {
    ``;
    try {
      const response = await fetchWithAuth(`${apiUrl}/users/me`, {
        method: "PUT",
        body: profileData,
      });

      return response as ApiResponse;
    } catch (error: any) {
      throw new Error(error.data?.detail || "Failed to update profile");
    }
  };

  const deleteUser = async (): Promise<ApiResponse> => {
    try {
      const response = await fetchWithAuth(`${apiUrl}/users/me`, {
        method: "DELETE",
      });

      return response as ApiResponse;
    } catch (error: any) {
      throw new Error(error.data?.detail || "Failed to delete user");
    }
  };

  return { getUser, getCurrentUser, updateProfile, deleteUser };
};
