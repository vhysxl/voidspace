package like

import (
	"context"
	"voidspace/posts/internal/domain"
)

// LikePost implements [domain.LikeUsecase].
func (l *likeUsecase) LikePost(ctx context.Context, like *domain.Like) error {
	ctx, cancel := context.WithTimeout(ctx, l.contextTimeout)
	defer cancel()

	err := l.likeRepository.LikePost(ctx, like)
	if err != nil {
		return err
	}

	return nil

}
