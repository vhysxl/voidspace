package error

import (
	"context"
	"errors"
	"voidspace/users/internal/domain"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

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

func HandleError(err error, logger *zap.Logger, operation string) error {
	if err == nil {
		return nil
	}

	logger.Error(ErrUsecase, zap.String("OPERATION", operation), zap.Error(err))

	switch {
	case errors.Is(err, context.DeadlineExceeded) || errors.Is(err, context.Canceled):
		return status.Error(codes.DeadlineExceeded, ErrRequestTimeout)
	case errors.Is(err, domain.ErrUserNotFound):
		return status.Error(codes.NotFound, err.Error())
	case errors.Is(err, domain.ErrUserExists):
		return status.Error(codes.AlreadyExists, err.Error())
	case errors.Is(err, domain.ErrInvalidCredentials):
		return status.Error(codes.Unauthenticated, err.Error())
	case errors.Is(err, domain.ErrAlreadyFollow):
		return status.Error(codes.AlreadyExists, err.Error())
	case errors.Is(err, domain.ErrSelfFollow):
		return status.Error(codes.InvalidArgument, err.Error())
	case errors.Is(err, domain.ErrUnauthorizedAction):
		return status.Error(codes.PermissionDenied, err.Error())
	default:
		return status.Error(codes.Internal, ErrInternalServer)
	}
}

func HandleAuthError(userID any, logger *zap.Logger) error {
	if userID == nil {
		logger.Error(ErrFailedGetUserID)
		return status.Error(codes.Unauthenticated, ErrFailedGetUserID)
	}

	if _, ok := userID.(int); !ok {
		logger.Error(ErrFailedGetUserID)
		return status.Error(codes.Unauthenticated, ErrFailedGetUserID)
	}

	return nil
}

func GetUserIDFromContext[T any](ctx context.Context, ctxKey T) (int, error) {
	userID, ok := ctx.Value(ctxKey).(int)
	if !ok {
		return 0, errors.New(ErrFailedGetUserID)
	}
	return userID, nil
}

func GetOptionalUserIDFromContext[T any](ctx context.Context, ctxKey T) (int, error) {
	value := ctx.Value(ctxKey)

	if value == nil {
		return 0, nil
	}

	userID, ok := value.(int)
	if !ok {
		return 0, errors.New(ErrFailedGetUserID)
	}

	return userID, nil
}
