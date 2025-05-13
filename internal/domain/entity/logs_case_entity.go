package domain

import (
    "time"
)

type Log struct {
    LogID      int64     `gorm:"primaryKey;autoIncrement" json:"log_id"`
    CaseID     int64     `gorm:"not null;index" json:"case_id"` // Liên kết với Case
    Action     string    `gorm:"not null" json:"action"`        // Ví dụ: "created", "updated", "closed"
    PerformedBy string   `gorm:"not null" json:"performed_by"`  // Người thực hiện hành động
    Details    string    `json:"details"`                      // Chi tiết về hành động
    CreatedAt  time.Time `gorm:"autoCreateTime" json:"created_at"`
}