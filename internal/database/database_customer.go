package database

import (
	"context"
	"errors"

	"github.com/victorlabussiere/go-echo-gorm-example/internal/database/dberrors"
	"github.com/victorlabussiere/go-echo-gorm-example/internal/model"
	"gorm.io/gorm"
)

func (c Client) GetAllCustomers(ctx context.Context) ([]model.Customer, error) {
	var customers []model.Customer
	result := c.DB.WithContext(ctx).Find(&customers)
	return customers, result.Error
}

func (c Client) GetCustomerById(ctx context.Context, ID uint) (*model.Customer, error) {
	var customer = &model.Customer{}
	result := c.DB.WithContext(ctx).
		Where(&model.Customer{ID: ID}).
		First(&customer)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, &dberrors.NotFoundError{
				Entity: "customer",
				ID:     uint(ID),
			}
		}
		return nil, result.Error
	}

	return customer, nil

}

func (c Client) AddCustomer(ctx context.Context, customer *model.Customer) (*model.Customer, error) {
	result := c.DB.WithContext(ctx).Create(&customer)

	if result.Error != nil {
		return nil, result.Error
	}

	return customer, nil

}

func (c Client) UpdateCustomer(ctx context.Context, customer *model.Customer) (*model.Customer, error) {
	result := c.DB.WithContext(ctx).
		Where("id = ?", customer.ID).
		Updates(model.Customer{
			Name:  customer.Name,
			Email: customer.Email,
		})

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			return nil, &dberrors.ConflictError{}
		}

		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, &dberrors.NotFoundError{Entity: "Customer", ID: customer.ID}
	}

	return customer, nil
}

func (c Client) DeleteCustomerById(ctx context.Context, ID uint) error {
	return c.DB.WithContext(ctx).Delete(&model.Customer{ID: ID}).Error
}
