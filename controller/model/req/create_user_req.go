package req

type CreateUserRequest struct {
	Name     string `json:"name" binding:"required"`
	UserType int    `json:"user_type"`
	Password string `json:"password" binding:"required"`
}
