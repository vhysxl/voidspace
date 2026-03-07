package helper

import (
	"context"
	"errors"

	"github.com/vhysxl/voidspace/shared/utils/constants"
)

func GetOptionalUserIDFromContext[T any](ctx context.Context, ctxKey T) (int, error) {
	value := ctx.Value(ctxKey)

	if value == nil {
		return 0, nil
	}

	userID, ok := value.(int)
	if !ok {
		return 0, errors.New(constants.FailedGetUserID)
	}

	return userID, nil
}
