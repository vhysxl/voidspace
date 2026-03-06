package models

type LikeRequest struct {
	PostID   int    `json:"post_id" validate:"required"`
	UserID   string `json:"user_id" validate:"required"`
	Username string `json:"username" validate:"required"`
}

type LikeResponse struct {
	NewLikesCount int `json:"new_likes_count"`
}
