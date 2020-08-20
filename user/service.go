package user

import (
	"github.com/dwadp/auth-example/models"
)

type Service interface {
	GetAll() ([]models.User, error)
	GetByID(id string) (models.User, error)
}
