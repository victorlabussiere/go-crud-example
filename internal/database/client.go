package database

import (
	"fmt"
	"log"
	"time"

	"github.com/victorlabussiere/go-echo-gorm-example/internal/database/migrations"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DatabaseClient interface {
	Ready() bool
}

type Client struct {
	DB *gorm.DB
}

const (
	host     = "localhost"
	port     = 5432
	user     = "admin"
	password = "senha123"
	dbname   = "shopping_db"
)

func NewDatabaseClient() (DatabaseClient, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=UTC",
		host, user, password, dbname, port,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
		QueryFields: true,
	})

	if err != nil {
		log.Fatal("Falha ao conectar ao banco de dados: ", err)
		return nil, err
	}

	client := Client{
		DB: db,
	}

	log.Println("Conexão om banco realizada com sucesso.")
	migrations.RunMigrations(db)
	// migrations.RollBackMigrations(db)

	return client, nil
}

func (c Client) Ready() bool {
	var ready string

	tx := c.DB.Raw("SELECT 1 AS READY").Scan(&ready)
	if tx.Error != nil {
		return false
	}

	if ready == "1" {
		return true
	}

	return false
}
