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

func NewFollowUsecase(userRepository domain.UserRepository, followRepository domain.FollowRepository, contextTimeout time.Duration) domain.FollowUsecase {
	return &followUsecase{
		userRepository:   userRepository,
		followRepository: followRepository,
		contextTimeout:   contextTimeout,
	}
}

// Follow implements FollowUsecase.
func (f *followUsecase) Follow(ctx context.Context, followerId int32, username string) error {
	ctx, cancel := context.WithTimeout(ctx, f.contextTimeout)
	defer cancel()

	targetUser, err := f.userRepository.GetUserByUsername(ctx, username)
	if err != nil {
		return err
	}

	if followerId == targetUser.Id {
		return domain.ErrSelfFollow
	}

	prepData := domain.Follow{
		UserId:       followerId,
		TargetUserId: targetUser.Id,
	}

	return f.followRepository.Follow(ctx, &prepData)
}

// Unfollow implements FollowUsecase.
func (f *followUsecase) Unfollow(ctx context.Context, followerId int32, username string) error {
	ctx, cancel := context.WithTimeout(ctx, f.contextTimeout)
	defer cancel()

	targetUser, err := f.userRepository.GetUserByUsername(ctx, username)
	if err != nil {
		return err
	}

	prepData := domain.Follow{
		UserId:       followerId,
		TargetUserId: targetUser.Id,
	}

	return f.followRepository.Unfollow(ctx, &prepData)
}
