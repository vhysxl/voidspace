package user

import (
	"context"
	"errors"
	"voidspace/users/internal/domain"
	"voidspace/users/internal/domain/views"
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
		if errors.Is(err, domain.ErrUserNotFound) {
			return nil, err
		}

		return nil, domain.ErrInternalServer
	}

	exist, err := u.followRepository.IsFollowing(ctx, authUserID, user.ID)
	if err != nil {
		return nil, domain.ErrInternalServer
	}

	user.IsFollowed = exist

	return user, nil
}
