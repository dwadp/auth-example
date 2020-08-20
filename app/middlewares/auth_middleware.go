package middlewares

import (
	"errors"
	"net/http"
	"strings"

	"github.com/dwadp/auth-example/auth"
	authRepository "github.com/dwadp/auth-example/auth"
	"github.com/dwadp/auth-example/pkg/jwt"
	"github.com/dwadp/auth-example/response"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type AuthMiddleware struct {
	jwt      jwt.JWT
	authRepo authRepository.Repository
}

func NewAuthMiddleware(jwt jwt.JWT, authRepo authRepository.Repository) *AuthMiddleware {
	return &AuthMiddleware{
		jwt:      jwt,
		authRepo: authRepo,
	}
}

func (a *AuthMiddleware) Handle() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorization := c.GetHeader("Authorization")

		if authorization == "" {
			c.Abort()
			c.JSON(
				http.StatusUnauthorized,
				response.Error("Unauthenticated", errors.New("bearer token is required")),
			)
			return
		}

		if !strings.Contains(authorization, "Bearer ") {
			c.Abort()
			c.JSON(
				http.StatusUnauthorized,
				response.Error("Unauthenticated", errors.New("bad formatted bearer token")),
			)
			return
		}

		header := strings.Split(authorization, " ")

		authUser, err := a.jwt.Parse(header[1], viper.GetString("auth.secret"))

		if err != nil {
			c.Abort()
			c.JSON(
				http.StatusUnauthorized,
				response.Error("Unauthenticated", err),
			)
			return
		}

		session, err := a.authRepo.Get(authUser.ID, auth.String)

		if err != nil {
			c.Abort()
			c.JSON(
				http.StatusUnauthorized,
				response.Error("Unauthenticated", err),
			)
			return
		}

		if session == nil {
			c.Abort()
			c.JSON(
				http.StatusUnauthorized,
				response.Error("Unauthenticated", errors.New("your session is over")),
			)
			return
		}

		c.Set("X-USER-ID", session)
		c.Set("X-SESSION-ID", authUser.ID)
	}
}
