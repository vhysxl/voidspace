package user

import (
	"context"
	temporal_dto "voidspaceGateway/temporal/dto"
	"voidspaceGateway/utils"

	"go.uber.org/zap"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/emptypb"
)

const DeleteUserActivity = "DeleteUserActivity"

func (ua *UserActivities) DeleteUserActivity(
	ctx context.Context,
	req temporal_dto.DeleteUserReq,
) error {
	ua.Logger.Info(
		"Starting Delete User Activity",
		zap.String("userID", req.UserID),
		zap.String("username", req.Username))

	md := utils.MetaDataHandler(req.UserID, req.Username)
	ctx = metadata.NewOutgoingContext(ctx, md)

	_, err := ua.UserClient.DeleteUser(ctx, &emptypb.Empty{})
	if err != nil {
		ua.Logger.Error("failed to call UserService.DeleteUser", zap.Error(err))
		return err
	}

	return nil
}
