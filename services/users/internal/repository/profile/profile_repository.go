package profile

import (
	"voidspace/users/internal/domain"

	"github.com/jackc/pgx/v5/pgxpool"
)

type ProfileRepository struct {
	db *pgxpool.Pool
}

func NewProfileRepository(db *pgxpool.Pool) domain.ProfileRepository {
	return &ProfileRepository{
		db: db,
	}
}
