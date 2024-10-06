package model

import "time"

type Product struct {
	ID        uint       `gorm:"primaryKey;size:50;AUTO_INCREMENT" json:"id"`
	Name      string     `gorm:"not null;unique;size:50" json:"name"`
	Value     string     `gorm:"not null;size:50" json:"value"`
	CreatedAt *time.Time `gorm:"defualt:current_timestamp" json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
}
