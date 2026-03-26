package like

import (
	"context"

	"github.com/georgysavva/scany/v2/pgxscan"
)

// IsPostsLikedByUser implements [domain.LikeRepository].
func (l *LikeRepository) IsPostsLikedByUser(ctx context.Context, userID int, postIDs []int) (map[int]bool, error) {
	if len(postIDs) == 0 {
		return map[int]bool{}, nil
	}

	query := `
		SELECT post_id
		FROM post_likes
		WHERE user_id = $1
		AND post_id = ANY($2)
		AND deleted_at IS NULL
	`

	var likedPostIDs []int
	err := pgxscan.Select(ctx, l.db, &likedPostIDs, query, userID, postIDs)
	if err != nil {
		return nil, err
	}

	result := make(map[int]bool, len(postIDs))
	for _, postID := range postIDs {
		result[postID] = false
	}

	for _, postID := range likedPostIDs {
		result[postID] = true
	}

	return result, nil
}
