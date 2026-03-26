package post

import (
	"context"
	"voidspace/posts/internal/domain"

	"github.com/georgysavva/scany/v2/pgxscan"
)

// GetAllUserPosts implements [domain.PostRepository].
func (p *PostRepository) GetByUserID(ctx context.Context, userID int) ([]domain.Post, error) {
	var posts []domain.Post

	query := `
		SELECT 
			p.id, 
			p.content, 
			p.user_id, 
			COALESCE(p.post_images, '[]'::jsonb) AS post_images, 
			p.created_at, 
			p.updated_at,
			(SELECT COUNT(*) FROM post_likes WHERE post_id = p.id) AS likes_count
		FROM posts p
		WHERE p.user_id = $1 AND p.deleted_at IS NULL
		ORDER BY p.created_at DESC
	`

	err := pgxscan.Select(ctx, p.db, &posts, query, userID)
	if err != nil {
		return nil, err
	}

	return posts, nil
}
