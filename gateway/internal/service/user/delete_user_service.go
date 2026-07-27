package user

import (
	"context"
	commentpb "voidspaceGateway/proto/generated/comments/v1"
	postpb "voidspaceGateway/proto/generated/posts/v1"
	"voidspaceGateway/utils"

	"go.uber.org/zap"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *UserService) DeleteUser(ctx context.Context, userID string, username string) error {
	ctx, cancel := context.WithTimeout(ctx, s.ContextTimeout)
	defer cancel()

	user, err := s.GetCurrentUser(ctx, userID, username)
	if err != nil {
		s.Logger.Error("failed to get user", zap.Error(err))
		return err
	}

	md := utils.MetaDataHandler(userID, username)
	ctx = metadata.NewOutgoingContext(ctx, md)

	// 1. Delete User via gRPC
	_, err = s.UserClient.DeleteUser(ctx, &emptypb.Empty{})
	if err != nil {
		s.Logger.Error("failed to call UserService.DeleteUser", zap.Error(err))
		return err
	}

	userIDInt := int64(user.ID)

	// 2. Delete User's Posts via gRPC
	_, err = s.PostClient.HandleAccountDeletion(ctx, &postpb.HandleAccountDeletionRequest{
		UserId: userIDInt,
	})
	if err != nil {
		s.Logger.Error("failed to delete user's posts", zap.Error(err))
		return err
	}

	// 3. Delete User's Comments via gRPC
	_, err = s.CommentClient.HandleAccountDeletion(ctx, &commentpb.HandleAccountDeletionRequest{
		UserId: userIDInt,
	})
	if err != nil {
		s.Logger.Error("failed to delete user's comments", zap.Error(err))
		return err
	}

	return nil
}
