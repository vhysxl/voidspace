package follow

import (
	"context"
	"errors"
	"voidspace/users/internal/domain"

	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
)

func (f *FollowRepository) Follow(
	ctx context.Context,
	updates *domain.Follow,
) error {
	query := `
        INSERT INTO user_follows (user_id, target_user_id)
        SELECT $1, $2
        WHERE EXISTS (
            SELECT 1 FROM users WHERE id = $2 AND deleted_at IS NULL
        )
    `

	cmdTag, err := f.db.Exec(ctx, query, updates.UserID, updates.TargetUserID)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == pgerrcode.UniqueViolation {
				return domain.ErrAlreadyFollow
			}
		}
		return err
	}

	if cmdTag.RowsAffected() == 0 {
		return domain.ErrUserNotFound
	}

	return nil
}
