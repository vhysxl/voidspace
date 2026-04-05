package comment

import (
	"context"
)

// DeleteByPostID implements [domain.CommentRepository].
func (c *CommentRepository) DeleteByPostID(ctx context.Context, postID int) error {
	query := `
		UPDATE comments 
		SET deleted_at = NOW() 
		WHERE post_id = $1 
		AND deleted_at IS NULL
	`

	_, err := c.db.Exec(ctx, query, postID)
	if err != nil {
		return err
	}

	return nil
}
