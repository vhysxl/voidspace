package post

import (
	"context"

	"github.com/vhysxl/voidspace/shared/utils/constants"
)

// SoftDeletePost implements [domain.PostRepository].
func (p *PostRepository) Delete(ctx context.Context, postID int) error {
	cmdTag, err := p.db.Exec(
		ctx,
		`DELETE FROM posts  WHERE id = $1`,
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
