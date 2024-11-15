package model

import (
	"time"

	"gorm.io/gorm"
)

type Status struct {
	ID        int64          `json:"id" gorm:"primaryKey;not null"`
	Name      string         `json:"name" gorm:"size:255;not null"`
	CreatedAt time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

func (Status) TableName() string {
	return "status"
}
