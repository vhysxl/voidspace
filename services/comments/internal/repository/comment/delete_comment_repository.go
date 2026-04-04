package comment

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/vhysxl/voidspace/shared/utils/constants"
)

// SoftDelete implements [domain.CommentRepository].
func (c *CommentRepository) Delete(
	ctx context.Context,
	commentID int) (int, error) {

	query := `DELETE FROM comments WHERE id = $1 RETURNING id`

	var id int
	err := c.db.QueryRow(ctx, query, commentID).Scan(&id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return 0, constants.ErrCommentNotFound
		}
		return 0, err
	}

	return id, nil
}
