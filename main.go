package main

import (
	"lab-manager-api/config"
	"lab-manager-api/docs"
	"lab-manager-api/models"
	"lab-manager-api/routes"
	"log"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Lab Manager API
// @version 1.0
// @description This is a sample server for a lab manager.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1
type CreateUserRequest struct {
	Name     string          `json:"name" binding:"required"`
	UserType models.UserType `json:"user_type"`
	Password string          `json:"password" binding:"required"`
}

func main() {
	config.ConnectDB()

	r := gin.Default()

	docs.SwaggerInfo.BasePath = "/"

	api := r.Group("/api/v1")
	routes.InitRoutes(api)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
