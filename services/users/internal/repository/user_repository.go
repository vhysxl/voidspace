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
		if mysqlErr, ok := err.(*mysql.MySQLError); ok && mysqlErr.Number == 1062 {
			return domain.ErrUserExists
		}
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	user.Id = int32(id) // Fixed: convert to int32

	_, err = tx.ExecContext(
		ctx,
		"INSERT INTO user_profile (user_id) VALUES (?)",
		user.Id,
	)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (u *userRepository) GetUserByUsername(ctx context.Context, username string) (*domain.User, error) {
	user := &domain.User{}
	err := u.db.QueryRowContext(ctx,
		`SELECT id, username, email, password_hash, created_at, updated_at 
		 FROM users 
		 WHERE username = ?`,
		username,
	).Scan(
		&user.Id,
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

func (u *userRepository) GetUserByIds(ctx context.Context, userIDs []int32) ([]*views.UserProfile, error) {
	if len(userIDs) == 0 {
		return []*views.UserProfile{}, nil
	}

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
		var displayName, bio, avatarUrl, bannerUrl, location sql.NullString

		err := rows.Scan(
			&user.Id,
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

func (u *userRepository) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	user := &domain.User{}
	err := u.db.QueryRowContext(ctx,
		`SELECT id, username, email, password_hash, created_at, updated_at 
		 FROM users 
		 WHERE email = ?`,
		email,
	).Scan(
		&user.Id,
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

func (u *userRepository) GetUserByCredentials(ctx context.Context, credentials string) (*domain.User, error) {
	user := &domain.User{}
	err := u.db.QueryRowContext(ctx,
		`SELECT id, username, email, password_hash, created_at, updated_at 
		 FROM users 
		 WHERE email = ? OR username = ?`,
		credentials, credentials,
	).Scan(
		&user.Id,
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

func (u *userRepository) GetUserProfile(ctx context.Context, userId int32) (*views.UserProfile, error) {
	var (
		id          int32
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
		userId,
	).Scan(
		&id,
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

	return &views.UserProfile{
		Id:          id,
		Username:    username,
		CreatedAt:   createdAt,
		DisplayName: displayName.String,
		Bio:         bio.String,
		AvatarUrl:   avatarUrl.String,
		BannerUrl:   bannerUrl.String,
		Location:    location.String,
		Following:   int32(follower.Int64),
		Followers:   int32(follower.Int64),
	}, nil
}

func (u *userRepository) IsFollowed(ctx context.Context, userId, targetUserId int32) (bool, error) {
	var dummy int
	err := u.db.QueryRowContext(ctx,
		`SELECT 1 FROM user_follows WHERE user_id = ? AND target_user_id = ?`,
		userId, targetUserId,
	).Scan(&dummy)

	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (u *userRepository) GetUserFollowedById(ctx context.Context, userId int32) ([]int32, error) {
	rows, err := u.db.QueryContext(ctx, `SELECT target_user_id FROM user_follows WHERE user_id = ?`, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var userIDs []int32 // Fixed: return int32 slice
	for rows.Next() {
		var id int32 // Fixed: use int32
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		userIDs = append(userIDs, id)
	}

	return userIDs, nil
}

func (u *userRepository) DeleteUser(ctx context.Context, userId int32) error {
	result, err := u.db.ExecContext(
		ctx,
		`DELETE FROM users WHERE id = ?`,
		userId,
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
