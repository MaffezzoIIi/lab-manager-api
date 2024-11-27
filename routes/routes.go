package routes

import (
	"lab-manager-api/controller"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup) {

	r.POST("/users/create", controller.CreateUser)

	r.POST("/labs/create", controller.CreateLab)
	r.GET("/labs/:id", controller.GetLab)
	r.GET("/labs", controller.GetLabs)
	r.PUT("/labs/:id", controller.UpdateLab)
	r.DELETE("/labs/:id", controller.DeleteLab)
}
