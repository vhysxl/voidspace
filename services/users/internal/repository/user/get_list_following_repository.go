package user

import (
	"context"
	"voidspace/users/internal/domain/views"

	"github.com/georgysavva/scany/v2/pgxscan"
)

func (u *UserRepository) ListFollowing(
	ctx context.Context,
	userID int,
) ([]views.UserBanner, error) {
	var users []views.UserBanner
	query := `
		SELECT uf.target_user_id as user_id, 
		u.username, 
		COALESCE(up.display_name, '') AS display_name, 
		COALESCE(up.avatar_url, '') AS avatar_url
		FROM user_follows uf
		JOIN users u ON u.id = uf.target_user_id
        JOIN user_profile up ON up.user_id = u.id
		WHERE uf.user_id = $1
		AND u.deleted_at IS NULL
	`
	err := pgxscan.Select(ctx, u.db, &users, query, userID)
	if err != nil {
		return nil, err
	}

	return users, nil
}
