package post

import (
	"context"
	"time"
	"voidspace/posts/internal/domain"

	"github.com/georgysavva/scany/v2/pgxscan"
)

// GetFollowFeed implements [domain.PostRepository].
func (p *PostRepository) GetFollowFeed(ctx context.Context, userIDs []int, cursorTime time.Time, cursorID int) ([]domain.Post, bool, error) {
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
		WHERE p.user_id = ANY($1) AND p.deleted_at IS NULL
		  AND ((p.created_at < $2) OR (p.created_at = $2 AND p.id < $3))
		ORDER BY p.created_at DESC, p.id DESC
		LIMIT $4
	`

	err := pgxscan.Select(ctx, p.db, &posts, query, userIDs, cursorTime, cursorID, 10+1)
	if err != nil {
		return nil, false, err
	}

	hasMore := len(posts) > 10
	if hasMore {
		posts = posts[:10]
	}

	return posts, hasMore, nil
}
