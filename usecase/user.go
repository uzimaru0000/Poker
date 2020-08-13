package usecase

import "github.com/uzimaru0000/poker/model"

type UserUsecase interface {
	CreateUser(*model.User) (*model.User, error)
	GetUserByID(string) (*model.User, error)
	GetUserByEmail(string) (*model.User, error)
	UpdateUser(*model.User) (*model.User, error)
	DeleteUser(*model.User) error
}
