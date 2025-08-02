package usecase

import (
	"context"
	"time"
	"voidspace/users/internal/domain"
)

type registerUsecase struct {
	userRepository domain.UserRepository
	contextTimeout time.Duration
}

type RegisterUsecase interface {
	Register(ctx context.Context, username, email, password string) (*domain.User, error)
}

func NewRegisterUsecase(userRepository domain.UserRepository, contextTimeout time.Duration) RegisterUsecase {
	return &registerUsecase{
		userRepository: userRepository,
		contextTimeout: contextTimeout,
	}
}

// Register implements RegisterUsecase.
func (r *registerUsecase) Register(ctx context.Context, username string, email string, passwordHash string) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var existingUser *domain.User

	existingUser, _ = r.userRepository.GetUserByEmail(ctx, email)
	if existingUser != nil {
		return nil, domain.ErrEmailExists
	}

	existingUser, _ = r.userRepository.GetUserByUsername(ctx, username)
	if existingUser != nil {
		return nil, domain.ErrUserExists
	}

	user := &domain.User{
		Username:     username,
		Email:        email,
		PasswordHash: passwordHash,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	err := r.userRepository.Create(ctx, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
