package domain

import (
	"time"
)

type Token struct {
	TokenID      int64     `gorm:"primaryKey;autoIncrement" json:"token_id"`
	AccessToken  string    `gorm:"not null" json:"access_token"`
	RefreshToken string    `gorm:"not null" json:"refresh_token"`
	Expiry       time.Time `gorm:"not null" json:"expiry"`
	UserID       int64     `gorm:"not null" json:"user_id"`
	CreatedAt    time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime" json:"updated_at"`

	User *User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
