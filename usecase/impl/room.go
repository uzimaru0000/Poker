package impl

import (
	"github.com/uzimaru0000/poker/lib"
	"github.com/uzimaru0000/poker/model"
	"github.com/uzimaru0000/poker/repository"
	"github.com/uzimaru0000/poker/usecase"
)

type roomUseCase struct {
	roomRepo repository.RoomRepository
	idGen lib.UUIDGenerator
}

func NewRoomUseCase(roomRepo repository.RoomRepository, idGen lib.UUIDGenerator) usecase.RoomUseCase {
	return &roomUseCase{
		roomRepo: roomRepo,
		idGen: idGen,
	}
}

func (r roomUseCase) CreateRoom(owner *model.User) (*model.Room, error) {
	id, err := r.idGen.Generate()
	if err != nil {
		return nil, err
	}

	room := &model.Room{
		ID: id,
		Owner: owner.ID,
		Member: []*model.User{ owner },
	}
	err = r.roomRepo.CreateRoom(room)
	if err != nil {
		return nil, err
	}

	return room, nil
}

func (r roomUseCase) GetRoom(id string) (*model.Room, error) {
	room, err := r.roomRepo.GetRoomByID(id)
	if err != nil {
		// TODO: UseCase層でのエラーに変換
		return nil, err
	}

	return room, nil
}

func (r roomUseCase) JoinRoom(id string, user *model.User) error {
	room, err := r.GetRoom(id)
	if err != nil {
		return err
	}

	room.Member = append(room.Member, user)
	err = r.roomRepo.UpdateRoom(room)
	if err != nil {
		return err
	}

	return nil
}

func (r roomUseCase) DeleteRoom(id string) error {
	err := r.roomRepo.DeleteRoomByID(id)
	if err != nil {
		return err
	}

	return nil
}
