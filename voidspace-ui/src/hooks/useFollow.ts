import { apiFetch } from "@/lib/api";
import { ApiResponse, User, FollowUserRequest } from "@/types";

export const useFollow = () => {
  const followUser = async (target_username: string) => {
    const data: FollowUserRequest = { target_username };
    return await apiFetch<ApiResponse<null>>("/follow", {
      method: "POST",
      body: JSON.stringify(data),
    });
  };

  const unfollowUser = async (target_username: string) => {
    const data: FollowUserRequest = { target_username };
    return await apiFetch<ApiResponse<null>>("/follow", {
      method: "DELETE",
      body: JSON.stringify(data),
    });
  };

  const getFollowers = async (username: string) => {
    return await apiFetch<ApiResponse<User[]>>(`/user/${username}/followers`);
  };

  const getFollowing = async (username: string) => {
    return await apiFetch<ApiResponse<User[]>>(`/user/${username}/following`);
  };

  return {
    followUser,
    unfollowUser,
    getFollowers,
    getFollowing,
  };
};
