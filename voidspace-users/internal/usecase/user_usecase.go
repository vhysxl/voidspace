package usecase

import (
	"context"
	"time"
	"voidspace/users/internal/domain"
	"voidspace/users/internal/domain/views"
)

type userUsecase struct {
	userRepository domain.UserRepository
	contextTimeout time.Duration
}

type UserUsecase interface {
	GetCurrentUser(ctx context.Context, ID int) (*views.UserProfile, error)
	GetUser(ctx context.Context, username string) (*views.UserProfile, error)
	UpdateUser(ctx context.Context, userIdRequester int, userIdTarget int, updateData *views.UserProfile) (*views.UserProfile, error)
	DeleteUser(ctx context.Context, ID int) error
}

func NewUserUsecase(userRepository domain.UserRepository, contextTimeout time.Duration) UserUsecase {
	return &userUsecase{
		userRepository: userRepository,
		contextTimeout: contextTimeout,
	}
}

// DeleteUser implements UserUsecase.
func (u *userUsecase) DeleteUser(ctx context.Context, ID int) error {
	panic("unimplemented")
}

// GetCurrentUser implements UserUsecase.
func (u *userUsecase) GetCurrentUser(ctx context.Context, ID int) (*views.UserProfile, error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	return u.userRepository.GetUserProfile(ctx, ID)
}

// GetUser implements UserUsecase.
func (u *userUsecase) GetUser(ctx context.Context, username string) (*views.UserProfile, error) {
	panic("unimplemented")
}

// UpdateUser implements UserUsecase.
func (u *userUsecase) UpdateUser(ctx context.Context, userIdRequester int, userIdTarget int, updateData *views.UserProfile) (*views.UserProfile, error) {
	panic("unimplemented")
}
