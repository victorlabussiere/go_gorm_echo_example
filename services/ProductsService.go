package services

import (
	"context"

	"github.com/victorlabussiere/go_gorm_echo_postgres_example/initializer"
	"github.com/victorlabussiere/go_gorm_echo_postgres_example/models"
)

func AddProduct(ctx context.Context, product *models.Product) (*models.Product, error) {
	result := initializer.DB.WithContext(ctx).Create(product)
	if result.Error != nil {
		return nil, result.Error
	}

	return product, nil
}

func GetAllProducts(ctx context.Context) ([]models.Product, error) {
	var products []models.Product
	result := initializer.DB.WithContext(ctx).Find(&products)
	return products, result.Error
}

func GetProductById(ctx context.Context, ID uint) (*models.Product, error) {
	var product = &models.Product{}
	result := initializer.DB.WithContext(ctx).Where("id = ?", ID).First(&product)
	if result.Error != nil {
		return nil, result.Error
	}

	return product, nil
}

func GetProductByCategoryId(ctx context.Context, ID uint) ([]models.Product, error) {
	var products []models.Product

	result := initializer.DB.WithContext(ctx).Where("category_id = ?", ID).Find(&products)
	if result.Error != nil {
		return nil, result.Error
	}

	return products, nil
}
