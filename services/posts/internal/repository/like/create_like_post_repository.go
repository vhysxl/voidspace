package like

import (
	"context"
	"errors"
	"voidspace/posts/internal/domain"

	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/vhysxl/voidspace/shared/utils/constants"
)

// LikePost implements [domain.LikeRepository].
func (l *LikeRepository) LikePost(ctx context.Context, like *domain.Like) error {
	query := `
        INSERT INTO post_likes (user_id, post_id, created_at)
        VALUES ($1, $2, $3)
    `

	cmdTag, err := l.db.Exec(ctx, query, like.UserID, like.PostID, like.CreatedAt)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			switch pgErr.Code {
			case pgerrcode.UniqueViolation:
				return constants.ErrAlreadyLiked
			case pgerrcode.ForeignKeyViolation:
				return constants.ErrUserOrPostNotFound
			}
		}
		return err
	}

	if cmdTag.RowsAffected() == 0 {
		return constants.ErrUserOrPostNotFound
	}

	return nil
}
