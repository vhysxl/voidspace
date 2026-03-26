package post

import (
	"context"
	"encoding/json"
	"voidspace/posts/internal/domain"

	"github.com/vhysxl/voidspace/shared/utils/constants"
)

// Update implements [domain.PostRepository].
func (p *PostRepository) Update(ctx context.Context, post *domain.Post) error {
	if len(post.PostImages) == 0 {
		post.PostImages = nil
	}

	jsonImages, err := json.Marshal(post.PostImages)
	if err != nil {
		return err
	}

	cmdTag, err := p.db.Exec(
		ctx,
		`UPDATE posts SET content = $1, post_images = $2, updated_at = NOW()
    	WHERE id = $3 AND deleted_at IS NULL`,
		post.Content,
		jsonImages,
		post.ID,
	)
	if err != nil {
		return err
	}

	if cmdTag.RowsAffected() == 0 {
		return constants.ErrPostNotFound
	}

	return nil
}
