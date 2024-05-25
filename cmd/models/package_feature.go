package models

import (
	"time"

	"gorm.io/gorm"
)

type PackageFeature struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	PackageID uint           `json:"package_id" gorm:"index"`
	Name      string         `json:"name" gorm:"size:50"`
	Value     string         `json:"value" gorm:"type:text"`
	CreatedAt time.Time      `json:"created_at" gorm:"default:CURENT_TIMESTAMP"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"default:CURENT_TIMESTAMP"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}
