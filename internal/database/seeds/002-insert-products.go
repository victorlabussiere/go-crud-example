package seeds

import (
	"log"

	"github.com/victorlabussiere/go-echo-gorm-example/internal/model"
	"gorm.io/gorm"
)

type InsertInitialProducts struct{}

func (c *InsertInitialProducts) Up(db *gorm.DB) error {
	log.Println("Inserindo seeds 002-insert-products")
	products := []model.Product{
		{Name: "Telefone", Value: 3000, CreatedAt: nil, UpdatedAt: nil},
		{Name: "Relogio", Value: 700, CreatedAt: nil, UpdatedAt: nil},
		{Name: "Tênis", Value: 300, CreatedAt: nil, UpdatedAt: nil},
	}

	for _, product := range products {
		if err := db.FirstOrCreate(&product, product.ID).Error; err != nil {
			log.Fatalln("Erro na inserção dos dados", err.Error())
			return err
		} else {
			log.Printf("Producto %s resolvido com sucesso", product.Name)
		}
	}
	return nil
}

func (c *InsertInitialProducts) Down(db *gorm.DB) error {
	log.Println("Revertendo seed 002")
	return db.Where("name IN ?", []string{"Telefone", "Relogio", "Tênis"}).Delete(&model.Product{}).Error
}
