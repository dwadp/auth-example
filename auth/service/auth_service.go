package service

import (
	"time"

	"github.com/dwadp/auth-example/auth"
	"github.com/dwadp/auth-example/models"
	"github.com/dwadp/auth-example/pkg/hash"
	"github.com/dwadp/auth-example/pkg/jwt"
	"github.com/dwadp/auth-example/pkg/random"
	"github.com/dwadp/auth-example/user"
	"go.mongodb.org/mongo-driver/mongo"
)

type authService struct {
	userRepo user.Repository
	authRepo auth.Repository
	hash     hash.Hash
	jwt      jwt.JWT
}

func NewAuthService(
	userRepo user.Repository,
	authRepo auth.Repository,
	hash hash.Hash,
	jwt jwt.JWT) auth.Service {
	return &authService{
		userRepo: userRepo,
		authRepo: authRepo,
		hash:     hash,
		jwt:      jwt,
	}
}

func (u *authService) Register(user models.User) error {
	hashedPassword, err := u.hash.Make(user.Password)

	if err != nil {
		return err
	}

	user.Password = hashedPassword

	return u.userRepo.Store(user)
}

func (u *authService) Login(userLogin models.UserLogin, secretKey string, expiry int) (models.AuthUser, error) {
	auth := models.AuthUser{}
	result, err := u.userRepo.GetByEmail(userLogin.Email)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return auth, user.NotFound
		}

		return auth, err
	}

	if passwordMatch := u.hash.Check(userLogin.Password, result.Password); !passwordMatch {
		return auth, user.WrongPassword
	}

	sessionID := random.String(20)

	token, err := u.jwt.Sign(models.AuthClaims{
		ID: sessionID,
	}, secretKey, expiry)

	if err != nil {
		return auth, err
	}

	auth.Token = token
	auth.User = result

	authExpiry := time.Duration(expiry) * time.Hour

	if err := u.authRepo.SetJSON(sessionID, result.ID, authExpiry); err != nil {
		return auth, err
	}

	return auth, nil
}

func (a *authService) Logout(sessionID string) error {
	return a.authRepo.Delete(sessionID)
}
