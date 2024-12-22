package response

import (
	"time"
)

type RespReseller struct {
	ID              int64     `json:"id" gorm:"primaryKey;not null"`
	Fullname        string    `json:"fullname" gorm:"size:255;not null"`
	WhatsappNo      string    `json:"whatsapp_no" gorm:"size:20;not null"`
	WhatsappLink    string    `json:"whatsapp_link" gorm:"type:text;not null"`
	Email           string    `json:"email" gorm:"size:255;not null"`
	Address         string    `json:"address" gorm:"type:text;null"`
	StatusID        int64     `json:"status_id" gorm:"not null"`
	StatusName      string    `json:"status_name" gorm:"size:50;not null"`
	WhereDidYouKnow string    `form:"where_did_you_know"`
	ReasonsToJoin   string    `form:"reasons_to_join"`
	CreatedAt       time.Time `json:"created_at" gorm:"autoCreateTime"`
}

type RespResellerDashboard struct {
	Total  int64 `json:"total" gorm:"not null"`
	Unread int64 `json:"unread" gorm:"not null"`
	Read   int64 `json:"read" gorm:"not null"`
}

type ResellerChart struct {
	Date  string
	Count int
}

type RespResellerChart struct {
	Categories []string `json:"categories"`
	Data       []int    `json:"data"`
}

type RespExportReseller struct {
	Filename string `json:"filename"`
}
