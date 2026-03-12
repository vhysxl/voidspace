package post

import (
	"context"
)

// SoftDeletePosts implements [domain.PostRepository].
func (p *PostRepository) SoftDeletePosts(ctx context.Context, userID int) error {
	_, err := p.db.Exec(
		ctx,
		`UPDATE posts SET deleted_at = NOW() WHERE user_id = $1 AND deleted_at IS NULL`,
		userID,
	)
	if err != nil {
		return err
	}

	return nil
}
