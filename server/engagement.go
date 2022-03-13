package server

import (
	"context"
	"encoding/json"
	"errors"
	"log"

	idl "github.com/ourspaceapp/carving/idl"
	"github.com/ourspaceapp/carving/kinesisclient"
	"github.com/ourspaceapp/carving/session"
	"google.golang.org/protobuf/types/known/emptypb"
)

type EngagementServiceServer struct {
	idl.EngagementServiceServer
}

func (s *EngagementServiceServer) LogScreenChange(ctx context.Context, req *idl.LogScreenChangeRequest) (*emptypb.Empty, error) {

	if req.SessionToken == "" {
		log.Printf("[LogScreenChange] INVALID_REQUEST session_token is missing\n")
		return nil, errors.New("invalid request: session_token is missing")
	} else if req.ScreenId == "" {
		log.Printf("[LogScreenChange] INVALID_REQUEST screen_id is missing\n")
		return nil, errors.New("invalid request: screen_id is missing")
	}

	token, err := session.NewToken(req.SessionToken)
	if err != nil {
		log.Printf("[LogScreenChange] TOKEN_ERROR session_token is invalid\n")
		return nil, errors.New("token error: session_token is invalid")
	}

	claims, ok := token.GetClaims([]string{"session_id", "user_id"})
	if !ok {
		log.Printf("[LogScreenChange] TOKEN_ERROR session_token does not include all required claims\n")
		return nil, errors.New("token error: session_token does not include all required claims")
	}

	if err := kinesisclient.Produce(
		claims["session_id"],
		"screen_change",
		map[string]interface{}{
			"session_id":     claims["session_id"],
			"user_id":        claims["user_id"],
			"screen_id":      req.ScreenId,
			"prev_screen_id": req.PrevScreenId,
		},
	); err != nil {
		log.Printf("[LogScreenChange] KINESIS_ERROR unable to produce kiensis record\n")
		return nil, errors.New("kinesis error: unable to produce kiensis record")
	}

	log.Printf("[LogScreenChange] session_id=%s user_id=%s screen_id=%s prev_screen_id=%s\n", claims["session_id"], claims["user_id"], req.ScreenId, req.PrevScreenId)

	return &emptypb.Empty{}, nil

}

func (s *EngagementServiceServer) LogEvent(ctx context.Context, req *idl.LogEventRequest) (*emptypb.Empty, error) {

	if req.SessionToken == "" {
		log.Printf("[LogEvent] INVALID_REQUEST session_token is missing\n")
		return nil, errors.New("invalid request: session_token is missing")
	} else if req.Event == "" {
		log.Printf("[LogEvent] INVALID_REQUEST event is missing\n")
		return nil, errors.New("invalid request: event is missing")
	}

	token, err := session.NewToken(req.SessionToken)
	if err != nil {
		log.Printf("[LogEvent] TOKEN_ERROR session_token is invalid\n")
		return nil, errors.New("token error: session_token is invalid")
	}

	claims, ok := token.GetClaims([]string{"session_id", "user_id"})
	if !ok {
		log.Printf("[LogEvent] TOKEN_ERROR session_token does not include all required claims\n")
		return nil, errors.New("token error: session_token does not include all required claims")
	}

	payload := make(map[string]interface{})
	if err := json.Unmarshal([]byte(req.Payload), &payload); err != nil {
		log.Printf("[LogEvent] INVALID_REQUEST payload is not in valid format\n")
		return nil, errors.New("invalid request: payload is not in valid format")
	}

	if err := kinesisclient.Produce(
		claims["session_id"],
		"event",
		map[string]interface{}{
			"session_id": claims["session_id"],
			"user_id":    claims["user_id"],
			"event":      req.Event,
			"payload":    payload,
		},
	); err != nil {
		log.Printf("[LogEvent] KINESIS_ERROR unable to produce kiensis record\n")
		return nil, errors.New("kinesis error: unable to produce kiensis record")
	}

	log.Printf("[LogEvent] session_id=%s user_id=%s event=%s\n", claims["session_id"], claims["user_id"], req.Event)

	return &emptypb.Empty{}, nil

}
