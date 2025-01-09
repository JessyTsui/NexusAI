package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupUserRouter(server *gin.Engine) {
	userRouter := server.Group("/user")
	userRouter.GET("/info", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello, This is user router!"})
	})
}
