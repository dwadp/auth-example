package user

import (
	"github.com/dwadp/auth-example/models"
)

type Repository interface {
	GetAll() ([]models.User, error)
	GetByEmail(email string) (models.User, error)
	GetByID(id string) (models.User, error)
	Store(user models.User) error
}
