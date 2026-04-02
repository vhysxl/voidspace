package user

import (
	"context"
	"voidspaceGateway/internal/models"
	"voidspaceGateway/utils"

	"go.uber.org/zap"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *UserService) RefreshToken(
	ctx context.Context,
	userID string,
	username string,
) (*models.AuthResponseService, error) {
	ctx, cancel := context.WithTimeout(ctx, s.ContextTimeout)
	defer cancel()

	md := utils.MetaDataHandler(userID, username)
	ctx = metadata.NewOutgoingContext(ctx, md)

	res, err := s.UserClient.RefreshToken(ctx, &emptypb.Empty{})
	if err != nil {
		s.Logger.Error("failed to call UserService.RefreshToken", zap.Error(err))
		return nil, err
	}

	return utils.AuthMapper(res), nil
}
