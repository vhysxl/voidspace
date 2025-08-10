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

type FollowRepository interface {
	Follow(ctx context.Context, updates *Follow) error
	Unfollow(ctx context.Context, updates *Follow) error
}
