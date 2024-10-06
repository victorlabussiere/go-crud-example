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
			Name:      "Victor Labussiere",
			Email:     "victor.dev@email.com",
			CreatedAt: nil,
			UpdatedAt: nil,
		},
		{
			Name:      "Ana Julia",
			Email:     "ana.julia@email.com",
			CreatedAt: nil,
			UpdatedAt: nil,
		},
	}
	for _, customer := range customers {
		// Usamos o e-mail como condição para a busca
		if err := db.Where("email = ?", customer.Email).First(&customer).Error; err != nil {
			if err == gorm.ErrRecordNotFound { // Se o cliente não foi encontrado, cria um novo
				if err := db.Create(&customer).Error; err != nil {
					log.Fatalln("Erro na inserção dos dados:", err.Error())
					return err
				}
				log.Printf("Customer %v inserido com sucesso", customer.Name)
			} else {
				log.Fatalln("Erro de validação de customer:", err.Error())
				return err
			}
		} else {
			log.Printf("Customer %v já existe, não inserido", customer.Name)
		}
	}

	return nil
}

func (c *InsertInitialCustomers) Down(db *gorm.DB) error {
	log.Println("Revertendo seed 001")
	return db.Where("email IN ?", []string{"victor.dev@email.com", "ana.julia@email.com"}).Delete(&model.Customer{}).Error
}
