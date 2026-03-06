package user

import (
	"context"
	"voidspace/users/internal/domain"
	"voidspace/users/internal/domain/views"
)

func (u *UserUsecase) ListFollowing(
	ctx context.Context,
	userID int,
) ([]views.UserBanner, error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	users, err := u.userRepository.ListFollowing(ctx, userID)
	if err != nil {
		return nil, domain.ErrInternalServer
	}

	return users, nil
}
