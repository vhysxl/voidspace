package like

import (
	"context"
	"voidspace/posts/internal/domain"
)

// UnlikePost implements [domain.LikeRepository].
func (l *LikeRepository) UnlikePost(ctx context.Context, like *domain.Like) error {
	query := `
		UPDATE post_likes
		SET deleted_at = NOW()
		WHERE post_id = $1
		AND user_id = $2
		AND deleted_at IS NULL
	`
	_, err := l.db.Exec(
		ctx,
		query,
		like.PostID,
		like.UserID,
	)
	if err != nil {
		return err
	}
	return nil
}
