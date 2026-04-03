package like

import (
	"context"
	"voidspace/posts/internal/domain"
)

// LikePost implements [domain.LikeUsecase].
func (l *likeUsecase) LikePost(ctx context.Context, like *domain.Like) error {
	err := l.likeRepository.LikePost(ctx, like)
	if err != nil {
		return err
	}

	return nil

}
