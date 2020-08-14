package repository

import "github.com/uzimaru0000/poker/model"

type RoomRepository interface {
	CreateRoom(*model.Room) error
	GetRoomByID(string) (*model.Room, error)
	UpdateRoom(*model.Room) error
	DeleteRoomByID(string) error
}
