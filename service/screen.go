package service

import (
	"context"
	"log"

	"github.com/bewannabe96/daytrip-datapipeline-grpc-idl/gen/go/carving"
)

type ScreenServiceServer struct {
	carving.ScreenServiceServer
}

func (s *ScreenServiceServer) LogScreenChanged(ctx context.Context, req *carving.LogScreenChangedRequest) (*carving.LogScreenChangedResponse, error) {
	log.Printf("[LogSessionStart] screen=%s\n", req.ScreenName)

	return &carving.LogScreenChangedResponse{
		Status:    "OK",
		SessionId: "THISISSESSIONID",
	}, nil
}
