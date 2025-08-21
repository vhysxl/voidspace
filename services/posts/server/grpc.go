package server

import (
	"voidspace/posts/bootstrap"
	"voidspace/posts/internal/service"
	pb "voidspace/posts/proto/generated/posts"
	"voidspace/posts/utils/interceptor"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func SetupGRPCServer(app *bootstrap.Application) *grpc.Server {
	s := grpc.NewServer(
		grpc.UnaryInterceptor(interceptor.AuthInterceptor()))

	likeHandler := service.NewLikeHandler(
		app.LikeUsecase,
		app.Validator,
		app.ContextTimeout,
		app.Logger,
	)

	postHandler := service.NewPostHandler(
		app.PostUsecase,
		app.Validator,
		app.ContextTimeout,
		app.Logger,
	)

	pb.RegisterPostServiceServer(s, postHandler)
	pb.RegisterLikesServiceServer(s, likeHandler)

	reflection.Register(s)

	return s
}
