package user

import (
	"context"
	"errors"
	"voidspace/users/internal/domain/views"

	"github.com/vhysxl/voidspace/shared/utils/constants"
)

func (u *UserUsecase) ListFollowers(
	ctx context.Context,
	userID int,
) ([]views.UserBanner, error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	users, err := u.userRepository.ListFollowers(ctx, userID)
	if err != nil {
		if errors.Is(err, constants.ErrUserNotFound) {
			return nil, err
		}
		return nil, constants.ErrInternalServer
	}

	return users, nil
}
