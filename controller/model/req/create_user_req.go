package req

import "lab-manager-api/models"

type CreateUserRequest struct {
	Name     string          `json:"name" binding:"required"`
	UserType models.UserType `json:"user_type"`
	Password string          `json:"password" binding:"required"`
}
