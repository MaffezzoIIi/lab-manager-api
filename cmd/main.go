package main

import (
	"lab-manager-api/config"
	"lab-manager-api/models"
	"lab-manager-api/routes"
	"log"

	"github.com/gin-gonic/gin"
)

type CreateUserRequest struct {
	Name     string          `json:"name" binding:"required"`
	UserType models.UserType `json:"user_type" binding:"required"`
	Password string          `json:"password" binding:"required"`
}

func main() {
	config.ConnectDB()

	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
