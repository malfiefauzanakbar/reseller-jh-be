package model

import (
	"time"

	"gorm.io/gorm"
)

type Reseller struct {
	ID           int64          `json:"id" gorm:"primaryKey;not null"`
	Fullname     string         `json:"fullname" gorm:"size:255;not null"`
	WhatsappNo   string         `json:"whatsapp_no" gorm:"size:20;not null"`
	WhatsappLink string         `json:"whatsapp_link" gorm:"type:text;not null"`
	Email        string         `json:"email" gorm:"size:255;not null"`
	NIK          string         `json:"nik" gorm:"not null"`
	Address      string         `json:"address" gorm:"type:text;null"`
	StatusID     int64          `json:"status_id" gorm:"not null"`
	CreatedAt    time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt    time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

func (Reseller) TableName() string {
	return "resellers"
}
