package domain

import (
	"context"
	"time"
	"voidspace/users/internal/domain/views"
)

type User struct {
	Id           int32
	Username     string
	Email        string
	PasswordHash string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type UserUsecase interface {
	GetCurrentUser(ctx context.Context, userId int32) (*views.UserProfile, error)
	GetUser(ctx context.Context, username string, userId int32) (*views.UserProfile, error)
	GetUserById(ctx context.Context, userId int32) (*views.UserProfile, error)
	GetUserByIds(ctx context.Context, userIds []int32) ([]*views.UserProfile, error)
	GetUserFollowedById(ctx context.Context, userId int32) ([]int32, error)
	DeleteUser(ctx context.Context, userId int32) error
}

type UserRepository interface {
	Create(ctx context.Context, user *User) error
	GetUserByUsername(ctx context.Context, username string) (*User, error)
	IsFollowed(ctx context.Context, userId, targetUserId int32) (bool, error)
	GetUserFollowedById(ctx context.Context, userId int32) ([]int32, error)
	GetUserByEmail(ctx context.Context, email string) (*User, error)
	GetUserByCredentials(ctx context.Context, credentials string) (*User, error)
	GetUserByIds(ctx context.Context, userIds []int32) ([]*views.UserProfile, error)
	GetUserProfile(ctx context.Context, userId int32) (*views.UserProfile, error)
	DeleteUser(ctx context.Context, userId int32) error
}
