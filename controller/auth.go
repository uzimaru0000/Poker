package controller

import "github.com/uzimaru0000/poker/model"

type AuthController interface {
	SignUp(email string, name string, password string) (*model.User, error)
	SignIn(email string, password string) (string, error)
}
