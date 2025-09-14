package domain

import "errors"

var (
	ErrUserNotFound       = errors.New("user not found")
	ErrUserExists         = errors.New("user already exists")
	ErrInvalidUserData    = errors.New("invalid user data")
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrUnauthorizedAction = errors.New("unauthorized action")
	ErrAlreadyFollow      = errors.New("follow relationship already exists")
	ErrSelfFollow         = errors.New("self-follow is not allowed")
)
