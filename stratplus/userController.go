package main

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type userController struct {
	db *gorm.DB
}

// Initialize coneccion db and migrations
func (uc *userController) init(db gorm.DB) {
	uc.db = &db
	uc.db.AutoMigrate(&User{})
}

// Verify user exists
func (uc *userController) exists(user User) (exist bool, err error) {
	var count int64

	res := uc.db.Model(&User{}).Where(
		"email  = ?", user.Email,
	).Or(
		"phone = ?", user.Phone,
	).Or(
		"username = ?", user.Username,
	).Count(&count)

	if res.Error != nil {
		return false, res.Error
	}

	return count > 0, nil
}

// Register new user return user
func (uc userController) create(user *User) (tx *gorm.DB) {
	return uc.db.Create(&user)
}

// Login user return token sesion
func (uc userController) login(usr UserLogin, lg *LoginData) error {
	usrDat := User{}
	res := uc.db.Model(&User{}).Where("email = ?", usr.Email).Or("username = ?", usr.Username).First(&usrDat)
	if res.RowsAffected > 0 {
		err := bcrypt.CompareHashAndPassword([]byte(usrDat.Password), []byte(usr.Password))
		if err == nil {
			tokenString := getToken(usrDat.Username, usrDat.ID)
			if tokenString != "" {
				lg.Token = tokenString
				return nil
			}
			return errors.New("there was a problem creating the session")
		}
	}
	return errors.New("wrong username/password")
}

// get user data from id, return user instance
func (uc userController) getUser(id uint64, usrDat *User) error {
	res := uc.db.Model(&User{}).First(&usrDat, id)
	if res.RowsAffected < 1 {
		return errors.New("there is no user matching the ID")
	}
	return nil
}
