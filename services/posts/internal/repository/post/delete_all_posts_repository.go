package post

import "context"

func (p *PostRepository) DeleteAllPosts(ctx context.Context, userID int) error {
	_, err := p.db.Exec(
		ctx,
		`DELETE FROM posts WHERE user_id = $1`,
		userID,
	)
	if err != nil {
		return err
	}

	return nil
}
