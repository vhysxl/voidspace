package usecase

import (
	"context"
)

// DeleteByPostID implements [domain.CommentUsecase].
func (c *commentUsecase) DeleteByPostID(ctx context.Context, postID int) error {
	return c.commentRepository.DeleteByPostID(ctx, postID)
}
