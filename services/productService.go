package services

import (
	"context"

	"github.com/AryaTabani/Dorivo/models"
	"github.com/AryaTabani/Dorivo/repository"
)

func SearchProducts(ctx context.Context, tenantID string, filters map[string][]string) ([]models.Product, error) {
	return repository.SearchProducts(ctx, tenantID, filters)
}
func GetTags(ctx context.Context, tenantID string) ([]models.Tag, error) {
	return repository.GetTags(ctx, tenantID)
}
