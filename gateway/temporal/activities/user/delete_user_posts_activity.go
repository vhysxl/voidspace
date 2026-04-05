package user

import (
	"context"

	postpb "voidspaceGateway/proto/generated/posts/v1"
	temporal_dto "voidspaceGateway/temporal/dto"
	"voidspaceGateway/utils"

	"go.uber.org/zap"
	"google.golang.org/grpc/metadata"
)

const DeleteUserPostsActivity = "DeleteUserPostsActivity"

func (ua *UserActivities) DeleteUserPostsActivity(
	ctx context.Context,
	req temporal_dto.DeleteUserReq,
) error {
	ua.Logger.Info(
		"Starting Delete User's Posts Activity",
		zap.String("userID", req.UserID),
	)

	md := utils.MetaDataHandler(req.UserID, req.Username)
	ctx = metadata.NewOutgoingContext(ctx, md)

	_, err := ua.PostClient.HandleAccountDeletion(ctx, &postpb.HandleAccountDeletionRequest{
		UserId: int64(req.UserIDInt),
	})
	if err != nil {
		ua.Logger.Error("Failed to delete user's posts", zap.Error(err))
		return err
	}

	ua.Logger.Info("User's posts deleted successfully")
	return nil
}
