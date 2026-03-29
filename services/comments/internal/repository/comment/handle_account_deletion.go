package comment

import (
	"context"
	"time"
)

// HandleAccountDeletion implements [domain.CommentRepository].
func (c *CommentRepository) HandleAccountDeletion(
	ctx context.Context,
	userID int) error {

	query := `
		UPDATE comments 
		SET deleted_at = $1 
		WHERE user_id = $2 
		AND deleted_at IS NULL
	`

	_, err := c.db.Exec(ctx, query, time.Now(), userID)
	if err != nil {
		return err
	}

	return nil
}
