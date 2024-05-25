package models

import "time"

type Swap struct {
	ID         uint      `json:"id" gorm:"primaryKey"`
	UserID     uint      `json:"user_id"`
	User       *User     `json:"user"`
	SwapUserID uint      `json:"swap_user_id"`
	SwapUser   *User     `json:"swap_user"`
	Like       bool      `json:"like"`
	CreatedAt  time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
}
