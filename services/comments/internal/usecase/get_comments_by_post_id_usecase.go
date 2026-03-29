package usecase

import (
	"context"
	"voidspace/comments/internal/domain"
)

// GetAllCommentsByPostID implements [domain.CommentUsecase].
func (c *commentUsecase) GetAllCommentsByPostID(
	ctx context.Context,
	postID int) (
	domain.CommentRes, error) {
	ctx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()

	comments, err := c.commentRepository.GetAllByPostID(ctx, postID)
	if err != nil {
		return domain.CommentRes{}, err
	}

	return domain.CommentRes{
		CommentsCount: len(comments),
		Comments:      comments,
	}, nil
}
