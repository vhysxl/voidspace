package comment

import (
	"context"
	"errors"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/vhysxl/voidspace/shared/utils/constants"
)

// SoftDelete implements [domain.CommentRepository].
func (c *CommentRepository) SoftDelete(
	ctx context.Context,
	commentID int) (int, error) {

	query := `
		UPDATE comments 
		SET deleted_at = $1 
		WHERE id = $2 
		AND deleted_at IS NULL
		RETURNING id
	`

	var id int
	err := c.db.QueryRow(ctx, query, time.Now(), commentID).Scan(&id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return 0, constants.ErrCommentNotFound
		}
		return 0, err
	}

	return id, nil
}
