package user

import (
	"context"
	userpb "voidspaceGateway/proto/generated/users/v1"
	"voidspaceGateway/utils"

	"go.uber.org/zap"
	"google.golang.org/grpc/metadata"
)

func (s *UserService) Unfollow(ctx context.Context, userID string, username string, targetUsername string) error {
	ctx, cancel := context.WithTimeout(ctx, s.ContextTimeout)
	defer cancel()

	// Resolve username to ID
	targetUser, err := s.GetUser(ctx, targetUsername, userID, username)
	if err != nil {
		return err
	}

	md := utils.MetaDataHandler(userID, username)
	ctx = metadata.NewOutgoingContext(ctx, md)

	_, err = s.UserClient.Unfollow(ctx, &userpb.UnfollowRequest{
		UserId: int64(targetUser.ID),
	})
	if err != nil {
		s.Logger.Error("failed to call UserService.Unfollow", zap.Error(err))
		return err
	}

	return nil
}
