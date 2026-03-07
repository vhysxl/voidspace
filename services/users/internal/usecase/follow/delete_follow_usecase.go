package follow

import (
	"context"
	"errors"
	"voidspace/users/internal/domain"

	"github.com/vhysxl/voidspace/shared/utils/constants"
)

func (f *FollowUsecase) Unfollow(
	ctx context.Context,
	authUserID int,
	targetUserID int,
) error {
	ctx, cancel := context.WithTimeout(ctx, f.contextTimeout)
	defer cancel()

	updates := domain.Follow{
		UserID:       authUserID,
		TargetUserID: targetUserID,
	}

	err := f.followRepository.Unfollow(ctx, &updates)
	if err != nil {
		if errors.Is(err, constants.ErrNotFollowing) {
			return err
		}

		return constants.ErrInternalServer
	}

	return nil
}
