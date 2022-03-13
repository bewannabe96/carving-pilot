package session

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

type Token struct {
	value  string
	claims jwt.MapClaims
}

func (t *Token) Get() string {

	return t.value

}

func (t *Token) GetClaims(keys []string) (map[string]string, bool) {

	claims := map[string]string{}
	ok := true

	for _, key := range keys {
		value, o := t.claims[key].(string)
		claims[key] = value
		ok = ok && o
	}

	return claims, ok

}

func issueToken(claims jwt.MapClaims) (*Token, error) {

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	sessionToken, err := token.SignedString([]byte(os.Getenv("SESSION_JWT_SECRET")))
	if err != nil {
		return nil, err
	}

	return &Token{
		value:  sessionToken,
		claims: claims,
	}, nil

}

func extractClaims(t string) (map[string]interface{}, error) {
	token, err := jwt.Parse(
		t,
		func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("token error: unable verify token")
			}

			return []byte(os.Getenv("SESSION_JWT_SECRET")), nil
		},
	)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, errors.New("token error: token is not valid")
	}

	return claims, nil
}

func NewToken(token string) (*Token, error) {

	claims, err := extractClaims(token)
	if err != nil {
		return nil, err
	}

	return &Token{
		value:  token,
		claims: claims,
	}, nil

}

func IssueToken(sessionId string) (*Token, error) {

	return issueToken(jwt.MapClaims{
		"session_id": sessionId,
		"iat":        time.Now().Unix(),
	})

}

func IssueTokenWithUserId(sessionId string, userId string) (*Token, error) {

	return issueToken(jwt.MapClaims{
		"session_id": sessionId,
		"user_id":    userId,
		"iat":        time.Now().Unix(),
	})

}
