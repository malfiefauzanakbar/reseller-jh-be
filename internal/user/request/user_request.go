package request

// import (
// 	"gorm.io/gorm"
// )

type ReqRegister struct {
	Username             string `json:"username"`
	Password             string `json:"password"`
	PasswordConfirmation string `json:"password_confirmation"`
	Email                string `json:"email"`
	CaptchaToken         string `form:"captcha_token"`
}

type ReqLogin struct {
	Username     string `json:"username"`
	Password     string `json:"password"`
	CaptchaToken string `form:"captcha_token"`
}
