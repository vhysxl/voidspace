package user

import (
	"context"
	"errors"
	"time"
	"voidspace/users/internal/domain"

	"golang.org/x/crypto/bcrypt"
)

func (u *UserUsecase) Register(
	ctx context.Context,
	username string,
	email string,
	password string,
) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return nil, domain.ErrInternalServer
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
		if errors.Is(err, domain.ErrUserExists) {
			return nil, err
		}

		return nil, domain.ErrInternalServer
	}

	return user, nil
}
