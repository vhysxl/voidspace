package helper

import (
	"github.com/vhysxl/voidspace/shared/utils/constants"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func HandleAuthError(userID any, logger *zap.Logger) error {
	if userID == nil {
		logger.Error(constants.FailedGetUserID)
		return status.Error(codes.Unauthenticated, constants.FailedGetUserID)
	}

	if _, ok := userID.(int); !ok {
		logger.Error(constants.FailedGetUserID)
		return status.Error(codes.Unauthenticated, constants.FailedGetUserID)
	}

	return nil
}
