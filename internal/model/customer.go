package model

import "time"

type Customer struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	Name      string     `gorm:"not null;size:50" json:"name"`
	Email     string     `gorm:"not null;unique;size:64" json:"email"`
	CreatedAt *time.Time `gorm:"default:current_timestamp" json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
}
