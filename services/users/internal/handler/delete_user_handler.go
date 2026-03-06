package handler

import (
	"context"
	pb "voidspace/users/proto/users/v1"
	errorutils "voidspace/users/utils/error"
	"voidspace/users/utils/interceptor"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (u *UserHandler) DeleteUser(
	ctx context.Context,
	_ *emptypb.Empty,
) (*pb.DeleteUserResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, u.ContextTimeout)
	defer cancel()

	userID, err := errorutils.GetUserIDFromContext(ctx, interceptor.CtxKeyUserID)
	if err != nil {
		return nil, errorutils.HandleAuthError(nil, u.Logger)
	}

	err = u.UserUsecase.DeleteUser(ctx, userID)
	if err != nil {
		return nil, errorutils.HandleError(err, u.Logger, "Delete User")
	}

	return &pb.DeleteUserResponse{}, nil
}
