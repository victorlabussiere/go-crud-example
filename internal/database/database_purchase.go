package database

import (
	"context"

	"github.com/victorlabussiere/go-echo-gorm-example/internal/model"
)

func (c Client) AddPurchase(ctx context.Context, purchase *model.Purchase) (*model.Purchase, error) {
	result := c.DB.WithContext(ctx).Create(purchase)
	if result.Error != nil {
		return nil, result.Error
	}

	return purchase, nil
}
