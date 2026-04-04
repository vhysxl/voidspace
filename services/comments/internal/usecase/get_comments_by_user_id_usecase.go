package usecase

import (
	"context"
	"voidspace/comments/internal/domain"
)

// GetAllCommentsByUserID implements [domain.CommentUsecase].
func (c *commentUsecase) GetAllCommentsByUserID(
	ctx context.Context,
	userID int) (
	domain.CommentRes, error) {

	comments, err := c.commentRepository.GetAllByUserID(ctx, userID)
	if err != nil {
		return domain.CommentRes{}, err
	}

	return domain.CommentRes{
		CommentsCount: len(comments),
		Comments:      comments,
	}, nil
}
