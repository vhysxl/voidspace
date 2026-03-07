package handler

import (
	"context"
	pb "voidspace/users/proto/users/v1"

	"github.com/vhysxl/voidspace/shared/utils/helper"
	"github.com/vhysxl/voidspace/shared/utils/interceptor"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (u *UserHandler) GetCurrentUser(
	ctx context.Context,
	_ *emptypb.Empty,
) (*pb.GetCurrentUserResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, u.ContextTimeout)
	defer cancel()

	userID, err := helper.GetUserIDFromContext(ctx, interceptor.CtxKeyUserID)
	if err != nil {
		return nil, helper.HandleAuthError(nil, u.Logger)
	}

	user, err := u.UserUsecase.GetCurrentUser(ctx, userID)
	if err != nil {
		return nil, helper.HandleError(err, u.Logger, "Get Current User")
	}

	return &pb.GetCurrentUserResponse{
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
