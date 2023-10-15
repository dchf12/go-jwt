package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique"`
	Password []byte `json:"-"`
}

// NewUser is a constructor for User
func NewUser(name, email string, password []byte) *User {
	return &User{
		Name:     name,
		Email:    email,
		Password: password,
	}
}
