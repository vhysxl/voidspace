package comment

import (
	"context"
	"voidspace/comments/internal/domain"

	"github.com/georgysavva/scany/v2/pgxscan"
)

func (c *CommentRepository) SearchComments(
	ctx context.Context,
	query string,
) ([]*domain.Comment, error) {
	var comments []domain.Comment

	sqlQuery := `
		SELECT 
			c.id,
			c.content,
			c.post_id,
			c.user_id,
			c.created_at
		FROM comments c
		WHERE c.content ILIKE '%' || $1 || '%'
		AND c.deleted_at IS NULL
		ORDER BY c.created_at DESC
		LIMIT 20
	`

	err := pgxscan.Select(ctx, c.db, &comments, sqlQuery, query)
	if err != nil {
		return nil, err
	}

	result := make([]*domain.Comment, len(comments))
	for i := range comments {
		result[i] = &comments[i]
	}

	return result, nil
}
