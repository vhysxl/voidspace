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
	GetUser(ctx context.Context, username string, userID int) (*views.UserProfile, error)
	GetUserByID(ctx context.Context, userID int32) (*views.UserProfile, error)
	GetUserByIds(ctx context.Context, UserIDs []int32) ([]*views.UserProfile, error)
	GetUserFollowedByID(ctx context.Context, userID int32) ([]int32, error)
	DeleteUser(ctx context.Context, ID int) error
}

func NewUserUsecase(userRepository domain.UserRepository, contextTimeout time.Duration) UserUsecase {
	return &userUsecase{
		userRepository: userRepository,
		contextTimeout: contextTimeout,
	}
}

// GetCurrentUser implements UserUsecase.
func (u *userUsecase) GetCurrentUser(ctx context.Context, ID int) (*views.UserProfile, error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	return u.userRepository.GetUserProfile(ctx, ID)
}

// GetUser implements UserUsecase.
func (u *userUsecase) GetUser(ctx context.Context, username string, userID int) (*views.UserProfile, error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	user, err := u.userRepository.GetUserByUsername(ctx, username)
	if err != nil {
		return nil, err
	}

	fullUserData, err := u.userRepository.GetUserProfile(ctx, user.ID)
	if err != nil {
		return nil, err
	}

	fullUserData.IsFollowed = false

	if userID > 0 {
		isFollowed, err := u.userRepository.IsFollowed(ctx, int32(userID), int32(user.ID))
		if err != nil {
			return nil, err
		}

		fullUserData.IsFollowed = isFollowed
	}

	return fullUserData, nil
}

// GetUserFollowedByID implements UserUsecase.
func (u *userUsecase) GetUserFollowedByID(ctx context.Context, userID int32) ([]int32, error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	return u.userRepository.GetUserFollowedById(ctx, userID)
}

// GetUserByID implements UserUsecase.
func (u *userUsecase) GetUserByID(ctx context.Context, userID int32) (*views.UserProfile, error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	return u.userRepository.GetUserProfile(ctx, int(userID))
}

// GetUserByIds implements UserUsecase.
func (u *userUsecase) GetUserByIds(ctx context.Context, UserIDs []int32) ([]*views.UserProfile, error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	users, err := u.userRepository.GetUserByIds(ctx, UserIDs)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (u *userUsecase) DeleteUser(ctx context.Context, ID int) error {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	return u.userRepository.DeleteUser(ctx, ID)
}
