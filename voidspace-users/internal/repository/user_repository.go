package repository

import (
	"context"
	"database/sql"
	"voidspace/users/internal/domain"
	"voidspace/users/internal/domain/views"
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

// GetUserProfile implements domain.UserRepository.
func (u *userRepository) GetUserProfile(ctx context.Context, ID int) (*views.UserProfile, error) {
	var ( //initializer
		userID      int
		username    string
		displayName sql.NullString
		bio         sql.NullString
		avatarUrl   sql.NullString
		bannerUrl   sql.NullString
		location    sql.NullString
	)

	err := u.db.QueryRowContext(ctx,
		`SELECT u.id ,u.username, up.display_name, up.bio, up.avatar_url, up.banner_url, up.location
    FROM users u
    JOIN user_profile up ON u.id = up.user_id 
    WHERE u.id = ?`,
		ID,
	).Scan(
		&userID,
		&username,
		&displayName,
		&bio,
		&avatarUrl,
		&bannerUrl,
		&location,
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
		DisplayName: displayName.String,
		Bio:         bio.String,
		AvatarUrl:   avatarUrl.String,
		BannerUrl:   bannerUrl.String,
		Location:    location.String,
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

// GetUserByID implements domain.UserRepository.
func (u *userRepository) GetUserByID(ctx context.Context, id int) (*domain.User, error) {
	panic("unimplemented")
}

// UpdateUser implements domain.UserRepository.
func (u *userRepository) UpdateUser(ctx context.Context, user *views.UserProfile, ID int) error {
	_, err := u.db.ExecContext(
		ctx,
		`UPDATE user_profile SET 
         display_name = ?, 
         bio = ?, 
         avatar_url = ?, 
         banner_url = ?, 
         location = ?, 
         updated_at = NOW() 
         WHERE user_id = ?`,
		user.DisplayName,
		user.Bio,
		user.AvatarUrl,
		user.BannerUrl,
		user.Location,
		ID,
	)
	return err
}

// DeleteUser implements domain.UserRepository.
func (u *userRepository) DeleteUser(ctx context.Context, id int) error {
	panic("unimplemented")
}
