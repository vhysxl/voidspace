package models

type LikeRequest struct {
	PostID   int    `json:"postid" validate:"required"`
	UserID   string `json:"userid" validate:"required"`
	Username string `json:"username" validate:"required"`
}

type LikeResponse struct {
	NewLikesCount int `json:"newlikescount"`
}
