package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/victorlabussiere/go-echo-gorm-example/internal/database/migrations"
	"github.com/victorlabussiere/go-echo-gorm-example/internal/database/seeds"
	"github.com/victorlabussiere/go-echo-gorm-example/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DatabaseClient interface {
	Ready() bool

	AddCustomer(ctx context.Context, customer *model.Customer) (*model.Customer, error)
	GetAllCustomers(ctx context.Context) ([]model.Customer, error)
	GetCustomerById(ctx context.Context, ID uint) (*model.Customer, error)
	UpdateCustomer(ctx context.Context, customer *model.Customer) (*model.Customer, error)
	DeleteCustomerById(ctx context.Context, ID uint) error

	GetAllProducts(ctx context.Context) ([]model.Product, error)
	AddProduct(ctx context.Context, product *model.Product) (*model.Product, error)
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
	seeds.RunSeeds(db)
	// seeds.RollbackSeeds(db)           // rollback seeds -> manter comentado
	// migrations.RollBackMigrations(db) // rollback migrations -> manter comentado

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
