package user

import (
	"context"

	"github.com/vhysxl/voidspace/shared/utils/constants"
)

func (u *UserRepository) SoftDelete(
	ctx context.Context,
	userID int,
) error {
	query := `
		UPDATE users 
		SET deleted_at = NOW()
		WHERE id = $1 
	`

	cmdTag, err := u.db.Exec(ctx, query, userID)
	if err != nil {
		return err
	}

	affected := cmdTag.RowsAffected()
	if affected == 0 {
		return constants.ErrUserNotFound
	}

	return nil
}
