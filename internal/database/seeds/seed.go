package seeds

import (
	"log"

	"gorm.io/gorm"
)

type Seed interface {
	Up(db *gorm.DB) error
	Down(db *gorm.DB) error
}

func RunSeeds(db *gorm.DB) error {
	log.Println("Inicialiando seeds")
	seeds := []Seed{
		&InsertInitialCustomers{},
		&InsertCategoryData{},
		&InsertInitialProducts{},
	}

	for _, seed := range seeds {
		if err := seed.Up(db); err != nil {
			log.Fatalln("Falha na criação da Seed.", err)
			return err
		}
	}

	log.Println("Seeds executadas com sucesso.")
	return nil
}

func RollbackSeeds(db *gorm.DB) error {
	log.Println("Inicializando rollback da seeds")
	seeds := []Seed{
		&InsertInitialCustomers{},
		&InsertInitialProducts{},
	}

	for _, seed := range seeds {
		if err := seed.Down(db); err != nil {
			log.Fatalln("Falha no rollback da seed", err)
			return err
		}
	}

	log.Println("Seeds executadas com sucesso")
	return nil
}
