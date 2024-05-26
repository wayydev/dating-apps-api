package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID             uint           `json:"id" gorm:"primaryKey"`
	Name           string         `json:"name" gorm:"size:50 not null"`
	Username       string         `json:"username" gorm:"size:100 not null unique"`
	Email          string         `json:"email" gorm:"type:text not null unique"`
	Password       string         `json:"password,omitempty" gorm:"type:text"`
	Phone          string         `json:"phone" gorm:"size:20"`
	Latitude       float32        `json:"latitude" gorm:"type:float"`
	Longitude      float32        `json:"longitude" gorm:"type:float"`
	Location       string         `json:"string" gorm:"type:geography(POINT,4326)"`
	PlaceOfBirth   string         `json:"place_of_birth" gorm:"size:50"`
	DateOfBirth    time.Time      `json:"date_of_birth"`
	Gender         string         `json:"gender" gorm:"size:20 not null"`
	PhotoProfile   string         `json:"photo_profile" gorm:"type:text"`
	UserPhoto      *UserPhoto     `json:"user_photos"`
	FindOnDistance uint           `json:"find_on_distance" gorm:"default:20"`
	PackageID      uint           `json:"package_id"`
	CreatedAt      time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt      time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt      gorm.DeletedAt `json:"deleted_at,omitempty"`
}
