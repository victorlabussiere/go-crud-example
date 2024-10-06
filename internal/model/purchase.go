package model

type Purchase struct {
	ID         string    `gorm:"primaryKey;size:50"`
	CustomerId string    `gorm:"not null;index"`
	Customer   Customer  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Products   []Product `gorm:"many2many:purchase_products;"`
}
