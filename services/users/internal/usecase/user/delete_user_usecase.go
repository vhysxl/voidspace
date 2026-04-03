package user

import (
	"context"
	"errors"

	"github.com/vhysxl/voidspace/shared/utils/constants"
)

func (u *UserUsecase) DeleteUser(
	ctx context.Context,
	userID int,
) error {
	err := u.userRepository.SoftDelete(ctx, userID)
	if err != nil {
		if errors.Is(err, constants.ErrUserNotFound) {
			return err
		}

		return constants.ErrInternalServer
	}

	return nil
}
