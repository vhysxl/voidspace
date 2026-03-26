package server

import (
	"voidspace/posts/bootstrap"
	service "voidspace/posts/internal/handler"
	pb "voidspace/posts/proto/generated/posts/v1"

	"github.com/vhysxl/voidspace/shared/utils/interceptor"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func SetupGRPCServer(app *bootstrap.Application) *grpc.Server {
	s := grpc.NewServer(
		grpc.UnaryInterceptor(interceptor.AuthInterceptor()))

	postHandler := service.NewPostHandler(
		app.PostUsecase,
		app.LikeUsecase,
		app.Logger,
		app.ContextTimeout,
	)

	pb.RegisterPostServiceServer(s, postHandler)

	reflection.Register(s)

	return s
}
