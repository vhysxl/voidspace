package domain

import (
	"context"
	"time"
	"voidspace/users/internal/domain/views"
)

type User struct {
	ID           int       `json:"id" db:"id"`
	Username     string    `json:"username" db:"username"`
	Email        string    `json:"email" db:"email"`
	PasswordHash string    `json:"-" db:"password_hash"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
}

type UserRepository interface {
	Create(ctx context.Context, user *User) error
	GetUsersByUsername(ctx context.Context, username string) ([]*User, error)
	GetUserByID(ctx context.Context, id int) (*User, error)
	GetUserByUsername(ctx context.Context, username string) (*User, error)
	GetUserByEmail(ctx context.Context, email string) (*User, error)
	GetUserProfile(ctx context.Context, userId int) (*views.UserProfile, error)
	UpdateUser(ctx context.Context, user *User) error
	DeleteUser(ctx context.Context, id int) error
	GetUserByCredentials(ctx context.Context, credentials string) (*User, error)
}
