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
}

type EchoServer struct {
	echo *echo.Echo
	DB   database.DatabaseClient
}

func NewEchoServer(db database.DatabaseClient) Server {
	server := &EchoServer{
		echo: echo.New(),
		DB:   db,
	}

	server.registerRoutes()
	return server
}

func (e *EchoServer) Start() error {
	if err := e.echo.Start("localhost:8080"); err != nil && err != http.ErrServerClosed {
		log.Fatalf("server shutdown occurred: %s", err)
		return err
	}

	return nil
}

func (e *EchoServer) registerRoutes() {
	// implementar routes
	e.echo.GET("/readiness", e.Readiness)
	e.echo.GET("/liveness", e.Liveness)
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
