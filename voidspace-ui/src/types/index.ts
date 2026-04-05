export interface ApiResponse<T = any> {
  success: boolean;
  detail: string;
  data?: T;
}

export interface AuthResponse {
  access_token: string;
  expires_in: number;
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

export interface User {
  id: number;
  username: string;
  email?: string;
  profile: Profile;
  created_at: string;
  is_followed: boolean;
}
