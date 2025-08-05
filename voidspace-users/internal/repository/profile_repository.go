package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
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
	setClauses := []string{}
	args := []any{}

	if profile.DisplayName != nil {
		setClauses = append(setClauses, "display_name = ?")
		args = append(args, *profile.DisplayName)
	}
	if profile.Bio != nil {
		setClauses = append(setClauses, "bio = ?")
		args = append(args, *profile.Bio)
	}
	if profile.AvatarUrl != nil {
		setClauses = append(setClauses, "avatar_url = ?")
		args = append(args, *profile.AvatarUrl)
	}
	if profile.BannerUrl != nil {
		setClauses = append(setClauses, "banner_url = ?")
		args = append(args, *profile.BannerUrl)
	}
	if profile.Location != nil {
		setClauses = append(setClauses, "location = ?")
		args = append(args, *profile.Location)
	}

	if len(setClauses) == 0 {
		// Nothing to update
		return nil
	}

	query := fmt.Sprintf("UPDATE user_profile SET %s WHERE user_id = ?", strings.Join(setClauses, ", "))
	args = append(args, userID)

	_, err := p.db.ExecContext(ctx, query, args...)
	return err
}
