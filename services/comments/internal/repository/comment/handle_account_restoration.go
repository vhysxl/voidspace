package comment

import (
	"context"
)

// HandleAccountRestoration implements [domain.CommentRepository].
func (c *CommentRepository) HandleAccountRestoration(
	ctx context.Context,
	userID int) error {

	query := `
		UPDATE comments 
		SET deleted_at = NULL 
		WHERE user_id = $1 
		AND deleted_at IS NOT NULL
	`

	_, err := c.db.Exec(ctx, query, userID)
	if err != nil {
		return err
	}

	return nil
}
