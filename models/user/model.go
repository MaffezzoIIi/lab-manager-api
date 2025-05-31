package user

import (
	"encoding/json"
	"fmt"
)

type UserType int

const (
	Admin = iota
	Professor
)

func (u *UserType) UnmarshalJSON(data []byte) error {
	var intValue int
	if err := json.Unmarshal(data, &intValue); err != nil {
		return fmt.Errorf("UserType should be an integer, got %s", string(data))
	}

	switch intValue {
	case int(Admin), int(Professor):
		*u = UserType(intValue)
	default:
		return fmt.Errorf("invalid UserType value: %d", intValue)
	}

	return nil
}

type User struct {
	ID       string   `json:"id"`
	Name     string   `json:"name"`
	UserType UserType `json:"user_type"`
	Password string   `json:"password"`
}
