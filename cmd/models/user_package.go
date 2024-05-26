package models

import (
	"time"

	"gorm.io/gorm"
)

type UserPackage struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	UserId    uint           `json:"user_id"`
	PackageId uint           `json:"package_id"`
	ExpiresAt time.Time      `json:"expires_at"`
	CreatedAt time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty"`
}
