package jwt

import (
	"errors"

	"github.com/dwadp/auth-example/models"
)

type JWT interface {
	Parse(token, secret string) (models.AuthClaims, error)
	Sign(claims models.AuthClaims, secret string, expiry int) (string, error)
}

var (
	InvalidToken   = errors.New("jwt token is invalid or bad formatted")
	MalformedToken = errors.New("jwt token is malformed")
	ExpiredToken   = errors.New("jwt token is expired")
	UnknownToken   = errors.New("unkown jwt token")
)
