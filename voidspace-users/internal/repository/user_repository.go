package repository

import (
	"context"
	"database/sql"
	"voidspace/users/internal/domain"
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
	_, err := u.db.ExecContext(
		ctx,
		"INSERT INTO users (username, email, password_hash, created_at, updated_at) VALUES (?, ?, ?, ?, ?)",
		user.Username,
		user.Email,
		user.PasswordHash,
		user.CreatedAt,
		user.UpdatedAt,
	)
	return err
}

// GetUsersByUsername implements domain.UserRepository.
func (u *userRepository) GetUsersByUsername(ctx context.Context, username string) ([]*domain.User, error) {
	panic("unimplemented")
}

// GetUserByEmail implements domain.UserRepository.
func (u *userRepository) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	user := &domain.User{}
	err := u.db.QueryRowContext(ctx,
		"SELECT id, username, email, password_hash, created_at, updated_at FROM users WHERE email = ?",
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
		"SELECT id, username, email, password_hash, created_at, updated_at FROM users WHERE username = ?",
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
func (u *userRepository) UpdateUser(ctx context.Context, user *domain.User) error {
	panic("unimplemented")
}

// DeleteUser implements domain.UserRepository.
func (u *userRepository) DeleteUser(ctx context.Context, id int) error {
	panic("unimplemented")
}
