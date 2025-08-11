package usecase

import (
	"context"
	"time"
	"voidspace/users/internal/domain"
)

type followUsecase struct {
	userRepository   domain.UserRepository
	followRepository domain.FollowRepository
	contextTimeout   time.Duration
}

type FollowUsecase interface {
	Follow(ctx context.Context, followingID int, username string) error
	Unfollow(ctx context.Context, followingID int, username string) error
}

func NewFollowUsecase(userRepository domain.UserRepository, followRepository domain.FollowRepository, contextTimeout time.Duration) FollowUsecase {
	return &followUsecase{
		userRepository:   userRepository,
		followRepository: followRepository,
		contextTimeout:   contextTimeout,
	}
}

// Follow implements FollowUsecase.
func (f *followUsecase) Follow(ctx context.Context, userID int, username string) error {
	ctx, cancel := context.WithTimeout(ctx, f.contextTimeout)
	defer cancel()

	targetUser, err := f.userRepository.GetUserByUsername(ctx, username)
	if err != nil {
		return err
	}

	if userID == targetUser.ID {
		return domain.ErrSelfFollow
	}

	prepData := domain.Follow{
		UserID:       userID,
		TargetUserID: targetUser.ID,
	}

	return f.followRepository.Follow(ctx, &prepData)
}

// Unfollow implements FollowUsecase.
func (f *followUsecase) Unfollow(ctx context.Context, userID int, username string) error {
	ctx, cancel := context.WithTimeout(ctx, f.contextTimeout)
	defer cancel()

	targetUser, err := f.userRepository.GetUserByUsername(ctx, username)
	if err != nil {
		return err
	}

	prepData := domain.Follow{
		UserID:       userID,
		TargetUserID: targetUser.ID,
	}

	return f.followRepository.Unfollow(ctx, &prepData)
}
