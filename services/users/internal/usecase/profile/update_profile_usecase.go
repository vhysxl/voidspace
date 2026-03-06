package profile

import (
	"context"
	"errors"
	"voidspace/users/internal/domain"
)

func (p *ProfileUsecase) UpdateProfile(
	ctx context.Context,
	userID int,
	updates *domain.Profile,
) error {
	ctx, cancel := context.WithTimeout(ctx, p.contextTimeout)
	defer cancel()

	err := p.profileRepository.Update(ctx, userID, updates)
	if err != nil {
		if errors.Is(err, domain.ErrUserNotFound) {
			return domain.ErrUserNotFound
		}

		return domain.ErrInternalServer
	}

	return nil
}
