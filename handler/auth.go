package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/uzimaru0000/poker/controller"
)

type AuthHandler struct {
	authController controller.AuthController
}

func NewAuthHandler(authController controller.AuthController) *AuthHandler {
	return &AuthHandler{
		authController: authController,
	}
}

func (h *AuthHandler) SignUp(c *gin.Context) {
	var body struct {
		Email    string `"json":email`
		Name     string `"json":name`
		Password string `"json":password`
	}

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "BadRequest"})
		return
	}

	user, err := h.authController.SignUp(body.Email, body.Name, body.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "InternalServerError"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *AuthHandler) SignIn(c *gin.Context) {
	var body struct {
		Email    string `"json":email`
		Password string `"json":email`
	}

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "BadRequest"})
		return
	}

	token, err := h.authController.SignIn(body.Email, body.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
