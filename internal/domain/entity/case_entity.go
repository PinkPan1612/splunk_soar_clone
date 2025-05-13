package domain

import (
	"time"
)

type Case struct {
	CaseID      int64     `gorm:"primaryKey;autoIncrement" json:"case_id"`
	Title       string    `gorm:"not null" json:"title"`
	Description string    `json:"description"`
	Status      string    `gorm:"default:'open'" json:"status"`  // Example statuses: open, in_progress, closed
	Severity    string    `gorm:"default:'low'" json:"severity"` // Example priorities: low, medium, high
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
