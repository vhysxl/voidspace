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

type LikeRepository interface {
	LikePost(ctx context.Context, like *Like) (int32, error)
	UnlikePost(ctx context.Context, like *Like) (int32, error)
	DeleteAllLikes(ctx context.Context, userID int32) error
}
