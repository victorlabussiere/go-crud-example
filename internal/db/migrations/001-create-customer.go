package migrations

import (
	"log"

	"github.com/victorlabussiere/go-echo-gorm-example/internal/model"
	"gorm.io/gorm"
)

type CreateCustomerMigration struct{}

func (m *CreateCustomerMigration) Up(db *gorm.DB) error {
	log.Println("Migrando tabela customer")
	return db.AutoMigrate(&model.Customer{})
}

func (m *CreateCustomerMigration) Down(db *gorm.DB) error {
	log.Println("Deletando tabela customer")
	return db.Migrator().DropTable(&model.Customer{})
}
