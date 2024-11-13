package main

import (
	"lab-manager-api/config"
	"lab-manager-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateUserRequest struct {
	Name     string          `json:"name" binding:"required"`
	UserType models.UserType `json:"user_type" binding:"required"`
	Password string          `json:"password" binding:"required"`
}

func main() {
	config.ConnectDB()

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Server is running!")
	})

	r.POST("/create-user", createUserHandler)

	r.Run(":8080")
}

func createUserHandler(c *gin.Context) {
	var req CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := models.HashPassword(req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	user := models.User{
		Name:     req.Name,
		UserType: req.UserType,
		Password: hashedPassword,
	}

}
