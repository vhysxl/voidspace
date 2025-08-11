package usecase

import (
	"context"
	"testing"
	"time"
	"voidspace/users/internal/domain"
	"voidspace/users/internal/domain/views"
	"voidspace/users/internal/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUserUsecase_GetCurrentUser(t *testing.T) {
	id := 1
	t.Run("Get Current User", func(t *testing.T) {
		userRepo := mocks.NewMockUserRepository(t)
		usecase := NewUserUsecase(userRepo, time.Second)

		userRepo.On("GetUserProfile", mock.Anything, id).
			Return(&views.UserProfile{ID: 1, Username: "someusername", DisplayName: "somedisplayname"}, nil)

		user, err := usecase.GetCurrentUser(context.Background(), id)

		assert.NoError(t, err)
		assert.NotNil(t, user)
		userRepo.AssertExpectations(t)
	})

	t.Run("User not found", func(t *testing.T) {
		userRepo := mocks.NewMockUserRepository(t)
		usecase := NewUserUsecase(userRepo, time.Second)

		userRepo.On("GetUserProfile", mock.Anything, id).
			Return(nil, domain.ErrUserNotFound)

		user, err := usecase.GetCurrentUser(context.Background(), id)

		assert.ErrorIs(t, err, domain.ErrUserNotFound)
		assert.Nil(t, user)
		userRepo.AssertExpectations(t)
	})
}

func TestUserUsecase_GetUser(t *testing.T) {
	username := "verycoolusername"
	t.Run("Get User by Username", func(t *testing.T) {
		userRepo := mocks.NewMockUserRepository(t)
		usecase := NewUserUsecase(userRepo, time.Second)

		userRepo.On("GetUserByUsername", mock.Anything, username).
			Return(&domain.User{ID: 1, Username: username}, nil)

		userRepo.On("GetUserProfile", mock.Anything, 1).
			Return(&views.UserProfile{ID: 1, Username: username, DisplayName: "somedisplayname"}, nil)

		userProfile, err := usecase.GetUser(context.Background(), username)

		assert.NoError(t, err)
		assert.NotNil(t, userProfile)
		userRepo.AssertExpectations(t)
	})

	t.Run("User not found", func(t *testing.T) {
		userRepo := mocks.NewMockUserRepository(t)
		usecase := NewUserUsecase(userRepo, time.Second)

		userRepo.On("GetUserByUsername", mock.Anything, username).
			Return(nil, domain.ErrUserNotFound)

		userProfile, err := usecase.GetUser(context.Background(), username)

		assert.ErrorIs(t, err, domain.ErrUserNotFound)
		assert.Nil(t, userProfile)
		userRepo.AssertExpectations(t)
	})
}

func TestUserUsecase_DeleteUser(t *testing.T) {
	id := 1
	t.Run("Delete User", func(t *testing.T) {
		userRepo := mocks.NewMockUserRepository(t)
		usecase := NewUserUsecase(userRepo, time.Second)

		userRepo.On("DeleteUser", mock.Anything, id).Return(nil)

		err := usecase.DeleteUser(context.Background(), id)

		assert.NoError(t, err)
		userRepo.AssertExpectations(t)
	})

	t.Run("Delete User Not Found", func(t *testing.T) {
		userRepo := mocks.NewMockUserRepository(t)
		usecase := NewUserUsecase(userRepo, time.Second)

		userRepo.On("DeleteUser", mock.Anything, id).Return(domain.ErrUserNotFound)

		err := usecase.DeleteUser(context.Background(), id)

		assert.ErrorIs(t, err, domain.ErrUserNotFound)
		userRepo.AssertExpectations(t)
	})
}
