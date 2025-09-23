package domain

import (
	"context"
	"time"
)

type Like struct {
	PostID    int32
	UserID    int32
	CreatedAt time.Time
}

type LikeUsecase interface {
	LikePost(ctx context.Context, like *Like) (int32, error)
	UnlikePost(ctx context.Context, like *Like) (int32, error)
}

type LikeRepository interface {
	LikePost(ctx context.Context, like *Like) (int32, error)
	UnlikePost(ctx context.Context, like *Like) (int32, error)
	DeleteAllLikes(ctx context.Context, userID int32) error
	IsPostLikedByUser(ctx context.Context, userID, postID int32) (bool, error)
	IsPostsLikedByUser(ctx context.Context, userID int32, postIDs []int32) (map[int32]bool, error)
}
