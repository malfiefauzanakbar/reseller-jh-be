package response

// import (	
// 	"gorm.io/gorm"
// )

type RespLogin struct {
	ID        int64          `json:"id"`
	Username  string         `json:"username"`
	Email     string         `json:"email"`
	Token     string         `json:"token"`	
}
