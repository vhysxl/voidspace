package post

import (
	"context"
	"voidspace/posts/internal/domain"
)

// Delete implements [domain.PostRepository].
func (p *PostRepository) Delete(ctx context.Context, postID int) error {
	cmdTag, err := p.db.Exec(
		ctx,
		`DELETE FROM posts WHERE id = $1`,
		postID,
	)
	if err != nil {
		return err
	}

	if cmdTag.RowsAffected() == 0 {
		return domain.ErrPostNotFound
	}

	return nil
}
