package impl

import (
	"github.com/uzimaru0000/poker/model"
	"github.com/uzimaru0000/poker/repository"
)

type inmemAuthRepository struct {
	table map[string]model.Auth
}

func NewInMemAuthRepository() repository.AuthRepository {
	return &inmemAuthRepository{
		table: make(map[string]model.Auth),
	}
}

func (r *inmemAuthRepository) GetHash(auth *model.Auth) (string, error) {
	target, ok := r.table[auth.ID]
	if ok {
		return target.Hash, nil
	}

	return "", &repository.NotFound{ID: auth.ID}
}

func (r *inmemAuthRepository) StoreHash(auth *model.Auth) error {
	_, ok := r.table[auth.ID]

	if ok {
		return &repository.AlreadyRegistered{ID: auth.ID}
	}

	r.table[auth.ID] = *auth

	return nil
}
