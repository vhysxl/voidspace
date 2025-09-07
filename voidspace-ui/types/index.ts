export interface ApiResponse<T> {
  success: boolean;
  detail: string;
  data?: T;
}

export interface AuthResponse {
  access_token: string;
  expires_in: number;
}
