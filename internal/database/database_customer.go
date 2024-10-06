package database

import (
	"context"

	"github.com/victorlabussiere/go-echo-gorm-example/internal/model"
)

func (c Client) GetAllCustomers(ctx context.Context) ([]model.Customer, error) {
	var customers []model.Customer
	result := c.DB.WithContext(ctx).Find(&customers)
	return customers, result.Error
}

func (c Client) AddCustomer(ctx context.Context, customer *model.Customer) (*model.Customer, error) {
	result := c.DB.WithContext(ctx).Create(&customer)

	if result.Error != nil {
		return nil, result.Error
	}

	return customer, nil

}
