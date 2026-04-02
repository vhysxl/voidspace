package user

import (
	"context"
	"voidspaceGateway/internal/models"
	userpb "voidspaceGateway/proto/generated/users/v1"
	"voidspaceGateway/utils"

	"go.uber.org/zap"
	"google.golang.org/grpc/metadata"
)

func (s *UserService) GetUser(ctx context.Context, username string, userID string, usernameRequester string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(ctx, s.ContextTimeout)
	defer cancel()

	md := utils.MetaDataHandler(userID, usernameRequester)

	ctx = metadata.NewOutgoingContext(ctx, md)

	res, err := s.UserClient.GetUser(ctx, &userpb.GetUserRequest{
		Username: username,
	})
	if err != nil {
		s.Logger.Error("failed to call UserService.GetUser", zap.Error(err))
		return nil, err
	}

	return utils.UserMapper(res.User), nil
}
