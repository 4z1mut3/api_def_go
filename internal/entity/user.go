package entity

import (
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       string `json:id`
	Name     string `json:name`
	Email    string `json:email`
	Password string `json:password`
}

func NewUser(name, email, password string) (*User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return &User{
		ID:       "id",
		Name:     name,
		Email:    email,
		Password: string(hash),
	}, nil
}
func (u *User) validatePassword(password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
