package main

import (
	"fmt"
	"log"
	"net"
	"voidspace/posts/bootstrap"
	"voidspace/posts/server"

	"go.uber.org/zap"
)

func main() {
	app, err := bootstrap.App()
	if err != nil {
		log.Fatal(err)
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", app.Config.Port))
	if err != nil {
		app.Logger.Fatal("listening error", zap.Error(err))
	}

	s := server.SetupGRPCServer(app)

	app.Logger.Info("gRPC server starting", zap.String("port", app.Config.Port))
	if err := s.Serve(lis); err != nil {
		app.Logger.Fatal("Serve error", zap.Error(err))
	}
}
