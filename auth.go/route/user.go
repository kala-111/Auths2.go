package routes

import (
	"gin-go-api/controller"

	"github.com/gin-gonic/gin"
)

func CalRoute(router *gin.Engine) {
	router.GET("/", controller.GetAuths)
	router.POST("/", controller.CreateAuths)
	router.DELETE("/:id", controller.DeleteUser)
	router.PUT("/:id", controller.UpdateUser)
}
