package models

import (
	"encoding/json"
	"fmt"

	"golang.org/x/crypto/bcrypt"
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

func NewUser(name string, userType UserType, password string) (User, error) {
	hasedPassword, err := HashPassword(password)
	if err != nil {
		return User{}, err
	}

	return User{Name: name, UserType: userType, Password: hasedPassword}, nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
