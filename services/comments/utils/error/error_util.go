package error

import (
	"context"
	"errors"
	"voidspace/comments/internal/domain"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	ErrRequestTimeout    = "Request Timeout"
	ErrInvalidRequest    = "Invalid request"
	ErrInternalServer    = "Internal server error"
	ErrUnauthorized      = "Unauthorized"
	ErrUsecase           = "Usecase error"
	ErrFailedGetUserID   = "failed to get user ID from context"
	ErrFailedGetUsername = "failed to get username from context"
)

func HandleError(err error, logger *zap.Logger, operation string) error {
	if err == nil {
		return nil
	}

	logger.Error(ErrUsecase, zap.String("operation", operation), zap.Error(err))

	switch {
	case errors.Is(err, context.DeadlineExceeded) || errors.Is(err, context.Canceled):
		return status.Error(codes.DeadlineExceeded, ErrRequestTimeout)
	case errors.Is(err, domain.ErrCommentsNotFound):
		return status.Error(codes.NotFound, err.Error())
	case errors.Is(err, domain.ErrUnauthorizedAction):
		return status.Error(codes.PermissionDenied, err.Error())
	default:
		return status.Error(codes.Internal, ErrInternalServer)
	}
}

func HandleAuthError(userID interface{}, logger *zap.Logger) error {
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

func GetUserIDFromContext(ctx context.Context, ctxKey interface{}) (int, error) {
	userID, ok := ctx.Value(ctxKey).(int)
	if !ok {
		return 0, errors.New(ErrFailedGetUserID)
	}
	return userID, nil
}
