package follow

import (
	"time"
	"voidspace/users/internal/domain"
)

type FollowUsecase struct {
	followRepository domain.FollowRepository
	contextTimeout   time.Duration
}

func NewFollowUsecase(
	followRepository domain.FollowRepository,
	contextTimeout time.Duration,
) domain.FollowUsecase {
	return &FollowUsecase{
		followRepository: followRepository,
		contextTimeout:   contextTimeout,
	}
}
