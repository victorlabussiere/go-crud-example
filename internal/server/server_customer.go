package server

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/victorlabussiere/go-echo-gorm-example/internal/model"
)

func (s *EchoServer) GetAllCustomers(ctx echo.Context) error {
	customers, err := s.DB.GetAllCustomers(ctx.Request().Context())
	if err != nil {
		log.Fatal(err.Error())
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusOK, customers)
}

func (s *EchoServer) AddCustomer(ctx echo.Context) error {
	customer := new(model.Customer)
	if err := ctx.Bind(customer); err != nil {
		return ctx.JSON(http.StatusUnsupportedMediaType, err)
	}

	customer, err := s.DB.AddCustomer(ctx.Request().Context(), customer)
	if err != nil {
		log.Fatal(err.Error())
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusCreated, customer)
}
