package user

import (
	"context"
	"errors"

	"github.com/vhysxl/voidspace/shared/utils/constants"
)

func (u *UserUsecase) RestoreUser(
	ctx context.Context,
	userID int,
) error {
	err := u.userRepository.RestoreUser(ctx, userID)
	if err != nil {
		if errors.Is(err, constants.ErrUserNotFound) {
			return err
		}

		return constants.ErrInternalServer
	}

	return nil
}
