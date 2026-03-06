package follow

import (
	"context"
	"errors"
	"voidspace/users/internal/domain"
)

func (f *FollowUsecase) Follow(
	ctx context.Context,
	authUserID int,
	targetUserID int,
) error {
	ctx, cancel := context.WithTimeout(ctx, f.contextTimeout)
	defer cancel()

	if authUserID == targetUserID {
		return domain.ErrSelfFollow
	}

	updates := domain.Follow{
		UserID:       authUserID,
		TargetUserID: targetUserID,
	}

	err := f.followRepository.Follow(ctx, &updates)
	if err != nil {
		switch {
		case errors.Is(err, domain.ErrUserNotFound):
			return domain.ErrUserNotFound
		case errors.Is(err, domain.ErrAlreadyFollow):
			return domain.ErrAlreadyFollow
		default:
			return domain.ErrInternalServer
		}
	}

	return nil
}
