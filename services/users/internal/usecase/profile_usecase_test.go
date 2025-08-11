package usecase

import (
	"context"
	"errors"
	"testing"
	"time"
	"voidspace/users/internal/domain"
	"voidspace/users/internal/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestProfileUsecase_Update(t *testing.T) {
	userID := 1
	testUpdates := &domain.Profile{
		DisplayName: "vhysxl",
		Location:    "Jakarta",
		Bio:         "Very cool Bio",
	}

	t.Run("successful update", func(t *testing.T) {
		profileRepo := mocks.NewMockProfileRepository(t)
		usecase := NewProfileUsecase(profileRepo, time.Second)

		profileRepo.On("Update", mock.Anything, userID, testUpdates).
			Return(nil)

		err := usecase.UpdateProfile(context.Background(), userID, testUpdates)

		assert.NoError(t, err)
		profileRepo.AssertExpectations(t)
	})

	t.Run("failed_update", func(t *testing.T) {
		profileRepo := mocks.NewMockProfileRepository(t)
		usecase := NewProfileUsecase(profileRepo, time.Second)

		profileRepo.On("Update", mock.Anything, userID, testUpdates).
			Return(errors.New("update failed"))

		err := usecase.UpdateProfile(context.Background(), userID, testUpdates)

		assert.Error(t, err)
		profileRepo.AssertExpectations(t)
	})
}
