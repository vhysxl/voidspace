package profile

import (
	"time"
	"voidspace/users/internal/domain"
)

type ProfileUsecase struct {
	profileRepository domain.ProfileRepository
	contextTimeout    time.Duration
}

func NewProfileUsecase(
	profileRepository domain.ProfileRepository,
	contextTimeout time.Duration,
) domain.ProfileUsecase {
	return &ProfileUsecase{
		profileRepository: profileRepository,
		contextTimeout:    contextTimeout,
	}
}
