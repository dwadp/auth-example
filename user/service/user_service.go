package service

import (
	"github.com/dwadp/auth-example/auth"
	"github.com/dwadp/auth-example/models"
	"github.com/dwadp/auth-example/pkg/hash"
	"github.com/dwadp/auth-example/pkg/jwt"
	"github.com/dwadp/auth-example/user"
	"go.mongodb.org/mongo-driver/mongo"
)

type userService struct {
	userRepo user.Repository
	authRepo auth.Repository
	hash     hash.Hash
	jwt      jwt.JWT
}

func NewUserService(userRepo user.Repository) user.Service {
	return &userService{
		userRepo: userRepo,
	}
}

func (u *userService) GetAll() ([]models.User, error) {
	return u.userRepo.GetAll()
}

func (u *userService) GetByID(id string) (models.User, error) {
	result, err := u.userRepo.GetByID(id)

	if err != nil {
		if err != mongo.ErrNoDocuments {
			return result, err
		}
	}

	return result, nil
}
