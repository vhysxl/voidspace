package domain

import (
	"context"
	"time"
)

type Follow struct {
	UserID       int
	TargetUserID int
	CreatedAt    time.Time
}

type FollowUsecase interface {
	Follow(ctx context.Context, authUserID int, targetUserID int) error
	Unfollow(ctx context.Context, authUserID int, targetUserID int) error
}

type FollowRepository interface {
	Follow(ctx context.Context, updates *Follow) error
	Unfollow(ctx context.Context, updates *Follow) error
	IsFollowing(ctx context.Context, userID, targetUserID int) (bool, error)
}
