package usecase

import (
	"github.com/abimo2020/book-store/models"
	"github.com/abimo2020/book-store/models/payload"
	"github.com/abimo2020/book-store/repositories/database"
	"golang.org/x/crypto/bcrypt"
)

func Register(req *payload.UserRequest) (resp *models.Users, err error) {
	password, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return
	}
	user := models.Users{
		Name:     req.Name,
		Email:    req.Email,
		Password: string(password),
	}
	if err = database.CreateUser(&user); err != nil {
		return
	}
	resp = &user
	return
}
func Login(req *payload.UserRequest) (resp *models.Users, err error) {
	user, err := database.GetUserByEmail(req.Email)
	if err != nil {
		return
	}
	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return
	}
	resp = user
	return
}
