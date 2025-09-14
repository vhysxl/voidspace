package domain

import "errors"

var (
	ErrCommentsNotFound   = errors.New("comments not found")
	ErrUnauthorizedAction = errors.New("unauthorized action")
	ErrUserNotFound       = errors.New("failed to get user ID from context")
)
