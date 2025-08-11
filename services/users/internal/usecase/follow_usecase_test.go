package usecase

import (
	"context"
	"testing"
	"time"
	"voidspace/users/internal/domain"
	"voidspace/users/internal/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestFollowUsecase_Follow(t *testing.T) {
	t.Run("successful follow", func(t *testing.T) {
		userRepo := mocks.NewMockUserRepository(t)
		followRepo := mocks.NewMockFollowRepository(t)
		usecase := NewFollowUsecase(userRepo, followRepo, time.Second)
		userRepo.On("GetUserByUsername", mock.Anything, "john").
			Return(&domain.User{ID: 1}, nil)
		followRepo.On("Follow", mock.Anything, mock.AnythingOfType("*domain.Follow")).
			Return(nil)

		err := usecase.Follow(context.Background(), 2, "john")

		assert.NoError(t, err)
		userRepo.AssertExpectations(t)
		followRepo.AssertExpectations(t)
	})

	t.Run("user not found", func(t *testing.T) {
		userRepo := mocks.NewMockUserRepository(t)
		followRepo := mocks.NewMockFollowRepository(t)
		usecase := NewFollowUsecase(userRepo, followRepo, time.Second)

		userRepo.On("GetUserByUsername", mock.Anything, "nobody").
			Return(nil, domain.ErrUserNotFound)

		err := usecase.Follow(context.Background(), 2, "nobody")

		assert.Equal(t, domain.ErrUserNotFound, err)
		userRepo.AssertExpectations(t)
	})
}

func TestFollowUsecase_Unfollow(t *testing.T) {
	t.Run("successful unfollow", func(t *testing.T) {
		userRepo := mocks.NewMockUserRepository(t)
		followRepo := mocks.NewMockFollowRepository(t)
		usecase := NewFollowUsecase(userRepo, followRepo, time.Second)
		userRepo.On("GetUserByUsername", mock.Anything, "john").
			Return(&domain.User{ID: 1}, nil)
		followRepo.On("Unfollow", mock.Anything, mock.AnythingOfType("*domain.Follow")).
			Return(nil)

		err := usecase.Unfollow(context.Background(), 2, "john")

		assert.NoError(t, err)
		userRepo.AssertExpectations(t)
		followRepo.AssertExpectations(t)
	})

	t.Run("user not found", func(t *testing.T) {
		userRepo := mocks.NewMockUserRepository(t)
		followRepo := mocks.NewMockFollowRepository(t)
		usecase := NewFollowUsecase(userRepo, followRepo, time.Second)

		userRepo.On("GetUserByUsername", mock.Anything, "nobody").
			Return(nil, domain.ErrUserNotFound)

		err := usecase.Unfollow(context.Background(), 2, "nobody")

		assert.Equal(t, domain.ErrUserNotFound, err)
		userRepo.AssertExpectations(t)
	})
}
