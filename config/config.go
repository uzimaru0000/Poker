package config

import (
	"github.com/uzimaru0000/poker/controller"
	ctrl "github.com/uzimaru0000/poker/controller/impl"
	"github.com/uzimaru0000/poker/lib"
	repo "github.com/uzimaru0000/poker/repository/impl"
	"github.com/uzimaru0000/poker/usecase"
	usecaseImpl "github.com/uzimaru0000/poker/usecase/impl"
)

// Context contains app dependency information.
type Context struct {
	AuthController controller.AuthController
	RoomController controller.RoomController
	Tokenizer usecase.Tokenizer
}

// NewContext is creating Context struct.
func NewContext() *Context {
	authRepo := repo.NewInMemAuthRepository()
	authUseCase := usecaseImpl.NewAuthUseCase(authRepo, &lib.Hash{})

	userRepo := repo.NewInMemUserRepository()
	userUsercase := usecaseImpl.NewUserUseCase(userRepo, &lib.UUID{})

	roomRepo := repo.NewInMemRoomRepository()
	roomUseCase := usecaseImpl.NewRoomUseCase(roomRepo, &lib.UUID{})

	tokenizer := usecaseImpl.NewTokenizer("SECRET")

	return &Context{
		AuthController: ctrl.NewAuthController(authUseCase, userUsercase, tokenizer),
		RoomController: ctrl.NewRoomController(userUsercase, roomUseCase),
		Tokenizer: tokenizer,
	}
}
