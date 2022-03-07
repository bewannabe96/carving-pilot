package service

import (
	"context"
	"log"

	"github.com/bewannabe96/daytrip-datapipeline-grpc-idl/gen/go/carving"
)

type SessionServiceServer struct {
	carving.SessionServiceServer
}

func (s *SessionServiceServer) LogSessionStart(ctx context.Context, req *carving.LogSessionStartRequest) (*carving.LogSessionStartResponse, error) {
	log.Printf("[LogSessionStart] screen=%s\n", req.ScreenName)

	return &carving.LogSessionStartResponse{
		Status:    "OK",
		SessionId: "THISISSESSIONID",
	}, nil
}
