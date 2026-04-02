package user

import (
	"context"
	userpb "voidspaceGateway/proto/generated/users/v1"

	"go.uber.org/zap"
)

func (s *UserService) RestoreUser(ctx context.Context, userID int64) error {
	ctx, cancel := context.WithTimeout(ctx, s.ContextTimeout)
	defer cancel()

	_, err := s.UserClient.RestoreUser(ctx, &userpb.RestoreUserRequest{
		UserId: userID,
	})
	if err != nil {
		s.Logger.Error("failed to call UserService.RestoreUser", zap.Error(err))
		return err
	}

	return nil
}
