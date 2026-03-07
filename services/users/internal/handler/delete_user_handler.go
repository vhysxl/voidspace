package handler

import (
	"context"
	pb "voidspace/users/proto/users/v1"

	"github.com/vhysxl/voidspace/shared/utils/helper"
	"github.com/vhysxl/voidspace/shared/utils/interceptor"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (u *UserHandler) DeleteUser(
	ctx context.Context,
	_ *emptypb.Empty,
) (*pb.DeleteUserResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, u.ContextTimeout)
	defer cancel()

	userID, err := helper.GetUserIDFromContext(ctx, interceptor.CtxKeyUserID)
	if err != nil {
		return nil, helper.HandleAuthError(nil, u.Logger)
	}

	err = u.UserUsecase.DeleteUser(ctx, userID)
	if err != nil {
		return nil, helper.HandleError(err, u.Logger, "Delete User")
	}

	return &pb.DeleteUserResponse{}, nil
}
