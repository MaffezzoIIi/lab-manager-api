package res

import "lab-manager-api/models"

type CreateUserResponse struct {
	ID       int             `json:"id"`
	Name     string          `json:"name"`
	UserType models.UserType `json:"user_type"`
}
