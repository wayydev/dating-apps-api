package models

import (
	"time"

	"gorm.io/gorm"
)

type UserPackage struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	UserId    uint           `json:"user_id" gorm:"index"`
	PackageId uint           `json:"package_id" gorm:"index"`
	ExpiresAt time.Time      `json:"expires_at"`
	CreatedAt time.Time      `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}
