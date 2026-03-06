package handler

import (
	"context"
	pb "voidspace/users/proto/users/v1"
	errorutils "voidspace/users/utils/error"
)

func (u *UserHandler) ListFollowing(
	ctx context.Context,
	req *pb.GetUserByIdRequest) (*pb.ListFollowingResponse, error) {

	ctx, cancel := context.WithTimeout(ctx, u.ContextTimeout)
	defer cancel()

	users, err := u.UserUsecase.ListFollowing(ctx, int(req.GetUserId()))
	if err != nil {
		return nil, errorutils.HandleError(err, u.Logger, "Get Follower")
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

	return &pb.ListFollowingResponse{
		Users: userBanners,
	}, nil
}
