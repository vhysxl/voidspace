package handler

import (
	"context"
	"voidspace/users/internal/domain"
	pb "voidspace/users/proto/users/v1"
	errorutils "voidspace/users/utils/error"
	"voidspace/users/utils/interceptor"
	"voidspace/users/utils/token"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (u *UserHandler) RefreshToken(ctx context.Context, _ *emptypb.Empty) (*pb.AuthResponse, error) {
	userID, err := errorutils.GetUserIDFromContext(ctx, interceptor.CtxKeyUserID)
	if err != nil {
		return nil, errorutils.HandleAuthError(userID, u.Logger)
	}

	username, ok := ctx.Value(interceptor.CtxKeyUsername).(string)
	if !ok {
		return nil, errorutils.HandleAuthError(nil, u.Logger)
	}

	user := &domain.User{ID: int(userID), Username: username}

	accessToken, err := token.CreateAccessToken(user, u.PrivateKey, u.AccessTokenDuration)
	if err != nil {
		return nil, errorutils.HandleError(err, u.Logger, "RefreshToken")
	}

	return &pb.AuthResponse{
		AccessToken: accessToken,
		ExpiresIn:   int64(u.AccessTokenDuration.Seconds()),
		Message:     "Token refreshed successfully",
	}, nil
}
