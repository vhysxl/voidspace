package handler

import (
	"context"
	pb "voidspace/users/proto/users/v1"

	"github.com/vhysxl/voidspace/shared/utils/helper"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func (u *UserHandler) GetUserById(
	ctx context.Context,
	req *pb.GetUserByIdRequest,
) (*pb.GetUserResponse, error) {
	user, err := u.UserUsecase.GetUserByID(ctx, int(req.GetUserId()))
	if err != nil {
		return nil, helper.HandleError(err, u.Logger, "Get User By ID")
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
