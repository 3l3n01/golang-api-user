package main

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// User model
type User struct {
	gorm.Model
	Username string `json:"username" binding:"required"`
	Phone    string `json:"phone" binding:"required,phone"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,password"`
}

// User login data
type UserLogin struct {
	Username string `json:"username" binding:"required_if=Email ''"`
	Email    string `json:"email" binding:"required_if=Username ''"`
	Password string `json:"password" binding:"required"`
}

// Error format
type ErrorFormatMsg struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// Token login
type LoginData struct {
	Token string `json:"token"`
}

// Header Auth
type AuthHeader struct {
	Authorization string `json:"Authorization"`
}

// Hook encrypt password
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(u.Password), 14)
	u.Password = string(bytes)

	if err != nil {
		return errors.New("it was not possible to generate the secure password")
	}
	return
}
