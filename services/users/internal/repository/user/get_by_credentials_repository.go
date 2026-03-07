package user

import (
	"context"
	"errors"
	"voidspace/users/internal/domain"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5"
	"github.com/vhysxl/voidspace/shared/utils/constants"
)

// GetByCredentials retrieves a user by their email or username.
//
// The email and username columns use the citext (case-insensitive text) type,
// eliminating the need for LOWER() calls that would prevent index usage and
// cause full table scans.

func (u *UserRepository) GetByCredentials(
	ctx context.Context,
	credentials string,
) (*domain.User, error) {
	user := domain.User{}

	query :=
		`
        SELECT id, username, email, password_hash, created_at, updated_at
        FROM users
        WHERE (email = $1 OR username = $1)
          AND deleted_at IS NULL
        LIMIT 1
    `
	err := pgxscan.Get(ctx, u.db, &user, query, credentials)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, constants.ErrUserNotFound
		}

		return nil, err
	}

	return &user, nil
}
