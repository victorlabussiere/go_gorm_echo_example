package services

import (
	"context"

	"github.com/victorlabussiere/go_gorm_echo_postgres_example/initializer"
	"github.com/victorlabussiere/go_gorm_echo_postgres_example/models"
)

func AddCategory(ctx context.Context, category *models.Category) (*models.Category, error) {
	result := initializer.DB.WithContext(ctx).Create(&category)
	if result.Error != nil {
		return nil, result.Error
	}

	return category, nil
}

func GetCategories(ctx context.Context) (*[]models.Category, error) {
	var categories = new([]models.Category)

	result := initializer.DB.WithContext(ctx).Find(&categories)
	if result.Error != nil {
		return nil, result.Error
	}

	return categories, nil
}

func GetCategoriesById(ctx context.Context, ID uint) (*models.Category, error) {
	var category = new(models.Category)
	result := initializer.DB.WithContext(ctx).Where("id = ?", ID).First(&category)
	if result.Error != nil {
		return nil, result.Error
	}

	return category, nil
}
