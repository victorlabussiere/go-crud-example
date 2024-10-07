package migrations

import (
	"log"

	"github.com/victorlabussiere/go-echo-gorm-example/internal/model"
	"gorm.io/gorm"
)

type CreateCategoryMigration struct{}

func (m *CreateCategoryMigration) Up(db *gorm.DB) error {
	log.Println("Criando tabela Category")
	return db.AutoMigrate(&model.Category{})
}

func (m *CreateCategoryMigration) Down(db *gorm.DB) error {
	log.Println("Deletando tabela product")
	return db.Migrator().DropTable(&model.Category{})
}
