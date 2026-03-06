package user

import (
	"context"
	"errors"
	"voidspace/users/internal/domain"
)

func (u *UserUsecase) DeleteUser(
	ctx context.Context,
	userID int,
) error {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	err := u.userRepository.SoftDelete(ctx, userID)
	if err != nil {
		if errors.Is(err, domain.ErrUserNotFound) {
			return err
		}

		return domain.ErrInternalServer
	}

	return nil
}
