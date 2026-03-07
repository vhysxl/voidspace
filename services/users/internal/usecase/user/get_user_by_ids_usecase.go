package user

import (
	"context"
	"voidspace/users/internal/domain/views"

	"github.com/vhysxl/voidspace/shared/utils/constants"
)

func (u *UserUsecase) GetUserByIDs(
	ctx context.Context,
	userIDs []int,
) ([]views.UserProfile, error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	user, err := u.userRepository.GetByIDs(ctx, userIDs)
	if err != nil {
		return nil, constants.ErrInternalServer
	}

	return user, nil
}
