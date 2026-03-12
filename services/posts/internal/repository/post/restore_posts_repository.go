package post

import "context"

// RestorePosts implements [domain.PostRepository].
func (p *PostRepository) RestorePosts(ctx context.Context, userID int) error {
	_, err := p.db.Exec(
		ctx,
		`UPDATE posts SET deleted_at = NULL WHERE user_id = $1 AND deleted_at IS NOT NULL`,
		userID,
	)
	if err != nil {
		return err
	}

	return nil
}
