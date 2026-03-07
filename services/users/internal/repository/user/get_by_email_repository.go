package user

import (
	"context"
	"database/sql"
	"voidspace/users/internal/domain"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/vhysxl/voidspace/shared/utils/constants"
)

func (u *UserRepository) GetByEmail(
	ctx context.Context,
	email string,
) (*domain.User, error) {
	// TODO:  IDK WHAT TO DO WITH THIS JUST KEEP IT HERE IN CASE I NEED SOMETHING

	var user domain.User

	query :=
		`SELECT id, username, email, password_hash, created_at, updated_at
	 FROM users
	 WHERE email = $1
	 AND deleted_at IS NULL
	`

	err := pgxscan.Get(ctx, u.db, &user, query, email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, constants.ErrUserNotFound
		}
		return nil, err
	}

	return &user, nil
}
