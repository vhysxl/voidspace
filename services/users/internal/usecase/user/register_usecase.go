package user

import (
	"context"
	"errors"
	"time"
	"voidspace/users/internal/domain"

	"github.com/vhysxl/voidspace/shared/utils/constants"
	"golang.org/x/crypto/bcrypt"
)

func (u *UserUsecase) Register(
	ctx context.Context,
	username string,
	email string,
	password string,
) (*domain.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return nil, constants.ErrInternalServer
	}

	now := time.Now()

	user := &domain.User{
		Username:     username,
		Email:        email,
		PasswordHash: string(hashedPassword),
		CreatedAt:    now,
		UpdatedAt:    now,
	}

	err = u.userRepository.Create(ctx, user)
	if err != nil {
		if errors.Is(err, constants.ErrUserExists) {
			return nil, err
		}

		return nil, constants.ErrInternalServer
	}

	return user, nil
}
