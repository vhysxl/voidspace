package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"
	"voidspace/users/internal/domain"
	"voidspace/users/internal/domain/views"

	"github.com/go-sql-driver/mysql"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) domain.UserRepository {
	return &userRepository{
		db: db,
	}
}

func (u *userRepository) Create(ctx context.Context, user *domain.User) error {
	tx, err := u.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	result, err := tx.ExecContext(
		ctx,
		`INSERT INTO users 
		(username, email, password_hash, created_at, updated_at) 
		VALUES (?, ?, ?, ?, ?)`,
		user.Username,
		user.Email,
		user.PasswordHash,
		user.CreatedAt,
		user.UpdatedAt,
	)
	if err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			if mysqlErr.Number == 1062 {
				return domain.ErrUserExists
			}
		}
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	user.ID = int(id)

	_, err = tx.ExecContext(
		ctx,
		"INSERT INTO user_profile (user_id) VALUES (?)",
		user.ID,
	)
	if err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

// GetUserByUsername implements domain.UserRepository.
func (u *userRepository) GetUserByUsername(ctx context.Context, username string) (*domain.User, error) {
	user := &domain.User{}
	err := u.db.QueryRowContext(ctx,
		`SELECT id, username, email, password_hash, created_at, updated_at 
		FROM users 
		WHERE username = ?`,
		username).
		Scan(
			&user.ID,
			&user.Username,
			&user.Email,
			&user.PasswordHash,
			&user.CreatedAt,
			&user.UpdatedAt,
		)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, domain.ErrUserNotFound
		}
		return nil, err
	}

	return user, nil
}

// GetUserByIds implements domain.UserRepository.
func (u *userRepository) GetUserByIds(ctx context.Context, userIDs []int32) ([]*views.UserProfile, error) {
	placeholders := make([]string, len(userIDs))
	args := make([]any, len(userIDs))
	for i, id := range userIDs {
		placeholders[i] = "?"
		args[i] = id
	}

	query := fmt.Sprintf(`
    SELECT 
        u.id, 
        u.username, 
        up.display_name, 
        up.avatar_url,
        up.bio, 
        up.banner_url, 
        up.location,
        (SELECT COUNT(*) FROM user_follows WHERE target_user_id = u.id) AS followers,
        (SELECT COUNT(*) FROM user_follows WHERE user_id = u.id) AS following,
        u.created_at
    FROM users u
    JOIN user_profile up ON u.id = up.user_id
    WHERE u.id IN (%s)
`, strings.Join(placeholders, ","))

	rows, err := u.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := make([]*views.UserProfile, 0)
	for rows.Next() {
		var user views.UserProfile
		var displayName sql.NullString
		var bio sql.NullString
		var avatarUrl sql.NullString
		var bannerUrl sql.NullString
		var location sql.NullString

		err := rows.Scan(
			&user.ID,
			&user.Username,
			&displayName,
			&avatarUrl,
			&bio,
			&bannerUrl,
			&location,
			&user.Followers,
			&user.Following,
			&user.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		user.DisplayName = displayName.String
		user.Bio = bio.String
		user.AvatarUrl = avatarUrl.String
		user.BannerUrl = bannerUrl.String
		user.Location = location.String

		users = append(users, &user)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

// GetUserByEmail implements domain.UserRepository.
func (u *userRepository) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	user := &domain.User{}
	err := u.db.QueryRowContext(ctx,
		`SELECT id, username, email, password_hash, created_at, updated_at 
		FROM users 
		WHERE email = ?`,
		email).
		Scan(
			&user.ID,
			&user.Username,
			&user.Email,
			&user.PasswordHash,
			&user.CreatedAt,
			&user.UpdatedAt,
		)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, domain.ErrUserNotFound
		}
		return nil, err
	}

	return user, nil
}

// GetUserByCredentials implements domain.UserRepository.
func (u *userRepository) GetUserByCredentials(ctx context.Context, credentials string) (*domain.User, error) {
	user := &domain.User{}
	err := u.db.QueryRowContext(ctx,
		`SELECT id, username, email, password_hash, created_at, updated_at 
		FROM users 
		WHERE email = ? OR username = ?`,
		credentials, credentials,
	).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.PasswordHash,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, domain.ErrUserNotFound
		}
		return nil, err
	}

	return user, nil
}

// GetUserProfile implements domain.UserRepository.
func (u *userRepository) GetUserProfile(ctx context.Context, ID int) (*views.UserProfile, error) {
	var ( //initializer
		userID      int
		username    string
		createdAt   time.Time
		displayName sql.NullString
		bio         sql.NullString
		avatarUrl   sql.NullString
		bannerUrl   sql.NullString
		location    sql.NullString
		follower    sql.NullInt64
		following   sql.NullInt64
	)

	err := u.db.QueryRowContext(ctx,
		`SELECT 
		u.id,
		u.username,
		u.created_at,
		up.display_name,
		up.bio,
		up.avatar_url,
		up.banner_url,
		up.location,
		(SELECT COUNT(*) FROM user_follows WHERE target_user_id = u.id) AS follower,
		(SELECT COUNT(*) FROM user_follows WHERE user_id = u.id) AS following
	FROM users u
	JOIN user_profile up ON u.id = up.user_id
	WHERE u.id = ?`,
		ID,
	).Scan(
		&userID,
		&username,
		&createdAt,
		&displayName,
		&bio,
		&avatarUrl,
		&bannerUrl,
		&location,
		&follower,
		&following,
	)
	if err != nil {

		if err == sql.ErrNoRows {
			return nil, domain.ErrUserNotFound
		}
		return nil, err
	}

	user := &views.UserProfile{
		ID:          userID,
		Username:    username,
		CreatedAt:   createdAt,
		DisplayName: displayName.String,
		Bio:         bio.String,
		AvatarUrl:   avatarUrl.String,
		BannerUrl:   bannerUrl.String,
		Location:    location.String,
		Following:   int(following.Int64),
		Followers:   int(follower.Int64),
	}

	return user, nil

}

// IsFollowed implements domain.UserRepository.
func (u *userRepository) IsFollowed(ctx context.Context, userID int32, targetUserID int32) (bool, error) {
	var dummy int
	err := u.db.QueryRowContext(ctx,
		`SELECT 1 FROM user_follows WHERE user_id = ? AND target_user_id = ?`,
		userID, targetUserID,
	).Scan(&dummy)

	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}

		return false, err
	}
	return true, nil
}

// GetUserFollowedById implements domain.UserRepository.
func (u *userRepository) GetUserFollowedById(ctx context.Context, userID int32) ([]int32, error) {
	rows, err := u.db.QueryContext(ctx, `SELECT target_user_id FROM user_follows WHERE user_id = ?`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var userIDs []int32
	for rows.Next() {
		var id int32
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		userIDs = append(userIDs, id)
	}

	return userIDs, nil
}

// DeleteUser implements domain.UserRepository.
func (u *userRepository) DeleteUser(ctx context.Context, id int) error {
	result, err := u.db.ExecContext(
		ctx,
		`DELETE FROM users WHERE id = ?`,
		id,
	)
	if err != nil {
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
