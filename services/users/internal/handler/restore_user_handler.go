package handler

import (
	"context"
	pb "voidspace/users/proto/users/v1"

	"github.com/vhysxl/voidspace/shared/utils/helper"
)

func (u *UserHandler) RestoreUser(
	ctx context.Context,
	req *pb.RestoreUserRequest,
) (*pb.RestoreUserResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, u.ContextTimeout)
	defer cancel()

	err := u.UserUsecase.RestoreUser(ctx, int(req.GetUserId()))
	if err != nil {
		return nil, helper.HandleError(err, u.Logger, "Restore User")
	}

	return &pb.RestoreUserResponse{}, nil
}
