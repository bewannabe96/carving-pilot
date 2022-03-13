package server

import (
	"context"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"time"

	idl "github.com/ourspaceapp/carving/idl"
	"github.com/ourspaceapp/carving/kinesisclient"
	"github.com/ourspaceapp/carving/session"
)

type SessionServiceServer struct {
	idl.SessionServiceServer
}

func (s *SessionServiceServer) LogSessionStart(ctx context.Context, req *idl.LogSessionStartRequest) (*idl.LogSessionStartResponse, error) {

	sessionId := fmt.Sprintf("%d_%02d", time.Now().UnixMicro(), rand.Intn(99))

	token, err := session.IssueToken(sessionId)
	if err != nil {
		log.Printf("[LogSessionStart] TOKEN_ERROR unable to issue new token\n")
		return nil, errors.New("token error: unable to issue new token")
	}

	if err := kinesisclient.Produce(
		sessionId,
		"session_start",
		map[string]interface{}{
			"session_id": sessionId,
		},
	); err != nil {
		log.Printf("[LogSessionStart] KINESIS_ERROR unable to produce kiensis record\n")
		return nil, errors.New("kinesis error: unable to produce kiensis record")
	}

	log.Printf("[LogSessionStart] session_id=%s\n", sessionId)

	return &idl.LogSessionStartResponse{
		SessionToken: token.Get(),
	}, nil

}
