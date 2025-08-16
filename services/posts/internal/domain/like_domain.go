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
	LikePost(ctx context.Context, like *Like) error
	UnlikePost(ctx context.Context, like *Like) error
}
