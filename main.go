package main

import (
	"log"
	"net"

	idl "github.com/ourspaceapp/carving/idl"
	"github.com/ourspaceapp/carving/server"
	"google.golang.org/grpc"
)

func main() {
	const host = ":3000"

	listener, err := net.Listen("tcp", host)
	if err != nil {
		log.Fatalf("Failed to listen: %v\n", err)
	} else {
		log.Printf("Start listening on %s\n", host)
	}

	grpcServer := grpc.NewServer()

	idl.RegisterSessionServiceServer(grpcServer, &server.SessionServiceServer{})
	idl.RegisterUserServiceServer(grpcServer, &server.UserServiceServer{})
	idl.RegisterEngagementServiceServer(grpcServer, &server.EngagementServiceServer{})

	grpcServer.Serve(listener)
}
