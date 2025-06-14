// internal/service/user_service.go
package service

import (
	"errors"
	"todo-app/internal/db"
	"todo-app/internal/model"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(user *model.User) error {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	user.Password = string(hashed)
	return db.DB.Create(user).Error
}

func LoginUser(email, password string) (string, error) {
	var user model.User
	if err := db.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return "", errors.New("user not found")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("invalid password")
	}
	return GenerateJWT(user.ID)
}
