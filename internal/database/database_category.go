package database

import (
	"context"
	"errors"

	"github.com/victorlabussiere/go-echo-gorm-example/internal/database/dberrors"
	"github.com/victorlabussiere/go-echo-gorm-example/internal/model"
	"gorm.io/gorm"
)

func (c Client) AddCategory(ctx context.Context, category *model.Category) (*model.Category, error) {
	result := c.DB.WithContext(ctx).Create(&category)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			return nil, &dberrors.ConflictError{}
		}

		return nil, result.Error
	}

	return category, nil
}

func (c Client) GetCategories(ctx context.Context) (*[]model.Category, error) {
	var categories = new([]model.Category)

	result := c.DB.WithContext(ctx).Find(&categories)
	if result.Error != nil {
		return nil, result.Error
	}

	return categories, nil
}

func (c Client) GetCategoriesById(ctx context.Context, ID uint) (*model.Category, error) {
	var category = new(model.Category)
	result := c.DB.WithContext(ctx).Where("id = ?", ID).First(&category)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, &dberrors.NotFoundError{Entity: "Category", ID: ID}
		}

		return nil, result.Error
	}

	return category, nil
}
