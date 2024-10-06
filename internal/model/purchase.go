package model

type Purchase struct {
	ID         uint      `gorm:"primaryKey;size:50;AUTO_INCREMENT" json:"id"`
	CustomerId int       `gorm:"not null;index" json:"customerId"`
	ProductId  int       `gorm:"not null; index" json:"productId"`
	Products   []Product `gorm:"many2many:purchase_products;" json:"customer_putchase_prodcut"`
}
