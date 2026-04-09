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

export interface PostImage {
  image_url: string;
  order: number;
  width: number;
  height: number;
}

export interface Post {
  id: number;
  content: string;
  post_images: PostImage[];
  author: User;
  created_at: string;
  likes_count: number;
  comments_count: number;
  is_liked: boolean;
}

export interface Comment {
  comment_id: number;
  content: string;
  author: User;
  created_at: string;
  post_id: number;
}

export interface FeedResponse {
  posts: Post[];
  has_more: boolean;
}

// Request Types
export interface UpdateProfileRequest {
  display_name?: string;
  bio?: string;
  location?: string;
  avatar_url?: string;
  banner_url?: string;
}

export interface CreatePostRequest {
  content: string;
  post_images?: PostImage[];
}

export interface UpdatePostRequest {
  content: string;
  post_images?: PostImage[];
}

export interface CreateCommentRequest {
  post_id: number;
  content: string;
}

export interface FollowUserRequest {
  target_username: string;
}

// Upload Types
export type UploadFolder = "posts" | "avatars" | "banners";

export interface SignedURLRequest {
  contentType: string;
  folder: UploadFolder;
}

export interface SignedURLResponse {
  signedUrl: string;
  publicUrl: string;
  fileName: string;
  folder: string;
}

export interface UserBanner {
  id: number;
  username: string
  display_name: string;
  avatar_url: string;
}

export interface PostCardProps {
  post: Post;
  isDetailed?: boolean;
  onCommentClick?: () => void;
}
