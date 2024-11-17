package request

// import (
// 	"gorm.io/gorm"
// )

type ReqRegister struct {
	Username             string `json:"username"`
	Password             string `json:"password"`
	PasswordConfirmation string `json:"password_confirmation"`
	Email                string `json:"email"`
}

type ReqLogin struct {
	Username             string `json:"username"`
	Password             string `json:"password"`	
}