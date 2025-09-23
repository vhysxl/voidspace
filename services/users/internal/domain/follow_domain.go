package domain

import (
	"context"
	"time"
)

type Follow struct {
	UserId       int32
	TargetUserId int32
	CreatedAt    time.Time
}

type FollowUsecase interface {
	Follow(ctx context.Context, followId int32, username string) error
	Unfollow(ctx context.Context, followerId int32, username string) error
}

type FollowRepository interface {
	Follow(ctx context.Context, updates *Follow) error
	Unfollow(ctx context.Context, updates *Follow) error
}
