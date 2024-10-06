package model

import "time"

type Customer struct {
	ID        string     `gorm:"primaryKey;size:50" json:"customer_id"`
	Name      string     `gorm:"not null;size:50"`
	Email     string     `gorm:"not null;unique;size:64"`
	CreatedAt *time.Time `gorm:"default:current_timestamp"`
	UpdatedAt *time.Time
}
