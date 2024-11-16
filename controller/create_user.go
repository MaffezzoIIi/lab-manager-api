package controller

import (
	"lab-manager-api/controller/model/req"
	"lab-manager-api/models"
	"lab-manager-api/rest_err"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var userReq req.CreateUserRequest

	if err := c.ShouldBindJSON(&userReq); err != nil {
		restErr := rest_err.NewRestErr("invalid json body", http.StatusBadRequest, "bad_request", nil)
		c.JSON(restErr.Status, restErr)
		return
	}

	user, err := models.NewUser(userReq.Name, userReq.UserType, userReq.Password)
	if err != nil {
		restErr := rest_err.NewRestErr("error creating user", http.StatusInternalServerError, "internal_server_error", nil)
		c.JSON(restErr.Status, restErr)
		return
	}

	var userRes = models.User{
		ID:       "1",
		Name:     user.Name,
		UserType: user.UserType,
	}

	c.JSON(http.StatusCreated, userRes)
}
