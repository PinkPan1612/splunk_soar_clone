package domain

import "time"

type Ticket struct {
	TicketID    string     `gorm:"primaryKey" json:"ticket_id"`
	Title       string     `gorm:"not null" json:"title"`
	Description string     `gorm:"not null" json:"description"`
	CreatedAt   time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   *time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	AssignedTo  string     `json:"assigned_user_id gorm:"foreignKey:UserID" `
	Status      string     `json:"status"`
	Case        string     `json:"case_id" gorm:"foreignKey:CaseID"`
}
