package user

import (
	"context"
	"voidspaceGateway/internal/models"
	userpb "voidspaceGateway/proto/generated/users/v1"
	"voidspaceGateway/utils"

	"go.uber.org/zap"
	"google.golang.org/grpc/metadata"
)

func (s *UserService) UpdateProfile(ctx context.Context, userID string, username string, req *models.UpdateProfileRequest) error {
	ctx, cancel := context.WithTimeout(ctx, s.ContextTimeout)
	defer cancel()

	md := utils.MetaDataHandler(userID, username)

	ctx = metadata.NewOutgoingContext(ctx, md)

	_, err := s.UserClient.UpdateProfile(ctx, &userpb.UpdateProfileRequest{
		DisplayName: &req.DisplayName,
		Bio:         &req.Bio,
		AvatarUrl:   &req.AvatarURL,
		BannerUrl:   &req.BannerURL,
		Location:    &req.Location,
	})
	if err != nil {
		s.Logger.Error("failed to call UserService.UpdateProfile", zap.Error(err))
		return err
	}

	return nil
}
