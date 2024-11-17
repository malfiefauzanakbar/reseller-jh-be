// /api/v1/users/user.go

package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        int64          `json:"id" gorm:"primaryKey;not null"`
	Username  string         `json:"username" gorm:"size:50;not null;unique"`
	Password  string         `json:"password" gorm:"size:255;not null"`
	Email     string         `json:"email" gorm:"size:100;not null;unique"`
	CreatedAt time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

func (User) TableName() string {
	return "users"
}
