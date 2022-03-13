package server

import (
	"context"
	"errors"
	"log"

	idl "github.com/ourspaceapp/carving/idl"
	"github.com/ourspaceapp/carving/kinesisclient"
	"github.com/ourspaceapp/carving/session"
)

type UserServiceServer struct {
	idl.UserServiceServer
}

func (s *UserServiceServer) SetSessionUser(ctx context.Context, req *idl.SetSessionUserRequest) (*idl.SetSessionUserResponse, error) {

	if req.SessionToken == "" {
		log.Printf("[SetSessionUser] INVALID_REQUEST session_token is missing\n")
		return nil, errors.New("invalid request: session_token is missing")
	} else if req.UserId == "" {
		log.Printf("[SetSessionUser] INVALID_REQUEST user_id is missing\n")
		return nil, errors.New("invalid request: user_id is missing")
	}

	token, err := session.NewToken(req.SessionToken)
	if err != nil {
		log.Printf("[SetSessionUser] TOKEN_ERROR session_token is invalid\n")
		return nil, errors.New("token error: session_token is invalid")
	}

	claims, ok := token.GetClaims([]string{"session_id"})
	if !ok {
		log.Printf("[SetSessionUser] TOKEN_ERROR session_token does not include all required claims\n")
		return nil, errors.New("token error: session_token does not include all required claims")
	}

	token, err = session.IssueTokenWithUserId(claims["session_id"], req.UserId)
	if err != nil {
		log.Printf("[SetSessionUser] TOKEN_ERROR unable to issue new token\n")
		return nil, errors.New("token error: unable to issue new token")
	}

	if err := kinesisclient.Produce(
		claims["session_id"],
		"session_user",
		map[string]interface{}{
			"session_id": claims["session_id"],
			"user_id":    req.UserId,
		},
	); err != nil {
		log.Printf("[SetSessionUser] KINESIS_ERROR unable to produce kiensis record\n")
		return nil, errors.New("kinesis error: unable to produce kiensis record")
	}

	log.Printf("[SetSessionUser] session_id=%s user_id=%s\n", claims["session_id"], req.UserId)

	return &idl.SetSessionUserResponse{
		SessionToken: token.Get(),
	}, nil

}
