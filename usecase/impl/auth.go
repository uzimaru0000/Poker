package impl

import (
	"errors"

	"github.com/uzimaru0000/poker/lib"
	"github.com/uzimaru0000/poker/model"
	"github.com/uzimaru0000/poker/repository"
	"github.com/uzimaru0000/poker/usecase"
)

type authUsecase struct {
	authRepo repository.AuthRepository
	hashGen  lib.HashGenerator
}

func NewAuthUseCase(authRepo repository.AuthRepository, hashGen lib.HashGenerator) usecase.AuthUseCase {
	return &authUsecase{
		authRepo: authRepo,
		hashGen:  hashGen,
	}
}

func (au *authUsecase) SignUp(id string, password string) error {
	hash, err := au.hashGen.Create(password)
	if err != nil {
		return err
	}

	return au.authRepo.StoreHash(&model.Auth{ID: id, Hash: hash})
}

func (au *authUsecase) SignIn(id string, password string) error {
	hash, err := au.authRepo.GetHash(&model.Auth{ID: id})
	if err != nil {
		return err
	}

	err = au.hashGen.Verify(hash, password)
	if err != nil {
		return errors.New("Unauthorized")
	}

	return nil
}
