package main

import (
	"fmt"
	"log"
	"net"
	"voidspace/posts/bootstrap"
	"voidspace/posts/server"
)

func main() {
	app, err := bootstrap.App()
	if err != nil {
		log.Fatal(err)
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", app.Config.Port))
	if err != nil {
		log.Fatalf("listening error: %v", err)
	}

	s := server.SetupGRPCServer(app)

	log.Printf("gRPC server starting on port: %v", app.Config.Port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Serve error: %v", err)
	}
}
