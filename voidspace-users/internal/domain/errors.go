package domain

import "errors"

var (
	ErrUserNotFound    = errors.New("user not found")
	ErrUserExists      = errors.New("user already exists")
	ErrInvalidUserData = errors.New("invalid user data")
	ErrEmailExists     = errors.New("email already registered")
)
