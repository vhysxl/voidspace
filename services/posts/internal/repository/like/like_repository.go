package like

import (
	"context"
	"voidspace/posts/internal/domain"

	"github.com/jackc/pgx/v5/pgxpool"
)

type LikeRepository struct {
	db *pgxpool.Pool
}

// IsPostLikedByUser implements [domain.LikeRepository].
func (l *LikeRepository) IsPostLikedByUser(ctx context.Context, userID int, postID int) (bool, error) {
	panic("unimplemented")
}

// IsPostsLikedByUser implements [domain.LikeRepository].
func (l *LikeRepository) IsPostsLikedByUser(ctx context.Context, userID int, postIDs []int) (map[int]bool, error) {
	panic("unimplemented")
}

// UnlikePost implements [domain.LikeRepository].
func (l *LikeRepository) UnlikePost(ctx context.Context, like *domain.Like) (int, error) {
	panic("unimplemented")
}

func NewLikeRepository(db *pgxpool.Pool) domain.LikeRepository {
	return &LikeRepository{
		db: db,
	}
}
