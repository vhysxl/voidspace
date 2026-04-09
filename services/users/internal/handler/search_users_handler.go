package handler

import (
	"context"
	pb "voidspace/users/proto/users/v1"

	"github.com/vhysxl/voidspace/shared/utils/helper"
)

func (u *UserHandler) SearchUsers(
	ctx context.Context,
	req *pb.SearchUsersRequest,
) (*pb.SearchUsersResponse, error) {
	users, err := u.UserUsecase.SearchUsers(ctx, req.GetQuery())
	if err != nil {
		return nil, helper.HandleError(err, u.Logger, "Search Users")
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

	return &pb.SearchUsersResponse{
		Users: userBanners,
	}, nil
}
