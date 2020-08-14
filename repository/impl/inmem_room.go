package impl

import (
	"github.com/uzimaru0000/poker/model"
	"github.com/uzimaru0000/poker/repository"
)

type inmemRoomRepository struct {
	table map[string]model.Room
}

func NewInMemRoomRepository() repository.RoomRepository {
	return &inmemRoomRepository{
		table: make(map[string]model.Room),
	}
}

func (i inmemRoomRepository) CreateRoom(room *model.Room) error {
	_, ok := i.table[room.ID]
	if ok {
		return &repository.AlreadyRegistered{
			ID: room.ID,
		}
	}

	i.table[room.ID] = *room
	return nil
}

func (i inmemRoomRepository) GetRoomByID(s string) (*model.Room, error) {
	room, ok := i.table[s]
	if !ok {
		return nil, &repository.NotFound{
			ID: s,
		}
	}

	return &room, nil
}

func (i inmemRoomRepository) UpdateRoom(room *model.Room) error {
	_, ok := i.table[room.ID]
	if !ok {
		return &repository.NotFound{
			ID: room.ID,
		}
	}

	i.table[room.ID] = *room

	return nil
}

func (i inmemRoomRepository) DeleteRoomByID(s string) error {
	_, ok := i.table[s]
	if !ok {
		return &repository.NotFound{
			ID: s,
		}
	}

	delete(i.table, s)
	return nil
}
