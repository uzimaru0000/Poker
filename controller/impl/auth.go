package impl

import (
	"fmt"

	"github.com/uzimaru0000/poker/controller"
	"github.com/uzimaru0000/poker/model"
	"github.com/uzimaru0000/poker/usecase"
	"github.com/uzimaru0000/poker/usecase/impl"
)

type authController struct {
	authUsecase usecase.AuthUsecase
	userUsecase usecase.UserUsecase
	tokenizer   usecase.Tokenizer
}

func NewAuthController(authUsecase usecase.AuthUsecase, userUsecase usecase.UserUsecase, tokenizer usecase.Tokenizer) controller.AuthController {
	return &authController{authUsecase: authUsecase, userUsecase: userUsecase, tokenizer: tokenizer}
}

func (ac *authController) SignUp(email string, name string, password string) (*model.User, error) {
	user, err := ac.userUsecase.CreateUser(&model.User{Name: name, Email: email})
	if err != nil {
		return nil, err
	}

	err = ac.authUsecase.SignUp(user.ID, password)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (ac *authController) SignIn(email string, password string) (string, error) {
	user, err := ac.userUsecase.GetUser(&model.User{Email: email})
	if err != nil {
		fmt.Printf("GetUser\n")
		return "", err
	}

	err = ac.authUsecase.SignIn(user.ID, password)
	if err != nil {
		fmt.Printf("SignIn\n")
		return "", err
	}

	token, err := ac.tokenizer.CreateToken(impl.Claim{Subject: user.ID})
	if err != nil {
		fmt.Printf("CreateToken\n")
		return "", err
	}

	return token, nil
}
