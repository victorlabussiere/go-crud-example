package database

import (
	"context"

	"github.com/victorlabussiere/go-echo-gorm-example/internal/model"
)

func (c Client) GetAllProducts(ctx context.Context) ([]model.Product, error) {
	var products []model.Product
	result := c.DB.WithContext(ctx).Find(&products)
	return products, result.Error
}
