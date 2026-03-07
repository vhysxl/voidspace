package handler

import (
	"context"
	pb "voidspace/users/proto/users/v1"

	"github.com/vhysxl/voidspace/shared/utils/helper"
	"github.com/vhysxl/voidspace/shared/utils/interceptor"
)

func (u *UserHandler) Unfollow(
	ctx context.Context,
	req *pb.UnfollowRequest,
) (*pb.UnfollowResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, u.ContextTimeout)
	defer cancel()

	userID, err := helper.GetUserIDFromContext(ctx, interceptor.CtxKeyUserID)
	if err != nil {
		return nil, helper.HandleAuthError(nil, u.Logger)
	}

	err = u.FollowUsecase.Unfollow(ctx, userID, int(req.GetUserId()))
	if err != nil {
		return nil, helper.HandleError(err, u.Logger, "Failed to Unfollow")
	}

	return &pb.UnfollowResponse{}, nil
}
