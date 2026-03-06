package user

import (
	"context"
	"database/sql"
	"errors"
	"voidspace/users/internal/domain"
	"voidspace/users/internal/domain/views"

	"github.com/georgysavva/scany/v2/pgxscan"
)

func (u *UserRepository) GetProfile(
	ctx context.Context,
	userID int,
) (*views.UserProfile, error) {
	var userProfile views.UserProfile

	query :=
		`SELECT 
			u.id,
			u.username,
			u.created_at,
      		COALESCE(up.display_name, '') AS display_name,
      		COALESCE(up.bio, '') AS bio,
      		COALESCE(up.avatar_url, '') AS avatar_url,
      		COALESCE(up.banner_url, '') AS banner_url,
      		COALESCE(up.location, '') AS location,
			(SELECT COUNT(*) FROM user_follows WHERE target_user_id = u.id) AS follower,
			(SELECT COUNT(*) FROM user_follows WHERE user_id = u.id) AS following
		FROM users u
		JOIN user_profile up ON u.id = up.user_id
		WHERE u.id = $1
		AND u.deleted_at IS NULL
		`

	err := pgxscan.Get(ctx, u.db, &userProfile, query, userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrUserNotFound
		}
		return nil, err
	}

	return &userProfile, nil
}
