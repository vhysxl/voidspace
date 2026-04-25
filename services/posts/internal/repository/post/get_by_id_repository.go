package post

import (
	"context"
	"errors"
	"voidspace/posts/internal/domain"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5"
	"github.com/vhysxl/voidspace/shared/utils/constants"
)

// GetByID implements [domain.PostRepository].
func (p *PostRepository) GetByID(ctx context.Context, postID int) (*domain.Post, error) {
	var post domain.Post

	query := `
		SELECT
			p.id,
			p.content,
			p.user_id,
			COALESCE(p.post_images, '[]'::jsonb) AS post_images,
		p.created_at,
			p.updated_at,
			(SELECT COUNT(*) FROM post_likes WHERE post_id = p.id
			AND deleted_at IS NULL
			) AS likes_count
        FROM posts p
		WHERE p.id = $1 AND p.deleted_at IS NULL
	`

	err := pgxscan.Get(ctx, p.db, &post, query, postID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, constants.ErrPostNotFound
		}
		return nil, err
	}

	return &post, nil
}
