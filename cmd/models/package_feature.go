package models

import (
	"time"

	"gorm.io/gorm"
)

type PackageFeature struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	PackageID uint           `json:"package_id"`
	Name      string         `json:"name" gorm:"size:50"`
	Value     string         `json:"value" gorm:"type:text"`
	CreatedAt time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty"`
}
