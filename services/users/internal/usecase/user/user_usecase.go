package user

import (
	"time"
	"voidspace/users/internal/domain"
)

type UserUsecase struct {
	userRepository   domain.UserRepository
	followRepository domain.FollowRepository
	contextTimeout   time.Duration
}

func NewUserUsecase(
	userRepository domain.UserRepository,
	followRepository domain.FollowRepository,
	contextTimeout time.Duration,
) domain.UserUsecase {
	return &UserUsecase{
		userRepository:   userRepository,
		followRepository: followRepository,
		contextTimeout:   contextTimeout,
	}
}
