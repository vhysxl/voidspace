package user

import (
	"context"
	"voidspace/users/internal/domain/views"

	"github.com/georgysavva/scany/v2/pgxscan"
)

func (u *UserRepository) SearchUsers(
	ctx context.Context,
	query string,
) ([]views.UserBanner, error) {
	var users []views.UserBanner

	sqlQuery := `
		SELECT 
			u.id as user_id, 
			u.username, 
			COALESCE(up.display_name, '') as display_name, 
			COALESCE(up.avatar_url, '') as avatar_url
		FROM users u
		JOIN user_profile up ON u.id = up.user_id
		WHERE (u.username ILIKE '%' || $1 || '%' OR up.display_name ILIKE '%' || $1 || '%')
		AND u.deleted_at IS NULL
		LIMIT 20
	`

	err := pgxscan.Select(ctx, u.db, &users, sqlQuery, query)
	if err != nil {
		return nil, err
	}

	return users, nil
}
