package post

import (
	"context"
	"time"
	"voidspace/posts/internal/domain"

	"github.com/georgysavva/scany/v2/pgxscan"
)

// GetGlobalFeed implements [domain.PostRepository].
func (p *PostRepository) GetGlobalFeed(ctx context.Context, cursorTime time.Time, cursorID int) ([]domain.Post, bool, error) {
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
		WHERE p.deleted_at IS NULL 
		  AND ((p.created_at < $1) OR (p.created_at = $1 AND p.id < $2))
		ORDER BY p.created_at DESC, p.id DESC
		LIMIT $3
	`

	err := pgxscan.Select(ctx, p.db, &posts, query, cursorTime, cursorID, 10+1)
	if err != nil {
		return nil, false, err
	}

	hasMore := len(posts) > 10
	if hasMore {
		posts = posts[:10]
	}

	return posts, hasMore, nil
}
