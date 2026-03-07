package user

import (
	"context"
	"errors"
	"voidspace/users/internal/domain"

	"github.com/vhysxl/voidspace/shared/utils/constants"
	"golang.org/x/crypto/bcrypt"
)

func (u *UserUsecase) Login(
	ctx context.Context,
	credentials string,
	password string,
) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	user, err := u.userRepository.GetByCredentials(ctx, credentials)
	if err != nil {
		if errors.Is(err, constants.ErrUserNotFound) {
			return nil, err
		}

		return nil, constants.ErrInternalServer
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return nil, constants.ErrInvalidCredentials
	}

	return user, nil
}
