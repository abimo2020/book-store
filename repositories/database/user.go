package database

import (
	"github.com/abimo2020/book-store/config"
	"github.com/abimo2020/book-store/models"
)

func CreateUser(user *models.Users) (err error) {
	if err = config.DB.Save(&user).Error; err != nil {
		return
	}
	return
}
func GetUsers() (users *models.Users, err error) {
	if err = config.DB.Find(&users).Error; err != nil {
		return
	}
	return
}
func GetUserByEmail(email string) (user *models.Users, err error) {
	if err = config.DB.Where("email", email).First(&user).Error; err != nil {
		return
	}
	return
}
