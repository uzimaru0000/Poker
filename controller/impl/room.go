package impl

import (
	"github.com/uzimaru0000/poker/controller"
	"github.com/uzimaru0000/poker/model"
	"github.com/uzimaru0000/poker/usecase"
)

type roomController struct {
	userUseCase usecase.UserUseCase
	roomUseCase usecase.RoomUseCase
}

func NewRoomController(userUseCase usecase.UserUseCase, roomUseCase usecase.RoomUseCase) controller.RoomController {
	return &roomController{
		userUseCase: userUseCase,
		roomUseCase: roomUseCase,
	}
}

func (r roomController) CreateRoom(owner_id string) (*model.Room, error) {
	user, err := r.userUseCase.GetUserByID(owner_id)
	if err != nil {
		return nil, err
	}

	room, err := r.roomUseCase.CreateRoom(user)
	if err != nil {
		return nil, err
	}

	return room, nil
}

func (r roomController) JoinRoom(id string, user_id string) error {
	user, err := r.userUseCase.GetUserByID(user_id)
	if err != nil {
		return err
	}

	err = r.roomUseCase.JoinRoom(id, user)
	if err != nil {
		return err
	}

	return nil
}


