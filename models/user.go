package models

import "gorm.io/gorm"

type Users struct {
	gorm.Model
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	IsAdmin  bool   `json:"is_admin" form:"is_admin" gorm:"type:bool;default:false"`
}
