package handler

import (
	"context"
	pb "voidspace/users/proto/users/v1"
	"voidspace/users/utils/token"

	"github.com/vhysxl/voidspace/shared/utils/helper"

	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
)

func (u *UserHandler) Login(ctx context.Context, req *pb.LoginRequest) (*pb.AuthResponse, error) {
	user, err := u.UserUsecase.Login(ctx, req.GetEmailOrUsername(), req.GetPassword())
	if err != nil {
		return nil, helper.HandleError(err, u.Logger, "Login")
	}

	g, _ := errgroup.WithContext(ctx)

	var (
		accessToken  string
		refreshToken string
	)

	g.Go(func() error {
		var err error
		accessToken, err = token.CreateAccessToken(user, u.PrivateKey, u.AccessTokenDuration)
		return err
	})

	g.Go(func() error {
		var err error
		refreshToken, err = token.CreateRefreshToken(user, u.PrivateKey, u.RefreshTokenDuration)
		return err
	})

	err = g.Wait()
	if err != nil {
		u.Logger.Error("failed to generate token", zap.Error(err))
		return nil, helper.HandleError(err, u.Logger, "Create Token")
	}

	return &pb.AuthResponse{
		RefreshToken: &refreshToken,
		AccessToken:  accessToken,
		ExpiresIn:    int64(u.AccessTokenDuration.Seconds()),
		Message:      "Login Success",
	}, nil
}
