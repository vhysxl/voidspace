package post

import "context"

// SoftDeletePost implements [domain.PostRepository].
func (p *PostRepository) SoftDeletePost(ctx context.Context, postID int) error {
	_, err := p.db.Exec(
		ctx,
		`UPDATE posts SET deleted_at = NOW() WHERE id = $1`,
		postID,
	)
	if err != nil {
		return err
	}

	return nil
}
