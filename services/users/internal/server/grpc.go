package server

import (
	"voidspace/users/bootstrap"
	handler "voidspace/users/internal/handler"
	user_pb "voidspace/users/proto/users/v1"

	"github.com/vhysxl/voidspace/shared/utils/interceptor"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func SetupGRPCServer(app *bootstrap.Application) *grpc.Server {
	s := grpc.NewServer(
		grpc.UnaryInterceptor(interceptor.AuthInterceptor()),
	)

	userHandler := handler.NewUserHandler(
		app.UserUsecase,
		app.ProfileUsecase,
		app.FollowUsecase,
		app.ContextTimeout,
		app.Logger,
		app.PrivateKey,
		app.AccessTokenDuration,
		app.RefreshTokenDuration,
	)

	user_pb.RegisterUserServiceServer(s, userHandler)

	reflection.Register(s)

	return s
}
