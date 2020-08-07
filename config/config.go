package config

import (
	"github.com/uzimaru0000/poker/controller"
	ctrl "github.com/uzimaru0000/poker/controller/impl"
	"github.com/uzimaru0000/poker/lib"
	repo "github.com/uzimaru0000/poker/repository/impl"
	usecase "github.com/uzimaru0000/poker/usecase/impl"
)

// Context contains app dependency information.
type Context struct {
	AuthController controller.AuthController
}

// NewContext is creating Context struct.
func NewContext() *Context {
	authRepo := repo.NewInMemAuthRepository()
	authUsecase := usecase.NewAuthUsecase(authRepo, &lib.Hash{})

	userRepo := repo.NewInMemUserRepository()
	userUsercase := usecase.NewUserUsecase(userRepo, &lib.UUID{})

	tokenizer := usecase.NewTokenizer("SECRET")

	return &Context{
		AuthController: ctrl.NewAuthController(authUsecase, userUsercase, tokenizer),
	}
}
