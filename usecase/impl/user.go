package impl

import (
	"errors"

	"github.com/uzimaru0000/poker/lib"
	"github.com/uzimaru0000/poker/model"
	"github.com/uzimaru0000/poker/repository"
	"github.com/uzimaru0000/poker/usecase"
)

type userUseCase struct {
	userRepo      repository.UserRepository
	uuidGenerator lib.UUIDGenerator
}

func NewUserUseCase(userRepo repository.UserRepository, uuidGenerator lib.UUIDGenerator) usecase.UserUseCase {
	return &userUseCase{
		userRepo:      userRepo,
		uuidGenerator: uuidGenerator,
	}
}

func (uu *userUseCase) CreateUser(user *model.User) (*model.User, error) {
	id, err := uu.uuidGenerator.Generate()
	if err != nil {
		return nil, err
	}

	user.ID = id
	err = uu.userRepo.CreateUser(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (uu *userUseCase) GetUserByID(id string) (*model.User, error) {
	return uu.userRepo.GetUserByID(id)
}

func (uu *userUseCase) GetUserByEmail(email string) (*model.User, error) {
	return uu.userRepo.GetUserByEmail(email)
}

func (uu *userUseCase) UpdateUser(user *model.User) (*model.User, error) {
	return nil, errors.New("Failed")
}

func (uu *userUseCase) DeleteUser(user *model.User) error {
	return errors.New("Failed")
}
