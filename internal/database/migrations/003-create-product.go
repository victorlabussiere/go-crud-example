package migrations

import (
	"log"

	"github.com/victorlabussiere/go-echo-gorm-example/internal/model"
	"gorm.io/gorm"
)

type CreateProductsMigration struct{}

func (m *CreateProductsMigration) Up(db *gorm.DB) error {
	log.Println("Migrando tabela product")
	return db.AutoMigrate(&model.Product{})
}

func (m *CreateProductsMigration) Down(db *gorm.DB) error {
	log.Println("Deletando tabela product")
	return db.Migrator().DropTable(&model.Product{})
}
