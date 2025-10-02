package services

import (
	"context"

	"github.com/AryaTabani/Dorivo/models"
	"github.com/AryaTabani/Dorivo/repository"
)

func CreateProduct(ctx context.Context, tenantID string, payload *models.ProductPayload) (int64, error) {
	return repository.CreateProduct(ctx, tenantID, payload)
}

func UpdateProduct(ctx context.Context, tenantID string, productID int64, payload *models.ProductPayload) error {
	return repository.UpdateProduct(ctx, tenantID, productID, payload)
}

func DeleteProduct(ctx context.Context, tenantID string, productID int64) error {
	return repository.DeleteProduct(ctx, tenantID, productID)
}

func UpdateTenantConfig(ctx context.Context, tenantID string, payload *models.TenantConfig) error {
	return repository.UpdateTenantConfig(ctx, tenantID, payload)
}
