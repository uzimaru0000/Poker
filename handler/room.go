package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/uzimaru0000/poker/controller"
	"github.com/uzimaru0000/poker/middleware"
	"github.com/uzimaru0000/poker/usecase/impl"
	"net/http"
)

type RoomHandler struct {
	roomController controller.RoomController
}

func NewRoomHandler(roomController controller.RoomController) *RoomHandler {
	return &RoomHandler{
		roomController: roomController,
	}
}

func (r *RoomHandler) CreateRoom(c *gin.Context) {
	rawClaim, exist := c.Get(middleware.ClaimKey)
	if !exist {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "InternalServerError"})
		return
	}

	claim, ok := rawClaim.(*impl.Claim)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "InternalServerError"})
		return
	}

	room, err := r.roomController.CreateRoom(claim.Subject)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed create room"})
		return
	}

	c.JSON(http.StatusOK, room)
}
