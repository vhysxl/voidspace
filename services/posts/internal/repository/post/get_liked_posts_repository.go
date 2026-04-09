package post

import (
	"context"
	"voidspace/posts/internal/domain"

	"github.com/georgysavva/scany/v2/pgxscan"
)

// GetLikedPosts implements [domain.PostRepository].
func (p *PostRepository) GetLikedByUserID(ctx context.Context, userID int) ([]domain.Post, error) {
	var posts []domain.Post

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
		JOIN post_likes pl ON p.id = pl.post_id
		WHERE pl.user_id = $1 
		AND p.deleted_at IS NULL
		AND pl.deleted_at IS NULL
		ORDER BY pl.created_at DESC
	`

	err := pgxscan.Select(ctx, p.db, &posts, query, userID)
	if err != nil {
		return nil, err
	}

	return posts, nil
}
