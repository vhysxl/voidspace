package handler

import (
	"context"
	pb "voidspace/users/proto/users/v1"
	errorutils "voidspace/users/utils/error"
	"voidspace/users/utils/interceptor"
)

func (u *UserHandler) Follow(
	ctx context.Context,
	req *pb.FollowRequest) (
	*pb.FollowResponse, error) {

	userID, err := errorutils.GetUserIDFromContext(ctx, interceptor.CtxKeyUserID)
	if err != nil {
		return nil, errorutils.HandleAuthError(nil, u.Logger)
	}

	err = u.FollowUsecase.Follow(ctx, userID, int(req.GetUserId()))
	if err != nil {
		return nil, errorutils.HandleError(err, u.Logger, "Failed to Follow")
	}

	return &pb.FollowResponse{}, nil
}
