package models

import (
	"time"

	"gorm.io/gorm"
)

type Package struct {
	ID             uint            `json:"id" gorm:"primaryKey"`
	Name           string          `json:"name" gorm:"size:50"`
	Price          uint            `json:"price"`
	PackageFeature *PackageFeature `json:"package_feature"`
	Description    string          `json:"description" gorm:"type:text"`
	CreatedAt      time.Time       `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt      time.Time       `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt      gorm.DeletedAt  `json:"deleted_at,omitempty"`
}
