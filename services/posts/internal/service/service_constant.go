package service

// common error messages used across multiple services
const (
	ErrRequestTimeout    = "Request Timeout"
	ErrInvalidRequest    = "Invalid request"
	ErrInternalServer    = "Internal server error"
	ErrUnauthorized      = "Unauthorized"
	ErrValidation        = "Validation failed"
	ErrUsecase           = "Usecase error"
	ErrFailedGetUserID   = "failed to get user ID from context"
	ErrFailedGetUsername = "failed to get username from context"
)
