package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/victorlabussiere/go-echo-gorm-example/internal/database/dberrors"
	"github.com/victorlabussiere/go-echo-gorm-example/internal/model"
)

func (e *EchoServer) AddPurchase(ctx echo.Context) error {
	purchase := new(model.Purchase)
	if err := ctx.Bind(purchase); err != nil {
		return ctx.JSON(http.StatusUnsupportedMediaType, err)
	}

	purchase, err := e.DB.AddPurchase(ctx.Request().Context(), purchase)
	if err != nil {
		switch err.(type) {
		case *dberrors.ConflictError:
			return ctx.JSON(http.StatusConflict, err)
		default:
			return ctx.JSON(http.StatusInternalServerError, err)
		}
	}

	return ctx.JSON(http.StatusCreated, purchase)
}
