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

type ProfileUsecase interface {
	UpdateProfile(ctx context.Context, userID int, updates *domain.Profile) error
}

func NewProfileUsecase(profileRepository domain.ProfileRepository, contextTimeout time.Duration) ProfileUsecase {
	return &profileUsecase{
		profileRepository: profileRepository,
		contextTimeout:    contextTimeout,
	}
}

func (p *profileUsecase) UpdateProfile(ctx context.Context, userID int, updates *domain.Profile) error {
	ctx, cancel := context.WithTimeout(ctx, p.contextTimeout)
	defer cancel()

	return p.profileRepository.Update(ctx, userID, updates)
}
