package helper

import (
	"context"
	"errors"

	"github.com/vhysxl/voidspace/shared/utils/constants"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func HandleError(err error, logger *zap.Logger, operation string) error {
	if err == nil {
		return nil
	}

	logger.Error(constants.Usecase, zap.String("OPERATION", operation), zap.Error(err))

	switch {
	case errors.Is(err, context.DeadlineExceeded) || errors.Is(err, context.Canceled):
		return status.Error(codes.DeadlineExceeded, constants.RequestTimeout)
	case errors.Is(err, constants.ErrUserNotFound):
		return status.Error(codes.NotFound, err.Error())
	case errors.Is(err, constants.ErrUserExists):
		return status.Error(codes.AlreadyExists, err.Error())
	case errors.Is(err, constants.ErrInvalidCredentials):
		return status.Error(codes.Unauthenticated, err.Error())
	case errors.Is(err, constants.ErrAlreadyFollowing):
		return status.Error(codes.AlreadyExists, err.Error())
	case errors.Is(err, constants.ErrCannotFollowSelf):
		return status.Error(codes.InvalidArgument, err.Error())
	case errors.Is(err, constants.ErrUnauthorized):
		return status.Error(codes.PermissionDenied, err.Error())
	default:
		return status.Error(codes.Internal, constants.InternalServer)
	}
}
