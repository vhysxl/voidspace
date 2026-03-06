package handler

import (
	"context"
	pb "voidspace/users/proto/users/v1"
	errorutils "voidspace/users/utils/error"
	"voidspace/users/utils/interceptor"
)

func (u *UserHandler) Unfollow(
	ctx context.Context,
	req *pb.UnfollowRequest,
) (*pb.UnfollowResponse, error) {

	userID, err := errorutils.GetUserIDFromContext[interceptor.CtxKey](ctx, interceptor.CtxKeyUserID)
	if err != nil {
		return nil, errorutils.HandleAuthError(nil, u.Logger)
	}
	err = u.FollowUsecase.Unfollow(ctx, userID, int(req.GetUserId()))
	if err != nil {
		return nil, errorutils.HandleError(err, u.Logger, "Failed to Follow")
	}

	return &pb.UnfollowResponse{}, nil

}
