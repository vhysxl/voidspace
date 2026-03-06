package user

import (
	"context"
	"errors"
	"voidspace/users/internal/domain"

	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

func (u *UserRepository) Create(
	ctx context.Context,
	user *domain.User,
) error {

	sqlUser := `INSERT INTO users (username, email, password_hash, created_at, updated_at)
			VALUES ($1, $2, $3, $4, $5) RETURNING id`

	sqlUserProfile := `INSERT INTO user_profile (user_id) VALUES ($1)`

	err := pgx.BeginFunc(ctx, u.db, func(tx pgx.Tx) error {
		var userID int

		// Insert User (Create User)
		err := tx.QueryRow(ctx,
			sqlUser,
			user.Username,
			user.Email,
			user.PasswordHash,
			user.CreatedAt,
			user.UpdatedAt,
		).Scan(&userID)
		if err != nil {
			return err
		}

		// Insert UserProfile (Create User Profile)
		_, err = tx.Exec(
			ctx,
			sqlUserProfile,
			userID,
		)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == pgerrcode.UniqueViolation {
				return domain.ErrUserExists
			}
		}
		return err
	}

	return nil
}
