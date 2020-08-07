package repository

import "github.com/uzimaru0000/poker/model"

type AuthRepository interface {
	GetHash(*model.Auth) (string, error)
	StoreHash(*model.Auth) error
}
