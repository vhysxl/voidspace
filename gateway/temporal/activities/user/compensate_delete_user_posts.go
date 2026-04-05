package user

import (
	"context"

	postpb "voidspaceGateway/proto/generated/posts/v1"
	temporal_dto "voidspaceGateway/temporal/dto"
	"voidspaceGateway/utils"

	"go.uber.org/zap"
	"google.golang.org/grpc/metadata"
)

const DeleteUserPostsCompensateActivity = "DeleteUserPostsCompensateActivity"

func (ua *UserActivities) DeleteUserPostsCompensateActivity(
	ctx context.Context,
	req temporal_dto.DeleteUserReq,
) error {
	ua.Logger.Info(
		"Starting Delete User's Posts Activity",
		zap.String("userID", req.UserID),
	)

	md := utils.MetaDataHandler(req.UserID, req.Username)
	ctx = metadata.NewOutgoingContext(ctx, md)

	_, err := ua.PostClient.HandleAccountRestoration(ctx, &postpb.HandleAccountRestorationRequest{
		UserId: int64(req.UserIDInt),
	})
	if err != nil {
		ua.Logger.Error("Failed to compensate user's posts", zap.Error(err))
		return err
	}

	ua.Logger.Info("User's posts compensated successfully")
	return nil
}
