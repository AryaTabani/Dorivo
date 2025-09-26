package services

import (
	"context"

	"github.com/AryaTabani/Dorivo/models"
	"github.com/AryaTabani/Dorivo/repository"
)

func GetFAQs(ctx context.Context, tenantID string, category string) ([]models.FAQ, error) {
	return repository.GetFAQsByTenant(ctx, tenantID, category)
}
