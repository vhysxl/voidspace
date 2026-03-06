package follow

import (
	"context"
	"voidspace/users/internal/domain"
)

func (f *FollowRepository) Unfollow(
	ctx context.Context,
	updates *domain.Follow,
) error {

	query := `
		DELETE FROM user_follows 
		WHERE user_id = $1
		AND target_user_id = $2
	`

	cmdTag, err := f.db.Exec(ctx, query, updates.UserID, updates.TargetUserID)
	if err != nil {
		return err
	}

	affected := cmdTag.RowsAffected()
	if affected == 0 {
		return nil
	}

	return nil

}
