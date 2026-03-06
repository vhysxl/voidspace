package handler

import (
	"context"
	pb "voidspace/users/proto/users/v1"
	errorutils "voidspace/users/utils/error"
)

func (u *UserHandler) RestoreUser(
	ctx context.Context,
	req *pb.RestoreUserRequest,
) (*pb.RestoreUserResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, u.ContextTimeout)
	defer cancel()

	err := u.UserUsecase.RestoreUser(ctx, int(req.GetUserId()))
	if err != nil {
		return nil, errorutils.HandleError(err, u.Logger, "Restore User")
	}

	return &pb.RestoreUserResponse{}, nil
}
