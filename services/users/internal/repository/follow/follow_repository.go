package follow

import (
	"voidspace/users/internal/domain"

	"github.com/jackc/pgx/v5/pgxpool"
)

type FollowRepository struct {
	db *pgxpool.Pool
}

func NewFollowRepository(db *pgxpool.Pool) domain.FollowRepository {
	return &FollowRepository{
		db: db,
	}
}
