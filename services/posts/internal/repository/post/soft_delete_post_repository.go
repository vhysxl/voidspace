package post

import (
	"context"

	"github.com/vhysxl/voidspace/shared/utils/constants"
)

// SoftDeletePost implements [domain.PostRepository].
func (p *PostRepository) SoftDelete(ctx context.Context, postID int) error {
	cmdTag, err := p.db.Exec(
		ctx,
		`UPDATE posts SET deleted_at = NOW() WHERE id = $1
		AND deleted_at IS NULL`,
		postID,
	)
	if err != nil {
		return err
	}

	rowsAffected := cmdTag.RowsAffected()
	if rowsAffected == 0 {
		return constants.ErrPostNotFound
	}

	return nil
}
