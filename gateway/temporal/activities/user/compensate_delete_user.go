package user

import (
	"context"

	userpb "voidspaceGateway/proto/generated/users/v1"
	temporal_dto "voidspaceGateway/temporal/dto"
	"voidspaceGateway/utils"

	"go.uber.org/zap"
	"google.golang.org/grpc/metadata"
)

const DeleteUserCompensateActivity = "DeleteUserCompensateActivity"

func (ua *UserActivities) DeleteUserCompensateActivity(
	ctx context.Context,
	req temporal_dto.DeleteUserReq,
) error {
	ua.Logger.Info(
		"Starting Delete User's Activity",
		zap.String("userID", req.UserID),
	)

	md := utils.MetaDataHandler(req.UserID, req.Username)
	ctx = metadata.NewOutgoingContext(ctx, md)

	_, err := ua.UserClient.RestoreUser(ctx, &userpb.RestoreUserRequest{
		UserId: int64(req.UserIDInt),
	})
	if err != nil {
		ua.Logger.Error("Failed to compensate user", zap.Error(err))
		return err
	}

	ua.Logger.Info("User's comments compensated successfully")
	return nil
}
