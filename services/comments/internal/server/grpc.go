package server

import (
	"voidspace/comments/bootstrap"
	handler "voidspace/comments/internal/handler"
	pb "voidspace/comments/proto/generated/comments"
	"voidspace/comments/utils/interceptor"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func SetupGRPCServer(app *bootstrap.Application) *grpc.Server {
	s := grpc.NewServer(grpc.UnaryInterceptor(interceptor.AuthInterceptor())) // interceptor here

	commentHandler := handler.NewCommentHandler(
		app.ContextTimeout,
		app.Logger,
		app.CommentUseCase,
	)

	pb.RegisterCommentServiceServer(s, commentHandler) // handler

	reflection.Register(s)

	return s
}
