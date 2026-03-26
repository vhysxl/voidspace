package like

import (
	"voidspace/posts/internal/domain"

	"github.com/jackc/pgx/v5/pgxpool"
)

type LikeRepository struct {
	db *pgxpool.Pool
}

func NewLikeRepository(db *pgxpool.Pool) domain.LikeRepository {
	return &LikeRepository{
		db: db,
	}
}
