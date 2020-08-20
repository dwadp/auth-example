package auth

import "github.com/dwadp/auth-example/models"

type Service interface {
	Register(user models.User) error
	Login(userLogin models.UserLogin, secretKey string, expiry int) (models.AuthUser, error)
	Logout(sessionID string) error
}
