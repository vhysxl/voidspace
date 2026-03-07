package handler

import (
	"context"
	pb "voidspace/users/proto/users/v1"

	"github.com/vhysxl/voidspace/shared/utils/helper"
	"github.com/vhysxl/voidspace/shared/utils/interceptor"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func (u *UserHandler) GetUser(
	ctx context.Context,
	req *pb.GetUserRequest,
) (*pb.GetUserResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, u.ContextTimeout)
	defer cancel()

	userID, err := helper.GetOptionalUserIDFromContext(ctx, interceptor.CtxKeyUserID)
	if err != nil {
		return nil, helper.HandleError(err, u.Logger, "Get User")
	}

	user, err := u.UserUsecase.GetUser(ctx, req.GetUsername(), userID)
	if err != nil {
		return nil, helper.HandleError(err, u.Logger, "Get User")
	}

	return &pb.GetUserResponse{
		User: &pb.UserProfile{
			Id:          int64(user.ID),
			Username:    user.Username,
			DisplayName: user.DisplayName,
			Bio:         user.Bio,
			AvatarUrl:   user.AvatarURL,
			BannerUrl:   user.BannerURL,
			Location:    user.Location,
			Followers:   int64(user.Follower),
			Following:   int64(user.Following),
			CreatedAt:   timestamppb.New(user.CreatedAt),
			IsFollowed:  user.IsFollowed,
		},
	}, nil

}
