package like

import "context"

// IsPostLikedByUser implements [domain.LikeRepository].
func (l *LikeRepository) IsPostLikedByUser(ctx context.Context, userID int, postID int) (bool, error) {
	query := `
		SELECT EXISTS(
			SELECT 1 FROM post_likes
			WHERE user_id = $1
			AND post_id = $2
			AND deleted_at IS NULL
		)
	`
	var isLiked bool
	err := l.db.QueryRow(ctx, query, userID, postID).Scan(&isLiked)
	if err != nil {
		return false, err
	}
	return isLiked, nil
}
