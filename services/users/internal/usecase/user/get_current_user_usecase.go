package user

import (
	"context"
	"errors"
	"voidspace/users/internal/domain/views"

	"github.com/vhysxl/voidspace/shared/utils/constants"
)

func (u *UserUsecase) GetCurrentUser(
	ctx context.Context,
	userID int,
) (*views.UserProfile, error) {
	user, err := u.userRepository.GetProfile(ctx, userID)
	if err != nil {
		if errors.Is(err, constants.ErrUserNotFound) {
			return nil, err
		}
		return nil, constants.ErrInternalServer
	}

	return user, nil

}
