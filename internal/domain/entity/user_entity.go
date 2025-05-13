package domain

import (
	"time"
)

type User struct {
	UserID       int64     `gorm:"primaryKey;autoIncrement" json:"user_id"`
	Username     string    `gorm:"unique;not null" json:"username"`
	Email        string    `gorm:"unique;not null" json:"email"`
	PasswordHash string    `gorm:"not null" json:"password_hash"`
	RoleID       string    `gorm:"not null;type:text" json:"role_id"`
	CreatedAt    time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime" json:"updated_at"`

	Role   *Role   `gorm:"foreignKey:RoleID;`
	Tokens []Token `gorm:"foreignKey:UserID"`
}
