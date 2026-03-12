package post

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func (u *PostRepository) HandleDeleteAccount(ctx context.Context, userID int) error {

	sqlPost := `
		UPDATE posts SET deleted_at = NOW()
		WHERE user_id = $1 AND deleted_at IS NULL
	`

	queryLike := `
		UPDATE post_likes SET deleted_at = NOW()
		WHERE user_id = $1 AND deleted_at IS NULL
	`

	return pgx.BeginFunc(ctx, u.db, func(tx pgx.Tx) error {

		_, err := tx.Exec(ctx, sqlPost, userID)
		if err != nil {
			return err
		}

		_, err = tx.Exec(ctx, queryLike, userID)
		if err != nil {
			return err
		}

		return nil

	})
}
