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
		INSERT INTO post_likes (user_id, post_id, created_at, deleted_at)
		VALUES ($1, $2, NOW(), NULL)
		ON CONFLICT (user_id, post_id)
		DO UPDATE SET
			deleted_at = NULL,
			created_at = NOW()
		WHERE post_likes.deleted_at IS NOT NULL
	`
	_, err := l.db.Exec(ctx, query, like.UserID, like.PostID)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == pgerrcode.ForeignKeyViolation {
			return constants.ErrUserOrPostNotFound
		}
		return err
	}
	return nil
}
