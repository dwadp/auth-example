package jwt

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	gojwt "github.com/dgrijalva/jwt-go"
	"github.com/dwadp/auth-example/models"
)

type (
	jwtAuth struct{}

	authClaim struct {
		models.AuthClaims
		gojwt.StandardClaims
	}
)

func NewJWTAuth() JWT {
	return &jwtAuth{}
}

func (j *jwtAuth) Parse(token string, secret string) (models.AuthClaims, error) {
	claims := authClaim{}

	parsedToken, err := gojwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if !parsedToken.Valid {
		return claims.AuthClaims, InvalidToken
	}

	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return claims.AuthClaims, MalformedToken
			} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
				return claims.AuthClaims, ExpiredToken
			} else {
				return claims.AuthClaims, UnknownToken
			}
		}
	}

	return claims.AuthClaims, nil
}

func (j *jwtAuth) Sign(claims models.AuthClaims, secret string, expiry int) (string, error) {
	authClaim := authClaim{
		claims,
		gojwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(expiry) * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "dwadp.com",
		},
	}

	claimInstance := gojwt.NewWithClaims(gojwt.SigningMethodHS256, authClaim)
	token, err := claimInstance.SignedString([]byte(secret))

	if err != nil {
		return "", err
	}

	return token, nil
}
