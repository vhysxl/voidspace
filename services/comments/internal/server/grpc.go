package server

import (
	"voidspace/comments/bootstrap"
	"voidspace/comments/internal/service"
	pb "voidspace/comments/proto/generated/comments"
	"voidspace/comments/utils/interceptor"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func SetupGRPCServer(app *bootstrap.Application) *grpc.Server {
	s := grpc.NewServer(grpc.UnaryInterceptor(interceptor.AuthInterceptor())) // interceptor here

	commentService := service.NewCommentService(
		app.ContextTimeout,
		app.Logger,
		app.CommentUseCase,
	)

	pb.RegisterCommentServiceServer(s, commentService) // handler

	reflection.Register(s)

	return s
}
