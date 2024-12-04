package res

type CreateUserResponse struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	UserType int    `json:"user_type"`
}
