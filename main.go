package main

import (
	"log"
	"net"

	"github.com/bewannabe96/carving-pilot/service"
	"github.com/bewannabe96/daytrip-datapipeline-grpc-idl/gen/go/carving"
	"google.golang.org/grpc"
)

func main() {
	const host = "localhost:3000"

	listener, err := net.Listen("tcp", host)
	if err != nil {
		log.Fatalf("Failed to listen: %v\n", err)
	} else {
		log.Printf("Start listening on %s\n", host)
	}

	grpcServer := grpc.NewServer()

	carving.RegisterSessionServiceServer(grpcServer, &service.SessionServiceServer{})
	carving.RegisterScreenServiceServer(grpcServer, &service.ScreenServiceServer{})

	grpcServer.Serve(listener)
}
