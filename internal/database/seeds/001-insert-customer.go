package seeds

import (
	"log"

	"github.com/victorlabussiere/go-echo-gorm-example/internal/model"
	"gorm.io/gorm"
)

type InsertInitialCustomers struct{}

func (c *InsertInitialCustomers) Up(db *gorm.DB) error {
	log.Println("Inserindo seed 001")
	customers := []model.Customer{
		{
			ID:        1,
			Name:      "Victor Labussiere",
			Email:     "victor.dev@email.com",
			CreatedAt: nil,
			UpdatedAt: nil,
		},
		{
			ID:        2,
			Name:      "Ana Julia",
			Email:     "ana.julia@email.com",
			CreatedAt: nil,
			UpdatedAt: nil,
		},
	}
	for _, customer := range customers {

		if err := db.FirstOrCreate(&customer, customer.ID).Error; err != nil {
			log.Fatalln("Erro na inserção dos dados", err.Error())
			return err
		} else {
			log.Printf("Customer %v resolvido com sucesso", customer.Name)
		}
	}

	return nil
}

func (c *InsertInitialCustomers) Down(db *gorm.DB) error {
	log.Println("Revertendo seed 001")
	return db.Where("id IN ?", []int{1, 2}).Delete(&model.Customer{}).Error
}
