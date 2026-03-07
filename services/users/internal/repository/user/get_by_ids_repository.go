package user

import (
	"context"
	"fmt"
	"voidspace/users/internal/domain/views"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/vhysxl/voidspace/shared/utils/helper"
)

// TODO: idk what to do with this function
func (u *UserRepository) GetByIDs(
	ctx context.Context,
	userIDs []int,
) ([]views.UserProfile, error) {

	if len(userIDs) == 0 {
		return []views.UserProfile{}, nil
	}

	queryBatch, args := helper.GenerateDBPlaceholders(userIDs)

	users := []views.UserProfile{}

	query := fmt.Sprintf(`
		SELECT 
			u.id, 
			u.username, 
      		COALESCE(up.display_name, '') AS display_name,
      		COALESCE(up.bio, '') AS bio,
      		COALESCE(up.avatar_url, '') AS avatar_url,
      		COALESCE(up.banner_url, '') AS banner_url,
      		COALESCE(up.location, '') AS location,
			(SELECT COUNT(*) FROM user_follows WHERE target_user_id = u.id) AS follower,
			(SELECT COUNT(*) FROM user_follows WHERE user_id = u.id) AS following,
			u.created_at
		FROM users u
		JOIN user_profile up ON u.id = up.user_id
		WHERE u.id IN (%s)
		AND u.deleted_at IS NULL
	`, queryBatch)

	err := pgxscan.Select(ctx, u.db, &users, query, args...)
	if err != nil {
		return nil, err
	}

	return users, nil
}
