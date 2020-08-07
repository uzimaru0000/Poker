package impl

import (
	"github.com/uzimaru0000/poker/model"
	"github.com/uzimaru0000/poker/repository"
)

type inmemUserRepository struct {
	table map[string]model.User
}

func NewInMemUserRepository() repository.UserRepository {
	return &inmemUserRepository{
		table: make(map[string]model.User),
	}
}

func (repo *inmemUserRepository) GetUser(user *model.User) (*model.User, error) {
	u, ok := repo.table[user.ID]
	if ok {
		return &u, nil
	}

	return repo.GetUserWithEmail(user)
}

func (repo *inmemUserRepository) GetUserWithEmail(user *model.User) (*model.User, error) {
	var result *model.User

	for _, value := range repo.table {
		if value.Email == user.Email {
			result = &value
			break
		}
	}

	if result == nil {
		return nil, &repository.NotFound{ID: user.ID}
	}

	return result, nil
}

func (repo *inmemUserRepository) CreateUser(user *model.User) error {
	if _, ok := repo.table[user.ID]; ok {
		return &repository.AlreadyRegistered{}
	}

	repo.table[user.ID] = *user
	return nil
}

func (repo *inmemUserRepository) UpdateUser(user *model.User) (*model.User, error) {
	if _, ok := repo.table[user.ID]; !ok {
		return nil, &repository.NotFound{}
	}

	repo.table[user.ID] = *user
	result := repo.table[user.ID]

	return &result, nil
}

func (repo *inmemUserRepository) DeleteUser(user *model.User) error {
	if _, ok := repo.table[user.ID]; !ok {
		return &repository.NotFound{}
	}

	delete(repo.table, user.ID)
	return nil
}
