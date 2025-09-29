package services

import (
	"context"
	"database/sql"
	"errors"

	"github.com/AryaTabani/Dorivo/models"
	"github.com/AryaTabani/Dorivo/repository"
)

var ErrProductNotFound = errors.New("product not found")

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
