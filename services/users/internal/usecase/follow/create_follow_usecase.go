package follow

import (
	"context"
	"errors"
	"voidspace/users/internal/domain"

	"github.com/vhysxl/voidspace/shared/utils/constants"
)

func (f *FollowUsecase) Follow(
	ctx context.Context,
	authUserID int,
	targetUserID int,
) error {
	if authUserID == targetUserID {
		return constants.ErrCannotFollowSelf
	}

	updates := domain.Follow{
		UserID:       authUserID,
		TargetUserID: targetUserID,
	}

	err := f.followRepository.Follow(ctx, &updates)
	if err != nil {
		switch {
		case errors.Is(err, constants.ErrUserNotFound):
			return constants.ErrUserNotFound
		case errors.Is(err, constants.ErrAlreadyFollowing):
			return constants.ErrAlreadyFollowing
		default:
			return constants.ErrInternalServer
		}
	}

	return nil
}
