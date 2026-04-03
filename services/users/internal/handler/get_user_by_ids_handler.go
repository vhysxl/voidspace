package handler

import (
	"context"
	pb "voidspace/users/proto/users/v1"

	"github.com/vhysxl/voidspace/shared/utils/helper"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (u *UserHandler) GetUsers(
	ctx context.Context,
	req *pb.GetUsersRequest) (
	*pb.GetUsersResponse, error) {
	convertedIDs := make([]int, 0, len(req.GetUserIds()))
	for _, ID := range req.GetUserIds() {
		convertedIDs = append(convertedIDs, int(ID))
	}

	users, err := u.UserUsecase.GetUserByIDs(ctx, convertedIDs)
	if err != nil {
		return nil, helper.HandleError(err, u.Logger, "GetUsersByIds")
	}

	usersRes := make([]*pb.UserProfile, 0, len(users))
	for _, user := range users {
		usersRes = append(usersRes,
			&pb.UserProfile{
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
			})
	}

	return &pb.GetUsersResponse{
		Users: usersRes,
	}, nil

}
