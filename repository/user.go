package repository

import "github.com/uzimaru0000/poker/model"

type UserRepository interface {
	GetUser(*model.User) (*model.User, error)
	CreateUser(*model.User) error
	UpdateUser(*model.User) (*model.User, error)
	DeleteUser(*model.User) error
}
