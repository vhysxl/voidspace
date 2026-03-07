package handler

import (
	"context"
	pb "voidspace/users/proto/users/v1"

	"github.com/vhysxl/voidspace/shared/utils/helper"
)

func (u *UserHandler) ListFollowers(
	ctx context.Context,
	req *pb.GetUserByIdRequest) (*pb.ListFollowersResponse, error) {

	ctx, cancel := context.WithTimeout(ctx, u.ContextTimeout)
	defer cancel()

	users, err := u.UserUsecase.ListFollowers(ctx, int(req.GetUserId()))
	if err != nil {
		return nil, helper.HandleError(err, u.Logger, "Get Follower")
	}

	userBanners := make([]*pb.UserBanner, 0, len(users))
	for _, user := range users {
		userBanners = append(userBanners, &pb.UserBanner{
			Id:          int64(user.ID),
			Username:    user.Username,
			DisplayName: user.DisplayName,
			AvatarUrl:   user.AvatarURL,
		})
	}

	return &pb.ListFollowersResponse{
		Users: userBanners,
	}, nil
}
