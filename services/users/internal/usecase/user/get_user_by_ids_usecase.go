package user

import (
	"context"
	"fmt"
	"voidspace/users/internal/domain"
	"voidspace/users/internal/domain/views"
)

func (u *UserUsecase) GetUserByIDs(
	ctx context.Context,
	userIDs []int,
) ([]views.UserProfile, error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	fmt.Println("user ids: ", userIDs)

	user, err := u.userRepository.GetByIDs(ctx, userIDs)
	if err != nil {
		return nil, domain.ErrInternalServer
	}

	return user, nil
}
