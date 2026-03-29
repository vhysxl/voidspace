package comment

import (
	"context"
	"voidspace/comments/internal/domain"
)

// Create implements [domain.CommentRepository].
func (c *CommentRepository) Create(
	ctx context.Context,
	comment *domain.Comment) (*domain.Comment, error) {

	query := `
		INSERT INTO comments (user_id, post_id, content)
		VALUES ($1, $2, $3)
		RETURNING id, user_id, post_id, content
	`

	err := c.db.QueryRow(
		ctx,
		query,
		comment.UserID,
		comment.PostID,
		comment.Content,
	).Scan(
		&comment.ID,
		&comment.UserID,
		&comment.PostID,
		&comment.Content,
	)
	if err != nil {
		return nil, err
	}

	return comment, nil

}
