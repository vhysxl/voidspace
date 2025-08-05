package usecase

import (
	"context"
	"time"
	"voidspace/users/internal/domain"
)

type profileUsecase struct {
	userRepository    domain.UserRepository
	profileRepository domain.ProfileRepository
	contextTimeout    time.Duration
}

type ProfileUsecase interface {
	UpdateProfile(ctx context.Context, userID int, updates *domain.Profile) error
}

func NewProfileUsecase(profileRepository domain.ProfileRepository, userRepository domain.UserRepository, contextTimeout time.Duration) ProfileUsecase {
	return &profileUsecase{
		userRepository:    userRepository,
		profileRepository: profileRepository,
		contextTimeout:    contextTimeout,
	}
}

func (p *profileUsecase) UpdateProfile(ctx context.Context, userID int, updates *domain.Profile) error {
	ctx, cancel := context.WithTimeout(ctx, p.contextTimeout)
	defer cancel()

	return p.profileRepository.Update(ctx, userID, updates)
}
