package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"voidspace/users/internal/domain"
	"voidspace/users/utils/common"
)

type profileRepository struct {
	db *sql.DB
}

func ProfileRepository(db *sql.DB) domain.ProfileRepository {
	return &profileRepository{
		db: db,
	}
}

func NewProfileRepository(db *sql.DB) domain.ProfileRepository {
	return &profileRepository{
		db: db,
	}
}

// GetProfileById implements domain.ProfileRepository.
func (p *profileRepository) GetProfileById(ctx context.Context, userID int) (*domain.Profile, error) {
	profile := &domain.Profile{}
	err := p.db.QueryRowContext(ctx,
		`SELECT user_id, display_name, bio, avatar_url, banner_url, location 
		FROM user_profile 
		WHERE user_id = ?`,
		userID).
		Scan(
			&profile.UserID,
			&profile.DisplayName,
			&profile.Bio,
			&profile.AvatarUrl,
			&profile.BannerUrl,
			&profile.Location,
		)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, domain.ErrUserNotFound
		}
		return nil, err
	}

	return profile, nil
}

// Update implements domain.ProfileRepository.
func (p *profileRepository) Update(ctx context.Context, userID int, profile *domain.Profile) error {
	setClauses := []string{
		"display_name = ?",
		"bio = ?",
		"avatar_url = ?",
		"banner_url = ?",
		"location = ?",
	}

	args := []any{
		common.NullIfEmpty(profile.DisplayName),
		common.NullIfEmpty(profile.Bio),
		common.NullIfEmpty(profile.AvatarUrl),
		common.NullIfEmpty(profile.BannerUrl),
		common.NullIfEmpty(profile.Location),
		userID,
	}

	query := fmt.Sprintf("UPDATE user_profile SET %s WHERE user_id = ?", strings.Join(setClauses, ", "))
	_, err := p.db.ExecContext(ctx, query, args...)
	return err
}
