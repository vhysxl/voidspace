package user

import (
	"context"
	"voidspace/users/internal/domain"
)

func (u *UserRepository) RestoreUser(
	ctx context.Context,
	userID int,
) error {

	query :=
		`
		UPDATE users
		SET deleted_at = NULL
		WHERE id = $1
		AND deleted_at IS NOT NULL
	`

	cmdTag, err := u.db.Exec(ctx, query, userID)
	if err != nil {
		return err
	}

	affected := cmdTag.RowsAffected()
	if affected == 0 {
		return domain.ErrUserNotFound
	}

	return nil
}
