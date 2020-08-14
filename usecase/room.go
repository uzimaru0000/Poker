package usecase

import "github.com/uzimaru0000/poker/model"

type RoomUseCase interface {
	CreateRoom(owner *model.User) (*model.Room, error)
	GetRoom(id string) (*model.Room, error)
	JoinRoom(id string, user *model.User) error
	DeleteRoom(id string) error
}
