package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/uzimaru0000/poker/config"
	"github.com/uzimaru0000/poker/handler"
)

var context *config.Context

func init() {
	context = config.NewContext()
}

func main() {
	r := gin.Default()

	authHandler := handler.NewAuthHandler(context.AuthController)

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "server is running"})
	})

	r.POST("/signup", authHandler.SignUp)
	r.POST("/signin", authHandler.SignIn)

	r.Run()
}
