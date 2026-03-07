package profile

import (
	"context"
	"voidspace/users/internal/domain"

	"github.com/vhysxl/voidspace/shared/utils/constants"
	"github.com/vhysxl/voidspace/shared/utils/helper"
)

func (p *ProfileRepository) Update(
	ctx context.Context,
	userID int,
	profile *domain.Profile,
) error {
	query := `
        UPDATE user_profile 
        SET display_name = $1, 
            bio = $2, 
            avatar_url = $3, 
            banner_url = $4, 
            location = $5
        WHERE user_id = $6
          AND EXISTS (
              SELECT 1 FROM users 
              WHERE id = $6 AND deleted_at IS NULL
          )
    `

	args := []any{
		helper.NullIfEmpty(profile.DisplayName),
		helper.NullIfEmpty(profile.Bio),
		helper.NullIfEmpty(profile.AvatarUrl),
		helper.NullIfEmpty(profile.BannerUrl),
		helper.NullIfEmpty(profile.Location),
		userID,
	}

	commandTag, err := p.db.Exec(ctx, query, args...)
	if err != nil {
		return err
	}

	if commandTag.RowsAffected() == 0 {
		return constants.ErrUserNotFound
	}

	return nil
}
