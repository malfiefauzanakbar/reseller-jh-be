package request

import (
	"gorm.io/gorm"
)

type ReqRegister struct {
	Username             string `json:"username" gorm:"size:50;not null;unique"`
	Password             string `json:"password" gorm:"size:255;not null"`
	PasswordConfirmation string `json:"password_confirmation" gorm:"size:255;not null"`
	Email                string `json:"email" gorm:"size:100;not null;unique"`
}

var DB *gorm.DB
