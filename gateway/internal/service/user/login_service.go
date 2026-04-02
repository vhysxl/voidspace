package user

import (
	"context"
	"voidspaceGateway/internal/models"
	userpb "voidspaceGateway/proto/generated/users/v1"
	"voidspaceGateway/utils"

	"go.uber.org/zap"
)

func (s *UserService) Login(
	ctx context.Context,
	req *models.LoginRequest,
) (*models.AuthResponseService, error) {
	ctx, cancel := context.WithTimeout(ctx, s.ContextTimeout)
	defer cancel()

	res, err := s.UserClient.Login(ctx, &userpb.LoginRequest{
		EmailOrUsername: req.UsernameOrEmail,
		Password:        req.Password,
	})
	if err != nil {
		s.Logger.Error("failed to call AuthService.Login", zap.Error(err))
		return nil, err
	}

	return utils.AuthMapper(res), nil
}
