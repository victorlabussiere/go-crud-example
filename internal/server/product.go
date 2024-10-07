package server

import (
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/victorlabussiere/go-echo-gorm-example/internal/database/dberrors"
	"github.com/victorlabussiere/go-echo-gorm-example/internal/model"
)

func (s *EchoServer) AddProduct(ctx echo.Context) error {
	product := new(model.Product)
	if err := ctx.Bind(&product); err != nil {
		return ctx.JSON(http.StatusUnsupportedMediaType, err)
	}

	product, err := s.DB.AddProduct(ctx.Request().Context(), product)
	if err != nil {
		switch err.(type) {
		case *dberrors.ConflictError:
			return ctx.JSON(http.StatusConflict, err)
		default:
			return ctx.JSON(http.StatusInternalServerError, err)
		}
	}

	return ctx.JSON(http.StatusCreated, product)
}
func (s *EchoServer) GetAllProducts(ctx echo.Context) error {
	products, err := s.DB.GetAllProducts(ctx.Request().Context())
	if err != nil {
		log.Fatal(err.Error())
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusOK, products)
}

func (s *EchoServer) GetProductById(ctx echo.Context) error {
	var paramId = ctx.Param("id")
	ID, err := strconv.Atoi(paramId)
	if err != nil {
		return ctx.JSON(http.StatusUnsupportedMediaType, err)
	}

	result, err := s.DB.GetProductById(ctx.Request().Context(), uint(ID))
	if err != nil {
		switch err.(type) {
		case *dberrors.ConflictError:
			return ctx.JSON(http.StatusConflict, err)
		default:
			return ctx.JSON(http.StatusInternalServerError, err)
		}
	}

	return ctx.JSON(http.StatusOK, result)
}
