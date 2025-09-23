package usecase

import (
	"context"
	"time"
	"voidspace/users/internal/domain"
)

type profileUsecase struct {
	profileRepository domain.ProfileRepository
	contextTimeout    time.Duration
}

func NewProfileUsecase(profileRepository domain.ProfileRepository, contextTimeout time.Duration) domain.ProfileUsecase {
	return &profileUsecase{
		profileRepository: profileRepository,
		contextTimeout:    contextTimeout,
	}
}

func (p *profileUsecase) UpdateProfile(ctx context.Context, userId int32, updates *domain.Profile) error {
	ctx, cancel := context.WithTimeout(ctx, p.contextTimeout)
	defer cancel()

	return p.profileRepository.Update(ctx, userId, updates)
}
