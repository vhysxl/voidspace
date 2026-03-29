package usecase

import (
	"context"
	"voidspace/comments/internal/domain"
)

// CreateComment implements [domain.CommentUsecase].
func (c *commentUsecase) CreateComment(
	ctx context.Context,
	comment *domain.Comment) (*domain.Comment, error) {
	ctx, cancel := context.WithTimeout(ctx, c.contextTimeout)
	defer cancel()

	comment, err := c.commentRepository.Create(ctx, comment)
	if err != nil {
		return nil, err
	}

	return comment, nil
}
