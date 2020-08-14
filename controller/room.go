package controller

import "github.com/uzimaru0000/poker/model"

type RoomController interface {
	CreateRoom(owner_id string) (*model.Room, error)
	JoinRoom(id string, user string) error
}
