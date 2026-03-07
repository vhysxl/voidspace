package profile

import (
	"context"
	"errors"
	"voidspace/users/internal/domain"

	"github.com/vhysxl/voidspace/shared/utils/constants"
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
		if errors.Is(err, constants.ErrUserNotFound) {
			return err
		}

		return constants.ErrInternalServer
	}

	return nil
}
