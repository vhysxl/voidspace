package comment

import (
	"context"
	"voidspace/comments/internal/domain"

	"github.com/georgysavva/scany/v2/pgxscan"
)

// GetAllByPostID implements [domain.CommentRepository].
func (c *CommentRepository) GetAllByPostID(
	ctx context.Context,
	postID int) ([]*domain.Comment, error) {

	var comments []domain.Comment

	query := `
		SELECT 
			c.id,
			c.content,
			c.post_id,
			c.user_id,
			c.created_at
		FROM comments c
		WHERE c.post_id = $1
		AND c.deleted_at IS NULL
		ORDER BY c.created_at DESC
	`

	err := pgxscan.Select(ctx, c.db, &comments, query, postID)
	if err != nil {
		return nil, err
	}

	result := make([]*domain.Comment, len(comments))
	for i := range comments {
		result[i] = &comments[i]
	}

	return result, nil
}
