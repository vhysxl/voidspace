package follow

import (
	"context"
	"errors"
	"voidspace/users/internal/domain"
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
		if errors.Is(err, domain.ErrNotFollowing) {
			return err
		}

		return domain.ErrInternalServer
	}

	return nil
}
