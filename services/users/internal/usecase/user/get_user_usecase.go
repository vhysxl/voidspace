package user

import (
	"context"
	"errors"
	"voidspace/users/internal/domain/views"

	"github.com/vhysxl/voidspace/shared/utils/constants"
)

func (u *UserUsecase) GetUser(
	ctx context.Context,
	username string,
	authUserID int,
) (*views.UserProfile, error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	user, err := u.userRepository.GetByUsername(ctx, username)
	if err != nil {
		if errors.Is(err, constants.ErrUserNotFound) {
			return nil, err
		}

		return nil, constants.ErrInternalServer
	}

	exist, err := u.followRepository.IsFollowing(ctx, authUserID, user.ID)
	if err != nil {
		return nil, constants.ErrInternalServer
	}

	user.IsFollowed = exist

	return user, nil
}
