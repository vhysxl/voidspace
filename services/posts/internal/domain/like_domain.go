package domain

import (
	"context"
	"time"
)

type Like struct {
	PostID    int
	UserID    int
	CreatedAt time.Time
}

type LikeUsecase interface {
	LikePost(ctx context.Context, like *Like) error
	UnlikePost(ctx context.Context, like *Like) error
}

type LikeRepository interface {
	LikePost(ctx context.Context, like *Like) error
	UnlikePost(ctx context.Context, like *Like) error

	IsPostLikedByUser(ctx context.Context, userID, postID int) (bool, error)
	IsPostsLikedByUser(ctx context.Context, userID int, postIDs []int) (map[int]bool, error)
}
