package user

import (
	"context"
	"errors"
	"voidspace/users/internal/domain"
)

func (u *UserUsecase) RestoreUser(
	ctx context.Context,
	userID int,
) error {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	err := u.userRepository.RestoreUser(ctx, userID)
	if err != nil {
		if errors.Is(err, domain.ErrUserNotFound) {
			return err
		}

		return domain.ErrInternalServer
	}

	return nil
}
