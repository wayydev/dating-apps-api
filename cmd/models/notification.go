package models

import (
	"time"

	"gorm.io/gorm"
)

type Notification struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	UserID    uint           `json:"user_id"`
	Title     string         `json:"title" gorm:"size:100"`
	Message   string         `json:"message" gorm:"size:356"`
	Data      interface{}    `json:"data" gorm:"type:json"`
	IsRead    bool           `json:"is_read"`
	CreatedAt time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty"`
}
