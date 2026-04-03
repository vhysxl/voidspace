package handler

import (
	"context"
	pb "voidspace/users/proto/users/v1"

	"github.com/vhysxl/voidspace/shared/utils/helper"
	"github.com/vhysxl/voidspace/shared/utils/interceptor"
)

func (u *UserHandler) Follow(
	ctx context.Context,
	req *pb.FollowRequest) (
	*pb.FollowResponse, error) {
	userID, err := helper.GetUserIDFromContext(ctx, interceptor.CtxKeyUserID)
	if err != nil {
		return nil, helper.HandleAuthError(nil, u.Logger)
	}

	err = u.FollowUsecase.Follow(ctx, userID, int(req.GetUserId()))
	if err != nil {
		return nil, helper.HandleError(err, u.Logger, "Failed to Follow")
	}

	return &pb.FollowResponse{}, nil
}
