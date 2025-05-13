package domain

import (
	"time"
)

type Role struct {
	RoleID      string     `gorm:"primaryKey" json:"role_id"`
	RoleName    string     `gorm:"not null" json:"role_name"`
	Description string     `json:"description"`
	CreatedAt   time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   *time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
