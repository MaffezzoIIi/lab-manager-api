package models

import "golang.org/x/crypto/bcrypt"

type UserType int

const (
	Admin UserType = iota
	Professor
)

type User struct {
	ID       string
	Name     string
	UserType UserType
	Password string
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
