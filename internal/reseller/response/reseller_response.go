package response

import (
	"time"
)

type RespReseller struct {
	ID           int64     `json:"id" gorm:"primaryKey;not null"`
	Fullname     string    `json:"fullname" gorm:"size:255;not null"`
	WhatsappNo   string    `json:"whatsapp_no" gorm:"size:20;not null"`
	WhatsappLink string    `json:"whatsapp_link" gorm:"type:text;not null"`
	Email        string    `json:"email" gorm:"size:255;not null"`
	NIK          string    `json:"nik" gorm:"not null"`
	Address      string    `json:"address" gorm:"type:text;null"`
	StatusID     int64     `json:"status_id" gorm:"not null"`
	StatusName   string    `json:"status_name" gorm:"size:50;not null"`
	CreatedAt    time.Time `json:"created_at" gorm:"autoCreateTime"`
}
