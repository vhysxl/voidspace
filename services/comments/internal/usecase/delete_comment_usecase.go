package usecase

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/vhysxl/voidspace/shared/utils/constants"
)

// DeleteComment implements [domain.CommentUsecase].
func (c *commentUsecase) DeleteComment(
	ctx context.Context,
	commentID int, userID int,
) error {
	comment, err := c.commentRepository.GetCommentByID(ctx, commentID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return constants.ErrCommentNotFound
		}

		return err
	}

	if comment.UserID != userID {
		return constants.ErrUnauthorized
	}

	_, err = c.commentRepository.Delete(ctx, commentID)
	if err != nil {
		return err
	}

	return nil
}
