package domain

import (
	"context"
	"time"
	"voidspace/users/internal/domain/views"
)

type User struct {
	ID           int
	Username     string
	Email        string
	PasswordHash string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type UserRepository interface {
	Create(ctx context.Context, user *User) error
	GetUserByUsername(ctx context.Context, username string) (*User, error)
	IsFollowed(ctx context.Context, userID, targetUserID int32) (bool, error)
	GetUserFollowedById(ctx context.Context, userID int32) ([]int32, error)
	GetUserByEmail(ctx context.Context, email string) (*User, error)
	GetUserByCredentials(ctx context.Context, credentials string) (*User, error)
	GetUserByIds(ctx context.Context, userIDs []int32) ([]*views.UserProfile, error)
	GetUserProfile(ctx context.Context, userId int) (*views.UserProfile, error)
	DeleteUser(ctx context.Context, id int) error
}
