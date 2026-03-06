package user

import (
	"context"
	"errors"
	"voidspace/users/internal/domain"
	"voidspace/users/internal/domain/views"
)

func (u *UserUsecase) GetCurrentUser(
	ctx context.Context,
	userID int,
) (*views.UserProfile, error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	user, err := u.userRepository.GetProfile(ctx, userID)
	if err != nil {
		if errors.Is(err, domain.ErrUserNotFound) {
			return nil, err
		}
		return nil, domain.ErrInternalServer
	}

	return user, nil

}
