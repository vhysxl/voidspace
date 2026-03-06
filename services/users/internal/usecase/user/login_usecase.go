package user

import (
	"context"
	"errors"
	"voidspace/users/internal/domain"

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
		if errors.Is(err, domain.ErrUserNotFound) {
			return nil, err
		}

		return nil, domain.ErrInternalServer
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return nil, domain.ErrInvalidCredentials
	}

	return user, nil
}
