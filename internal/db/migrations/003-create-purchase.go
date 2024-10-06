package migrations

import (
	"log"

	"github.com/victorlabussiere/go-echo-gorm-example/internal/model"
	"gorm.io/gorm"
)

type CreatePurchaseMigration struct{}

func (m *CreatePurchaseMigration) Up(db *gorm.DB) error {
	log.Println("Migrando tabelas purchase e purchase_products")
	return db.AutoMigrate(&model.Purchase{})
}

func (m *CreatePurchaseMigration) Down(db *gorm.DB) error {
	log.Println("Deletando tabelas purchase e purchase_products")
	return db.Migrator().DropTable(&model.Purchase{})
}
