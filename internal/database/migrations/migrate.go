package migrations

import (
	"log"

	"gorm.io/gorm"
)

type Migration interface {
	Up(db *gorm.DB) error
	Down(db *gorm.DB) error
}

func RunMigrations(db *gorm.DB) error {
	log.Println("Iniciando migrations.")

	migrations := []Migration{
		&CreateCustomerMigration{},
		&CreateProductsMigration{},
		&CreatePurchaseMigration{},
	}

	for _, migration := range migrations {
		if err := migration.Up(db); err != nil {
			log.Println("Migration falhou:\n", err.Error())
			return err
		}
	}

	log.Println("Migrations executadas com sucesso.")
	return nil
}

func RollBackMigrations(db *gorm.DB) error {
	log.Println("Iniciando rollback no banco de dados.")
	migrations := []Migration{
		&CreatePurchaseMigration{},
		&CreateProductsMigration{},
		&CreateCustomerMigration{},
	}

	for _, migration := range migrations {
		if err := migration.Down(db); err != nil {
			log.Println("Migration falhou:\n", err.Error())
			return err
		}
	}

	log.Println("Rollback realizado com sucesso.")
	return nil
}
