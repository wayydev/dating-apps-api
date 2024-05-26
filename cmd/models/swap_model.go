package models

import (
	"time"

	"gorm.io/gorm"
)

type Swap struct {
	ID         uint           `json:"id" gorm:"primaryKey"`
	UserID     uint           `json:"user_id"`
	User       *User          `json:"user"`
	SwapUserID uint           `json:"swap_user_id"`
	SwapUser   *User          `json:"swap_user"`
	Like       bool           `json:"like"`
	CreatedAt  time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt  time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at,omitempty"`
}
