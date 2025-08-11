package usecase

import (
	"context"
	"testing"
	"time"
	"voidspace/users/internal/domain"
	"voidspace/users/internal/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
)

// IMPORTANT: Create fresh mock instance for each test case to avoid:
// 1. State contamination from previous test expectations
// 2. Leftover mock calls interfering with current test
// 3. AssertExpectations failures due to accumulated state

func TestAuthUsecase_Login(t *testing.T) {
	passwordPlain := "supersecret password"
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(passwordPlain), bcrypt.DefaultCost)
	if err != nil {
		t.Fatalf("failed to generate password hash: %v", err)
	}

	t.Run("successful login", func(t *testing.T) {
		userRepo := mocks.NewMockUserRepository(t)
		usecase := NewAuthUsecase(userRepo, time.Second)

		userRepo.On("GetUserByCredentials", mock.Anything, "vhysxl").
			Return(&domain.User{ID: 1, PasswordHash: string(passwordHash)}, nil)

		user, err := usecase.Login(context.Background(), "vhysxl", passwordPlain)

		assert.NoError(t, err)
		assert.NotNil(t, user)
		userRepo.AssertExpectations(t)
	})

	t.Run("wrong password", func(t *testing.T) {
		userRepo := mocks.NewMockUserRepository(t)
		usecase := NewAuthUsecase(userRepo, time.Second)

		userRepo.On("GetUserByCredentials", mock.Anything, "vhysxl").
			Return(&domain.User{ID: 1, PasswordHash: string(passwordHash)}, nil)

		user, err := usecase.Login(context.Background(), "vhysxl", "wrongpassword")

		assert.ErrorIs(t, err, domain.ErrInvalidCredentials)
		assert.Nil(t, user)
		userRepo.AssertExpectations(t)
	})

	t.Run("user not found", func(t *testing.T) {
		userRepo := mocks.NewMockUserRepository(t)
		usecase := NewAuthUsecase(userRepo, time.Second)

		userRepo.On("GetUserByCredentials", mock.Anything, "nonexistentuser").
			Return(nil, domain.ErrInvalidCredentials)

		user, err := usecase.Login(context.Background(), "nonexistentuser", "secretpassword")

		assert.ErrorIs(t, err, domain.ErrInvalidCredentials)
		assert.Nil(t, user)
		userRepo.AssertExpectations(t)
	})
}

func TestAuthUsecase_Register(t *testing.T) {

	email := "test@gmail.com"
	username := "supercoolusername"
	passwordPlain := "supersecret password"

	t.Run("successful register", func(t *testing.T) {
		userRepo := mocks.NewMockUserRepository(t)
		usecase := NewAuthUsecase(userRepo, time.Second)
		userRepo.On("Create", mock.Anything, mock.AnythingOfType("*domain.User")).Return(nil)

		user, err := usecase.Register(context.Background(), username, email, passwordPlain)

		assert.NoError(t, err)
		assert.NotNil(t, user)
		userRepo.AssertExpectations(t)
	})

	t.Run("user exist", func(t *testing.T) {
		userRepo := mocks.NewMockUserRepository(t)
		usecase := NewAuthUsecase(userRepo, time.Second)
		userRepo.On("Create", mock.Anything, mock.AnythingOfType("*domain.User")).
			Return(domain.ErrUserExists)

		user, err := usecase.Register(context.Background(), username, email, passwordPlain)

		assert.ErrorIs(t, err, domain.ErrUserExists)
		assert.Nil(t, user)
		userRepo.AssertExpectations(t)
	})
}
