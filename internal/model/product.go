package model

import "time"

type Product struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	Name      string     `gorm:"not null;unique;size:50" json:"name"`
	Value     float64    `gorm:"not null;size:50" json:"value"`
	CreatedAt *time.Time `gorm:"defualt:current_timestamp" json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
}
