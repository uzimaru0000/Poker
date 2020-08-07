package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/uzimaru0000/poker/config"
)

var context *config.Context

func init() {
	context = config.NewContext()
}

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "server is running"})
	})

	r.POST("/signup", signUp)
	r.POST("/signin", signIn)

	r.Run()
}

func signUp(c *gin.Context) {
	var body struct {
		Email    string `"json":email`
		Name     string `"json":name`
		Password string `"json":password`
	}

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "BadRequest"})
		return
	}

	user, err := context.AuthController.SignUp(body.Email, body.Name, body.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "InternalServerError"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func signIn(c *gin.Context) {
	var body struct {
		Email    string `"json":email`
		Password string `"json":email`
	}

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "BadRequest"})
		return
	}

	token, err := context.AuthController.SignIn(body.Email, body.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
