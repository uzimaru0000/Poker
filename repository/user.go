package repository

import (
	"github.com/uzimaru0000/poker/model"
)

type UserRepository interface {
	GetUserByID(string) (*model.User, error)
	GetUserByEmail(string) (*model.User, error)
	CreateUser(*model.User) error
	UpdateUser(*model.User) (*model.User, error)
	DeleteUser(*model.User) error
}
