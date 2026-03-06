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

type UserUsecase interface {
	// Auth
	Login(ctx context.Context, credentials, password string) (*User, error)
	Register(ctx context.Context, username, email, password string) (*User, error)

	GetCurrentUser(ctx context.Context, userID int) (*views.UserProfile, error)
	GetUser(ctx context.Context, username string, authUserID int) (*views.UserProfile, error)
	GetUserByIDs(ctx context.Context, userIDs []int) ([]views.UserProfile, error)

	ListFollowers(ctx context.Context, userID int) ([]views.UserBanner, error)
	ListFollowing(ctx context.Context, userID int) ([]views.UserBanner, error)

	DeleteUser(ctx context.Context, userID int) error
	RestoreUser(ctx context.Context, userID int) error
}

type UserRepository interface {
	Create(ctx context.Context, user *User) error
	Delete(ctx context.Context, userID int) error
	SoftDelete(ctx context.Context, userID int) error

	GetByUsername(ctx context.Context, username string) (*views.UserProfile, error)
	GetByEmail(ctx context.Context, email string) (*User, error)
	GetByCredentials(ctx context.Context, credentials string) (*User, error)

	GetByIDs(ctx context.Context, userIDs []int) ([]views.UserProfile, error)

	GetProfile(ctx context.Context, userID int) (*views.UserProfile, error)

	ListFollowers(ctx context.Context, userID int) ([]views.UserBanner, error)
	ListFollowing(ctx context.Context, userID int) ([]views.UserBanner, error)

	// Compensation
	RestoreUser(ctx context.Context, userID int) error
}
