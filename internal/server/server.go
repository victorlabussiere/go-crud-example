package server

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/victorlabussiere/go-echo-gorm-example/internal/database"
	"github.com/victorlabussiere/go-echo-gorm-example/internal/model"
)

type Server interface {
	Start() error
	Readiness(ctx echo.Context) error
	Liveness(ctx echo.Context) error

	GetAllCustomers(ctx echo.Context) error
	AddCustomer(ctx echo.Context) error
	GetAllProducts(ctx echo.Context) error
	AddProduct(ctx echo.Context) error

	AddPurchase(ctx echo.Context) error
}

type EchoServer struct {
	echo *echo.Echo
	DB   database.DatabaseClient
}

// A inicialização do servidor depende de uma entidade que implementa a interface de métodos do Server
func NewEchoServer(db database.DatabaseClient) Server {
	server := &EchoServer{ // EchoServer deve implementar Server corretamente
		echo: echo.New(),
		DB:   db,
	}

	server.registerRoutes()
	return server
}

// métodos protegidos
func (e *EchoServer) registerRoutes() {
	// implementar routes
	e.echo.GET("/readiness", e.Readiness)
	e.echo.GET("/liveness", e.Liveness)

	customersGroup := e.echo.Group("/customers")
	customersGroup.GET("", e.GetAllCustomers)
	customersGroup.POST("", e.AddCustomer)

	productsGroup := e.echo.Group("/products")
	productsGroup.GET("", e.GetAllProducts)
	productsGroup.POST("", e.AddProduct)

	purchaseGroup := e.echo.Group("/purchase")
	purchaseGroup.POST("", e.AddPurchase)
}

// implementações da interface Server
func (e *EchoServer) Start() error {
	if err := e.echo.Start("localhost:8080"); err != nil && err != http.ErrServerClosed {
		log.Fatalf("server shutdown occurred: %s", err)
		return err
	}

	return nil
}

func (e EchoServer) Readiness(ctx echo.Context) error {
	ready := e.DB.Ready()
	if ready {
		return ctx.JSON(http.StatusOK, model.Health{
			Status: "OK",
		})
	}

	return ctx.JSON(http.StatusInternalServerError, model.Health{
		Status: "Failure",
	})
}

func (e *EchoServer) Liveness(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, model.Health{Status: "OK"})
}
