package user

import (
	"context"
	"voidspaceGateway/internal/models"
	userpb "voidspaceGateway/proto/generated/users/v1"
	"voidspaceGateway/utils"

	"go.uber.org/zap"
)

func (s *UserService) GetUsers(ctx context.Context, userIDs []int64) ([]*models.User, error) {
	ctx, cancel := context.WithTimeout(ctx, s.ContextTimeout)
	defer cancel()

	res, err := s.UserClient.GetUsers(ctx, &userpb.GetUsersRequest{
		UserIds: userIDs,
	})
	if err != nil {
		s.Logger.Error("failed to call UserService.GetUsers", zap.Error(err))
		return nil, err
	}

	users := make([]*models.User, 0, len(res.Users))
	for _, u := range res.Users {
		users = append(users, utils.UserMapper(u))
	}

	return users, nil
}
