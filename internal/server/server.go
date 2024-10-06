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

	cg := e.echo.Group("/customers")
	cg.GET("", e.GetAllCustomers)
	cg.POST("", e.AddCustomer)

	pg := e.echo.Group("/products")
	pg.GET("", e.GetAllProducts)
	pg.POST("", e.AddProduct)
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
