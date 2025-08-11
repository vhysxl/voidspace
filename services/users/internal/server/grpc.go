package server

import (
	"voidspace/users/bootstrap"
	handler "voidspace/users/internal/service"
	pb "voidspace/users/proto/generated/voidspace/users/proto/users/v1"
	"voidspace/users/utils/interceptor"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func SetupGRPCServer(app *bootstrap.Application) *grpc.Server {
	s := grpc.NewServer(
		grpc.UnaryInterceptor(interceptor.AuthInterceptor()),
	)

	authHandler := handler.NewAuthHandler(
		app.AuthUsecase,
		app.Validator,
		app.PrivateKey,
		app.ContextTimeout,
		app.AccessTokenDuration,
		app.RefreshTokenDuration,
		app.Logger,
	)

	userHandler := handler.NewUserHandler(
		app.UserUsecase,
		app.ProfileUsecase,
		app.FollowUsecase,
		app.Validator,
		app.ContextTimeout,
		app.Logger,
	)

	pb.RegisterAuthServiceServer(s, authHandler)
	pb.RegisterUserServiceServer(s, userHandler)

	reflection.Register(s)

	return s
}
