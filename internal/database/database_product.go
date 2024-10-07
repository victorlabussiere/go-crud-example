package database

import (
	"context"
	"errors"

	"github.com/victorlabussiere/go-echo-gorm-example/internal/database/dberrors"
	"github.com/victorlabussiere/go-echo-gorm-example/internal/model"
	"gorm.io/gorm"
)

func (c Client) AddProduct(ctx context.Context, product *model.Product) (*model.Product, error) {
	result := c.DB.WithContext(ctx).Create(product)
	if result.Error != nil {
		return nil, result.Error
	}

	return product, nil
}

func (c Client) GetAllProducts(ctx context.Context) ([]model.Product, error) {
	var products []model.Product
	result := c.DB.WithContext(ctx).Find(&products)
	return products, result.Error
}

func (c Client) GetProductById(ctx context.Context, ID uint) (*model.Product, error) {
	var product = &model.Product{}
	result := c.DB.WithContext(ctx).Where("id = ?", ID).First(&product)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, &dberrors.NotFoundError{Entity: "Product", ID: ID}
		}
		return nil, result.Error
	}

	return product, nil
}
