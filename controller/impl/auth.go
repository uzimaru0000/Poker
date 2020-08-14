package impl

import (
	"fmt"

	"github.com/uzimaru0000/poker/controller"
	"github.com/uzimaru0000/poker/model"
	"github.com/uzimaru0000/poker/usecase"
	"github.com/uzimaru0000/poker/usecase/impl"
)

type authController struct {
	authUseCase usecase.AuthUseCase
	userUseCase usecase.UserUseCase
	tokenizer   usecase.Tokenizer
}

func NewAuthController(authUseCase usecase.AuthUseCase, userUseCase usecase.UserUseCase, tokenizer usecase.Tokenizer) controller.AuthController {
	return &authController{authUseCase: authUseCase, userUseCase: userUseCase, tokenizer: tokenizer}
}

func (ac *authController) SignUp(email string, name string, password string) (*model.User, error) {
	user, err := ac.userUseCase.CreateUser(&model.User{Name: name, Email: email})
	if err != nil {
		return nil, err
	}

	err = ac.authUseCase.SignUp(user.ID, password)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (ac *authController) SignIn(email string, password string) (string, error) {
	user, err := ac.userUseCase.GetUserByEmail(email)
	if err != nil {
		fmt.Printf("GetUser\n")
		return "", err
	}

	err = ac.authUseCase.SignIn(user.ID, password)
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
