package repository

import (
	"context"
	"database/sql"
	"voidspace/users/internal/domain"

	"github.com/go-sql-driver/mysql"
)

type followRepository struct {
	db *sql.DB
}

func NewFollowRepository(db *sql.DB) domain.FollowRepository {
	return &followRepository{
		db: db,
	}
}

// Follow implements domain.FollowRepository.
func (f *followRepository) Follow(ctx context.Context, updates *domain.Follow) error {
	result, err := f.db.ExecContext(
		ctx,
		`INSERT into user_follows 
		(user_id, target_user_id) VALUES (?, ?)`,
		updates.UserID,
		updates.TargetUserID,
	)

	// Handle specific MySQL error such as duplicate entry or foreign key violation
	if err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			switch mysqlErr.Number {
			case 1062:
				return domain.ErrAlreadyFollow
			case 1452:
				return domain.ErrUserNotFound
			}
		}
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return domain.ErrUserNotFound
	}

	return nil
}

// Unfollow implements domain.FollowRepository.
func (f *followRepository) Unfollow(ctx context.Context, updates *domain.Follow) error {
	_, err := f.db.ExecContext(
		ctx,
		`DELETE FROM user_follows WHERE user_id = ? AND target_user_id = ?`,
		updates.UserID,
		updates.TargetUserID,
	)

	if err != nil {
		return err
	}

	return nil
}
