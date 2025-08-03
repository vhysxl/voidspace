package usecase

import (
	"context"
	"time"
	"voidspace/users/internal/domain"

	"golang.org/x/crypto/bcrypt"
)

type loginUsecase struct {
	userRepository domain.UserRepository
	contextTimeout time.Duration
}

type LoginUsecase interface {
	Login(ctx context.Context, credentials, password string) (*domain.User, error)
}

func NewLoginUsecase(userRepository domain.UserRepository, contextTimeout time.Duration) LoginUsecase {
	return &loginUsecase{
		userRepository: userRepository,
		contextTimeout: contextTimeout,
	}
}

// Login implements LoginUsecase.
func (l *loginUsecase) Login(ctx context.Context, credentials string, password string) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, l.contextTimeout)
	defer cancel()

	user, err := l.userRepository.GetUserByCredentials(ctx, credentials)
	if err != nil {
		if err == domain.ErrUserNotFound {
			return nil, domain.ErrInvalidCredentials
		}
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return nil, domain.ErrInvalidCredentials
	}

	return user, nil
}
