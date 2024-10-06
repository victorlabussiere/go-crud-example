package model

import "time"

type Product struct {
	ID        string     `gorm:"primaryKey;size:50" json:"product_id"`
	Name      string     `gorm:"not null;unique;size:50"`
	Value     string     `gorm:"not null;size:50"`
	CreatedAt *time.Time `gorm:"defualt:current_timestamp"`
}
