package repository

import (
	"context"
	"database/sql"
	"voidspace/users/internal/domain"
)

type profileRepository struct {
	db *sql.DB
}

func ProfileRepository(db *sql.DB) domain.ProfileRepository {
	return &profileRepository{
		db: db,
	}
}

// Update implements domain.ProfileRepository.
func (p *profileRepository) Update(ctx context.Context, userID int, profile *domain.Profile) error {
	panic("unimplemented")
}
