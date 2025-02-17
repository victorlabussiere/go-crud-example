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
	GetCustomerById(ctx echo.Context) error
	AddCustomer(ctx echo.Context) error
	UpdateCustomer(ctx echo.Context) error
	DeleteCustomerById(ctx echo.Context) error

	AddProduct(ctx echo.Context) error
	GetAllProducts(ctx echo.Context) error
	GetProductById(ctx echo.Context) error
	GetProductByCategoryId(ctx echo.Context) error

	AddCategory(ctx echo.Context) error
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
	customersGroup.GET("/:id", e.GetCustomerById)
	customersGroup.POST("", e.AddCustomer)
	customersGroup.PUT("/:id", e.UpdateCustomer)
	customersGroup.DELETE("/:id", e.DeleteCustomerById)

	productsGroup := e.echo.Group("/products")
	productsGroup.POST("", e.AddProduct)
	productsGroup.GET("", e.GetAllProducts)
	productsGroup.GET("/:id", e.GetProductById)
	productsGroup.GET("/category/:id", e.GetProductByCategoryId)

	categoriesGroup := e.echo.Group("/categories")
	categoriesGroup.POST("", e.AddCategory)
	categoriesGroup.GET("", e.GetAllCategories)
	categoriesGroup.GET("/:id", e.GetCategoryById)
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
