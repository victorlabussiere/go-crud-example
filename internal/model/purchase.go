package model

type Purchase struct {
	ID         string    `gorm:"primaryKey;size:50" json:"id"`
	CustomerId string    `gorm:"not null;index" json:"customerId"`
	ProductId  string    `gorm:"not null; index" json:"productId"`
	Products   []Product `gorm:"many2many:purchase_products;" json:"customer_putchase_prodcut"`
}
