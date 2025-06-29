package controller

import (
	"lab-manager-api/controller/model/req"

	"lab-manager-api/models/user"
	"lab-manager-api/rest_err"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateUser godoc
// @Summary Create a new user
// @Description Create a new user
// @Tags users
// @Accept  json
// @Produce  json
// @Param user body req.CreateUserRequest true "User object that needs to be created"
// @Success 201 {object} CreateUserResponse
// @Failure 400 {object} rest_err.RestErr
// @Router /api/v1/users/create [post]
func CreateUser(c *gin.Context) {
	var userReq req.CreateUserRequest

	if err := c.ShouldBindJSON(&userReq); err != nil {
		log.Printf("Error binding JSON: %v\n", err)
		restErr := rest_err.NewRestErr(err.Error(), http.StatusBadRequest, "bad_request", nil)
		c.JSON(restErr.Status, restErr)
		return
	}

	user, err := user.NewUser(userReq.Name, user.UserType(userReq.UserType), userReq.Password)
	if err != nil {
		restErr := rest_err.NewRestErr("error creating user", http.StatusInternalServerError, "internal_server_error", []rest_err.Causes{{Message: err.Error()}})
		c.JSON(restErr.Status, restErr)
		return
	}

	c.JSON(http.StatusCreated, ModelUserToResponse(user))
}

type CreateUserResponse struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	UserType int    `json:"user_type"`
}

func ModelUserToResponse(user user.User) CreateUserResponse {
	userType := int(user.UserType)

	return CreateUserResponse{
		ID:       user.ID,
		Name:     user.Name,
		UserType: userType,
	}
}
