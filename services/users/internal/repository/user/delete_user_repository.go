package user

import (
	"context"

	"github.com/vhysxl/voidspace/shared/utils/constants"
)

func (u *UserRepository) Delete(
	ctx context.Context,
	userID int,
) error {
	query := `
		DELETE FROM users WHERE id = $1
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
