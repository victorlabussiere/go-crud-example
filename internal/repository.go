package repository

import (
	"fmt"
	"log"

	"github.com/victorlabussiere/go-echo-gorm-example/internal/db/migrations"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "admin"
	password = "senha123"
	dbname   = "shopping_db"
)

var DB *gorm.DB

func InitDB() {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=UTC",
		host, user, password, dbname, port,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Falha ao conectar ao banco de dados: ", err)
	}

	DB = db
	log.Println("Conexão om banco realizada com sucesso.")
	migrations.RunMigrations(db)
	// migrations.RollBackMigrations(db)
}
