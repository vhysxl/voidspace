package usecase

import (
	"context"
	"time"
	"voidspace/users/internal/domain"

	"golang.org/x/crypto/bcrypt"
)

type authUsecase struct {
	userRepository domain.UserRepository
	contextTimeout time.Duration
}

type AuthUsecase interface {
	Login(ctx context.Context, credentials, password string) (*domain.User, error)
	Register(ctx context.Context, username, email, password string) (*domain.User, error)
}

func NewAuthUsecase(userRepository domain.UserRepository, contextTimeout time.Duration) AuthUsecase {
	return &authUsecase{
		userRepository: userRepository,
		contextTimeout: contextTimeout,
	}
}

// Login implements LoginUsecase.
func (a *authUsecase) Login(ctx context.Context, credentials string, password string) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, a.contextTimeout)
	defer cancel()

	user, err := a.userRepository.GetUserByCredentials(ctx, credentials)
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

func (a *authUsecase) Register(ctx context.Context, username string, email string, passwordHash string) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, a.contextTimeout)
	defer cancel()

	user := &domain.User{
		Username:     username,
		Email:        email,
		PasswordHash: passwordHash,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	err := a.userRepository.Create(ctx, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
