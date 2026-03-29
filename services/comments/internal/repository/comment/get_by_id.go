package comment

import (
	"context"
	"errors"
	"voidspace/comments/internal/domain"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5"
	"github.com/vhysxl/voidspace/shared/utils/constants"
)

// GetCommentByID implements [domain.CommentRepository].
func (c *CommentRepository) GetCommentByID(ctx context.Context, commentID int) (*domain.Comment, error) {
	var comment domain.Comment

	query := `
		SELECT 
			c.id,
			c.content,
			c.post_id,
			c.user_id,
			c.created_at
		FROM comments c
		WHERE c.id = $1 
		AND c.deleted_at IS NULL
	`

	err := pgxscan.Get(ctx, c.db, &comment, query, commentID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, constants.ErrCommentNotFound
		}
		return nil, err

	}

	return &comment, nil
}
