package constants

import "errors"

// Global error messages
const (
	RequestTimeout    = "Request timeout"
	InvalidRequest    = "Invalid request"
	InternalServer    = "Internal server error"
	Unauthorized      = "Unauthorized"
	Validation        = "Validation failed"
	Usecase           = "Usecase error"
	FailedGetUserID   = "Failed to get user ID from context"
	FailedGetUsername = "Failed to get username from context"
)

// Global errors
var (
	ErrUnauthorized   = errors.New(Unauthorized)
	ErrInvalidData    = errors.New(InvalidRequest)
	ErrInternalServer = errors.New(InternalServer)
)

// User-related errors
var (
	ErrUserNotFound       = errors.New("User not found")
	ErrUserExists         = errors.New("User already exists")
	ErrInvalidUserData    = errors.New("Invalid user data")
	ErrInvalidCredentials = errors.New("Invalid credentials")
)

// Follow-related errors
var (
	ErrAlreadyFollowing = errors.New("Already following this user")
	ErrCannotFollowSelf = errors.New("Cannot follow or unfollow yourself")
	ErrNotFollowing     = errors.New("Not following this user")
)

// Post-like related errors
var (
	ErrAlreadyLiked       = errors.New("already liked this post")
	ErrUserOrPostNotFound = errors.New("user or post not found")
	ErrPostNotFound       = errors.New("post not found")
)

// Comment related errors
var (
	ErrCommentNotFound = errors.New("Comment not found")
)
