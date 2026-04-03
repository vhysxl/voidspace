package like

import (
	"context"
	"voidspace/posts/internal/domain"
)

// UnlikePost implements [domain.LikeUsecase].
func (l *likeUsecase) UnlikePost(ctx context.Context, like *domain.Like) error {
	err := l.likeRepository.UnlikePost(ctx, like)
	if err != nil {
		return err
	}

	return nil

}
