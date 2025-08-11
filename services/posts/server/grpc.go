package server

import (
	"voidspace/posts/bootstrap"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func SetupGRPCServer(app *bootstrap.Application) *grpc.Server {
	s := grpc.NewServer()

	reflection.Register(s)

	return s
}
