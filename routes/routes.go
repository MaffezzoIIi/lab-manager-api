package routes

import (
	"lab-manager-api/controller"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup) {

	r.POST("/users/create", controller.CreateUser)
}
