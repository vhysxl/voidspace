package post

import (
	"context"
	"voidspace/posts/internal/domain"

	"github.com/georgysavva/scany/v2/pgxscan"
)

func (p *PostRepository) SearchPosts(
	ctx context.Context,
	query string,
) ([]domain.Post, error) {
	var posts []domain.Post

	sqlQuery := `
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
		WHERE p.content ILIKE '%' || $1 || '%'
		AND p.deleted_at IS NULL
		ORDER BY p.created_at DESC
		LIMIT 20
	`

	err := pgxscan.Select(ctx, p.db, &posts, sqlQuery, query)
	if err != nil {
		return nil, err
	}

	return posts, nil
}
