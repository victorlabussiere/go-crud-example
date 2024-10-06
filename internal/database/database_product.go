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

func (c Client) AddProduct(ctx context.Context, product *model.Product) (*model.Product, error) {
	result := c.DB.WithContext(ctx).Create(product)
	if result.Error != nil {
		return nil, result.Error
	}

	return product, nil
}
