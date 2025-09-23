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

func NewUserUsecase(userRepository domain.UserRepository, contextTimeout time.Duration) domain.UserUsecase {
	return &userUsecase{
		userRepository: userRepository,
		contextTimeout: contextTimeout,
	}
}

// GetCurrentUser implements UserUsecase.
func (u *userUsecase) GetCurrentUser(ctx context.Context, userId int32) (*views.UserProfile, error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	return u.userRepository.GetUserProfile(ctx, userId)
}

// GetUser implements UserUsecase.
func (u *userUsecase) GetUser(ctx context.Context, username string, userId int32) (*views.UserProfile, error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	user, err := u.userRepository.GetUserByUsername(ctx, username)
	if err != nil {
		return nil, err
	}

	fullUserData, err := u.userRepository.GetUserProfile(ctx, user.Id)
	if err != nil {
		return nil, err
	}

	// Initialize IsFollowed to false
	fullUserData.IsFollowed = false
	if userId > 0 {
		isFollowed, err := u.userRepository.IsFollowed(ctx, userId, user.Id)
		if err != nil {
			return nil, err
		}
		fullUserData.IsFollowed = isFollowed
	}

	return fullUserData, nil
}

// GetUserById implements UserUsecase.
func (u *userUsecase) GetUserById(ctx context.Context, userId int32) (*views.UserProfile, error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	return u.userRepository.GetUserProfile(ctx, userId)
}

// GetUserByIds implements UserUsecase.
func (u *userUsecase) GetUserByIds(ctx context.Context, userIds []int32) ([]*views.UserProfile, error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	return u.userRepository.GetUserByIds(ctx, userIds)
}

// GetUserFollowedById implements UserUsecase.
func (u *userUsecase) GetUserFollowedById(ctx context.Context, userId int32) ([]int32, error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	return u.userRepository.GetUserFollowedById(ctx, userId)
}

// DeleteUser implements UserUsecase.
func (u *userUsecase) DeleteUser(ctx context.Context, userId int32) error {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	return u.userRepository.DeleteUser(ctx, userId)
}
