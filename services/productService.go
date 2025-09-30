package services

import (
	"context"
	"database/sql"
	"errors"

	"github.com/AryaTabani/Dorivo/models"
	"github.com/AryaTabani/Dorivo/repository"
)

var ErrProductNotFound = errors.New("product not found")

const BEST_SELLER_LIMIT = 10

func SearchProducts(ctx context.Context, tenantID string, filters map[string][]string) ([]models.Product, error) {
	return repository.SearchProducts(ctx, tenantID, filters)
}
func GetTags(ctx context.Context, tenantID string) ([]models.Tag, error) {
	return repository.GetTags(ctx, tenantID)
}

func GetProductDetails(ctx context.Context, tenantID string, productID int64) (*models.Product, error) {
	product, err := repository.GetProductDetails(ctx, tenantID, productID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrProductNotFound
		}
		return nil, err
	}
	return product, nil
}
func GetBestSellers(ctx context.Context, tenantID string) ([]models.Product, error) {
	return repository.GetBestSellers(ctx, tenantID, BEST_SELLER_LIMIT)
}

func GetFeaturedProduct(ctx context.Context, tenantID string) (*models.Product, error) {
	product, err := repository.GetFeaturedProduct(ctx, tenantID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrProductNotFound
		}
		return nil, err
	}
	return product, nil
}

func GetRecommendedProducts(ctx context.Context, tenantID string) ([]models.Product, error) {
	return repository.GetRecommendedProducts(ctx, tenantID)
}
