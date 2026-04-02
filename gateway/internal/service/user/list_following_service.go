package user

// import (
// 	"context"
// 	"voidspaceGateway/internal/models"
// 	userpb "voidspaceGateway/proto/generated/users/v1"
// 	"voidspaceGateway/utils"

// 	"go.uber.org/zap"
// )

// func (s *UserService) ListFollowing(ctx context.Context, userID int64) ([]*models.UserBanner, error) {
// 	ctx, cancel := context.WithTimeout(ctx, s.ContextTimeout)
// 	defer cancel()

// 	res, err := s.UserClient.ListFollowing(ctx, &userpb.GetUserByIdRequest{
// 		UserId: userID,
// 	})
// 	if err != nil {
// 		s.Logger.Error("failed to call UserService.ListFollowing", zap.Error(err))
// 		return nil, err
// 	}

// 	banners := make([]*models.UserBanner, 0, len(res.Users))
// 	for _, u := range res.Users {
// 		banners = append(banners, utils.UserBannerMapper(u))
// 	}

// 	return banners, nil
// }
