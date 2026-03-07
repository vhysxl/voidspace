package helper

import (
	"context"
	"errors"

	"github.com/vhysxl/voidspace/shared/utils/constants"
)

func GetUserIDFromContext[T any](ctx context.Context, ctxKey T) (int, error) {
	userID, ok := ctx.Value(ctxKey).(int)
	if !ok {
		return 0, errors.New(constants.FailedGetUserID)
	}
	return userID, nil
}
