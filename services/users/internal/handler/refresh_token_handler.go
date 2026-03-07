package handler

import (
	"context"
	"voidspace/users/internal/domain"
	pb "voidspace/users/proto/users/v1"
	"voidspace/users/utils/token"

	"github.com/vhysxl/voidspace/shared/utils/helper"
	"github.com/vhysxl/voidspace/shared/utils/interceptor"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (u *UserHandler) RefreshToken(ctx context.Context, _ *emptypb.Empty) (*pb.AuthResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, u.ContextTimeout)
	defer cancel()

	userID, err := helper.GetUserIDFromContext(ctx, interceptor.CtxKeyUserID)
	if err != nil {
		return nil, helper.HandleAuthError(nil, u.Logger)
	}

	username, ok := ctx.Value(interceptor.CtxKeyUsername).(string)
	if !ok {
		return nil, helper.HandleAuthError(nil, u.Logger)
	}

	user := &domain.User{ID: int(userID), Username: username}

	accessToken, err := token.CreateAccessToken(user, u.PrivateKey, u.AccessTokenDuration)
	if err != nil {
		return nil, helper.HandleError(err, u.Logger, "RefreshToken")
	}

	return &pb.AuthResponse{
		AccessToken: accessToken,
		ExpiresIn:   int64(u.AccessTokenDuration.Seconds()),
		Message:     "Token refreshed successfully",
	}, nil
}
