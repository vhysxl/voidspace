package user

import (
	"context"

	commentpb "voidspaceGateway/proto/generated/comments/v1"
	temporal_dto "voidspaceGateway/temporal/dto"
	"voidspaceGateway/utils"

	"go.uber.org/zap"
	"google.golang.org/grpc/metadata"
)

const DeleteUserCommentsActivity = "DeleteUserCommentsActivity"

func (ua *UserActivities) DeleteUserCommentsActivity(
	ctx context.Context,
	req temporal_dto.DeleteUserReq,
) error {
	ua.Logger.Info(
		"Starting Delete User's Comments Activity",
		zap.String("userID", req.UserID),
	)

	md := utils.MetaDataHandler(req.UserID, req.Username)
	ctx = metadata.NewOutgoingContext(ctx, md)

	_, err := ua.CommentClient.HandleAccountDeletion(ctx, &commentpb.HandleAccountDeletionRequest{
		UserId: int64(req.UserIDInt),
	})
	if err != nil {
		ua.Logger.Error("Failed to delete user's comments", zap.Error(err))
		return err
	}

	ua.Logger.Info("User's comments deleted successfully")
	return nil
}
