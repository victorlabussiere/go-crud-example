package server

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/victorlabussiere/go-echo-gorm-example/internal/database/dberrors"
	"github.com/victorlabussiere/go-echo-gorm-example/internal/model"
)

func (s *EchoServer) AddCategory(ctx echo.Context) error {
	var category = new(model.Category)

	if err := ctx.Bind(category); err != nil {
		return ctx.JSON(http.StatusUnsupportedMediaType, err)
	}

	category, err := s.DB.AddCategory(ctx.Request().Context(), category)
	if err != nil {
		switch err.(type) {
		case *dberrors.ConflictError:
			return ctx.JSON(http.StatusConflict, err)
		default:
			return ctx.JSON(http.StatusInternalServerError, err)
		}
	}

	return ctx.JSON(http.StatusCreated, category)
}

func (s *EchoServer) GetAllCategories(ctx echo.Context) error {
	categories, err := s.DB.GetCategories(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusOK, categories)
}

func (s *EchoServer) GetCategoryById(ctx echo.Context) error {
	paramId := ctx.Param("id")
	ID, err := strconv.Atoi(paramId)
	if err != nil {
		return ctx.JSON(http.StatusUnsupportedMediaType, err)
	}

	result, err := s.DB.GetCategoriesById(ctx.Request().Context(), uint(ID))
	if err != nil {
		switch err.(type) {
		case *dberrors.NotFoundError:
			return ctx.JSON(http.StatusNotFound, err)
		default:
			return ctx.JSON(http.StatusInternalServerError, err)
		}
	}

	var response struct {
		Id   uint
		Name string
	}
	response.Id = result.ID
	response.Name = result.Name

	return ctx.JSON(http.StatusOK, response)
}
