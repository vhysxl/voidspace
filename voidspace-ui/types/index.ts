// ============================================
// Core API Types
// ============================================

export interface ApiResponse<T = never> {
  success: boolean;
  detail: string;
  data?: T;
}

// ============================================
// Authentication Types
// ============================================

export interface AuthResponse {
  access_token: string;
  expires_in: number;
}

// ============================================
// User & Profile Types
// ============================================

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
  profile: Profile;
  created_at: string;
  is_followed: boolean;
}

export interface UpdateProfileRequest {
  display_name?: string;
  bio?: string;
  avatar_url?: string;
  banner_url?: string;
  location?: string;
}

// ============================================
// Comment Types
// ============================================

export interface CommentType {
  id: number;
  content: string;
  post_id: number;
  user_id: number;
  created_at: string;
  updated_at: string;
  author: User;
}

export interface CreateCommentReq {
  content: string;
  post_id: number;
}

// ============================================
// Post Types
// ============================================

export type Post = {
  id: number;
  content: string;
  user_id: number;
  post_images: string[] | null;
  likes_count: number;
  comments_count: number;
  created_at: string;
  updated_at: string;
  author: User;
  is_liked: boolean;
};

export interface CreatePostReq {
  content: string;
  post_images: string[] | null;
}

// ============================================
// Feed Types
// ============================================

export type Feed = {
  posts: Post[];
  hasmore: boolean;
};

// ============================================
// Upload Types
// ============================================

export interface UploadResponse {
  data: {
    publicUrl: string;
    signedUrl: string;
  };
  publicUrl: string;
  detail: string;
  success: boolean;
}
