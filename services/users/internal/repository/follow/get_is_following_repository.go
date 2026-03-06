package follow

import (
	"context"

	"github.com/georgysavva/scany/v2/pgxscan"
)

func (f *FollowRepository) IsFollowing(
	ctx context.Context,
	userID, targetUserID int,
) (bool, error) {
	var exists bool

	query := `
    SELECT EXISTS (
        SELECT 1 
        FROM user_follows f
        JOIN users u1 ON f.user_id = u1.id
        JOIN users u2 ON f.target_user_id = u2.id
        WHERE f.user_id = $1 
          AND f.target_user_id = $2
          AND u1.deleted_at IS NULL 
          AND u2.deleted_at IS NULL
    )
`
	err := pgxscan.Get(ctx, f.db, &exists, query, userID, targetUserID)
	if err != nil {
		return false, err
	}
	return exists, nil
}
